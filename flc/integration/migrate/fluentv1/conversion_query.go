// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluentv1

import (
	"context"
	"fmt"
	"math"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/migrate/fluentv1/conversion"
	"github.com/usalko/fluent/flc/integration/migrate/fluentv1/predicate"
	"github.com/usalko/fluent/schema/field"
)

// ConversionQuery is the builder for querying Conversion entities.
type ConversionQuery struct {
	config
	ctx        *QueryContext
	order      []conversion.OrderOption
	inters     []Interceptor
	predicates []predicate.Conversion
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConversionQuery builder.
func (cq *ConversionQuery) Where(ps ...predicate.Conversion) *ConversionQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ConversionQuery) Limit(limit int) *ConversionQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ConversionQuery) Offset(offset int) *ConversionQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ConversionQuery) Unique(unique bool) *ConversionQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ConversionQuery) Order(o ...conversion.OrderOption) *ConversionQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// First returns the first Conversion entity from the query.
// Returns a *NotFoundError when no Conversion was found.
func (cq *ConversionQuery) First(ctx context.Context) (*Conversion, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{conversion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ConversionQuery) FirstX(ctx context.Context) *Conversion {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Conversion ID from the query.
// Returns a *NotFoundError when no Conversion ID was found.
func (cq *ConversionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, fluent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{conversion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ConversionQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Conversion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Conversion entity is found.
// Returns a *NotFoundError when no Conversion entities are found.
func (cq *ConversionQuery) Only(ctx context.Context) (*Conversion, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{conversion.Label}
	default:
		return nil, &NotSingularError{conversion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ConversionQuery) OnlyX(ctx context.Context) *Conversion {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Conversion ID in the query.
// Returns a *NotSingularError when more than one Conversion ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ConversionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, fluent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{conversion.Label}
	default:
		err = &NotSingularError{conversion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ConversionQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Conversions.
func (cq *ConversionQuery) All(ctx context.Context) ([]*Conversion, error) {
	ctx = setContextOp(ctx, cq.ctx, fluent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Conversion, *ConversionQuery]()
	return withInterceptors[[]*Conversion](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ConversionQuery) AllX(ctx context.Context) []*Conversion {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Conversion IDs.
func (cq *ConversionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, fluent.OpQueryIDs)
	if err = cq.Select(conversion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ConversionQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ConversionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, fluent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ConversionQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ConversionQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ConversionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, fluent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("fluentv1: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ConversionQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConversionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ConversionQuery) Clone() *ConversionQuery {
	if cq == nil {
		return nil
	}
	return &ConversionQuery{
		config:     cq.config,
		ctx:        cq.ctx.Clone(),
		order:      append([]conversion.OrderOption{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Conversion{}, cq.predicates...),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
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
//	client.Conversion.Query().
//		GroupBy(conversion.FieldName).
//		Aggregate(fluentv1.Count()).
//		Scan(ctx, &v)
func (cq *ConversionQuery) GroupBy(field string, fields ...string) *ConversionGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ConversionGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = conversion.Label
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
//	client.Conversion.Query().
//		Select(conversion.FieldName).
//		Scan(ctx, &v)
func (cq *ConversionQuery) Select(fields ...string) *ConversionSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ConversionSelect{ConversionQuery: cq}
	sbuild.label = conversion.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ConversionSelect configured with the given aggregations.
func (cq *ConversionQuery) Aggregate(fns ...AggregateFunc) *ConversionSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ConversionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("fluentv1: uninitialized interceptor (forgotten import fluentv1/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !conversion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("fluentv1: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ConversionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Conversion, error) {
	var (
		nodes = []*Conversion{}
		_spec = cq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Conversion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Conversion{config: cq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cq *ConversionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ConversionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(conversion.Table, conversion.Columns, sqlgraph.NewFieldSpec(conversion.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, conversion.FieldID)
		for i := range fields {
			if fields[i] != conversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ConversionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(conversion.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = conversion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ConversionGroupBy is the group-by builder for Conversion entities.
type ConversionGroupBy struct {
	selector
	build *ConversionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ConversionGroupBy) Aggregate(fns ...AggregateFunc) *ConversionGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ConversionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, fluent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConversionQuery, *ConversionGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ConversionGroupBy) sqlScan(ctx context.Context, root *ConversionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ConversionSelect is the builder for selecting fields of Conversion entities.
type ConversionSelect struct {
	*ConversionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ConversionSelect) Aggregate(fns ...AggregateFunc) *ConversionSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ConversionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, fluent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConversionQuery, *ConversionSelect](ctx, cs.ConversionQuery, cs, cs.inters, v)
}

func (cs *ConversionSelect) sqlScan(ctx context.Context, root *ConversionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
