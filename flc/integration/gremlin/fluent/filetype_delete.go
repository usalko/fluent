// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"

	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/__"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/filetype"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
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
	return withHooks(ctx, ftd.gremlinExec, ftd.mutation, ftd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FileTypeDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FileTypeDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftd.gremlin().Query()
	if err := ftd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	ftd.mutation.done = true
	return res.ReadInt()
}

func (ftd *FileTypeDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(filetype.Label)
	for _, p := range ftd.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
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
		return &NotFoundError{filetype.Label}
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