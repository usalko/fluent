// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"

	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/api"
)

// APICreate is the builder for creating a Api entity.
type APICreate struct {
	config
	mutation *APIMutation
	hooks    []Hook
}

// Mutation returns the APIMutation object of the builder.
func (ac *APICreate) Mutation() *APIMutation {
	return ac.mutation
}

// Save creates the Api in the database.
func (ac *APICreate) Save(ctx context.Context) (*Api, error) {
	return withHooks(ctx, ac.gremlinSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *APICreate) SaveX(ctx context.Context) *Api {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *APICreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *APICreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *APICreate) check() error {
	return nil
}

func (ac *APICreate) gremlinSave(ctx context.Context) (*Api, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := ac.gremlin().Query()
	if err := ac.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Api{config: ac.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	ac.mutation.id = &rnode.ID
	ac.mutation.done = true
	return rnode, nil
}

func (ac *APICreate) gremlin() *dsl.Traversal {
	v := g.AddV(api.Label)
	return v.ValueMap(true)
}

// APICreateBulk is the builder for creating many Api entities in bulk.
type APICreateBulk struct {
	config
	err      error
	builders []*APICreate
}
