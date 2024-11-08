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
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/filetype"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
)

// FileTypeQuery is the builder for querying FileType entities.
type FileTypeQuery struct {
	config
	ctx        *QueryContext
	order      []filetype.OrderOption
	inters     []Interceptor
	predicates []predicate.FileType
	withFiles  *FileQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the FileTypeQuery builder.
func (ftq *FileTypeQuery) Where(ps ...predicate.FileType) *FileTypeQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit the number of records to be returned by this query.
func (ftq *FileTypeQuery) Limit(limit int) *FileTypeQuery {
	ftq.ctx.Limit = &limit
	return ftq
}

// Offset to start from.
func (ftq *FileTypeQuery) Offset(offset int) *FileTypeQuery {
	ftq.ctx.Offset = &offset
	return ftq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ftq *FileTypeQuery) Unique(unique bool) *FileTypeQuery {
	ftq.ctx.Unique = &unique
	return ftq
}

// Order specifies how the records should be ordered.
func (ftq *FileTypeQuery) Order(o ...filetype.OrderOption) *FileTypeQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// QueryFiles chains the current query on the "files" edge.
func (ftq *FileTypeQuery) QueryFiles() *FileQuery {
	query := (&FileClient{config: ftq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := ftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := ftq.gremlinQuery(ctx)
		fromU = gremlin.OutE(filetype.FilesLabel).InV()
		return fromU, nil
	}
	return query
}

// First returns the first FileType entity from the query.
// Returns a *NotFoundError when no FileType was found.
func (ftq *FileTypeQuery) First(ctx context.Context) (*FileType, error) {
	nodes, err := ftq.Limit(1).All(setContextOp(ctx, ftq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{filetype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FileTypeQuery) FirstX(ctx context.Context) *FileType {
	node, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FileType ID from the query.
// Returns a *NotFoundError when no FileType ID was found.
func (ftq *FileTypeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ftq.Limit(1).IDs(setContextOp(ctx, ftq.ctx, fluent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{filetype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ftq *FileTypeQuery) FirstIDX(ctx context.Context) string {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FileType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FileType entity is found.
// Returns a *NotFoundError when no FileType entities are found.
func (ftq *FileTypeQuery) Only(ctx context.Context) (*FileType, error) {
	nodes, err := ftq.Limit(2).All(setContextOp(ctx, ftq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{filetype.Label}
	default:
		return nil, &NotSingularError{filetype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FileTypeQuery) OnlyX(ctx context.Context) *FileType {
	node, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FileType ID in the query.
// Returns a *NotSingularError when more than one FileType ID is found.
// Returns a *NotFoundError when no entities are found.
func (ftq *FileTypeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ftq.Limit(2).IDs(setContextOp(ctx, ftq.ctx, fluent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = &NotSingularError{filetype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ftq *FileTypeQuery) OnlyIDX(ctx context.Context) string {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FileTypes.
func (ftq *FileTypeQuery) All(ctx context.Context) ([]*FileType, error) {
	ctx = setContextOp(ctx, ftq.ctx, fluent.OpQueryAll)
	if err := ftq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FileType, *FileTypeQuery]()
	return withInterceptors[[]*FileType](ctx, ftq, qr, ftq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ftq *FileTypeQuery) AllX(ctx context.Context) []*FileType {
	nodes, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FileType IDs.
func (ftq *FileTypeQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ftq.ctx.Unique == nil && ftq.path != nil {
		ftq.Unique(true)
	}
	ctx = setContextOp(ctx, ftq.ctx, fluent.OpQueryIDs)
	if err = ftq.Select(filetype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FileTypeQuery) IDsX(ctx context.Context) []string {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FileTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ftq.ctx, fluent.OpQueryCount)
	if err := ftq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ftq, querierCount[*FileTypeQuery](), ftq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FileTypeQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FileTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ftq.ctx, fluent.OpQueryExist)
	switch _, err := ftq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FileTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftq *FileTypeQuery) Clone() *FileTypeQuery {
	if ftq == nil {
		return nil
	}
	return &FileTypeQuery{
		config:     ftq.config,
		ctx:        ftq.ctx.Clone(),
		order:      append([]filetype.OrderOption{}, ftq.order...),
		inters:     append([]Interceptor{}, ftq.inters...),
		predicates: append([]predicate.FileType{}, ftq.predicates...),
		withFiles:  ftq.withFiles.Clone(),
		// clone intermediate query.
		gremlin: ftq.gremlin.Clone(),
		path:    ftq.path,
	}
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (ftq *FileTypeQuery) WithFiles(opts ...func(*FileQuery)) *FileTypeQuery {
	query := (&FileClient{config: ftq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ftq.withFiles = query
	return ftq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FileType.Query().
//		GroupBy(filetype.FieldName).
//		Aggregate(fluent.Count()).
//		Scan(ctx, &v)
func (ftq *FileTypeQuery) GroupBy(field string, fields ...string) *FileTypeGroupBy {
	ftq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FileTypeGroupBy{build: ftq}
	grbuild.flds = &ftq.ctx.Fields
	grbuild.label = filetype.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.FileType.Query().
//		Select(filetype.FieldName).
//		Scan(ctx, &v)
func (ftq *FileTypeQuery) Select(fields ...string) *FileTypeSelect {
	ftq.ctx.Fields = append(ftq.ctx.Fields, fields...)
	sbuild := &FileTypeSelect{FileTypeQuery: ftq}
	sbuild.label = filetype.Label
	sbuild.flds, sbuild.scan = &ftq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FileTypeSelect configured with the given aggregations.
func (ftq *FileTypeQuery) Aggregate(fns ...AggregateFunc) *FileTypeSelect {
	return ftq.Select().Aggregate(fns...)
}

func (ftq *FileTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ftq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ftq); err != nil {
				return err
			}
		}
	}
	if ftq.path != nil {
		prev, err := ftq.path(ctx)
		if err != nil {
			return err
		}
		ftq.gremlin = prev
	}
	return nil
}

func (ftq *FileTypeQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*FileType, error) {
	res := &gremlin.Response{}
	traversal := ftq.gremlinQuery(ctx)
	if len(ftq.ctx.Fields) > 0 {
		fields := make([]any, len(ftq.ctx.Fields))
		for i, f := range ftq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var fts FileTypes
	if err := fts.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range fts {
		fts[i].config = ftq.config
	}
	return fts, nil
}

func (ftq *FileTypeQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery(ctx).Count().Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (ftq *FileTypeQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(filetype.Label)
	if ftq.gremlin != nil {
		v = ftq.gremlin.Clone()
	}
	for _, p := range ftq.predicates {
		p(v)
	}
	if len(ftq.order) > 0 {
		v.Order()
		for _, p := range ftq.order {
			p(v)
		}
	}
	switch limit, offset := ftq.ctx.Limit, ftq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := ftq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// FileTypeGroupBy is the group-by builder for FileType entities.
type FileTypeGroupBy struct {
	selector
	build *FileTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FileTypeGroupBy) Aggregate(fns ...AggregateFunc) *FileTypeGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the selector query and scans the result into the given value.
func (ftgb *FileTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ftgb.build.ctx, fluent.OpQueryGroupBy)
	if err := ftgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileTypeQuery, *FileTypeGroupBy](ctx, ftgb.build, ftgb, ftgb.build.inters, v)
}

func (ftgb *FileTypeGroupBy) gremlinScan(ctx context.Context, root *FileTypeQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range ftgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *ftgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*ftgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := ftgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*ftgb.flds)+len(ftgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// FileTypeSelect is the builder for selecting fields of FileType entities.
type FileTypeSelect struct {
	*FileTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fts *FileTypeSelect) Aggregate(fns ...AggregateFunc) *FileTypeSelect {
	fts.fns = append(fts.fns, fns...)
	return fts
}

// Scan applies the selector query and scans the result into the given value.
func (fts *FileTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fts.ctx, fluent.OpQuerySelect)
	if err := fts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileTypeQuery, *FileTypeSelect](ctx, fts.FileTypeQuery, fts, fts.inters, v)
}

func (fts *FileTypeSelect) gremlinScan(ctx context.Context, root *FileTypeQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := fts.ctx.Fields; len(fields) == 1 {
		if fields[0] != filetype.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(fts.ctx.Fields))
		for i, f := range fts.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := fts.driver.Exec(ctx, query, bindings, res); err != nil {
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
