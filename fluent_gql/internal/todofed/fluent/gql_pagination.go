// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by flc, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/usalko/fluent/fluent_gql"
	"github.com/usalko/fluent/fluent_gql/internal/todofed/fluent/category"
	"github.com/usalko/fluent/fluent_gql/internal/todofed/fluent/todo"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common fluent_gql types.
type (
	Cursor         = fluent_gql.Cursor[int]
	PageInfo       = fluent_gql.PageInfo[int]
	OrderDirection = fluent_gql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == fluent_gql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// CategoryEdge is the edge representation of Category.
type CategoryEdge struct {
	Node   *Category `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// CategoryConnection is the connection containing edges to Category.
type CategoryConnection struct {
	Edges      []*CategoryEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

func (c *CategoryConnection) build(nodes []*Category, pager *categoryPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Category
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Category {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Category {
			return nodes[i]
		}
	}
	c.Edges = make([]*CategoryEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &CategoryEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CategoryPaginateOption enables pagination customization.
type CategoryPaginateOption func(*categoryPager) error

// WithCategoryOrder configures pagination ordering.
func WithCategoryOrder(order *CategoryOrder) CategoryPaginateOption {
	if order == nil {
		order = DefaultCategoryOrder
	}
	o := *order
	return func(pager *categoryPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCategoryOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCategoryFilter configures pagination filter.
func WithCategoryFilter(filter func(*CategoryQuery) (*CategoryQuery, error)) CategoryPaginateOption {
	return func(pager *categoryPager) error {
		if filter == nil {
			return errors.New("CategoryQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type categoryPager struct {
	reverse bool
	order   *CategoryOrder
	filter  func(*CategoryQuery) (*CategoryQuery, error)
}

func newCategoryPager(opts []CategoryPaginateOption, reverse bool) (*categoryPager, error) {
	pager := &categoryPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCategoryOrder
	}
	return pager, nil
}

func (p *categoryPager) applyFilter(query *CategoryQuery) (*CategoryQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *categoryPager) toCursor(c *Category) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *categoryPager) applyCursors(query *CategoryQuery, after, before *Cursor) (*CategoryQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range fluent_gql.CursorsPredicate(after, before, DefaultCategoryOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *categoryPager) applyOrder(query *CategoryQuery) *CategoryQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultCategoryOrder.Field {
		query = query.Order(DefaultCategoryOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *categoryPager) orderExpr(query *CategoryQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultCategoryOrder.Field {
			b.Comma().Ident(DefaultCategoryOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Category.
func (c *CategoryQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CategoryPaginateOption,
) (*CategoryConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCategoryPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}
	conn := &CategoryConnection{Edges: []*CategoryEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := c.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if c, err = pager.applyCursors(c, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		c.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := c.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	c = pager.applyOrder(c)
	nodes, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// CategoryOrderFieldText orders Category by text.
	CategoryOrderFieldText = &CategoryOrderField{
		Value: func(c *Category) (fluent.Value, error) {
			return c.Text, nil
		},
		column: category.FieldText,
		toTerm: category.ByText,
		toCursor: func(c *Category) Cursor {
			return Cursor{
				ID:    c.ID,
				Value: c.Text,
			}
		},
	}
	// CategoryOrderFieldDuration orders Category by duration.
	CategoryOrderFieldDuration = &CategoryOrderField{
		Value: func(c *Category) (fluent.Value, error) {
			return c.Duration, nil
		},
		column: category.FieldDuration,
		toTerm: category.ByDuration,
		toCursor: func(c *Category) Cursor {
			return Cursor{
				ID:    c.ID,
				Value: c.Duration,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f CategoryOrderField) String() string {
	var str string
	switch f.column {
	case CategoryOrderFieldText.column:
		str = "TEXT"
	case CategoryOrderFieldDuration.column:
		str = "DURATION"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f CategoryOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *CategoryOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("CategoryOrderField %T must be a string", v)
	}
	switch str {
	case "TEXT":
		*f = *CategoryOrderFieldText
	case "DURATION":
		*f = *CategoryOrderFieldDuration
	default:
		return fmt.Errorf("%s is not a valid CategoryOrderField", str)
	}
	return nil
}

// CategoryOrderField defines the ordering field of Category.
type CategoryOrderField struct {
	// Value extracts the ordering value from the given Category.
	Value    func(*Category) (fluent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) category.OrderOption
	toCursor func(*Category) Cursor
}

// CategoryOrder defines the ordering of Category.
type CategoryOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *CategoryOrderField `json:"field"`
}

// DefaultCategoryOrder is the default ordering of Category.
var DefaultCategoryOrder = &CategoryOrder{
	Direction: fluent_gql.OrderDirectionAsc,
	Field: &CategoryOrderField{
		Value: func(c *Category) (fluent.Value, error) {
			return c.ID, nil
		},
		column: category.FieldID,
		toTerm: category.ByID,
		toCursor: func(c *Category) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Category into CategoryEdge.
func (c *Category) ToEdge(order *CategoryOrder) *CategoryEdge {
	if order == nil {
		order = DefaultCategoryOrder
	}
	return &CategoryEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// TodoEdge is the edge representation of Todo.
type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// TodoConnection is the connection containing edges to Todo.
type TodoConnection struct {
	Edges      []*TodoEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *TodoConnection) build(nodes []*Todo, pager *todoPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Todo
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Todo {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Todo {
			return nodes[i]
		}
	}
	c.Edges = make([]*TodoEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &TodoEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// TodoPaginateOption enables pagination customization.
type TodoPaginateOption func(*todoPager) error

// WithTodoOrder configures pagination ordering.
func WithTodoOrder(order *TodoOrder) TodoPaginateOption {
	if order == nil {
		order = DefaultTodoOrder
	}
	o := *order
	return func(pager *todoPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultTodoOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithTodoFilter configures pagination filter.
func WithTodoFilter(filter func(*TodoQuery) (*TodoQuery, error)) TodoPaginateOption {
	return func(pager *todoPager) error {
		if filter == nil {
			return errors.New("TodoQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type todoPager struct {
	reverse bool
	order   *TodoOrder
	filter  func(*TodoQuery) (*TodoQuery, error)
}

func newTodoPager(opts []TodoPaginateOption, reverse bool) (*todoPager, error) {
	pager := &todoPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultTodoOrder
	}
	return pager, nil
}

func (p *todoPager) applyFilter(query *TodoQuery) (*TodoQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *todoPager) toCursor(t *Todo) Cursor {
	return p.order.Field.toCursor(t)
}

func (p *todoPager) applyCursors(query *TodoQuery, after, before *Cursor) (*TodoQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range fluent_gql.CursorsPredicate(after, before, DefaultTodoOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *todoPager) applyOrder(query *TodoQuery) *TodoQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultTodoOrder.Field {
		query = query.Order(DefaultTodoOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *todoPager) orderExpr(query *TodoQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultTodoOrder.Field {
			b.Comma().Ident(DefaultTodoOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Todo.
func (t *TodoQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...TodoPaginateOption,
) (*TodoConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newTodoPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if t, err = pager.applyFilter(t); err != nil {
		return nil, err
	}
	conn := &TodoConnection{Edges: []*TodoEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := t.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if t, err = pager.applyCursors(t, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		t.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := t.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	t = pager.applyOrder(t)
	nodes, err := t.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// TodoOrderFieldCreatedAt orders Todo by created_at.
	TodoOrderFieldCreatedAt = &TodoOrderField{
		Value: func(t *Todo) (fluent.Value, error) {
			return t.CreatedAt, nil
		},
		column: todo.FieldCreatedAt,
		toTerm: todo.ByCreatedAt,
		toCursor: func(t *Todo) Cursor {
			return Cursor{
				ID:    t.ID,
				Value: t.CreatedAt,
			}
		},
	}
	// TodoOrderFieldStatus orders Todo by status.
	TodoOrderFieldStatus = &TodoOrderField{
		Value: func(t *Todo) (fluent.Value, error) {
			return t.Status, nil
		},
		column: todo.FieldStatus,
		toTerm: todo.ByStatus,
		toCursor: func(t *Todo) Cursor {
			return Cursor{
				ID:    t.ID,
				Value: t.Status,
			}
		},
	}
	// TodoOrderFieldPriority orders Todo by priority.
	TodoOrderFieldPriority = &TodoOrderField{
		Value: func(t *Todo) (fluent.Value, error) {
			return t.Priority, nil
		},
		column: todo.FieldPriority,
		toTerm: todo.ByPriority,
		toCursor: func(t *Todo) Cursor {
			return Cursor{
				ID:    t.ID,
				Value: t.Priority,
			}
		},
	}
	// TodoOrderFieldText orders Todo by text.
	TodoOrderFieldText = &TodoOrderField{
		Value: func(t *Todo) (fluent.Value, error) {
			return t.Text, nil
		},
		column: todo.FieldText,
		toTerm: todo.ByText,
		toCursor: func(t *Todo) Cursor {
			return Cursor{
				ID:    t.ID,
				Value: t.Text,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f TodoOrderField) String() string {
	var str string
	switch f.column {
	case TodoOrderFieldCreatedAt.column:
		str = "CREATED_AT"
	case TodoOrderFieldStatus.column:
		str = "STATUS"
	case TodoOrderFieldPriority.column:
		str = "PRIORITY"
	case TodoOrderFieldText.column:
		str = "TEXT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f TodoOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *TodoOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("TodoOrderField %T must be a string", v)
	}
	switch str {
	case "CREATED_AT":
		*f = *TodoOrderFieldCreatedAt
	case "STATUS":
		*f = *TodoOrderFieldStatus
	case "PRIORITY":
		*f = *TodoOrderFieldPriority
	case "TEXT":
		*f = *TodoOrderFieldText
	default:
		return fmt.Errorf("%s is not a valid TodoOrderField", str)
	}
	return nil
}

// TodoOrderField defines the ordering field of Todo.
type TodoOrderField struct {
	// Value extracts the ordering value from the given Todo.
	Value    func(*Todo) (fluent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) todo.OrderOption
	toCursor func(*Todo) Cursor
}

// TodoOrder defines the ordering of Todo.
type TodoOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *TodoOrderField `json:"field"`
}

// DefaultTodoOrder is the default ordering of Todo.
var DefaultTodoOrder = &TodoOrder{
	Direction: fluent_gql.OrderDirectionAsc,
	Field: &TodoOrderField{
		Value: func(t *Todo) (fluent.Value, error) {
			return t.ID, nil
		},
		column: todo.FieldID,
		toTerm: todo.ByID,
		toCursor: func(t *Todo) Cursor {
			return Cursor{ID: t.ID}
		},
	},
}

// ToEdge converts Todo into TodoEdge.
func (t *Todo) ToEdge(order *TodoOrder) *TodoEdge {
	if order == nil {
		order = DefaultTodoOrder
	}
	return &TodoEdge{
		Node:   t,
		Cursor: order.Field.toCursor(t),
	}
}
