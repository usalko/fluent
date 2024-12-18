// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/attached_file"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// AttachedFileDelete is the builder for deleting a AttachedFile entity.
type AttachedFileDelete struct {
	config
	hooks    []Hook
	mutation *AttachedFileMutation
}

// Where appends a list predicates to the AttachedFileDelete builder.
func (afd *AttachedFileDelete) Where(ps ...predicate.AttachedFile) *AttachedFileDelete {
	afd.mutation.Where(ps...)
	return afd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (afd *AttachedFileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, afd.sqlExec, afd.mutation, afd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (afd *AttachedFileDelete) ExecX(ctx context.Context) int {
	n, err := afd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (afd *AttachedFileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(attached_file.Table, sqlgraph.NewFieldSpec(attached_file.FieldID, field.TypeInt))
	if ps := afd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, afd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	afd.mutation.done = true
	return affected, err
}

// AttachedFileDeleteOne is the builder for deleting a single AttachedFile entity.
type AttachedFileDeleteOne struct {
	afd *AttachedFileDelete
}

// Where appends a list predicates to the AttachedFileDelete builder.
func (afdo *AttachedFileDeleteOne) Where(ps ...predicate.AttachedFile) *AttachedFileDeleteOne {
	afdo.afd.mutation.Where(ps...)
	return afdo
}

// Exec executes the deletion query.
func (afdo *AttachedFileDeleteOne) Exec(ctx context.Context) error {
	n, err := afdo.afd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{attached_file.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (afdo *AttachedFileDeleteOne) ExecX(ctx context.Context) {
	if err := afdo.Exec(ctx); err != nil {
		panic(err)
	}
}
