// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"fmt"
	"math"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/examples/view_schema/fluent/petusername"
	"github.com/usalko/fluent/examples/view_schema/fluent/predicate"
)

// PetUserNameQuery is the builder for querying PetUserName entities.
type PetUserNameQuery struct {
	config
	ctx        *QueryContext
	order      []petusername.OrderOption
	inters     []Interceptor
	predicates []predicate.PetUserName
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PetUserNameQuery builder.
func (punq *PetUserNameQuery) Where(ps ...predicate.PetUserName) *PetUserNameQuery {
	punq.predicates = append(punq.predicates, ps...)
	return punq
}

// Limit the number of records to be returned by this query.
func (punq *PetUserNameQuery) Limit(limit int) *PetUserNameQuery {
	punq.ctx.Limit = &limit
	return punq
}

// Offset to start from.
func (punq *PetUserNameQuery) Offset(offset int) *PetUserNameQuery {
	punq.ctx.Offset = &offset
	return punq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (punq *PetUserNameQuery) Unique(unique bool) *PetUserNameQuery {
	punq.ctx.Unique = &unique
	return punq
}

// Order specifies how the records should be ordered.
func (punq *PetUserNameQuery) Order(o ...petusername.OrderOption) *PetUserNameQuery {
	punq.order = append(punq.order, o...)
	return punq
}

// First returns the first PetUserName entity from the query.
// Returns a *NotFoundError when no PetUserName was found.
func (punq *PetUserNameQuery) First(ctx context.Context) (*PetUserName, error) {
	nodes, err := punq.Limit(1).All(setContextOp(ctx, punq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{petusername.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (punq *PetUserNameQuery) FirstX(ctx context.Context) *PetUserName {
	node, err := punq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single PetUserName entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PetUserName entity is found.
// Returns a *NotFoundError when no PetUserName entities are found.
func (punq *PetUserNameQuery) Only(ctx context.Context) (*PetUserName, error) {
	nodes, err := punq.Limit(2).All(setContextOp(ctx, punq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{petusername.Label}
	default:
		return nil, &NotSingularError{petusername.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (punq *PetUserNameQuery) OnlyX(ctx context.Context) *PetUserName {
	node, err := punq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of PetUserNames.
func (punq *PetUserNameQuery) All(ctx context.Context) ([]*PetUserName, error) {
	ctx = setContextOp(ctx, punq.ctx, fluent.OpQueryAll)
	if err := punq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PetUserName, *PetUserNameQuery]()
	return withInterceptors[[]*PetUserName](ctx, punq, qr, punq.inters)
}

// AllX is like All, but panics if an error occurs.
func (punq *PetUserNameQuery) AllX(ctx context.Context) []*PetUserName {
	nodes, err := punq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (punq *PetUserNameQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, punq.ctx, fluent.OpQueryCount)
	if err := punq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, punq, querierCount[*PetUserNameQuery](), punq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (punq *PetUserNameQuery) CountX(ctx context.Context) int {
	count, err := punq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (punq *PetUserNameQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, punq.ctx, fluent.OpQueryExist)
	switch _, err := punq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("fluent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (punq *PetUserNameQuery) ExistX(ctx context.Context) bool {
	exist, err := punq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PetUserNameQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (punq *PetUserNameQuery) Clone() *PetUserNameQuery {
	if punq == nil {
		return nil
	}
	return &PetUserNameQuery{
		config:     punq.config,
		ctx:        punq.ctx.Clone(),
		order:      append([]petusername.OrderOption{}, punq.order...),
		inters:     append([]Interceptor{}, punq.inters...),
		predicates: append([]predicate.PetUserName{}, punq.predicates...),
		// clone intermediate query.
		sql:  punq.sql.Clone(),
		path: punq.path,
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
//	client.PetUserName.Query().
//		GroupBy(petusername.FieldName).
//		Aggregate(fluent.Count()).
//		Scan(ctx, &v)
func (punq *PetUserNameQuery) GroupBy(field string, fields ...string) *PetUserNameGroupBy {
	punq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PetUserNameGroupBy{build: punq}
	grbuild.flds = &punq.ctx.Fields
	grbuild.label = petusername.Label
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
//	client.PetUserName.Query().
//		Select(petusername.FieldName).
//		Scan(ctx, &v)
func (punq *PetUserNameQuery) Select(fields ...string) *PetUserNameSelect {
	punq.ctx.Fields = append(punq.ctx.Fields, fields...)
	sbuild := &PetUserNameSelect{PetUserNameQuery: punq}
	sbuild.label = petusername.Label
	sbuild.flds, sbuild.scan = &punq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PetUserNameSelect configured with the given aggregations.
func (punq *PetUserNameQuery) Aggregate(fns ...AggregateFunc) *PetUserNameSelect {
	return punq.Select().Aggregate(fns...)
}

func (punq *PetUserNameQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range punq.inters {
		if inter == nil {
			return fmt.Errorf("fluent: uninitialized interceptor (forgotten import fluent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, punq); err != nil {
				return err
			}
		}
	}
	for _, f := range punq.ctx.Fields {
		if !petusername.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
		}
	}
	if punq.path != nil {
		prev, err := punq.path(ctx)
		if err != nil {
			return err
		}
		punq.sql = prev
	}
	return nil
}

func (punq *PetUserNameQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PetUserName, error) {
	var (
		nodes = []*PetUserName{}
		_spec = punq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PetUserName).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PetUserName{config: punq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, punq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (punq *PetUserNameQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := punq.querySpec()
	_spec.Node.Columns = punq.ctx.Fields
	if len(punq.ctx.Fields) > 0 {
		_spec.Unique = punq.ctx.Unique != nil && *punq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, punq.driver, _spec)
}

func (punq *PetUserNameQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(petusername.Table, petusername.Columns, nil)
	_spec.From = punq.sql
	if unique := punq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if punq.path != nil {
		_spec.Unique = true
	}
	if fields := punq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
	}
	if ps := punq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := punq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := punq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := punq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (punq *PetUserNameQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(punq.driver.Dialect())
	t1 := builder.Table(petusername.Table)
	columns := punq.ctx.Fields
	if len(columns) == 0 {
		columns = petusername.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if punq.sql != nil {
		selector = punq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if punq.ctx.Unique != nil && *punq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range punq.predicates {
		p(selector)
	}
	for _, p := range punq.order {
		p(selector)
	}
	if offset := punq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := punq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PetUserNameGroupBy is the group-by builder for PetUserName entities.
type PetUserNameGroupBy struct {
	selector
	build *PetUserNameQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pungb *PetUserNameGroupBy) Aggregate(fns ...AggregateFunc) *PetUserNameGroupBy {
	pungb.fns = append(pungb.fns, fns...)
	return pungb
}

// Scan applies the selector query and scans the result into the given value.
func (pungb *PetUserNameGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pungb.build.ctx, fluent.OpQueryGroupBy)
	if err := pungb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PetUserNameQuery, *PetUserNameGroupBy](ctx, pungb.build, pungb, pungb.build.inters, v)
}

func (pungb *PetUserNameGroupBy) sqlScan(ctx context.Context, root *PetUserNameQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pungb.fns))
	for _, fn := range pungb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pungb.flds)+len(pungb.fns))
		for _, f := range *pungb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pungb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pungb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PetUserNameSelect is the builder for selecting fields of PetUserName entities.
type PetUserNameSelect struct {
	*PetUserNameQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (puns *PetUserNameSelect) Aggregate(fns ...AggregateFunc) *PetUserNameSelect {
	puns.fns = append(puns.fns, fns...)
	return puns
}

// Scan applies the selector query and scans the result into the given value.
func (puns *PetUserNameSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, puns.ctx, fluent.OpQuerySelect)
	if err := puns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PetUserNameQuery, *PetUserNameSelect](ctx, puns.PetUserNameQuery, puns, puns.inters, v)
}

func (puns *PetUserNameSelect) sqlScan(ctx context.Context, root *PetUserNameQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(puns.fns))
	for _, fn := range puns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*puns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := puns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
