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

	"github.com/99designs/gqlgen/graphql"
)

func (c *Category) Todos(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy []*TodoOrder, where *TodoWhereInput,
) (*TodoConnection, error) {
	opts := []TodoPaginateOption{
		WithTodoOrder(orderBy),
		WithTodoFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[0][alias]
	if nodes, err := c.NamedTodos(alias); err == nil || hasTotalCount {
		pager, err := newTodoPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &TodoConnection{Edges: []*TodoEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QueryTodos().Paginate(ctx, after, first, before, last, opts...)
}

func (c *Category) SubCategories(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy []*CategoryOrder, where *CategoryWhereInput,
) (*CategoryConnection, error) {
	opts := []CategoryPaginateOption{
		WithCategoryOrder(orderBy),
		WithCategoryFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := c.Edges.totalCount[1][alias]
	if nodes, err := c.NamedSubCategories(alias); err == nil || hasTotalCount {
		pager, err := newCategoryPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &CategoryConnection{Edges: []*CategoryEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return c.QuerySubCategories().Paginate(ctx, after, first, before, last, opts...)
}

func (f *Friendship) User(ctx context.Context) (*User, error) {
	result, err := f.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = f.QueryUser().Only(ctx)
	}
	return result, err
}

func (f *Friendship) Friend(ctx context.Context) (*User, error) {
	result, err := f.Edges.FriendOrErr()
	if IsNotLoaded(err) {
		result, err = f.QueryFriend().Only(ctx)
	}
	return result, err
}

func (gr *Group) Users(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *UserOrder, where *UserWhereInput,
) (*UserConnection, error) {
	opts := []UserPaginateOption{
		WithUserOrder(orderBy),
		WithUserFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := gr.Edges.totalCount[0][alias]
	if nodes, err := gr.NamedUsers(alias); err == nil || hasTotalCount {
		pager, err := newUserPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &UserConnection{Edges: []*UserEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return gr.QueryUsers().Paginate(ctx, after, first, before, last, opts...)
}

func (otm *OneToMany) Parent(ctx context.Context) (*OneToMany, error) {
	result, err := otm.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = otm.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (otm *OneToMany) Children(ctx context.Context) (result []*OneToMany, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = otm.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = otm.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = otm.QueryChildren().All(ctx)
	}
	return result, err
}

func (pr *Project) Todos(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy []*TodoOrder, where *TodoWhereInput,
) (*TodoConnection, error) {
	opts := []TodoPaginateOption{
		WithTodoOrder(orderBy),
		WithTodoFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := pr.Edges.totalCount[0][alias]
	if nodes, err := pr.NamedTodos(alias); err == nil || hasTotalCount {
		pager, err := newTodoPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &TodoConnection{Edges: []*TodoEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return pr.QueryTodos().Paginate(ctx, after, first, before, last, opts...)
}

func (t *Todo) Parent(ctx context.Context) (*Todo, error) {
	result, err := t.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Todo) Children(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy []*TodoOrder, where *TodoWhereInput,
) (*TodoConnection, error) {
	opts := []TodoPaginateOption{
		WithTodoOrder(orderBy),
		WithTodoFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := t.Edges.totalCount[1][alias]
	if nodes, err := t.NamedChildren(alias); err == nil || hasTotalCount {
		pager, err := newTodoPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &TodoConnection{Edges: []*TodoEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return t.QueryChildren().Paginate(ctx, after, first, before, last, opts...)
}

func (t *Todo) Category(ctx context.Context) (*Category, error) {
	result, err := t.Edges.CategoryOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryCategory().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Groups(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, where *GroupWhereInput,
) (*GroupConnection, error) {
	opts := []GroupPaginateOption{
		WithGroupFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[0][alias]
	if nodes, err := u.NamedGroups(alias); err == nil || hasTotalCount {
		pager, err := newGroupPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &GroupConnection{Edges: []*GroupEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryGroups().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) Friends(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *UserOrder, where *UserWhereInput,
) (*UserConnection, error) {
	opts := []UserPaginateOption{
		WithUserOrder(orderBy),
		WithUserFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[1][alias]
	if nodes, err := u.NamedFriends(alias); err == nil || hasTotalCount {
		pager, err := newUserPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &UserConnection{Edges: []*UserEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryFriends().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) Friendships(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, where *FriendshipWhereInput,
) (*FriendshipConnection, error) {
	opts := []FriendshipPaginateOption{
		WithFriendshipFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[2][alias]
	if nodes, err := u.NamedFriendships(alias); err == nil || hasTotalCount {
		pager, err := newFriendshipPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &FriendshipConnection{Edges: []*FriendshipEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryFriendships().Paginate(ctx, after, first, before, last, opts...)
}
