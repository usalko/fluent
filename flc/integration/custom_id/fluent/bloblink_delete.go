// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/bloblink"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/predicate"
)

// BlobLinkDelete is the builder for deleting a BlobLink entity.
type BlobLinkDelete struct {
	config
	hooks    []Hook
	mutation *BlobLinkMutation
}

// Where appends a list predicates to the BlobLinkDelete builder.
func (bld *BlobLinkDelete) Where(ps ...predicate.BlobLink) *BlobLinkDelete {
	bld.mutation.Where(ps...)
	return bld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bld *BlobLinkDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bld.sqlExec, bld.mutation, bld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bld *BlobLinkDelete) ExecX(ctx context.Context) int {
	n, err := bld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bld *BlobLinkDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bloblink.Table, nil)
	if ps := bld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bld.mutation.done = true
	return affected, err
}

// BlobLinkDeleteOne is the builder for deleting a single BlobLink entity.
type BlobLinkDeleteOne struct {
	bld *BlobLinkDelete
}

// Where appends a list predicates to the BlobLinkDelete builder.
func (bldo *BlobLinkDeleteOne) Where(ps ...predicate.BlobLink) *BlobLinkDeleteOne {
	bldo.bld.mutation.Where(ps...)
	return bldo
}

// Exec executes the deletion query.
func (bldo *BlobLinkDeleteOne) Exec(ctx context.Context) error {
	n, err := bldo.bld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bloblink.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bldo *BlobLinkDeleteOne) ExecX(ctx context.Context) {
	if err := bldo.Exec(ctx); err != nil {
		panic(err)
	}
}
