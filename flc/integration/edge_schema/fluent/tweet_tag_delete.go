// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/tweet_tag"
	"github.com/usalko/fluent/schema/field"
)

// TweetTagDelete is the builder for deleting a TweetTag entity.
type TweetTagDelete struct {
	config
	hooks    []Hook
	mutation *TweetTagMutation
}

// Where appends a list predicates to the TweetTagDelete builder.
func (ttd *TweetTagDelete) Where(ps ...predicate.TweetTag) *TweetTagDelete {
	ttd.mutation.Where(ps...)
	return ttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttd *TweetTagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ttd.sqlExec, ttd.mutation, ttd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ttd *TweetTagDelete) ExecX(ctx context.Context) int {
	n, err := ttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttd *TweetTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tweet_tag.Table, sqlgraph.NewFieldSpec(tweet_tag.FieldID, field.TypeUUID))
	if ps := ttd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ttd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ttd.mutation.done = true
	return affected, err
}

// TweetTagDeleteOne is the builder for deleting a single TweetTag entity.
type TweetTagDeleteOne struct {
	ttd *TweetTagDelete
}

// Where appends a list predicates to the TweetTagDelete builder.
func (ttdo *TweetTagDeleteOne) Where(ps ...predicate.TweetTag) *TweetTagDeleteOne {
	ttdo.ttd.mutation.Where(ps...)
	return ttdo
}

// Exec executes the deletion query.
func (ttdo *TweetTagDeleteOne) Exec(ctx context.Context) error {
	n, err := ttdo.ttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tweet_tag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttdo *TweetTagDeleteOne) ExecX(ctx context.Context) {
	if err := ttdo.Exec(ctx); err != nil {
		panic(err)
	}
}
