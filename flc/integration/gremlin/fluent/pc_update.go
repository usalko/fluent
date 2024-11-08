// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"

	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/pc"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
)

// PCUpdate is the builder for updating PC entities.
type PCUpdate struct {
	config
	hooks    []Hook
	mutation *PCMutation
}

// Where appends a list predicates to the PCUpdate builder.
func (pu *PCUpdate) Where(ps ...predicate.PC) *PCUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// Mutation returns the PCMutation object of the builder.
func (pu *PCUpdate) Mutation() *PCMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PCUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.gremlinSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PCUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PCUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PCUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PCUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := pu.gremlin().Query()
	if err := pu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	pu.mutation.done = true
	return res.ReadInt()
}

func (pu *PCUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(pc.Label)
	for _, p := range pu.mutation.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// PCUpdateOne is the builder for updating a single PC entity.
type PCUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PCMutation
}

// Mutation returns the PCMutation object of the builder.
func (puo *PCUpdateOne) Mutation() *PCMutation {
	return puo.mutation
}

// Where appends a list predicates to the PCUpdate builder.
func (puo *PCUpdateOne) Where(ps ...predicate.PC) *PCUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PCUpdateOne) Select(field string, fields ...string) *PCUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated PC entity.
func (puo *PCUpdateOne) Save(ctx context.Context) (*PC, error) {
	return withHooks(ctx, puo.gremlinSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PCUpdateOne) SaveX(ctx context.Context) *PC {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PCUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PCUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PCUpdateOne) gremlinSave(ctx context.Context) (*PC, error) {
	res := &gremlin.Response{}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PC.id" for update`)}
	}
	query, bindings := puo.gremlin(id).Query()
	if err := puo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	puo.mutation.done = true
	_pc := &PC{config: puo.config}
	if err := _pc.FromResponse(res); err != nil {
		return nil, err
	}
	return _pc, nil
}

func (puo *PCUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if len(puo.fields) > 0 {
		fields := make([]any, 0, len(puo.fields)+1)
		fields = append(fields, true)
		for _, f := range puo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
