// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/mixin_id"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// MixinIDDelete is the builder for deleting a MixinID entity.
type MixinIDDelete struct {
	config
	hooks    []Hook
	mutation *MixinIDMutation
}

// Where appends a list predicates to the MixinIDDelete builder.
func (mi *MixinIDDelete) Where(ps ...predicate.MixinID) *MixinIDDelete {
	mi.mutation.Where(ps...)
	return mi
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mi *MixinIDDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mi.sqlExec, mi.mutation, mi.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mi *MixinIDDelete) ExecX(ctx context.Context) int {
	n, err := mi.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mi *MixinIDDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(mixin_id.Table, sqlgraph.NewFieldSpec(mixin_id.FieldID, field.TypeUUID))
	if ps := mi.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mi.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mi.mutation.done = true
	return affected, err
}

// MixinIDDeleteOne is the builder for deleting a single MixinID entity.
type MixinIDDeleteOne struct {
	mi *MixinIDDelete
}

// Where appends a list predicates to the MixinIDDelete builder.
func (mio *MixinIDDeleteOne) Where(ps ...predicate.MixinID) *MixinIDDeleteOne {
	mio.mi.mutation.Where(ps...)
	return mio
}

// Exec executes the deletion query.
func (mio *MixinIDDeleteOne) Exec(ctx context.Context) error {
	n, err := mio.mi.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mixin_id.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mio *MixinIDDeleteOne) ExecX(ctx context.Context) {
	if err := mio.Exec(ctx); err != nil {
		panic(err)
	}
}
