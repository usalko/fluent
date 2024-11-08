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

package todopulid

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"fmt"
	"time"

	"github.com/usalko/fluent/fluent_gql"
	"github.com/usalko/fluent/fluent_gql/internal/todopulid/fluent"
	"github.com/usalko/fluent/fluent_gql/internal/todopulid/fluent/schema/pulid"
	"github.com/usalko/fluent/fluent_gql/internal/todopulid/fluent/todo"
)

// TodosCount is the resolver for the todosCount field.
func (r *categoryResolver) TodosCount(ctx context.Context, obj *fluent.Category) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input fluent.CreateCategoryInput) (*fluent.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input fluent.CreateTodoInput) (*fluent.Todo, error) {
	return fluent.FromContext(ctx).Todo.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id pulid.ID, input fluent.UpdateTodoInput) (*fluent.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClearTodos is the resolver for the clearTodos field.
func (r *mutationResolver) ClearTodos(ctx context.Context) (int, error) {
	client := fluent.FromContext(ctx)
	return client.Todo.
		Delete().
		Exec(ctx)
}

// UpdateFriendship is the resolver for the updateFriendship field.
func (r *mutationResolver) UpdateFriendship(ctx context.Context, id pulid.ID, input UpdateFriendshipInput) (*fluent.Friendship, error) {
	panic(fmt.Errorf("not implemented"))
}

// Ping is the resolver for the ping field.
func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// TodosWithJoins is the resolver for the todosWithJoins field.
func (r *queryResolver) TodosWithJoins(ctx context.Context, after *fluent_gql.Cursor[pulid.ID], first *int, before *fluent_gql.Cursor[pulid.ID], last *int, orderBy []*fluent.TodoOrder, where *fluent.TodoWhereInput) (*fluent.TodoConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// ExtendedField is the resolver for the extendedField field.
func (r *todoResolver) ExtendedField(ctx context.Context, obj *fluent.Todo) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateTodos is the resolver for the createTodos field.
func (r *createCategoryInputResolver) CreateTodos(ctx context.Context, obj *fluent.CreateCategoryInput, data []*fluent.CreateTodoInput) error {
	panic(fmt.Errorf("not implemented"))
}

// CreatedToday is the resolver for the createdToday field.
func (r *todoWhereInputResolver) CreatedToday(ctx context.Context, obj *fluent.TodoWhereInput, data *bool) error {
	if data == nil {
		return nil
	}

	startOfDay := time.Now().Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24*time.Hour - 1)
	if *data {
		obj.AddPredicates(todo.And(todo.CreatedAtGTE(startOfDay), todo.CreatedAtLTE(endOfDay)))
	} else {
		obj.AddPredicates(todo.Or(todo.CreatedAtLT(startOfDay), todo.CreatedAtGT(endOfDay)))
	}

	return nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
