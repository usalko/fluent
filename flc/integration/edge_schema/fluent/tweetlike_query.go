// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/tweet"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/tweetlike"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/user"
)

// TweetLikeQuery is the builder for querying TweetLike entities.
type TweetLikeQuery struct {
	config
	ctx        *QueryContext
	order      []tweetlike.OrderOption
	inters     []Interceptor
	predicates []predicate.TweetLike
	withTweet  *TweetQuery
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TweetLikeQuery builder.
func (tlq *TweetLikeQuery) Where(ps ...predicate.TweetLike) *TweetLikeQuery {
	tlq.predicates = append(tlq.predicates, ps...)
	return tlq
}

// Limit the number of records to be returned by this query.
func (tlq *TweetLikeQuery) Limit(limit int) *TweetLikeQuery {
	tlq.ctx.Limit = &limit
	return tlq
}

// Offset to start from.
func (tlq *TweetLikeQuery) Offset(offset int) *TweetLikeQuery {
	tlq.ctx.Offset = &offset
	return tlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tlq *TweetLikeQuery) Unique(unique bool) *TweetLikeQuery {
	tlq.ctx.Unique = &unique
	return tlq
}

// Order specifies how the records should be ordered.
func (tlq *TweetLikeQuery) Order(o ...tweetlike.OrderOption) *TweetLikeQuery {
	tlq.order = append(tlq.order, o...)
	return tlq
}

// QueryTweet chains the current query on the "tweet" edge.
func (tlq *TweetLikeQuery) QueryTweet() *TweetQuery {
	query := (&TweetClient{config: tlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweetlike.Table, tweetlike.TweetColumn, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tweetlike.TweetTable, tweetlike.TweetColumn),
		)
		fromU = sqlgraph.SetNeighbors(tlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (tlq *TweetLikeQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: tlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweetlike.Table, tweetlike.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tweetlike.UserTable, tweetlike.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TweetLike entity from the query.
// Returns a *NotFoundError when no TweetLike was found.
func (tlq *TweetLikeQuery) First(ctx context.Context) (*TweetLike, error) {
	nodes, err := tlq.Limit(1).All(setContextOp(ctx, tlq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tweetlike.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tlq *TweetLikeQuery) FirstX(ctx context.Context) *TweetLike {
	node, err := tlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single TweetLike entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TweetLike entity is found.
// Returns a *NotFoundError when no TweetLike entities are found.
func (tlq *TweetLikeQuery) Only(ctx context.Context) (*TweetLike, error) {
	nodes, err := tlq.Limit(2).All(setContextOp(ctx, tlq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tweetlike.Label}
	default:
		return nil, &NotSingularError{tweetlike.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tlq *TweetLikeQuery) OnlyX(ctx context.Context) *TweetLike {
	node, err := tlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of TweetLikes.
func (tlq *TweetLikeQuery) All(ctx context.Context) ([]*TweetLike, error) {
	ctx = setContextOp(ctx, tlq.ctx, fluent.OpQueryAll)
	if err := tlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TweetLike, *TweetLikeQuery]()
	return withInterceptors[[]*TweetLike](ctx, tlq, qr, tlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tlq *TweetLikeQuery) AllX(ctx context.Context) []*TweetLike {
	nodes, err := tlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (tlq *TweetLikeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tlq.ctx, fluent.OpQueryCount)
	if err := tlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tlq, querierCount[*TweetLikeQuery](), tlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tlq *TweetLikeQuery) CountX(ctx context.Context) int {
	count, err := tlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tlq *TweetLikeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tlq.ctx, fluent.OpQueryExist)
	switch _, err := tlq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("fluent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tlq *TweetLikeQuery) ExistX(ctx context.Context) bool {
	exist, err := tlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TweetLikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tlq *TweetLikeQuery) Clone() *TweetLikeQuery {
	if tlq == nil {
		return nil
	}
	return &TweetLikeQuery{
		config:     tlq.config,
		ctx:        tlq.ctx.Clone(),
		order:      append([]tweetlike.OrderOption{}, tlq.order...),
		inters:     append([]Interceptor{}, tlq.inters...),
		predicates: append([]predicate.TweetLike{}, tlq.predicates...),
		withTweet:  tlq.withTweet.Clone(),
		withUser:   tlq.withUser.Clone(),
		// clone intermediate query.
		sql:  tlq.sql.Clone(),
		path: tlq.path,
	}
}

// WithTweet tells the query-builder to eager-load the nodes that are connected to
// the "tweet" edge. The optional arguments are used to configure the query builder of the edge.
func (tlq *TweetLikeQuery) WithTweet(opts ...func(*TweetQuery)) *TweetLikeQuery {
	query := (&TweetClient{config: tlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tlq.withTweet = query
	return tlq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tlq *TweetLikeQuery) WithUser(opts ...func(*UserQuery)) *TweetLikeQuery {
	query := (&UserClient{config: tlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tlq.withUser = query
	return tlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LikedAt time.Time `json:"liked_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TweetLike.Query().
//		GroupBy(tweetlike.FieldLikedAt).
//		Aggregate(fluent.Count()).
//		Scan(ctx, &v)
func (tlq *TweetLikeQuery) GroupBy(field string, fields ...string) *TweetLikeGroupBy {
	tlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TweetLikeGroupBy{build: tlq}
	grbuild.flds = &tlq.ctx.Fields
	grbuild.label = tweetlike.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LikedAt time.Time `json:"liked_at,omitempty"`
//	}
//
//	client.TweetLike.Query().
//		Select(tweetlike.FieldLikedAt).
//		Scan(ctx, &v)
func (tlq *TweetLikeQuery) Select(fields ...string) *TweetLikeSelect {
	tlq.ctx.Fields = append(tlq.ctx.Fields, fields...)
	sbuild := &TweetLikeSelect{TweetLikeQuery: tlq}
	sbuild.label = tweetlike.Label
	sbuild.flds, sbuild.scan = &tlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TweetLikeSelect configured with the given aggregations.
func (tlq *TweetLikeQuery) Aggregate(fns ...AggregateFunc) *TweetLikeSelect {
	return tlq.Select().Aggregate(fns...)
}

func (tlq *TweetLikeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tlq.inters {
		if inter == nil {
			return fmt.Errorf("fluent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tlq); err != nil {
				return err
			}
		}
	}
	for _, f := range tlq.ctx.Fields {
		if !tweetlike.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
		}
	}
	if tlq.path != nil {
		prev, err := tlq.path(ctx)
		if err != nil {
			return err
		}
		tlq.sql = prev
	}
	if tweetlike.Policy == nil {
		return errors.New("fluent: uninitialized tweetlike.Policy (forgotten import ent/runtime?)")
	}
	if err := tweetlike.Policy.EvalQuery(ctx, tlq); err != nil {
		return err
	}
	return nil
}

func (tlq *TweetLikeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TweetLike, error) {
	var (
		nodes       = []*TweetLike{}
		_spec       = tlq.querySpec()
		loadedTypes = [2]bool{
			tlq.withTweet != nil,
			tlq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TweetLike).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TweetLike{config: tlq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tlq.withTweet; query != nil {
		if err := tlq.loadTweet(ctx, query, nodes, nil,
			func(n *TweetLike, e *Tweet) { n.Edges.Tweet = e }); err != nil {
			return nil, err
		}
	}
	if query := tlq.withUser; query != nil {
		if err := tlq.loadUser(ctx, query, nodes, nil,
			func(n *TweetLike, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tlq *TweetLikeQuery) loadTweet(ctx context.Context, query *TweetQuery, nodes []*TweetLike, init func(*TweetLike), assign func(*TweetLike, *Tweet)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TweetLike)
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
func (tlq *TweetLikeQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*TweetLike, init func(*TweetLike), assign func(*TweetLike, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TweetLike)
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

func (tlq *TweetLikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tlq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, tlq.driver, _spec)
}

func (tlq *TweetLikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tweetlike.Table, tweetlike.Columns, nil)
	_spec.From = tlq.sql
	if unique := tlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tlq.path != nil {
		_spec.Unique = true
	}
	if fields := tlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if tlq.withTweet != nil {
			_spec.Node.AddColumnOnce(tweetlike.FieldTweetID)
		}
		if tlq.withUser != nil {
			_spec.Node.AddColumnOnce(tweetlike.FieldUserID)
		}
	}
	if ps := tlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tlq *TweetLikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tlq.driver.Dialect())
	t1 := builder.Table(tweetlike.Table)
	columns := tlq.ctx.Fields
	if len(columns) == 0 {
		columns = tweetlike.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tlq.sql != nil {
		selector = tlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tlq.ctx.Unique != nil && *tlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tlq.predicates {
		p(selector)
	}
	for _, p := range tlq.order {
		p(selector)
	}
	if offset := tlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TweetLikeGroupBy is the group-by builder for TweetLike entities.
type TweetLikeGroupBy struct {
	selector
	build *TweetLikeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tlgb *TweetLikeGroupBy) Aggregate(fns ...AggregateFunc) *TweetLikeGroupBy {
	tlgb.fns = append(tlgb.fns, fns...)
	return tlgb
}

// Scan applies the selector query and scans the result into the given value.
func (tlgb *TweetLikeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tlgb.build.ctx, fluent.OpQueryGroupBy)
	if err := tlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TweetLikeQuery, *TweetLikeGroupBy](ctx, tlgb.build, tlgb, tlgb.build.inters, v)
}

func (tlgb *TweetLikeGroupBy) sqlScan(ctx context.Context, root *TweetLikeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tlgb.fns))
	for _, fn := range tlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tlgb.flds)+len(tlgb.fns))
		for _, f := range *tlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TweetLikeSelect is the builder for selecting fields of TweetLike entities.
type TweetLikeSelect struct {
	*TweetLikeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tls *TweetLikeSelect) Aggregate(fns ...AggregateFunc) *TweetLikeSelect {
	tls.fns = append(tls.fns, fns...)
	return tls
}

// Scan applies the selector query and scans the result into the given value.
func (tls *TweetLikeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tls.ctx, fluent.OpQuerySelect)
	if err := tls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TweetLikeQuery, *TweetLikeSelect](ctx, tls.TweetLikeQuery, tls, tls.inters, v)
}

func (tls *TweetLikeSelect) sqlScan(ctx context.Context, root *TweetLikeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tls.fns))
	for _, fn := range tls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
