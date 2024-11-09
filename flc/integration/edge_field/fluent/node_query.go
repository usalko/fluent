// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/edge_field/fluent/node"
	"github.com/usalko/fluent/flc/integration/edge_field/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// NodeQuery is the builder for querying Node entities.
type NodeQuery struct {
	config
	ctx        *QueryContext
	order      []node.OrderOption
	inters     []Interceptor
	predicates []predicate.Node
	withPrev   *NodeQuery
	withNext   *NodeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NodeQuery builder.
func (nq *NodeQuery) Where(ps ...predicate.Node) *NodeQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NodeQuery) Limit(limit int) *NodeQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NodeQuery) Offset(offset int) *NodeQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NodeQuery) Unique(unique bool) *NodeQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NodeQuery) Order(o ...node.OrderOption) *NodeQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryPrev chains the current query on the "prev" edge.
func (nq *NodeQuery) QueryPrev() *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, node.PrevTable, node.PrevColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNext chains the current query on the "next" edge.
func (nq *NodeQuery) QueryNext() *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, node.NextTable, node.NextColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Node entity from the query.
// Returns a *NotFoundError when no Node was found.
func (nq *NodeQuery) First(ctx context.Context) (*Node, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{node.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NodeQuery) FirstX(ctx context.Context) *Node {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Node ID from the query.
// Returns a *NotFoundError when no Node ID was found.
func (nq *NodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, fluent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{node.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NodeQuery) FirstIDX(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Node entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Node entity is found.
// Returns a *NotFoundError when no Node entities are found.
func (nq *NodeQuery) Only(ctx context.Context) (*Node, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{node.Label}
	default:
		return nil, &NotSingularError{node.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NodeQuery) OnlyX(ctx context.Context) *Node {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Node ID in the query.
// Returns a *NotSingularError when more than one Node ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, fluent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{node.Label}
	default:
		err = &NotSingularError{node.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Nodes.
func (nq *NodeQuery) All(ctx context.Context) ([]*Node, error) {
	ctx = setContextOp(ctx, nq.ctx, fluent.OpQueryAll)
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Node, *NodeQuery]()
	return withInterceptors[[]*Node](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NodeQuery) AllX(ctx context.Context) []*Node {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Node IDs.
func (nq *NodeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, fluent.OpQueryIDs)
	if err = nq.Select(node.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NodeQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, fluent.OpQueryCount)
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NodeQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NodeQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, fluent.OpQueryExist)
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("fluent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NodeQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NodeQuery) Clone() *NodeQuery {
	if nq == nil {
		return nil
	}
	return &NodeQuery{
		config:     nq.config,
		ctx:        nq.ctx.Clone(),
		order:      append([]node.OrderOption{}, nq.order...),
		inters:     append([]Interceptor{}, nq.inters...),
		predicates: append([]predicate.Node{}, nq.predicates...),
		withPrev:   nq.withPrev.Clone(),
		withNext:   nq.withNext.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithPrev tells the query-builder to eager-load the nodes that are connected to
// the "prev" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithPrev(opts ...func(*NodeQuery)) *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withPrev = query
	return nq
}

// WithNext tells the query-builder to eager-load the nodes that are connected to
// the "next" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithNext(opts ...func(*NodeQuery)) *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withNext = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value int `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Node.Query().
//		GroupBy(node.FieldValue).
//		Aggregate(fluent.Count()).
//		Scan(ctx, &v)
func (nq *NodeQuery) GroupBy(field string, fields ...string) *NodeGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NodeGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = node.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value int `json:"value,omitempty"`
//	}
//
//	client.Node.Query().
//		Select(node.FieldValue).
//		Scan(ctx, &v)
func (nq *NodeQuery) Select(fields ...string) *NodeSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NodeSelect{NodeQuery: nq}
	sbuild.label = node.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NodeSelect configured with the given aggregations.
func (nq *NodeQuery) Aggregate(fns ...AggregateFunc) *NodeSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("fluent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !node.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Node, error) {
	var (
		nodes       = []*Node{}
		_spec       = nq.querySpec()
		loadedTypes = [2]bool{
			nq.withPrev != nil,
			nq.withNext != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Node).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Node{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withPrev; query != nil {
		if err := nq.loadPrev(ctx, query, nodes, nil,
			func(n *Node, e *Node) { n.Edges.Prev = e }); err != nil {
			return nil, err
		}
	}
	if query := nq.withNext; query != nil {
		if err := nq.loadNext(ctx, query, nodes, nil,
			func(n *Node, e *Node) { n.Edges.Next = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NodeQuery) loadPrev(ctx context.Context, query *NodeQuery, nodes []*Node, init func(*Node), assign func(*Node, *Node)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Node)
	for i := range nodes {
		fk := nodes[i].PrevID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(node.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "prev_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (nq *NodeQuery) loadNext(ctx context.Context, query *NodeQuery, nodes []*Node, init func(*Node), assign func(*Node, *Node)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Node)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(node.FieldPrevID)
	}
	query.Where(predicate.Node(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(node.NextColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PrevID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "prev_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (nq *NodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(node.Table, node.Columns, sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, node.FieldID)
		for i := range fields {
			if fields[i] != node.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if nq.withPrev != nil {
			_spec.Node.AddColumnOnce(node.FieldPrevID)
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(node.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = node.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NodeGroupBy is the group-by builder for Node entities.
type NodeGroupBy struct {
	selector
	build *NodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NodeGroupBy) Aggregate(fns ...AggregateFunc) *NodeGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, fluent.OpQueryGroupBy)
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NodeQuery, *NodeGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NodeGroupBy) sqlScan(ctx context.Context, root *NodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NodeSelect is the builder for selecting fields of Node entities.
type NodeSelect struct {
	*NodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NodeSelect) Aggregate(fns ...AggregateFunc) *NodeSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, fluent.OpQuerySelect)
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NodeQuery, *NodeSelect](ctx, ns.NodeQuery, ns, ns.inters, v)
}

func (ns *NodeSelect) sqlScan(ctx context.Context, root *NodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
