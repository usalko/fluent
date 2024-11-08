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
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/goods"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
)

// GoodsUpdate is the builder for updating Goods entities.
type GoodsUpdate struct {
	config
	hooks    []Hook
	mutation *GoodsMutation
}

// Where appends a list predicates to the GoodsUpdate builder.
func (gu *GoodsUpdate) Where(ps ...predicate.Goods) *GoodsUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// Mutation returns the GoodsMutation object of the builder.
func (gu *GoodsUpdate) Mutation() *GoodsMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GoodsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gu.gremlinSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GoodsUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GoodsUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GoodsUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GoodsUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := gu.gremlin().Query()
	if err := gu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	gu.mutation.done = true
	return res.ReadInt()
}

func (gu *GoodsUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(goods.Label)
	for _, p := range gu.mutation.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// GoodsUpdateOne is the builder for updating a single Goods entity.
type GoodsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoodsMutation
}

// Mutation returns the GoodsMutation object of the builder.
func (guo *GoodsUpdateOne) Mutation() *GoodsMutation {
	return guo.mutation
}

// Where appends a list predicates to the GoodsUpdate builder.
func (guo *GoodsUpdateOne) Where(ps ...predicate.Goods) *GoodsUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GoodsUpdateOne) Select(field string, fields ...string) *GoodsUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Goods entity.
func (guo *GoodsUpdateOne) Save(ctx context.Context) (*Goods, error) {
	return withHooks(ctx, guo.gremlinSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GoodsUpdateOne) SaveX(ctx context.Context) *Goods {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GoodsUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GoodsUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GoodsUpdateOne) gremlinSave(ctx context.Context) (*Goods, error) {
	res := &gremlin.Response{}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Goods.id" for update`)}
	}
	query, bindings := guo.gremlin(id).Query()
	if err := guo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	guo.mutation.done = true
	_go := &Goods{config: guo.config}
	if err := _go.FromResponse(res); err != nil {
		return nil, err
	}
	return _go, nil
}

func (guo *GoodsUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if len(guo.fields) > 0 {
		fields := make([]any, 0, len(guo.fields)+1)
		fields = append(fields, true)
		for _, f := range guo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
