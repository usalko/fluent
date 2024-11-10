// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"fmt"
	"math"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/tweet"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/user"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/usertweet"
	"github.com/usalko/fluent/schema/field"
)

// UserTweetQuery is the builder for querying UserTweet entities.
type UserTweetQuery struct {
	config
	ctx        *QueryContext
	order      []usertweet.OrderOption
	inters     []Interceptor
	predicates []predicate.UserTweet
	withUser   *UserQuery
	withTweet  *TweetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserTweetQuery builder.
func (utq *UserTweetQuery) Where(ps ...predicate.UserTweet) *UserTweetQuery {
	utq.predicates = append(utq.predicates, ps...)
	return utq
}

// Limit the number of records to be returned by this query.
func (utq *UserTweetQuery) Limit(limit int) *UserTweetQuery {
	utq.ctx.Limit = &limit
	return utq
}

// Offset to start from.
func (utq *UserTweetQuery) Offset(offset int) *UserTweetQuery {
	utq.ctx.Offset = &offset
	return utq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (utq *UserTweetQuery) Unique(unique bool) *UserTweetQuery {
	utq.ctx.Unique = &unique
	return utq
}

// Order specifies how the records should be ordered.
func (utq *UserTweetQuery) Order(o ...usertweet.OrderOption) *UserTweetQuery {
	utq.order = append(utq.order, o...)
	return utq
}

// QueryUser chains the current query on the "user" edge.
func (utq *UserTweetQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: utq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usertweet.Table, usertweet.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usertweet.UserTable, usertweet.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(utq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTweet chains the current query on the "tweet" edge.
func (utq *UserTweetQuery) QueryTweet() *TweetQuery {
	query := (&TweetClient{config: utq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usertweet.Table, usertweet.FieldID, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usertweet.TweetTable, usertweet.TweetColumn),
		)
		fromU = sqlgraph.SetNeighbors(utq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserTweet entity from the query.
// Returns a *NotFoundError when no UserTweet was found.
func (utq *UserTweetQuery) First(ctx context.Context) (*UserTweet, error) {
	nodes, err := utq.Limit(1).All(setContextOp(ctx, utq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usertweet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (utq *UserTweetQuery) FirstX(ctx context.Context) *UserTweet {
	node, err := utq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserTweet ID from the query.
// Returns a *NotFoundError when no UserTweet ID was found.
func (utq *UserTweetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = utq.Limit(1).IDs(setContextOp(ctx, utq.ctx, fluent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usertweet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (utq *UserTweetQuery) FirstIDX(ctx context.Context) int {
	id, err := utq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserTweet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserTweet entity is found.
// Returns a *NotFoundError when no UserTweet entities are found.
func (utq *UserTweetQuery) Only(ctx context.Context) (*UserTweet, error) {
	nodes, err := utq.Limit(2).All(setContextOp(ctx, utq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usertweet.Label}
	default:
		return nil, &NotSingularError{usertweet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (utq *UserTweetQuery) OnlyX(ctx context.Context) *UserTweet {
	node, err := utq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserTweet ID in the query.
// Returns a *NotSingularError when more than one UserTweet ID is found.
// Returns a *NotFoundError when no entities are found.
func (utq *UserTweetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = utq.Limit(2).IDs(setContextOp(ctx, utq.ctx, fluent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usertweet.Label}
	default:
		err = &NotSingularError{usertweet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (utq *UserTweetQuery) OnlyIDX(ctx context.Context) int {
	id, err := utq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserTweets.
func (utq *UserTweetQuery) All(ctx context.Context) ([]*UserTweet, error) {
	ctx = setContextOp(ctx, utq.ctx, fluent.OpQueryAll)
	if err := utq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserTweet, *UserTweetQuery]()
	return withInterceptors[[]*UserTweet](ctx, utq, qr, utq.inters)
}

// AllX is like All, but panics if an error occurs.
func (utq *UserTweetQuery) AllX(ctx context.Context) []*UserTweet {
	nodes, err := utq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserTweet IDs.
func (utq *UserTweetQuery) IDs(ctx context.Context) (ids []int, err error) {
	if utq.ctx.Unique == nil && utq.path != nil {
		utq.Unique(true)
	}
	ctx = setContextOp(ctx, utq.ctx, fluent.OpQueryIDs)
	if err = utq.Select(usertweet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (utq *UserTweetQuery) IDsX(ctx context.Context) []int {
	ids, err := utq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (utq *UserTweetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, utq.ctx, fluent.OpQueryCount)
	if err := utq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, utq, querierCount[*UserTweetQuery](), utq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (utq *UserTweetQuery) CountX(ctx context.Context) int {
	count, err := utq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (utq *UserTweetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, utq.ctx, fluent.OpQueryExist)
	switch _, err := utq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("fluent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (utq *UserTweetQuery) ExistX(ctx context.Context) bool {
	exist, err := utq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserTweetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (utq *UserTweetQuery) Clone() *UserTweetQuery {
	if utq == nil {
		return nil
	}
	return &UserTweetQuery{
		config:     utq.config,
		ctx:        utq.ctx.Clone(),
		order:      append([]usertweet.OrderOption{}, utq.order...),
		inters:     append([]Interceptor{}, utq.inters...),
		predicates: append([]predicate.UserTweet{}, utq.predicates...),
		withUser:   utq.withUser.Clone(),
		withTweet:  utq.withTweet.Clone(),
		// clone intermediate query.
		sql:  utq.sql.Clone(),
		path: utq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (utq *UserTweetQuery) WithUser(opts ...func(*UserQuery)) *UserTweetQuery {
	query := (&UserClient{config: utq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	utq.withUser = query
	return utq
}

// WithTweet tells the query-builder to eager-load the nodes that are connected to
// the "tweet" edge. The optional arguments are used to configure the query builder of the edge.
func (utq *UserTweetQuery) WithTweet(opts ...func(*TweetQuery)) *UserTweetQuery {
	query := (&TweetClient{config: utq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	utq.withTweet = query
	return utq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserTweet.Query().
//		GroupBy(usertweet.FieldCreatedAt).
//		Aggregate(fluent.Count()).
//		Scan(ctx, &v)
func (utq *UserTweetQuery) GroupBy(field string, fields ...string) *UserTweetGroupBy {
	utq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserTweetGroupBy{build: utq}
	grbuild.flds = &utq.ctx.Fields
	grbuild.label = usertweet.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.UserTweet.Query().
//		Select(usertweet.FieldCreatedAt).
//		Scan(ctx, &v)
func (utq *UserTweetQuery) Select(fields ...string) *UserTweetSelect {
	utq.ctx.Fields = append(utq.ctx.Fields, fields...)
	sbuild := &UserTweetSelect{UserTweetQuery: utq}
	sbuild.label = usertweet.Label
	sbuild.flds, sbuild.scan = &utq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserTweetSelect configured with the given aggregations.
func (utq *UserTweetQuery) Aggregate(fns ...AggregateFunc) *UserTweetSelect {
	return utq.Select().Aggregate(fns...)
}

func (utq *UserTweetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range utq.inters {
		if inter == nil {
			return fmt.Errorf("fluent: uninitialized interceptor (forgotten import fluent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, utq); err != nil {
				return err
			}
		}
	}
	for _, f := range utq.ctx.Fields {
		if !usertweet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
		}
	}
	if utq.path != nil {
		prev, err := utq.path(ctx)
		if err != nil {
			return err
		}
		utq.sql = prev
	}
	return nil
}

func (utq *UserTweetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserTweet, error) {
	var (
		nodes       = []*UserTweet{}
		_spec       = utq.querySpec()
		loadedTypes = [2]bool{
			utq.withUser != nil,
			utq.withTweet != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserTweet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserTweet{config: utq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, utq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := utq.withUser; query != nil {
		if err := utq.loadUser(ctx, query, nodes, nil,
			func(n *UserTweet, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := utq.withTweet; query != nil {
		if err := utq.loadTweet(ctx, query, nodes, nil,
			func(n *UserTweet, e *Tweet) { n.Edges.Tweet = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (utq *UserTweetQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserTweet, init func(*UserTweet), assign func(*UserTweet, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserTweet)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (utq *UserTweetQuery) loadTweet(ctx context.Context, query *TweetQuery, nodes []*UserTweet, init func(*UserTweet), assign func(*UserTweet, *Tweet)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserTweet)
	for i := range nodes {
		fk := nodes[i].TweetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tweet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tweet_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (utq *UserTweetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := utq.querySpec()
	_spec.Node.Columns = utq.ctx.Fields
	if len(utq.ctx.Fields) > 0 {
		_spec.Unique = utq.ctx.Unique != nil && *utq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, utq.driver, _spec)
}

func (utq *UserTweetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usertweet.Table, usertweet.Columns, sqlgraph.NewFieldSpec(usertweet.FieldID, field.TypeInt))
	_spec.From = utq.sql
	if unique := utq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if utq.path != nil {
		_spec.Unique = true
	}
	if fields := utq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usertweet.FieldID)
		for i := range fields {
			if fields[i] != usertweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if utq.withUser != nil {
			_spec.Node.AddColumnOnce(usertweet.FieldUserID)
		}
		if utq.withTweet != nil {
			_spec.Node.AddColumnOnce(usertweet.FieldTweetID)
		}
	}
	if ps := utq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := utq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := utq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := utq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (utq *UserTweetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(utq.driver.Dialect())
	t1 := builder.Table(usertweet.Table)
	columns := utq.ctx.Fields
	if len(columns) == 0 {
		columns = usertweet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if utq.sql != nil {
		selector = utq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if utq.ctx.Unique != nil && *utq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range utq.predicates {
		p(selector)
	}
	for _, p := range utq.order {
		p(selector)
	}
	if offset := utq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := utq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserTweetGroupBy is the group-by builder for UserTweet entities.
type UserTweetGroupBy struct {
	selector
	build *UserTweetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (utgb *UserTweetGroupBy) Aggregate(fns ...AggregateFunc) *UserTweetGroupBy {
	utgb.fns = append(utgb.fns, fns...)
	return utgb
}

// Scan applies the selector query and scans the result into the given value.
func (utgb *UserTweetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, utgb.build.ctx, fluent.OpQueryGroupBy)
	if err := utgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserTweetQuery, *UserTweetGroupBy](ctx, utgb.build, utgb, utgb.build.inters, v)
}

func (utgb *UserTweetGroupBy) sqlScan(ctx context.Context, root *UserTweetQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(utgb.fns))
	for _, fn := range utgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*utgb.flds)+len(utgb.fns))
		for _, f := range *utgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*utgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := utgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserTweetSelect is the builder for selecting fields of UserTweet entities.
type UserTweetSelect struct {
	*UserTweetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uts *UserTweetSelect) Aggregate(fns ...AggregateFunc) *UserTweetSelect {
	uts.fns = append(uts.fns, fns...)
	return uts
}

// Scan applies the selector query and scans the result into the given value.
func (uts *UserTweetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uts.ctx, fluent.OpQuerySelect)
	if err := uts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserTweetQuery, *UserTweetSelect](ctx, uts.UserTweetQuery, uts, uts.inters, v)
}

func (uts *UserTweetSelect) sqlScan(ctx context.Context, root *UserTweetQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uts.fns))
	for _, fn := range uts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
