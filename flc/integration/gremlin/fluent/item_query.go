// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"fmt"
	"math"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/__"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/item"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
)

// ItemQuery is the builder for querying Item entities.
type ItemQuery struct {
	config
	ctx        *QueryContext
	order      []item.OrderOption
	inters     []Interceptor
	predicates []predicate.Item
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the ItemQuery builder.
func (iq *ItemQuery) Where(ps ...predicate.Item) *ItemQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *ItemQuery) Limit(limit int) *ItemQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *ItemQuery) Offset(offset int) *ItemQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *ItemQuery) Unique(unique bool) *ItemQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *ItemQuery) Order(o ...item.OrderOption) *ItemQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// First returns the first Item entity from the query.
// Returns a *NotFoundError when no Item was found.
func (iq *ItemQuery) First(ctx context.Context) (*Item, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{item.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *ItemQuery) FirstX(ctx context.Context) *Item {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Item ID from the query.
// Returns a *NotFoundError when no Item ID was found.
func (iq *ItemQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, fluent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{item.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *ItemQuery) FirstIDX(ctx context.Context) string {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Item entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Item entity is found.
// Returns a *NotFoundError when no Item entities are found.
func (iq *ItemQuery) Only(ctx context.Context) (*Item, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{item.Label}
	default:
		return nil, &NotSingularError{item.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *ItemQuery) OnlyX(ctx context.Context) *Item {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Item ID in the query.
// Returns a *NotSingularError when more than one Item ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *ItemQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, fluent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{item.Label}
	default:
		err = &NotSingularError{item.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *ItemQuery) OnlyIDX(ctx context.Context) string {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Items.
func (iq *ItemQuery) All(ctx context.Context) ([]*Item, error) {
	ctx = setContextOp(ctx, iq.ctx, fluent.OpQueryAll)
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Item, *ItemQuery]()
	return withInterceptors[[]*Item](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *ItemQuery) AllX(ctx context.Context) []*Item {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Item IDs.
func (iq *ItemQuery) IDs(ctx context.Context) (ids []string, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, fluent.OpQueryIDs)
	if err = iq.Select(item.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *ItemQuery) IDsX(ctx context.Context) []string {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *ItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, fluent.OpQueryCount)
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*ItemQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *ItemQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *ItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, fluent.OpQueryExist)
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("fluent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *ItemQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *ItemQuery) Clone() *ItemQuery {
	if iq == nil {
		return nil
	}
	return &ItemQuery{
		config:     iq.config,
		ctx:        iq.ctx.Clone(),
		order:      append([]item.OrderOption{}, iq.order...),
		inters:     append([]Interceptor{}, iq.inters...),
		predicates: append([]predicate.Item{}, iq.predicates...),
		// clone intermediate query.
		gremlin: iq.gremlin.Clone(),
		path:    iq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Item.Query().
//		GroupBy(item.FieldText).
//		Aggregate(fluent.Count()).
//		Scan(ctx, &v)
func (iq *ItemQuery) GroupBy(field string, fields ...string) *ItemGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ItemGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = item.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Item.Query().
//		Select(item.FieldText).
//		Scan(ctx, &v)
func (iq *ItemQuery) Select(fields ...string) *ItemSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &ItemSelect{ItemQuery: iq}
	sbuild.label = item.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ItemSelect configured with the given aggregations.
func (iq *ItemQuery) Aggregate(fns ...AggregateFunc) *ItemSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *ItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("fluent: uninitialized interceptor (forgotten import fluent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.gremlin = prev
	}
	return nil
}

func (iq *ItemQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Item, error) {
	res := &gremlin.Response{}
	traversal := iq.gremlinQuery(ctx)
	if len(iq.ctx.Fields) > 0 {
		fields := make([]any, len(iq.ctx.Fields))
		for i, f := range iq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := iq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var is Items
	if err := is.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range is {
		is[i].config = iq.config
	}
	return is, nil
}

func (iq *ItemQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := iq.gremlinQuery(ctx).Count().Query()
	if err := iq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (iq *ItemQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(item.Label)
	if iq.gremlin != nil {
		v = iq.gremlin.Clone()
	}
	for _, p := range iq.predicates {
		p(v)
	}
	if len(iq.order) > 0 {
		v.Order()
		for _, p := range iq.order {
			p(v)
		}
	}
	switch limit, offset := iq.ctx.Limit, iq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := iq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// ItemGroupBy is the group-by builder for Item entities.
type ItemGroupBy struct {
	selector
	build *ItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *ItemGroupBy) Aggregate(fns ...AggregateFunc) *ItemGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *ItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, fluent.OpQueryGroupBy)
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemQuery, *ItemGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *ItemGroupBy) gremlinScan(ctx context.Context, root *ItemQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range igb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *igb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*igb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := igb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*igb.flds)+len(igb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// ItemSelect is the builder for selecting fields of Item entities.
type ItemSelect struct {
	*ItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *ItemSelect) Aggregate(fns ...AggregateFunc) *ItemSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *ItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, fluent.OpQuerySelect)
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemQuery, *ItemSelect](ctx, is.ItemQuery, is, is.inters, v)
}

func (is *ItemSelect) gremlinScan(ctx context.Context, root *ItemQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := is.ctx.Fields; len(fields) == 1 {
		if fields[0] != item.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(is.ctx.Fields))
		for i, f := range is.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := is.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.ctx.Fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
