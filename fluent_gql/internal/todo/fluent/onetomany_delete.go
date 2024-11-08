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
// Code generated by entc, DO NOT EDIT.

package fluent

import (
	"context"

	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/onetomany"
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/predicate"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/schema/field"
)

// OneToManyDelete is the builder for deleting a OneToMany entity.
type OneToManyDelete struct {
	config
	hooks    []Hook
	mutation *OneToManyMutation
}

// Where appends a list predicates to the OneToManyDelete builder.
func (otmd *OneToManyDelete) Where(ps ...predicate.OneToMany) *OneToManyDelete {
	otmd.mutation.Where(ps...)
	return otmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (otmd *OneToManyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, otmd.sqlExec, otmd.mutation, otmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (otmd *OneToManyDelete) ExecX(ctx context.Context) int {
	n, err := otmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (otmd *OneToManyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(onetomany.Table, sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt))
	if ps := otmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, otmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	otmd.mutation.done = true
	return affected, err
}

// OneToManyDeleteOne is the builder for deleting a single OneToMany entity.
type OneToManyDeleteOne struct {
	otmd *OneToManyDelete
}

// Where appends a list predicates to the OneToManyDelete builder.
func (otmdo *OneToManyDeleteOne) Where(ps ...predicate.OneToMany) *OneToManyDeleteOne {
	otmdo.otmd.mutation.Where(ps...)
	return otmdo
}

// Exec executes the deletion query.
func (otmdo *OneToManyDeleteOne) Exec(ctx context.Context) error {
	n, err := otmdo.otmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{onetomany.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (otmdo *OneToManyDeleteOne) ExecX(ctx context.Context) {
	if err := otmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
