// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/fluent/groupinfo"
	"github.com/usalko/fluent/flc/integration/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// GroupInfoDelete is the builder for deleting a GroupInfo entity.
type GroupInfoDelete struct {
	config
	hooks    []Hook
	mutation *GroupInfoMutation
}

// Where appends a list predicates to the GroupInfoDelete builder.
func (gid *GroupInfoDelete) Where(ps ...predicate.GroupInfo) *GroupInfoDelete {
	gid.mutation.Where(ps...)
	return gid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gid *GroupInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gid.sqlExec, gid.mutation, gid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gid *GroupInfoDelete) ExecX(ctx context.Context) int {
	n, err := gid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gid *GroupInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(groupinfo.Table, sqlgraph.NewFieldSpec(groupinfo.FieldID, field.TypeInt))
	if ps := gid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gid.mutation.done = true
	return affected, err
}

// GroupInfoDeleteOne is the builder for deleting a single GroupInfo entity.
type GroupInfoDeleteOne struct {
	gid *GroupInfoDelete
}

// Where appends a list predicates to the GroupInfoDelete builder.
func (gido *GroupInfoDeleteOne) Where(ps ...predicate.GroupInfo) *GroupInfoDeleteOne {
	gido.gid.mutation.Where(ps...)
	return gido
}

// Exec executes the deletion query.
func (gido *GroupInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := gido.gid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{groupinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gido *GroupInfoDeleteOne) ExecX(ctx context.Context) {
	if err := gido.Exec(ctx); err != nil {
		panic(err)
	}
}
