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
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/car"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/pet"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	hooks    []Hook
	mutation *CarMutation
}

// Where appends a list predicates to the CarUpdate builder.
func (cu *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetBeforeID sets the "before_id" field.
func (cu *CarUpdate) SetBeforeID(f float64) *CarUpdate {
	cu.mutation.ResetBeforeID()
	cu.mutation.SetBeforeID(f)
	return cu
}

// SetNillableBeforeID sets the "before_id" field if the given value is not nil.
func (cu *CarUpdate) SetNillableBeforeID(f *float64) *CarUpdate {
	if f != nil {
		cu.SetBeforeID(*f)
	}
	return cu
}

// AddBeforeID adds f to the "before_id" field.
func (cu *CarUpdate) AddBeforeID(f float64) *CarUpdate {
	cu.mutation.AddBeforeID(f)
	return cu
}

// ClearBeforeID clears the value of the "before_id" field.
func (cu *CarUpdate) ClearBeforeID() *CarUpdate {
	cu.mutation.ClearBeforeID()
	return cu
}

// SetAfterID sets the "after_id" field.
func (cu *CarUpdate) SetAfterID(f float64) *CarUpdate {
	cu.mutation.ResetAfterID()
	cu.mutation.SetAfterID(f)
	return cu
}

// SetNillableAfterID sets the "after_id" field if the given value is not nil.
func (cu *CarUpdate) SetNillableAfterID(f *float64) *CarUpdate {
	if f != nil {
		cu.SetAfterID(*f)
	}
	return cu
}

// AddAfterID adds f to the "after_id" field.
func (cu *CarUpdate) AddAfterID(f float64) *CarUpdate {
	cu.mutation.AddAfterID(f)
	return cu
}

// ClearAfterID clears the value of the "after_id" field.
func (cu *CarUpdate) ClearAfterID() *CarUpdate {
	cu.mutation.ClearAfterID()
	return cu
}

// SetModel sets the "model" field.
func (cu *CarUpdate) SetModel(s string) *CarUpdate {
	cu.mutation.SetModel(s)
	return cu
}

// SetNillableModel sets the "model" field if the given value is not nil.
func (cu *CarUpdate) SetNillableModel(s *string) *CarUpdate {
	if s != nil {
		cu.SetModel(*s)
	}
	return cu
}

// SetOwnerID sets the "owner" edge to the Pet entity by ID.
func (cu *CarUpdate) SetOwnerID(id string) *CarUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the "owner" edge to the Pet entity by ID if the given value is not nil.
func (cu *CarUpdate) SetNillableOwnerID(id *string) *CarUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the "owner" edge to the Pet entity.
func (cu *CarUpdate) SetOwner(p *Pet) *CarUpdate {
	return cu.SetOwnerID(p.ID)
}

// Mutation returns the CarMutation object of the builder.
func (cu *CarUpdate) Mutation() *CarMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the Pet entity.
func (cu *CarUpdate) ClearOwner() *CarUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CarUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CarUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CarUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CarUpdate) check() error {
	if v, ok := cu.mutation.BeforeID(); ok {
		if err := car.BeforeIDValidator(v); err != nil {
			return &ValidationError{Name: "before_id", err: fmt.Errorf(`fluent: validator failed for field "Car.before_id": %w`, err)}
		}
	}
	if v, ok := cu.mutation.AfterID(); ok {
		if err := car.AfterIDValidator(v); err != nil {
			return &ValidationError{Name: "after_id", err: fmt.Errorf(`fluent: validator failed for field "Car.after_id": %w`, err)}
		}
	}
	return nil
}

