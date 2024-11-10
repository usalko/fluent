// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqlgraph

import (
	"fmt"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/fluent_ql"
)

type (
	// A Schema holds a representation of fluent/schema at runtime. Each Node
	// represents a single schema-type and its relations in the graph (storage).
	//
	// It is used for translating common graph traversal operations to the
	// underlying SQL storage. For example, an operation like `has_edge(E)`,
	// will be translated to an SQL lookup based on the relation type and the
	// FK configuration.
	Schema struct {
		Nodes []*Node
	}

	// A Node in the graph holds the SQL information for an fluent/schema.
	Node struct {
		NodeSpec

		// Type holds the node type (schema name).
		Type string

		// Fields maps from field names to their spec.
		Fields map[string]*FieldSpec

		// Edges maps from edge names to their spec.
		Edges map[string]struct {
			To   *Node
			Spec *EdgeSpec
		}
	}
)

// AddE adds an edge to the graph. It fails, if one of the node
// types is missing.
//
//	g.AddE("pets", spec, "user", "pet")
//	g.AddE("friends", spec, "user", "user")
func (g *Schema) AddE(name string, spec *EdgeSpec, from, to string) error {
	var fromT, toT *Node
	for i := range g.Nodes {
		t := g.Nodes[i].Type
		if t == from {
			fromT = g.Nodes[i]
		}
		if t == to {
			toT = g.Nodes[i]
		}
	}
	if fromT == nil || toT == nil {
		return fmt.Errorf("from/to type was not found")
	}
	if fromT.Edges == nil {
		fromT.Edges = make(map[string]struct {
			To   *Node
			Spec *EdgeSpec
		})
	}
	fromT.Edges[name] = struct {
		To   *Node
		Spec *EdgeSpec
	}{
		To:   toT,
		Spec: spec,
	}
	return nil
}

// MustAddE is like AddE but panics if the edge can be added to the graph.
func (g *Schema) MustAddE(name string, spec *EdgeSpec, from, to string) {
	if err := g.AddE(name, spec, from, to); err != nil {
		panic(err)
	}
}

// EvalP evaluates the fluent_ql predicate on the given selector (query builder).
func (g *Schema) EvalP(nodeType string, p fluent_ql.P, selector *sql.Selector) error {
	var node *Node
	for i := range g.Nodes {
		if g.Nodes[i].Type == nodeType {
			node = g.Nodes[i]
			break
		}
	}
	if node == nil {
		return fmt.Errorf("node %s was not found in the graph schema", nodeType)
	}
	pr, err := evalExpr(node, selector, p)
	if err != nil {
		return err
	}
	selector.Where(pr)
	return nil
}

// FuncSelector represents a selector function to be used as an fluent_ql foreign-function.
const FuncSelector fluent_ql.Func = "func_selector"

// wrappedFunc wraps the selector-function to an ent-expression.
type wrappedFunc struct {
	fluent_ql.Expr
	Func func(*sql.Selector)
}

// WrapFunc wraps a selector-func with an fluent_ql call expression.
func WrapFunc(s func(*sql.Selector)) *fluent_ql.CallExpr {
	return &fluent_ql.CallExpr{
		Func: FuncSelector,
		Args: []fluent_ql.Expr{wrappedFunc{Func: s}},
	}
}

var (
	binary = [...]sql.Op{
		fluent_ql.OpEQ:    sql.OpEQ,
		fluent_ql.OpNEQ:   sql.OpNEQ,
		fluent_ql.OpGT:    sql.OpGT,
		fluent_ql.OpGTE:   sql.OpGTE,
		fluent_ql.OpLT:    sql.OpLT,
		fluent_ql.OpLTE:   sql.OpLTE,
		fluent_ql.OpIn:    sql.OpIn,
		fluent_ql.OpNotIn: sql.OpNotIn,
	}
	nary = [...]func(...*sql.Predicate) *sql.Predicate{
		fluent_ql.OpAnd: sql.And,
		fluent_ql.OpOr:  sql.Or,
	}
	strFunc = map[fluent_ql.Func]func(string, string) *sql.Predicate{
		fluent_ql.FuncContains:     sql.Contains,
		fluent_ql.FuncContainsFold: sql.ContainsFold,
		fluent_ql.FuncEqualFold:    sql.EqualFold,
		fluent_ql.FuncHasPrefix:    sql.HasPrefix,
		fluent_ql.FuncHasSuffix:    sql.HasSuffix,
	}
	nullFunc = [...]func(string) *sql.Predicate{
		fluent_ql.OpEQ:  sql.IsNull,
		fluent_ql.OpNEQ: sql.NotNull,
	}
)

