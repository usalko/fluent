// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/fluent/item"
	"github.com/usalko/fluent/flc/integration/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks     []Hook
	mutation  *ItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetText sets the "text" field.
func (iu *ItemUpdate) SetText(s string) *ItemUpdate {
	iu.mutation.SetText(s)
	return iu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableText(s *string) *ItemUpdate {
	if s != nil {
		iu.SetText(*s)
	}
	return iu
}

// ClearText clears the value of the "text" field.
func (iu *ItemUpdate) ClearText() *ItemUpdate {
	iu.mutation.ClearText()
	return iu
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ItemUpdate) check() error {
	if v, ok := iu.mutation.Text(); ok {
		if err := item.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Item.text": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iu *ItemUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ItemUpdate {
	iu.modifiers = append(iu.modifiers, modifiers...)
	return iu
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeString))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Text(); ok {
		_spec.SetField(item.FieldText, field.TypeString, value)
	}
	if iu.mutation.TextCleared() {
		_spec.ClearField(item.FieldText, field.TypeString)
	}
	_spec.AddModifiers(iu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetText sets the "text" field.
func (iuo *ItemUpdateOne) SetText(s string) *ItemUpdateOne {
	iuo.mutation.SetText(s)
	return iuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableText(s *string) *ItemUpdateOne {
	if s != nil {
		iuo.SetText(*s)
	}
	return iuo
}

// ClearText clears the value of the "text" field.
func (iuo *ItemUpdateOne) ClearText() *ItemUpdateOne {
	iuo.mutation.ClearText()
	return iuo
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iuo *ItemUpdateOne) Where(ps ...predicate.Item) *ItemUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ItemUpdateOne) check() error {
	if v, ok := iuo.mutation.Text(); ok {
		if err := item.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Item.text": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iuo *ItemUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ItemUpdateOne {
	iuo.modifiers = append(iuo.modifiers, modifiers...)
	return iuo
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeString))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Text(); ok {
		_spec.SetField(item.FieldText, field.TypeString, value)
	}
	if iuo.mutation.TextCleared() {
		_spec.ClearField(item.FieldText, field.TypeString)
	}
	_spec.AddModifiers(iuo.modifiers...)
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
