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
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/pet"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/user"
)

// PetQuery is the builder for querying Pet entities.
type PetQuery struct {
	config
	ctx        *QueryContext
	order      []pet.OrderOption
	inters     []Interceptor
	predicates []predicate.Pet
	withTeam   *UserQuery
	withOwner  *UserQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the PetQuery builder.
func (pq *PetQuery) Where(ps ...predicate.Pet) *PetQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PetQuery) Limit(limit int) *PetQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PetQuery) Offset(offset int) *PetQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PetQuery) Unique(unique bool) *PetQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PetQuery) Order(o ...pet.OrderOption) *PetQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryTeam chains the current query on the "team" edge.
func (pq *PetQuery) QueryTeam() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := pq.gremlinQuery(ctx)
		fromU = gremlin.InE(user.TeamLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (pq *PetQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := pq.gremlinQuery(ctx)
		fromU = gremlin.InE(user.PetsLabel).OutV()
		return fromU, nil
	}
	return query
}

// First returns the first Pet entity from the query.
// Returns a *NotFoundError when no Pet was found.
func (pq *PetQuery) First(ctx context.Context) (*Pet, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PetQuery) FirstX(ctx context.Context) *Pet {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pet ID from the query.
// Returns a *NotFoundError when no Pet ID was found.
func (pq *PetQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, fluent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PetQuery) FirstIDX(ctx context.Context) string {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pet entity is found.
// Returns a *NotFoundError when no Pet entities are found.
func (pq *PetQuery) Only(ctx context.Context) (*Pet, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pet.Label}
	default:
		return nil, &NotSingularError{pet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PetQuery) OnlyX(ctx context.Context) *Pet {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pet ID in the query.
// Returns a *NotSingularError when more than one Pet ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PetQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, fluent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = &NotSingularError{pet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PetQuery) OnlyIDX(ctx context.Context) string {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pets.
func (pq *PetQuery) All(ctx context.Context) ([]*Pet, error) {
	ctx = setContextOp(ctx, pq.ctx, fluent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Pet, *PetQuery]()
	return withInterceptors[[]*Pet](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PetQuery) AllX(ctx context.Context) []*Pet {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pet IDs.
func (pq *PetQuery) IDs(ctx context.Context) (ids []string, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, fluent.OpQueryIDs)
	if err = pq.Select(pet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PetQuery) IDsX(ctx context.Context) []string {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, fluent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PetQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PetQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, fluent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("fluent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PetQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PetQuery) Clone() *PetQuery {
	if pq == nil {
		return nil
	}
	return &PetQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]pet.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Pet{}, pq.predicates...),
		withTeam:   pq.withTeam.Clone(),
		withOwner:  pq.withOwner.Clone(),
		// clone intermediate query.
		gremlin: pq.gremlin.Clone(),
		path:    pq.path,
	}
}

// WithTeam tells the query-builder to eager-load the nodes that are connected to
// the "team" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PetQuery) WithTeam(opts ...func(*UserQuery)) *PetQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withTeam = query
	return pq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PetQuery) WithOwner(opts ...func(*UserQuery)) *PetQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withOwner = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Age float64 `json:"age,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Pet.Query().
//		GroupBy(pet.FieldAge).
//		Aggregate(fluent.Count()).
//		Scan(ctx, &v)
func (pq *PetQuery) GroupBy(field string, fields ...string) *PetGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PetGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = pet.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Age float64 `json:"age,omitempty"`
//	}
//
//	client.Pet.Query().
//		Select(pet.FieldAge).
//		Scan(ctx, &v)
func (pq *PetQuery) Select(fields ...string) *PetSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PetSelect{PetQuery: pq}
	sbuild.label = pet.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PetSelect configured with the given aggregations.
func (pq *PetQuery) Aggregate(fns ...AggregateFunc) *PetSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("fluent: uninitialized interceptor (forgotten import fluent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.gremlin = prev
	}
	return nil
}

func (pq *PetQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Pet, error) {
	res := &gremlin.Response{}
	traversal := pq.gremlinQuery(ctx)
	if len(pq.ctx.Fields) > 0 {
		fields := make([]any, len(pq.ctx.Fields))
		for i, f := range pq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var pes Pets
	if err := pes.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range pes {
		pes[i].config = pq.config
	}
	return pes, nil
}

func (pq *PetQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := pq.gremlinQuery(ctx).Count().Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (pq *PetQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(pet.Label)
	if pq.gremlin != nil {
		v = pq.gremlin.Clone()
	}
	for _, p := range pq.predicates {
		p(v)
	}
	if len(pq.order) > 0 {
		v.Order()
		for _, p := range pq.order {
			p(v)
		}
	}
	switch limit, offset := pq.ctx.Limit, pq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := pq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// PetGroupBy is the group-by builder for Pet entities.
type PetGroupBy struct {
	selector
	build *PetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PetGroupBy) Aggregate(fns ...AggregateFunc) *PetGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, fluent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PetQuery, *PetGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PetGroupBy) gremlinScan(ctx context.Context, root *PetQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range pgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *pgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*pgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := pgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*pgb.flds)+len(pgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// PetSelect is the builder for selecting fields of Pet entities.
type PetSelect struct {
	*PetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PetSelect) Aggregate(fns ...AggregateFunc) *PetSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, fluent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PetQuery, *PetSelect](ctx, ps.PetQuery, ps, ps.inters, v)
}

func (ps *PetSelect) gremlinScan(ctx context.Context, root *PetQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := ps.ctx.Fields; len(fields) == 1 {
		if fields[0] != pet.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(ps.ctx.Fields))
		for i, f := range ps.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ps.driver.Exec(ctx, query, bindings, res); err != nil {
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