// state represents the state of a predicate evaluation.
// Note that, the evaluation output is a predicate to be
// applied on the database.
type state struct {
	sql.Builder
	context  *Node
	selector *sql.Selector
}

// evalExpr evaluates the fluent_ql expression and returns a new SQL predicate to be applied on the database.
func evalExpr(context *Node, selector *sql.Selector, expr fluent_ql.Expr) (p *sql.Predicate, err error) {
	ex := &state{
		context:  context,
		selector: selector,
	}
	defer catch(&err)
	p = ex.evalExpr(expr)
	return
}

// evalExpr evaluates any expression.
func (e *state) evalExpr(expr fluent_ql.Expr) *sql.Predicate {
	switch expr := expr.(type) {
	case *fluent_ql.BinaryExpr:
		return e.evalBinary(expr)
	case *fluent_ql.UnaryExpr:
		return sql.Not(e.evalExpr(expr.X))
	case *fluent_ql.NaryExpr:
		ps := make([]*sql.Predicate, len(expr.Xs))
		for i, x := range expr.Xs {
			ps[i] = e.evalExpr(x)
		}
		return nary[expr.Op](ps...)
	case *fluent_ql.CallExpr:
		switch expr.Func {
		case fluent_ql.FuncHasPrefix, fluent_ql.FuncHasSuffix, fluent_ql.FuncContains, fluent_ql.FuncEqualFold, fluent_ql.FuncContainsFold:
			expect(len(expr.Args) == 2, "invalid number of arguments for %s", expr.Func)
			f, ok := expr.Args[0].(*fluent_ql.Field)
			expect(ok, "*fluent_ql.Field, got %T", expr.Args[0])
			v, ok := expr.Args[1].(*fluent_ql.Value)
			expect(ok, "*fluent_ql.Value, got %T", expr.Args[1])
			s, ok := v.V.(string)
			expect(ok, "string value, got %T", v.V)
			return strFunc[expr.Func](e.field(f), s)
		case fluent_ql.FuncHasEdge:
			expect(len(expr.Args) > 0, "invalid number of arguments for %s", expr.Func)
			edge, ok := expr.Args[0].(*fluent_ql.Edge)
			expect(ok, "*fluent_ql.Edge, got %T", expr.Args[0])
			return e.evalEdge(edge.Name, expr.Args[1:]...)
		}
	}
	panic("invalid")
}

// evalBinary evaluates binary expressions.
func (e *state) evalBinary(expr *fluent_ql.BinaryExpr) *sql.Predicate {
	switch expr.Op {
	case fluent_ql.OpOr:
		return sql.Or(e.evalExpr(expr.X), e.evalExpr(expr.Y))
	case fluent_ql.OpAnd:
		return sql.And(e.evalExpr(expr.X), e.evalExpr(expr.Y))
	case fluent_ql.OpEQ, fluent_ql.OpNEQ:
		if expr.Y == (*fluent_ql.Value)(nil) {
			f, ok := expr.X.(*fluent_ql.Field)
			expect(ok, "*fluent_ql.Field, got %T", expr.Y)
			return nullFunc[expr.Op](e.field(f))
		}
		fallthrough
	default:
		field, ok := expr.X.(*fluent_ql.Field)
		expect(ok, "expr.X to be *fluent_ql.Field (got %T)", expr.X)
		_, ok = expr.Y.(*fluent_ql.Field)
		if !ok {
			_, ok = expr.Y.(*fluent_ql.Value)
		}
		expect(ok, "expr.Y to be *fluent_ql.Field or *fluent_ql.Value (got %T)", expr.X)
		switch x := expr.Y.(type) {
		case *fluent_ql.Field:
			return sql.ColumnsOp(e.field(field), e.field(x), binary[expr.Op])
		case *fluent_ql.Value:
			c := e.field(field)
			return sql.P(func(b *sql.Builder) {
				b.Ident(c).WriteOp(binary[expr.Op])
				args(b, x)
			})
		default:
			panic("unreachable")
		}
	}
}

