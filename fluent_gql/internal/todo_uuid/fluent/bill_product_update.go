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

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/fluent_gql/internal/todo_uuid/fluent/bill_product"
	"github.com/usalko/fluent/fluent_gql/internal/todo_uuid/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// BillProductUpdate is the builder for updating BillProduct entities.
type BillProductUpdate struct {
	config
	hooks    []Hook
	mutation *BillProductMutation
}

// Where appends a list predicates to the BillProductUpdate builder.
func (bpu *BillProductUpdate) Where(ps ...predicate.BillProduct) *BillProductUpdate {
	bpu.mutation.Where(ps...)
	return bpu
}

// SetName sets the "name" field.
func (bpu *BillProductUpdate) SetName(s string) *BillProductUpdate {
	bpu.mutation.SetName(s)
	return bpu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bpu *BillProductUpdate) SetNillableName(s *string) *BillProductUpdate {
	if s != nil {
		bpu.SetName(*s)
	}
	return bpu
}

// SetSku sets the "sku" field.
func (bpu *BillProductUpdate) SetSku(s string) *BillProductUpdate {
	bpu.mutation.SetSku(s)
	return bpu
}

// SetNillableSku sets the "sku" field if the given value is not nil.
func (bpu *BillProductUpdate) SetNillableSku(s *string) *BillProductUpdate {
	if s != nil {
		bpu.SetSku(*s)
	}
	return bpu
}

// SetQuantity sets the "quantity" field.
func (bpu *BillProductUpdate) SetQuantity(u uint64) *BillProductUpdate {
	bpu.mutation.ResetQuantity()
	bpu.mutation.SetQuantity(u)
	return bpu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (bpu *BillProductUpdate) SetNillableQuantity(u *uint64) *BillProductUpdate {
	if u != nil {
		bpu.SetQuantity(*u)
	}
	return bpu
}

// AddQuantity adds u to the "quantity" field.
func (bpu *BillProductUpdate) AddQuantity(u int64) *BillProductUpdate {
	bpu.mutation.AddQuantity(u)
	return bpu
}

// Mutation returns the BillProductMutation object of the builder.
func (bpu *BillProductUpdate) Mutation() *BillProductMutation {
	return bpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bpu *BillProductUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bpu.sqlSave, bpu.mutation, bpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bpu *BillProductUpdate) SaveX(ctx context.Context) int {
	affected, err := bpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bpu *BillProductUpdate) Exec(ctx context.Context) error {
	_, err := bpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bpu *BillProductUpdate) ExecX(ctx context.Context) {
	if err := bpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bpu *BillProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bill_product.Table, bill_product.Columns, sqlgraph.NewFieldSpec(bill_product.FieldID, field.TypeUUID))
	if ps := bpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bpu.mutation.Name(); ok {
		_spec.SetField(bill_product.FieldName, field.TypeString, value)
	}
	if value, ok := bpu.mutation.Sku(); ok {
		_spec.SetField(bill_product.FieldSku, field.TypeString, value)
	}
	if value, ok := bpu.mutation.Quantity(); ok {
		_spec.SetField(bill_product.FieldQuantity, field.TypeUint64, value)
	}
	if value, ok := bpu.mutation.AddedQuantity(); ok {
		_spec.AddField(bill_product.FieldQuantity, field.TypeUint64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bill_product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bpu.mutation.done = true
	return n, nil
}

// BillProductUpdateOne is the builder for updating a single BillProduct entity.
type BillProductUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BillProductMutation
}

// SetName sets the "name" field.
func (bpuo *BillProductUpdateOne) SetName(s string) *BillProductUpdateOne {
	bpuo.mutation.SetName(s)
	return bpuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bpuo *BillProductUpdateOne) SetNillableName(s *string) *BillProductUpdateOne {
	if s != nil {
		bpuo.SetName(*s)
	}
	return bpuo
}

// SetSku sets the "sku" field.
func (bpuo *BillProductUpdateOne) SetSku(s string) *BillProductUpdateOne {
	bpuo.mutation.SetSku(s)
	return bpuo
}

// SetNillableSku sets the "sku" field if the given value is not nil.
func (bpuo *BillProductUpdateOne) SetNillableSku(s *string) *BillProductUpdateOne {
	if s != nil {
		bpuo.SetSku(*s)
	}
	return bpuo
}

// SetQuantity sets the "quantity" field.
func (bpuo *BillProductUpdateOne) SetQuantity(u uint64) *BillProductUpdateOne {
	bpuo.mutation.ResetQuantity()
	bpuo.mutation.SetQuantity(u)
	return bpuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (bpuo *BillProductUpdateOne) SetNillableQuantity(u *uint64) *BillProductUpdateOne {
	if u != nil {
		bpuo.SetQuantity(*u)
	}
	return bpuo
}

// AddQuantity adds u to the "quantity" field.
func (bpuo *BillProductUpdateOne) AddQuantity(u int64) *BillProductUpdateOne {
	bpuo.mutation.AddQuantity(u)
	return bpuo
}

// Mutation returns the BillProductMutation object of the builder.
func (bpuo *BillProductUpdateOne) Mutation() *BillProductMutation {
	return bpuo.mutation
}

// Where appends a list predicates to the BillProductUpdate builder.
func (bpuo *BillProductUpdateOne) Where(ps ...predicate.BillProduct) *BillProductUpdateOne {
	bpuo.mutation.Where(ps...)
	return bpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bpuo *BillProductUpdateOne) Select(field string, fields ...string) *BillProductUpdateOne {
	bpuo.fields = append([]string{field}, fields...)
	return bpuo
}

// Save executes the query and returns the updated BillProduct entity.
func (bpuo *BillProductUpdateOne) Save(ctx context.Context) (*BillProduct, error) {
	return withHooks(ctx, bpuo.sqlSave, bpuo.mutation, bpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bpuo *BillProductUpdateOne) SaveX(ctx context.Context) *BillProduct {
	node, err := bpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bpuo *BillProductUpdateOne) Exec(ctx context.Context) error {
	_, err := bpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bpuo *BillProductUpdateOne) ExecX(ctx context.Context) {
	if err := bpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bpuo *BillProductUpdateOne) sqlSave(ctx context.Context) (_node *BillProduct, err error) {
	_spec := sqlgraph.NewUpdateSpec(bill_product.Table, bill_product.Columns, sqlgraph.NewFieldSpec(bill_product.FieldID, field.TypeUUID))
	id, ok := bpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`fluent: missing "BillProduct.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bill_product.FieldID)
		for _, f := range fields {
			if !bill_product.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != bill_product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bpuo.mutation.Name(); ok {
		_spec.SetField(bill_product.FieldName, field.TypeString, value)
	}
	if value, ok := bpuo.mutation.Sku(); ok {
		_spec.SetField(bill_product.FieldSku, field.TypeString, value)
	}
	if value, ok := bpuo.mutation.Quantity(); ok {
		_spec.SetField(bill_product.FieldQuantity, field.TypeUint64, value)
	}
	if value, ok := bpuo.mutation.AddedQuantity(); ok {
		_spec.AddField(bill_product.FieldQuantity, field.TypeUint64, value)
	}
	_node = &BillProduct{config: bpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bill_product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bpuo.mutation.done = true
	return _node, nil
}
