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
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/ex_value_scan"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
)

// ExValueScanDelete is the builder for deleting a ExValueScan entity.
type ExValueScanDelete struct {
	config
	hooks    []Hook
	mutation *ExValueScanMutation
}

// Where appends a list predicates to the ExValueScanDelete builder.
func (evsd *ExValueScanDelete) Where(ps ...predicate.ExValueScan) *ExValueScanDelete {
	evsd.mutation.Where(ps...)
	return evsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (evsd *ExValueScanDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, evsd.gremlinExec, evsd.mutation, evsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (evsd *ExValueScanDelete) ExecX(ctx context.Context) int {
	n, err := evsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (evsd *ExValueScanDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := evsd.gremlin().Query()
	if err := evsd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	evsd.mutation.done = true
	return res.ReadInt()
}

func (evsd *ExValueScanDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(ex_value_scan.Label)
	for _, p := range evsd.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// ExValueScanDeleteOne is the builder for deleting a single ExValueScan entity.
type ExValueScanDeleteOne struct {
	evsd *ExValueScanDelete
}

// Where appends a list predicates to the ExValueScanDelete builder.
func (evsdo *ExValueScanDeleteOne) Where(ps ...predicate.ExValueScan) *ExValueScanDeleteOne {
	evsdo.evsd.mutation.Where(ps...)
	return evsdo
}

// Exec executes the deletion query.
func (evsdo *ExValueScanDeleteOne) Exec(ctx context.Context) error {
	n, err := evsdo.evsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ex_value_scan.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (evsdo *ExValueScanDeleteOne) ExecX(ctx context.Context) {
	if err := evsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