// evalEdge evaluates has-edge and has-edge-with calls.
func (e *state) evalEdge(name string, exprs ...fluent_ql.Expr) *sql.Predicate {
	edge, ok := e.context.Edges[name]
	expect(ok, "edge %q was not found for node %q", name, e.context.Type)
	var fromC, toC string
	switch {
	case edge.To.ID != nil:
		toC = edge.To.ID.Column
	// Edge-owner points to its edge schema.
	case edge.To.CompositeID != nil && !edge.Spec.Inverse:
		toC = edge.To.CompositeID[0].Column
	// Edge-backref points to its edge schema.
	case edge.To.CompositeID != nil && edge.Spec.Inverse:
		toC = edge.To.CompositeID[1].Column
	default:
		panic(evalError{fmt.Sprintf("expect id definition for edge %q", name)})
	}
	switch {
	case e.context.ID != nil:
		fromC = e.context.ID.Column
	case e.context.CompositeID != nil && (edge.Spec.Rel == M2O || (edge.Spec.Rel == O2O && edge.Spec.Inverse)):
		// An edge-schema with a composite id can query
		// only edges that it owns (holds the foreign-key).
	default:
		panic(evalError{fmt.Sprintf("unexpected edge-query from an edge-schema %q", e.context.Type)})
	}
	step := NewStep(
		From(e.context.Table, fromC),
		To(edge.To.Table, toC),
		Edge(edge.Spec.Rel, edge.Spec.Inverse, edge.Spec.Table, edge.Spec.Columns...),
	)
	selector := e.selector.Clone().SetP(nil)
	selector.SetTotal(e.Total())
	if len(exprs) == 0 {
		HasNeighbors(selector, step)
		return selector.P()
	}
	HasNeighborsWith(selector, step, func(s *sql.Selector) {
		for _, expr := range exprs {
			if cx, ok := expr.(*fluent_ql.CallExpr); ok && cx.Func == FuncSelector {
				expect(len(cx.Args) == 1, "invalid number of arguments for %s", FuncSelector)
				wrapped, ok := cx.Args[0].(wrappedFunc)
				expect(ok, "invalid argument for %s: %T", FuncSelector, cx.Args[0])
				wrapped.Func(s)
			} else {
				p, err := evalExpr(edge.To, s, expr)
				expect(err == nil, "edge evaluation failed for %s->%s: %s", e.context.Type, name, err)
				s.Where(p)
			}
		}
	})
	return selector.P()
}

func (e *state) field(f *fluent_ql.Field) string {
	_, ok := e.context.Fields[f.Name]
	expect(ok || e.context.ID.Column == f.Name, "field %q was not found for node %q", f.Name, e.context.Type)
	return e.selector.C(f.Name)
}

func args(b *sql.Builder, v *fluent_ql.Value) {
	vs, ok := v.V.([]any)
	if !ok {
		b.Arg(v.V)
		return
	}
	b.WriteByte('(').Args(vs...).WriteByte(')')
}

// expect panics if the condition is false.
func expect(cond bool, msg string, args ...any) {
	if !cond {
		panic(evalError{fmt.Sprintf("expect "+msg, args...)})
	}
}

type evalError struct {
	msg string
}

func (p evalError) Error() string {
	return fmt.Sprintf("sqlgraph: %s", p.msg)
}

func catch(err *error) {
	if e := recover(); e != nil {
		xerr, ok := e.(evalError)
		if !ok {
			panic(e)
		}
		*err = xerr
	}
}
