// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"time"

	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/__"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/p"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/node"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
)

// NodeUpdate is the builder for updating Node entities.
type NodeUpdate struct {
	config
	hooks    []Hook
	mutation *NodeMutation
}

// Where appends a list predicates to the NodeUpdate builder.
func (nu *NodeUpdate) Where(ps ...predicate.Node) *NodeUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetValue sets the "value" field.
func (nu *NodeUpdate) SetValue(i int) *NodeUpdate {
	nu.mutation.ResetValue()
	nu.mutation.SetValue(i)
	return nu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableValue(i *int) *NodeUpdate {
	if i != nil {
		nu.SetValue(*i)
	}
	return nu
}

// AddValue adds i to the "value" field.
func (nu *NodeUpdate) AddValue(i int) *NodeUpdate {
	nu.mutation.AddValue(i)
	return nu
}

// ClearValue clears the value of the "value" field.
func (nu *NodeUpdate) ClearValue() *NodeUpdate {
	nu.mutation.ClearValue()
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NodeUpdate) SetUpdatedAt(t time.Time) *NodeUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (nu *NodeUpdate) ClearUpdatedAt() *NodeUpdate {
	nu.mutation.ClearUpdatedAt()
	return nu
}

// SetPrevID sets the "prev" edge to the Node entity by ID.
func (nu *NodeUpdate) SetPrevID(id string) *NodeUpdate {
	nu.mutation.SetPrevID(id)
	return nu
}

// SetNillablePrevID sets the "prev" edge to the Node entity by ID if the given value is not nil.
func (nu *NodeUpdate) SetNillablePrevID(id *string) *NodeUpdate {
	if id != nil {
		nu = nu.SetPrevID(*id)
	}
	return nu
}

// SetPrev sets the "prev" edge to the Node entity.
func (nu *NodeUpdate) SetPrev(n *Node) *NodeUpdate {
	return nu.SetPrevID(n.ID)
}

// SetNextID sets the "next" edge to the Node entity by ID.
func (nu *NodeUpdate) SetNextID(id string) *NodeUpdate {
	nu.mutation.SetNextID(id)
	return nu
}

// SetNillableNextID sets the "next" edge to the Node entity by ID if the given value is not nil.
func (nu *NodeUpdate) SetNillableNextID(id *string) *NodeUpdate {
	if id != nil {
		nu = nu.SetNextID(*id)
	}
	return nu
}

// SetNext sets the "next" edge to the Node entity.
func (nu *NodeUpdate) SetNext(n *Node) *NodeUpdate {
	return nu.SetNextID(n.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (nu *NodeUpdate) Mutation() *NodeMutation {
	return nu.mutation
}

// ClearPrev clears the "prev" edge to the Node entity.
func (nu *NodeUpdate) ClearPrev() *NodeUpdate {
	nu.mutation.ClearPrev()
	return nu
}

// ClearNext clears the "next" edge to the Node entity.
func (nu *NodeUpdate) ClearNext() *NodeUpdate {
	nu.mutation.ClearNext()
	return nu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NodeUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks(ctx, nu.gremlinSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NodeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NodeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NodeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NodeUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok && !nu.mutation.UpdatedAtCleared() {
		v := node.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

func (nu *NodeUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := nu.gremlin().Query()
	if err := nu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	nu.mutation.done = true
	return res.ReadInt()
}

func (nu *NodeUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(node.Label)
	for _, p := range nu.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := nu.mutation.Value(); ok {
		v.Property(dsl.Single, node.FieldValue, value)
	}
	if value, ok := nu.mutation.AddedValue(); ok {
		v.Property(dsl.Single, node.FieldValue, __.Union(__.Values(node.FieldValue), __.Constant(value)).Sum())
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		v.Property(dsl.Single, node.FieldUpdatedAt, value)
	}
	var properties []any
	if nu.mutation.ValueCleared() {
		properties = append(properties, node.FieldValue)
	}
	if nu.mutation.UpdatedAtCleared() {
		properties = append(properties, node.FieldUpdatedAt)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if nu.mutation.PrevCleared() {
		tr := rv.Clone().InE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range nu.mutation.PrevIDs() {
		v.AddE(node.NextLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	if nu.mutation.NextCleared() {
		tr := rv.Clone().OutE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range nu.mutation.NextIDs() {
		v.AddE(node.NextLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// NodeUpdateOne is the builder for updating a single Node entity.
type NodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodeMutation
}

// SetValue sets the "value" field.
func (nuo *NodeUpdateOne) SetValue(i int) *NodeUpdateOne {
	nuo.mutation.ResetValue()
	nuo.mutation.SetValue(i)
	return nuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableValue(i *int) *NodeUpdateOne {
	if i != nil {
		nuo.SetValue(*i)
	}
	return nuo
}

// AddValue adds i to the "value" field.
func (nuo *NodeUpdateOne) AddValue(i int) *NodeUpdateOne {
	nuo.mutation.AddValue(i)
	return nuo
}

// ClearValue clears the value of the "value" field.
func (nuo *NodeUpdateOne) ClearValue() *NodeUpdateOne {
	nuo.mutation.ClearValue()
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NodeUpdateOne) SetUpdatedAt(t time.Time) *NodeUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (nuo *NodeUpdateOne) ClearUpdatedAt() *NodeUpdateOne {
	nuo.mutation.ClearUpdatedAt()
	return nuo
}

// SetPrevID sets the "prev" edge to the Node entity by ID.
func (nuo *NodeUpdateOne) SetPrevID(id string) *NodeUpdateOne {
	nuo.mutation.SetPrevID(id)
	return nuo
}

// SetNillablePrevID sets the "prev" edge to the Node entity by ID if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillablePrevID(id *string) *NodeUpdateOne {
	if id != nil {
		nuo = nuo.SetPrevID(*id)
	}
	return nuo
}

// SetPrev sets the "prev" edge to the Node entity.
func (nuo *NodeUpdateOne) SetPrev(n *Node) *NodeUpdateOne {
	return nuo.SetPrevID(n.ID)
}

// SetNextID sets the "next" edge to the Node entity by ID.
func (nuo *NodeUpdateOne) SetNextID(id string) *NodeUpdateOne {
	nuo.mutation.SetNextID(id)
	return nuo
}

// SetNillableNextID sets the "next" edge to the Node entity by ID if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableNextID(id *string) *NodeUpdateOne {
	if id != nil {
		nuo = nuo.SetNextID(*id)
	}
	return nuo
}

// SetNext sets the "next" edge to the Node entity.
func (nuo *NodeUpdateOne) SetNext(n *Node) *NodeUpdateOne {
	return nuo.SetNextID(n.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (nuo *NodeUpdateOne) Mutation() *NodeMutation {
	return nuo.mutation
}

// ClearPrev clears the "prev" edge to the Node entity.
func (nuo *NodeUpdateOne) ClearPrev() *NodeUpdateOne {
	nuo.mutation.ClearPrev()
	return nuo
}

// ClearNext clears the "next" edge to the Node entity.
func (nuo *NodeUpdateOne) ClearNext() *NodeUpdateOne {
	nuo.mutation.ClearNext()
	return nuo
}

// Where appends a list predicates to the NodeUpdate builder.
func (nuo *NodeUpdateOne) Where(ps ...predicate.Node) *NodeUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NodeUpdateOne) Select(field string, fields ...string) *NodeUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Node entity.
func (nuo *NodeUpdateOne) Save(ctx context.Context) (*Node, error) {
	nuo.defaults()
	return withHooks(ctx, nuo.gremlinSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NodeUpdateOne) SaveX(ctx context.Context) *Node {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NodeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NodeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NodeUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok && !nuo.mutation.UpdatedAtCleared() {
		v := node.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

func (nuo *NodeUpdateOne) gremlinSave(ctx context.Context) (*Node, error) {
	res := &gremlin.Response{}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Node.id" for update`)}
	}
	query, bindings := nuo.gremlin(id).Query()
	if err := nuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	nuo.mutation.done = true
	n := &Node{config: nuo.config}
	if err := n.FromResponse(res); err != nil {
		return nil, err
	}
	return n, nil
}

func (nuo *NodeUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := nuo.mutation.Value(); ok {
		v.Property(dsl.Single, node.FieldValue, value)
	}
	if value, ok := nuo.mutation.AddedValue(); ok {
		v.Property(dsl.Single, node.FieldValue, __.Union(__.Values(node.FieldValue), __.Constant(value)).Sum())
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		v.Property(dsl.Single, node.FieldUpdatedAt, value)
	}
	var properties []any
	if nuo.mutation.ValueCleared() {
		properties = append(properties, node.FieldValue)
	}
	if nuo.mutation.UpdatedAtCleared() {
		properties = append(properties, node.FieldUpdatedAt)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if nuo.mutation.PrevCleared() {
		tr := rv.Clone().InE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range nuo.mutation.PrevIDs() {
		v.AddE(node.NextLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	if nuo.mutation.NextCleared() {
		tr := rv.Clone().OutE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range nuo.mutation.NextIDs() {
		v.AddE(node.NextLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	if len(nuo.fields) > 0 {
		fields := make([]any, 0, len(nuo.fields)+1)
		fields = append(fields, true)
		for _, f := range nuo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}