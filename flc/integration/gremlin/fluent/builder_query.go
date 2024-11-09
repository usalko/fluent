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
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/builder"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
)

// BuilderQuery is the builder for querying Builder entities.
type BuilderQuery struct {
	config
	ctx        *QueryContext
	order      []builder.OrderOption
	inters     []Interceptor
	predicates []predicate.Builder
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the BuilderQuery builder.
func (bq *BuilderQuery) Where(ps ...predicate.Builder) *BuilderQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BuilderQuery) Limit(limit int) *BuilderQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BuilderQuery) Offset(offset int) *BuilderQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BuilderQuery) Unique(unique bool) *BuilderQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BuilderQuery) Order(o ...builder.OrderOption) *BuilderQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// First returns the first Builder entity from the query.
// Returns a *NotFoundError when no Builder was found.
func (bq *BuilderQuery) First(ctx context.Context) (*Builder, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{builder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BuilderQuery) FirstX(ctx context.Context) *Builder {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Builder ID from the query.
// Returns a *NotFoundError when no Builder ID was found.
func (bq *BuilderQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, fluent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{builder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BuilderQuery) FirstIDX(ctx context.Context) string {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Builder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Builder entity is found.
// Returns a *NotFoundError when no Builder entities are found.
func (bq *BuilderQuery) Only(ctx context.Context) (*Builder, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{builder.Label}
	default:
		return nil, &NotSingularError{builder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BuilderQuery) OnlyX(ctx context.Context) *Builder {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Builder ID in the query.
// Returns a *NotSingularError when more than one Builder ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BuilderQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, fluent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{builder.Label}
	default:
		err = &NotSingularError{builder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BuilderQuery) OnlyIDX(ctx context.Context) string {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Builders.
func (bq *BuilderQuery) All(ctx context.Context) ([]*Builder, error) {
	ctx = setContextOp(ctx, bq.ctx, fluent.OpQueryAll)
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Builder, *BuilderQuery]()
	return withInterceptors[[]*Builder](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BuilderQuery) AllX(ctx context.Context) []*Builder {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Builder IDs.
func (bq *BuilderQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, fluent.OpQueryIDs)
	if err = bq.Select(builder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BuilderQuery) IDsX(ctx context.Context) []string {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BuilderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, fluent.OpQueryCount)
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BuilderQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BuilderQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BuilderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, fluent.OpQueryExist)
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("fluent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BuilderQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BuilderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BuilderQuery) Clone() *BuilderQuery {
	if bq == nil {
		return nil
	}
	return &BuilderQuery{
		config:     bq.config,
		ctx:        bq.ctx.Clone(),
		order:      append([]builder.OrderOption{}, bq.order...),
		inters:     append([]Interceptor{}, bq.inters...),
		predicates: append([]predicate.Builder{}, bq.predicates...),
		// clone intermediate query.
		gremlin: bq.gremlin.Clone(),
		path:    bq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (bq *BuilderQuery) GroupBy(field string, fields ...string) *BuilderGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BuilderGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = builder.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (bq *BuilderQuery) Select(fields ...string) *BuilderSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BuilderSelect{BuilderQuery: bq}
	sbuild.label = builder.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BuilderSelect configured with the given aggregations.
func (bq *BuilderQuery) Aggregate(fns ...AggregateFunc) *BuilderSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BuilderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("fluent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.gremlin = prev
	}
	return nil
}

func (bq *BuilderQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Builder, error) {
	res := &gremlin.Response{}
	traversal := bq.gremlinQuery(ctx)
	if len(bq.ctx.Fields) > 0 {
		fields := make([]any, len(bq.ctx.Fields))
		for i, f := range bq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := bq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var bs Builders
	if err := bs.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range bs {
		bs[i].config = bq.config
	}
	return bs, nil
}

func (bq *BuilderQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := bq.gremlinQuery(ctx).Count().Query()
	if err := bq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (bq *BuilderQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(builder.Label)
	if bq.gremlin != nil {
		v = bq.gremlin.Clone()
	}
	for _, p := range bq.predicates {
		p(v)
	}
	if len(bq.order) > 0 {
		v.Order()
		for _, p := range bq.order {
			p(v)
		}
	}
	switch limit, offset := bq.ctx.Limit, bq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := bq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// BuilderGroupBy is the group-by builder for Builder entities.
type BuilderGroupBy struct {
	selector
	build *BuilderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BuilderGroupBy) Aggregate(fns ...AggregateFunc) *BuilderGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BuilderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, fluent.OpQueryGroupBy)
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BuilderQuery, *BuilderGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BuilderGroupBy) gremlinScan(ctx context.Context, root *BuilderQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range bgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *bgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*bgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := bgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*bgb.flds)+len(bgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// BuilderSelect is the builder for selecting fields of Builder entities.
type BuilderSelect struct {
	*BuilderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BuilderSelect) Aggregate(fns ...AggregateFunc) *BuilderSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BuilderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, fluent.OpQuerySelect)
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BuilderQuery, *BuilderSelect](ctx, bs.BuilderQuery, bs, bs.inters, v)
}

func (bs *BuilderSelect) gremlinScan(ctx context.Context, root *BuilderQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := bs.ctx.Fields; len(fields) == 1 {
		if fields[0] != builder.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(bs.ctx.Fields))
		for i, f := range bs.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := bs.driver.Exec(ctx, query, bindings, res); err != nil {
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
