// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"

	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/builder"
)

// BuilderCreate is the builder for creating a Builder entity.
type BuilderCreate struct {
	config
	mutation *BuilderMutation
	hooks    []Hook
}

// Mutation returns the BuilderMutation object of the builder.
func (bc *BuilderCreate) Mutation() *BuilderMutation {
	return bc.mutation
}

// Save creates the Builder in the database.
func (bc *BuilderCreate) Save(ctx context.Context) (*Builder, error) {
	return withHooks(ctx, bc.gremlinSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BuilderCreate) SaveX(ctx context.Context) *Builder {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BuilderCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BuilderCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BuilderCreate) check() error {
	return nil
}

func (bc *BuilderCreate) gremlinSave(ctx context.Context) (*Builder, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := bc.gremlin().Query()
	if err := bc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Builder{config: bc.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	bc.mutation.id = &rnode.ID
	bc.mutation.done = true
	return rnode, nil
}

func (bc *BuilderCreate) gremlin() *dsl.Traversal {
	v := g.AddV(builder.Label)
	return v.ValueMap(true)
}

// BuilderCreateBulk is the builder for creating many Builder entities in bulk.
type BuilderCreateBulk struct {
	config
	err      error
	builders []*BuilderCreate
}