func (cu *CarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(car.Table, car.Columns, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.BeforeID(); ok {
		_spec.SetField(car.FieldBeforeID, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedBeforeID(); ok {
		_spec.AddField(car.FieldBeforeID, field.TypeFloat64, value)
	}
	if cu.mutation.BeforeIDCleared() {
		_spec.ClearField(car.FieldBeforeID, field.TypeFloat64)
	}
	if value, ok := cu.mutation.AfterID(); ok {
		_spec.SetField(car.FieldAfterID, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedAfterID(); ok {
		_spec.AddField(car.FieldAfterID, field.TypeFloat64, value)
	}
	if cu.mutation.AfterIDCleared() {
		_spec.ClearField(car.FieldAfterID, field.TypeFloat64)
	}
	if value, ok := cu.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CarMutation
}

// SetBeforeID sets the "before_id" field.
func (cuo *CarUpdateOne) SetBeforeID(f float64) *CarUpdateOne {
	cuo.mutation.ResetBeforeID()
	cuo.mutation.SetBeforeID(f)
	return cuo
}

// SetNillableBeforeID sets the "before_id" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableBeforeID(f *float64) *CarUpdateOne {
	if f != nil {
		cuo.SetBeforeID(*f)
	}
	return cuo
}

// AddBeforeID adds f to the "before_id" field.
func (cuo *CarUpdateOne) AddBeforeID(f float64) *CarUpdateOne {
	cuo.mutation.AddBeforeID(f)
	return cuo
}

// ClearBeforeID clears the value of the "before_id" field.
func (cuo *CarUpdateOne) ClearBeforeID() *CarUpdateOne {
	cuo.mutation.ClearBeforeID()
	return cuo
}

// SetAfterID sets the "after_id" field.
func (cuo *CarUpdateOne) SetAfterID(f float64) *CarUpdateOne {
	cuo.mutation.ResetAfterID()
	cuo.mutation.SetAfterID(f)
	return cuo
}

// SetNillableAfterID sets the "after_id" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableAfterID(f *float64) *CarUpdateOne {
	if f != nil {
		cuo.SetAfterID(*f)
	}
	return cuo
}

// AddAfterID adds f to the "after_id" field.
func (cuo *CarUpdateOne) AddAfterID(f float64) *CarUpdateOne {
	cuo.mutation.AddAfterID(f)
	return cuo
}

// ClearAfterID clears the value of the "after_id" field.
func (cuo *CarUpdateOne) ClearAfterID() *CarUpdateOne {
	cuo.mutation.ClearAfterID()
	return cuo
}

// SetModel sets the "model" field.
func (cuo *CarUpdateOne) SetModel(s string) *CarUpdateOne {
	cuo.mutation.SetModel(s)
	return cuo
}

// SetNillableModel sets the "model" field if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableModel(s *string) *CarUpdateOne {
	if s != nil {
		cuo.SetModel(*s)
	}
	return cuo
}

// SetOwnerID sets the "owner" edge to the Pet entity by ID.
func (cuo *CarUpdateOne) SetOwnerID(id string) *CarUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the "owner" edge to the Pet entity by ID if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableOwnerID(id *string) *CarUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the "owner" edge to the Pet entity.
func (cuo *CarUpdateOne) SetOwner(p *Pet) *CarUpdateOne {
	return cuo.SetOwnerID(p.ID)
}

// Mutation returns the CarMutation object of the builder.
func (cuo *CarUpdateOne) Mutation() *CarMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the Pet entity.
func (cuo *CarUpdateOne) ClearOwner() *CarUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// Where appends a list predicates to the CarUpdate builder.
func (cuo *CarUpdateOne) Where(ps ...predicate.Car) *CarUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CarUpdateOne) Select(field string, fields ...string) *CarUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Car entity.
func (cuo *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CarUpdateOne) SaveX(ctx context.Context) *Car {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CarUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CarUpdateOne) check() error {
	if v, ok := cuo.mutation.BeforeID(); ok {
		if err := car.BeforeIDValidator(v); err != nil {
			return &ValidationError{Name: "before_id", err: fmt.Errorf(`fluent: validator failed for field "Car.before_id": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.AfterID(); ok {
		if err := car.AfterIDValidator(v); err != nil {
			return &ValidationError{Name: "after_id", err: fmt.Errorf(`fluent: validator failed for field "Car.after_id": %w`, err)}
		}
	}
	return nil
}

func (cuo *CarUpdateOne) sqlSave(ctx context.Context) (_node *Car, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(car.Table, car.Columns, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`fluent: missing "Car.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, car.FieldID)
		for _, f := range fields {
			if !car.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != car.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.BeforeID(); ok {
		_spec.SetField(car.FieldBeforeID, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedBeforeID(); ok {
		_spec.AddField(car.FieldBeforeID, field.TypeFloat64, value)
	}
	if cuo.mutation.BeforeIDCleared() {
		_spec.ClearField(car.FieldBeforeID, field.TypeFloat64)
	}
	if value, ok := cuo.mutation.AfterID(); ok {
		_spec.SetField(car.FieldAfterID, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedAfterID(); ok {
		_spec.AddField(car.FieldAfterID, field.TypeFloat64, value)
	}
	if cuo.mutation.AfterIDCleared() {
		_spec.ClearField(car.FieldAfterID, field.TypeFloat64)
	}
	if value, ok := cuo.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Car{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
