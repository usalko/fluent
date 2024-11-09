// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by flc, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/fluent_gql/internal/todouuid/fluent/billproduct"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/schema/field"
	"github.com/google/uuid"
)

// BillProductCreate is the builder for creating a BillProduct entity.
type BillProductCreate struct {
	config
	mutation *BillProductMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bpc *BillProductCreate) SetName(s string) *BillProductCreate {
	bpc.mutation.SetName(s)
	return bpc
}

// SetSku sets the "sku" field.
func (bpc *BillProductCreate) SetSku(s string) *BillProductCreate {
	bpc.mutation.SetSku(s)
	return bpc
}

// SetQuantity sets the "quantity" field.
func (bpc *BillProductCreate) SetQuantity(u uint64) *BillProductCreate {
	bpc.mutation.SetQuantity(u)
	return bpc
}

// SetID sets the "id" field.
func (bpc *BillProductCreate) SetID(u uuid.UUID) *BillProductCreate {
	bpc.mutation.SetID(u)
	return bpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bpc *BillProductCreate) SetNillableID(u *uuid.UUID) *BillProductCreate {
	if u != nil {
		bpc.SetID(*u)
	}
	return bpc
}

// Mutation returns the BillProductMutation object of the builder.
func (bpc *BillProductCreate) Mutation() *BillProductMutation {
	return bpc.mutation
}

// Save creates the BillProduct in the database.
func (bpc *BillProductCreate) Save(ctx context.Context) (*BillProduct, error) {
	bpc.defaults()
	return withHooks(ctx, bpc.sqlSave, bpc.mutation, bpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bpc *BillProductCreate) SaveX(ctx context.Context) *BillProduct {
	v, err := bpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bpc *BillProductCreate) Exec(ctx context.Context) error {
	_, err := bpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bpc *BillProductCreate) ExecX(ctx context.Context) {
	if err := bpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bpc *BillProductCreate) defaults() {
	if _, ok := bpc.mutation.ID(); !ok {
		v := billproduct.DefaultID()
		bpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bpc *BillProductCreate) check() error {
	if _, ok := bpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "BillProduct.name"`)}
	}
	if _, ok := bpc.mutation.Sku(); !ok {
		return &ValidationError{Name: "sku", err: errors.New(`ent: missing required field "BillProduct.sku"`)}
	}
	if _, ok := bpc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "BillProduct.quantity"`)}
	}
	return nil
}

func (bpc *BillProductCreate) sqlSave(ctx context.Context) (*BillProduct, error) {
	if err := bpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	bpc.mutation.id = &_node.ID
	bpc.mutation.done = true
	return _node, nil
}

func (bpc *BillProductCreate) createSpec() (*BillProduct, *sqlgraph.CreateSpec) {
	var (
		_node = &BillProduct{config: bpc.config}
		_spec = sqlgraph.NewCreateSpec(billproduct.Table, sqlgraph.NewFieldSpec(billproduct.FieldID, field.TypeUUID))
	)
	if id, ok := bpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bpc.mutation.Name(); ok {
		_spec.SetField(billproduct.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := bpc.mutation.Sku(); ok {
		_spec.SetField(billproduct.FieldSku, field.TypeString, value)
		_node.Sku = value
	}
	if value, ok := bpc.mutation.Quantity(); ok {
		_spec.SetField(billproduct.FieldQuantity, field.TypeUint64, value)
		_node.Quantity = value
	}
	return _node, _spec
}

// BillProductCreateBulk is the builder for creating many BillProduct entities in bulk.
type BillProductCreateBulk struct {
	config
	err      error
	builders []*BillProductCreate
}

// Save creates the BillProduct entities in the database.
func (bpcb *BillProductCreateBulk) Save(ctx context.Context) ([]*BillProduct, error) {
	if bpcb.err != nil {
		return nil, bpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bpcb.builders))
	nodes := make([]*BillProduct, len(bpcb.builders))
	mutators := make([]Mutator, len(bpcb.builders))
	for i := range bpcb.builders {
		func(i int, root context.Context) {
			builder := bpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BillProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bpcb *BillProductCreateBulk) SaveX(ctx context.Context) []*BillProduct {
	v, err := bpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bpcb *BillProductCreateBulk) Exec(ctx context.Context) error {
	_, err := bpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bpcb *BillProductCreateBulk) ExecX(ctx context.Context) {
	if err := bpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
