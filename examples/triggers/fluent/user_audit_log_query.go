// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"fmt"
	"math"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/examples/triggers/fluent/predicate"
	"github.com/usalko/fluent/examples/triggers/fluent/user_audit_log"
	"github.com/usalko/fluent/schema/field"
)

// UserAuditLogQuery is the builder for querying UserAuditLog entities.
type UserAuditLogQuery struct {
	config
	ctx        *QueryContext
	order      []user_audit_log.OrderOption
	inters     []Interceptor
	predicates []predicate.UserAuditLog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserAuditLogQuery builder.
func (ualq *UserAuditLogQuery) Where(ps ...predicate.UserAuditLog) *UserAuditLogQuery {
	ualq.predicates = append(ualq.predicates, ps...)
	return ualq
}

// Limit the number of records to be returned by this query.
func (ualq *UserAuditLogQuery) Limit(limit int) *UserAuditLogQuery {
	ualq.ctx.Limit = &limit
	return ualq
}

// Offset to start from.
func (ualq *UserAuditLogQuery) Offset(offset int) *UserAuditLogQuery {
	ualq.ctx.Offset = &offset
	return ualq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ualq *UserAuditLogQuery) Unique(unique bool) *UserAuditLogQuery {
	ualq.ctx.Unique = &unique
	return ualq
}

// Order specifies how the records should be ordered.
func (ualq *UserAuditLogQuery) Order(o ...user_audit_log.OrderOption) *UserAuditLogQuery {
	ualq.order = append(ualq.order, o...)
	return ualq
}

// First returns the first UserAuditLog entity from the query.
// Returns a *NotFoundError when no UserAuditLog was found.
func (ualq *UserAuditLogQuery) First(ctx context.Context) (*UserAuditLog, error) {
	nodes, err := ualq.Limit(1).All(setContextOp(ctx, ualq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user_audit_log.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ualq *UserAuditLogQuery) FirstX(ctx context.Context) *UserAuditLog {
	node, err := ualq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserAuditLog ID from the query.
// Returns a *NotFoundError when no UserAuditLog ID was found.
func (ualq *UserAuditLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ualq.Limit(1).IDs(setContextOp(ctx, ualq.ctx, fluent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user_audit_log.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ualq *UserAuditLogQuery) FirstIDX(ctx context.Context) int {
	id, err := ualq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserAuditLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserAuditLog entity is found.
// Returns a *NotFoundError when no UserAuditLog entities are found.
func (ualq *UserAuditLogQuery) Only(ctx context.Context) (*UserAuditLog, error) {
	nodes, err := ualq.Limit(2).All(setContextOp(ctx, ualq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user_audit_log.Label}
	default:
		return nil, &NotSingularError{user_audit_log.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ualq *UserAuditLogQuery) OnlyX(ctx context.Context) *UserAuditLog {
	node, err := ualq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserAuditLog ID in the query.
// Returns a *NotSingularError when more than one UserAuditLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (ualq *UserAuditLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ualq.Limit(2).IDs(setContextOp(ctx, ualq.ctx, fluent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user_audit_log.Label}
	default:
		err = &NotSingularError{user_audit_log.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ualq *UserAuditLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := ualq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserAuditLogs.
func (ualq *UserAuditLogQuery) All(ctx context.Context) ([]*UserAuditLog, error) {
	ctx = setContextOp(ctx, ualq.ctx, fluent.OpQueryAll)
	if err := ualq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserAuditLog, *UserAuditLogQuery]()
	return withInterceptors[[]*UserAuditLog](ctx, ualq, qr, ualq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ualq *UserAuditLogQuery) AllX(ctx context.Context) []*UserAuditLog {
	nodes, err := ualq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserAuditLog IDs.
func (ualq *UserAuditLogQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ualq.ctx.Unique == nil && ualq.path != nil {
		ualq.Unique(true)
	}
	ctx = setContextOp(ctx, ualq.ctx, fluent.OpQueryIDs)
	if err = ualq.Select(user_audit_log.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ualq *UserAuditLogQuery) IDsX(ctx context.Context) []int {
	ids, err := ualq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ualq *UserAuditLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ualq.ctx, fluent.OpQueryCount)
	if err := ualq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ualq, querierCount[*UserAuditLogQuery](), ualq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ualq *UserAuditLogQuery) CountX(ctx context.Context) int {
	count, err := ualq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ualq *UserAuditLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ualq.ctx, fluent.OpQueryExist)
	switch _, err := ualq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("fluent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ualq *UserAuditLogQuery) ExistX(ctx context.Context) bool {
	exist, err := ualq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserAuditLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ualq *UserAuditLogQuery) Clone() *UserAuditLogQuery {
	if ualq == nil {
		return nil
	}
	return &UserAuditLogQuery{
		config:     ualq.config,
		ctx:        ualq.ctx.Clone(),
		order:      append([]user_audit_log.OrderOption{}, ualq.order...),
		inters:     append([]Interceptor{}, ualq.inters...),
		predicates: append([]predicate.UserAuditLog{}, ualq.predicates...),
		// clone intermediate query.
		sql:  ualq.sql.Clone(),
		path: ualq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OperationType string `json:"operation_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserAuditLog.Query().
//		GroupBy(user_audit_log.FieldOperationType).
//		Aggregate(fluent.Count()).
//		Scan(ctx, &v)
func (ualq *UserAuditLogQuery) GroupBy(field string, fields ...string) *UserAuditLogGroupBy {
	ualq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserAuditLogGroupBy{build: ualq}
	grbuild.flds = &ualq.ctx.Fields
	grbuild.label = user_audit_log.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OperationType string `json:"operation_type,omitempty"`
//	}
//
//	client.UserAuditLog.Query().
//		Select(user_audit_log.FieldOperationType).
//		Scan(ctx, &v)
func (ualq *UserAuditLogQuery) Select(fields ...string) *UserAuditLogSelect {
	ualq.ctx.Fields = append(ualq.ctx.Fields, fields...)
	sbuild := &UserAuditLogSelect{UserAuditLogQuery: ualq}
	sbuild.label = user_audit_log.Label
	sbuild.flds, sbuild.scan = &ualq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserAuditLogSelect configured with the given aggregations.
func (ualq *UserAuditLogQuery) Aggregate(fns ...AggregateFunc) *UserAuditLogSelect {
	return ualq.Select().Aggregate(fns...)
}

func (ualq *UserAuditLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ualq.inters {
		if inter == nil {
			return fmt.Errorf("fluent: uninitialized interceptor (forgotten import fluent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ualq); err != nil {
				return err
			}
		}
	}
	for _, f := range ualq.ctx.Fields {
		if !user_audit_log.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
		}
	}
	if ualq.path != nil {
		prev, err := ualq.path(ctx)
		if err != nil {
			return err
		}
		ualq.sql = prev
	}
	return nil
}

func (ualq *UserAuditLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserAuditLog, error) {
	var (
		nodes = []*UserAuditLog{}
		_spec = ualq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserAuditLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserAuditLog{config: ualq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ualq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ualq *UserAuditLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ualq.querySpec()
	_spec.Node.Columns = ualq.ctx.Fields
	if len(ualq.ctx.Fields) > 0 {
		_spec.Unique = ualq.ctx.Unique != nil && *ualq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ualq.driver, _spec)
}

func (ualq *UserAuditLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user_audit_log.Table, user_audit_log.Columns, sqlgraph.NewFieldSpec(user_audit_log.FieldID, field.TypeInt))
	_spec.From = ualq.sql
	if unique := ualq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ualq.path != nil {
		_spec.Unique = true
	}
	if fields := ualq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user_audit_log.FieldID)
		for i := range fields {
			if fields[i] != user_audit_log.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ualq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ualq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ualq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ualq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ualq *UserAuditLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ualq.driver.Dialect())
	t1 := builder.Table(user_audit_log.Table)
	columns := ualq.ctx.Fields
	if len(columns) == 0 {
		columns = user_audit_log.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ualq.sql != nil {
		selector = ualq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ualq.ctx.Unique != nil && *ualq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ualq.predicates {
		p(selector)
	}
	for _, p := range ualq.order {
		p(selector)
	}
	if offset := ualq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ualq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserAuditLogGroupBy is the group-by builder for UserAuditLog entities.
type UserAuditLogGroupBy struct {
	selector
	build *UserAuditLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ualgb *UserAuditLogGroupBy) Aggregate(fns ...AggregateFunc) *UserAuditLogGroupBy {
	ualgb.fns = append(ualgb.fns, fns...)
	return ualgb
}

// Scan applies the selector query and scans the result into the given value.
func (ualgb *UserAuditLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ualgb.build.ctx, fluent.OpQueryGroupBy)
	if err := ualgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserAuditLogQuery, *UserAuditLogGroupBy](ctx, ualgb.build, ualgb, ualgb.build.inters, v)
}

func (ualgb *UserAuditLogGroupBy) sqlScan(ctx context.Context, root *UserAuditLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ualgb.fns))
	for _, fn := range ualgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ualgb.flds)+len(ualgb.fns))
		for _, f := range *ualgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ualgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ualgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserAuditLogSelect is the builder for selecting fields of UserAuditLog entities.
type UserAuditLogSelect struct {
	*UserAuditLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uals *UserAuditLogSelect) Aggregate(fns ...AggregateFunc) *UserAuditLogSelect {
	uals.fns = append(uals.fns, fns...)
	return uals
}

// Scan applies the selector query and scans the result into the given value.
func (uals *UserAuditLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uals.ctx, fluent.OpQuerySelect)
	if err := uals.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserAuditLogQuery, *UserAuditLogSelect](ctx, uals.UserAuditLogQuery, uals, uals.inters, v)
}

func (uals *UserAuditLogSelect) sqlScan(ctx context.Context, root *UserAuditLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uals.fns))
	for _, fn := range uals.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uals.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uals.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
