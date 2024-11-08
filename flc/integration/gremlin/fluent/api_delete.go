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
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/api"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
)

// APIDelete is the builder for deleting a Api entity.
type APIDelete struct {
	config
	hooks    []Hook
	mutation *APIMutation
}

// Where appends a list predicates to the APIDelete builder.
func (ad *APIDelete) Where(ps ...predicate.Api) *APIDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *APIDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.gremlinExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *APIDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *APIDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ad.gremlin().Query()
	if err := ad.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	ad.mutation.done = true
	return res.ReadInt()
}

func (ad *APIDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(api.Label)
	for _, p := range ad.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// APIDeleteOne is the builder for deleting a single Api entity.
type APIDeleteOne struct {
	ad *APIDelete
}

// Where appends a list predicates to the APIDelete builder.
func (ado *APIDeleteOne) Where(ps ...predicate.Api) *APIDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *APIDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{api.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *APIDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
