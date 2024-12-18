// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/fluent/file_type"
	"github.com/usalko/fluent/flc/integration/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// FileTypeDelete is the builder for deleting a FileType entity.
type FileTypeDelete struct {
	config
	hooks    []Hook
	mutation *FileTypeMutation
}

// Where appends a list predicates to the FileTypeDelete builder.
func (ftd *FileTypeDelete) Where(ps ...predicate.FileType) *FileTypeDelete {
	ftd.mutation.Where(ps...)
	return ftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftd *FileTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ftd.sqlExec, ftd.mutation, ftd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FileTypeDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FileTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(file_type.Table, sqlgraph.NewFieldSpec(file_type.FieldID, field.TypeInt))
	if ps := ftd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ftd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ftd.mutation.done = true
	return affected, err
}

// FileTypeDeleteOne is the builder for deleting a single FileType entity.
type FileTypeDeleteOne struct {
	ftd *FileTypeDelete
}

// Where appends a list predicates to the FileTypeDelete builder.
func (ftdo *FileTypeDeleteOne) Where(ps ...predicate.FileType) *FileTypeDeleteOne {
	ftdo.ftd.mutation.Where(ps...)
	return ftdo
}

// Exec executes the deletion query.
func (ftdo *FileTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ftdo.ftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{file_type.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FileTypeDeleteOne) ExecX(ctx context.Context) {
	if err := ftdo.Exec(ctx); err != nil {
		panic(err)
	}
}
