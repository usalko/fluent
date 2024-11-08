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
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/comment"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
)

// CommentDelete is the builder for deleting a Comment entity.
type CommentDelete struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentDelete builder.
func (cd *CommentDelete) Where(ps ...predicate.Comment) *CommentDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CommentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cd.gremlinExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CommentDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CommentDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := cd.gremlin().Query()
	if err := cd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	cd.mutation.done = true
	return res.ReadInt()
}

func (cd *CommentDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(comment.Label)
	for _, p := range cd.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// CommentDeleteOne is the builder for deleting a single Comment entity.
type CommentDeleteOne struct {
	cd *CommentDelete
}

// Where appends a list predicates to the CommentDelete builder.
func (cdo *CommentDeleteOne) Where(ps ...predicate.Comment) *CommentDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *CommentDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{comment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CommentDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}
