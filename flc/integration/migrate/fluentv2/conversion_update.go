// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluentv2

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/migrate/fluentv2/conversion"
	"github.com/usalko/fluent/flc/integration/migrate/fluentv2/predicate"
	"github.com/usalko/fluent/schema/field"
)

// ConversionUpdate is the builder for updating Conversion entities.
type ConversionUpdate struct {
	config
	hooks    []Hook
	mutation *ConversionMutation
}

// Where appends a list predicates to the ConversionUpdate builder.
func (cu *ConversionUpdate) Where(ps ...predicate.Conversion) *ConversionUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ConversionUpdate) SetName(s string) *ConversionUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *ConversionUpdate) SetNillableName(s *string) *ConversionUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of the "name" field.
func (cu *ConversionUpdate) ClearName() *ConversionUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetInt8ToString sets the "int8_to_string" field.
func (cu *ConversionUpdate) SetInt8ToString(s string) *ConversionUpdate {
	cu.mutation.SetInt8ToString(s)
	return cu
}

// SetNillableInt8ToString sets the "int8_to_string" field if the given value is not nil.
func (cu *ConversionUpdate) SetNillableInt8ToString(s *string) *ConversionUpdate {
	if s != nil {
		cu.SetInt8ToString(*s)
	}
	return cu
}

// ClearInt8ToString clears the value of the "int8_to_string" field.
func (cu *ConversionUpdate) ClearInt8ToString() *ConversionUpdate {
	cu.mutation.ClearInt8ToString()
	return cu
}

// SetUint8ToString sets the "uint8_to_string" field.
func (cu *ConversionUpdate) SetUint8ToString(s string) *ConversionUpdate {
	cu.mutation.SetUint8ToString(s)
	return cu
}

// SetNillableUint8ToString sets the "uint8_to_string" field if the given value is not nil.
func (cu *ConversionUpdate) SetNillableUint8ToString(s *string) *ConversionUpdate {
	if s != nil {
		cu.SetUint8ToString(*s)
	}
	return cu
}

// ClearUint8ToString clears the value of the "uint8_to_string" field.
func (cu *ConversionUpdate) ClearUint8ToString() *ConversionUpdate {
	cu.mutation.ClearUint8ToString()
	return cu
}

// SetInt16ToString sets the "int16_to_string" field.
func (cu *ConversionUpdate) SetInt16ToString(s string) *ConversionUpdate {
	cu.mutation.SetInt16ToString(s)
	return cu
}

// SetNillableInt16ToString sets the "int16_to_string" field if the given value is not nil.
func (cu *ConversionUpdate) SetNillableInt16ToString(s *string) *ConversionUpdate {
	if s != nil {
		cu.SetInt16ToString(*s)
	}
	return cu
}

// ClearInt16ToString clears the value of the "int16_to_string" field.
func (cu *ConversionUpdate) ClearInt16ToString() *ConversionUpdate {
	cu.mutation.ClearInt16ToString()
	return cu
}

// SetUint16ToString sets the "uint16_to_string" field.
func (cu *ConversionUpdate) SetUint16ToString(s string) *ConversionUpdate {
	cu.mutation.SetUint16ToString(s)
	return cu
}

// SetNillableUint16ToString sets the "uint16_to_string" field if the given value is not nil.
func (cu *ConversionUpdate) SetNillableUint16ToString(s *string) *ConversionUpdate {
	if s != nil {
		cu.SetUint16ToString(*s)
	}
	return cu
}

// ClearUint16ToString clears the value of the "uint16_to_string" field.
func (cu *ConversionUpdate) ClearUint16ToString() *ConversionUpdate {
	cu.mutation.ClearUint16ToString()
	return cu
}

// SetInt32ToString sets the "int32_to_string" field.
func (cu *ConversionUpdate) SetInt32ToString(s string) *ConversionUpdate {
	cu.mutation.SetInt32ToString(s)
	return cu
}

// SetNillableInt32ToString sets the "int32_to_string" field if the given value is not nil.
func (cu *ConversionUpdate) SetNillableInt32ToString(s *string) *ConversionUpdate {
	if s != nil {
		cu.SetInt32ToString(*s)
	}
	return cu
}

// ClearInt32ToString clears the value of the "int32_to_string" field.
func (cu *ConversionUpdate) ClearInt32ToString() *ConversionUpdate {
	cu.mutation.ClearInt32ToString()
	return cu
}

// SetUint32ToString sets the "uint32_to_string" field.
func (cu *ConversionUpdate) SetUint32ToString(s string) *ConversionUpdate {
	cu.mutation.SetUint32ToString(s)
	return cu
}

// SetNillableUint32ToString sets the "uint32_to_string" field if the given value is not nil.
func (cu *ConversionUpdate) SetNillableUint32ToString(s *string) *ConversionUpdate {
	if s != nil {
		cu.SetUint32ToString(*s)
	}
	return cu
}

// ClearUint32ToString clears the value of the "uint32_to_string" field.
func (cu *ConversionUpdate) ClearUint32ToString() *ConversionUpdate {
	cu.mutation.ClearUint32ToString()
	return cu
}

// SetInt64ToString sets the "int64_to_string" field.
func (cu *ConversionUpdate) SetInt64ToString(s string) *ConversionUpdate {
	cu.mutation.SetInt64ToString(s)
	return cu
}

// SetNillableInt64ToString sets the "int64_to_string" field if the given value is not nil.
func (cu *ConversionUpdate) SetNillableInt64ToString(s *string) *ConversionUpdate {
	if s != nil {
		cu.SetInt64ToString(*s)
	}
	return cu
}

// ClearInt64ToString clears the value of the "int64_to_string" field.
func (cu *ConversionUpdate) ClearInt64ToString() *ConversionUpdate {
	cu.mutation.ClearInt64ToString()
	return cu
}

// SetUint64ToString sets the "uint64_to_string" field.
func (cu *ConversionUpdate) SetUint64ToString(s string) *ConversionUpdate {
	cu.mutation.SetUint64ToString(s)
	return cu
}

// SetNillableUint64ToString sets the "uint64_to_string" field if the given value is not nil.
func (cu *ConversionUpdate) SetNillableUint64ToString(s *string) *ConversionUpdate {
	if s != nil {
		cu.SetUint64ToString(*s)
	}
	return cu
}

// ClearUint64ToString clears the value of the "uint64_to_string" field.
func (cu *ConversionUpdate) ClearUint64ToString() *ConversionUpdate {
	cu.mutation.ClearUint64ToString()
	return cu
}

// Mutation returns the ConversionMutation object of the builder.
func (cu *ConversionUpdate) Mutation() *ConversionMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConversionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConversionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConversionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConversionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ConversionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(conversion.Table, conversion.Columns, sqlgraph.NewFieldSpec(conversion.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(conversion.FieldName, field.TypeString, value)
	}
	if cu.mutation.NameCleared() {
		_spec.ClearField(conversion.FieldName, field.TypeString)
	}
	if value, ok := cu.mutation.Int8ToString(); ok {
		_spec.SetField(conversion.FieldInt8ToString, field.TypeString, value)
	}
	if cu.mutation.Int8ToStringCleared() {
		_spec.ClearField(conversion.FieldInt8ToString, field.TypeString)
	}
	if value, ok := cu.mutation.Uint8ToString(); ok {
		_spec.SetField(conversion.FieldUint8ToString, field.TypeString, value)
	}
	if cu.mutation.Uint8ToStringCleared() {
		_spec.ClearField(conversion.FieldUint8ToString, field.TypeString)
	}
	if value, ok := cu.mutation.Int16ToString(); ok {
		_spec.SetField(conversion.FieldInt16ToString, field.TypeString, value)
	}
	if cu.mutation.Int16ToStringCleared() {
		_spec.ClearField(conversion.FieldInt16ToString, field.TypeString)
	}
	if value, ok := cu.mutation.Uint16ToString(); ok {
		_spec.SetField(conversion.FieldUint16ToString, field.TypeString, value)
	}
	if cu.mutation.Uint16ToStringCleared() {
		_spec.ClearField(conversion.FieldUint16ToString, field.TypeString)
	}
	if value, ok := cu.mutation.Int32ToString(); ok {
		_spec.SetField(conversion.FieldInt32ToString, field.TypeString, value)
	}
	if cu.mutation.Int32ToStringCleared() {
		_spec.ClearField(conversion.FieldInt32ToString, field.TypeString)
	}
	if value, ok := cu.mutation.Uint32ToString(); ok {
		_spec.SetField(conversion.FieldUint32ToString, field.TypeString, value)
	}
	if cu.mutation.Uint32ToStringCleared() {
		_spec.ClearField(conversion.FieldUint32ToString, field.TypeString)
	}
	if value, ok := cu.mutation.Int64ToString(); ok {
		_spec.SetField(conversion.FieldInt64ToString, field.TypeString, value)
	}
	if cu.mutation.Int64ToStringCleared() {
		_spec.ClearField(conversion.FieldInt64ToString, field.TypeString)
	}
	if value, ok := cu.mutation.Uint64ToString(); ok {
		_spec.SetField(conversion.FieldUint64ToString, field.TypeString, value)
	}
	if cu.mutation.Uint64ToStringCleared() {
		_spec.ClearField(conversion.FieldUint64ToString, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConversionUpdateOne is the builder for updating a single Conversion entity.
type ConversionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConversionMutation
}

// SetName sets the "name" field.
func (cuo *ConversionUpdateOne) SetName(s string) *ConversionUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *ConversionUpdateOne) SetNillableName(s *string) *ConversionUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of the "name" field.
func (cuo *ConversionUpdateOne) ClearName() *ConversionUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetInt8ToString sets the "int8_to_string" field.
func (cuo *ConversionUpdateOne) SetInt8ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetInt8ToString(s)
	return cuo
}

// SetNillableInt8ToString sets the "int8_to_string" field if the given value is not nil.
func (cuo *ConversionUpdateOne) SetNillableInt8ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		cuo.SetInt8ToString(*s)
	}
	return cuo
}

// ClearInt8ToString clears the value of the "int8_to_string" field.
func (cuo *ConversionUpdateOne) ClearInt8ToString() *ConversionUpdateOne {
	cuo.mutation.ClearInt8ToString()
	return cuo
}

// SetUint8ToString sets the "uint8_to_string" field.
func (cuo *ConversionUpdateOne) SetUint8ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetUint8ToString(s)
	return cuo
}

// SetNillableUint8ToString sets the "uint8_to_string" field if the given value is not nil.
func (cuo *ConversionUpdateOne) SetNillableUint8ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		cuo.SetUint8ToString(*s)
	}
	return cuo
}

// ClearUint8ToString clears the value of the "uint8_to_string" field.
func (cuo *ConversionUpdateOne) ClearUint8ToString() *ConversionUpdateOne {
	cuo.mutation.ClearUint8ToString()
	return cuo
}

// SetInt16ToString sets the "int16_to_string" field.
func (cuo *ConversionUpdateOne) SetInt16ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetInt16ToString(s)
	return cuo
}

// SetNillableInt16ToString sets the "int16_to_string" field if the given value is not nil.
func (cuo *ConversionUpdateOne) SetNillableInt16ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		cuo.SetInt16ToString(*s)
	}
	return cuo
}

// ClearInt16ToString clears the value of the "int16_to_string" field.
func (cuo *ConversionUpdateOne) ClearInt16ToString() *ConversionUpdateOne {
	cuo.mutation.ClearInt16ToString()
	return cuo
}

// SetUint16ToString sets the "uint16_to_string" field.
func (cuo *ConversionUpdateOne) SetUint16ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetUint16ToString(s)
	return cuo
}

// SetNillableUint16ToString sets the "uint16_to_string" field if the given value is not nil.
func (cuo *ConversionUpdateOne) SetNillableUint16ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		cuo.SetUint16ToString(*s)
	}
	return cuo
}

// ClearUint16ToString clears the value of the "uint16_to_string" field.
func (cuo *ConversionUpdateOne) ClearUint16ToString() *ConversionUpdateOne {
	cuo.mutation.ClearUint16ToString()
	return cuo
}

// SetInt32ToString sets the "int32_to_string" field.
func (cuo *ConversionUpdateOne) SetInt32ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetInt32ToString(s)
	return cuo
}

// SetNillableInt32ToString sets the "int32_to_string" field if the given value is not nil.
func (cuo *ConversionUpdateOne) SetNillableInt32ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		cuo.SetInt32ToString(*s)
	}
	return cuo
}

// ClearInt32ToString clears the value of the "int32_to_string" field.
func (cuo *ConversionUpdateOne) ClearInt32ToString() *ConversionUpdateOne {
	cuo.mutation.ClearInt32ToString()
	return cuo
}

// SetUint32ToString sets the "uint32_to_string" field.
func (cuo *ConversionUpdateOne) SetUint32ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetUint32ToString(s)
	return cuo
}

// SetNillableUint32ToString sets the "uint32_to_string" field if the given value is not nil.
func (cuo *ConversionUpdateOne) SetNillableUint32ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		cuo.SetUint32ToString(*s)
	}
	return cuo
}

// ClearUint32ToString clears the value of the "uint32_to_string" field.
func (cuo *ConversionUpdateOne) ClearUint32ToString() *ConversionUpdateOne {
	cuo.mutation.ClearUint32ToString()
	return cuo
}

// SetInt64ToString sets the "int64_to_string" field.
func (cuo *ConversionUpdateOne) SetInt64ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetInt64ToString(s)
	return cuo
}

// SetNillableInt64ToString sets the "int64_to_string" field if the given value is not nil.
func (cuo *ConversionUpdateOne) SetNillableInt64ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		cuo.SetInt64ToString(*s)
	}
	return cuo
}

// ClearInt64ToString clears the value of the "int64_to_string" field.
func (cuo *ConversionUpdateOne) ClearInt64ToString() *ConversionUpdateOne {
	cuo.mutation.ClearInt64ToString()
	return cuo
}

// SetUint64ToString sets the "uint64_to_string" field.
func (cuo *ConversionUpdateOne) SetUint64ToString(s string) *ConversionUpdateOne {
	cuo.mutation.SetUint64ToString(s)
	return cuo
}

// SetNillableUint64ToString sets the "uint64_to_string" field if the given value is not nil.
func (cuo *ConversionUpdateOne) SetNillableUint64ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		cuo.SetUint64ToString(*s)
	}
	return cuo
}

// ClearUint64ToString clears the value of the "uint64_to_string" field.
func (cuo *ConversionUpdateOne) ClearUint64ToString() *ConversionUpdateOne {
	cuo.mutation.ClearUint64ToString()
	return cuo
}

// Mutation returns the ConversionMutation object of the builder.
func (cuo *ConversionUpdateOne) Mutation() *ConversionMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ConversionUpdate builder.
func (cuo *ConversionUpdateOne) Where(ps ...predicate.Conversion) *ConversionUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConversionUpdateOne) Select(field string, fields ...string) *ConversionUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Conversion entity.
func (cuo *ConversionUpdateOne) Save(ctx context.Context) (*Conversion, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConversionUpdateOne) SaveX(ctx context.Context) *Conversion {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConversionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConversionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ConversionUpdateOne) sqlSave(ctx context.Context) (_node *Conversion, err error) {
	_spec := sqlgraph.NewUpdateSpec(conversion.Table, conversion.Columns, sqlgraph.NewFieldSpec(conversion.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entv2: missing "Conversion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, conversion.FieldID)
		for _, f := range fields {
			if !conversion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entv2: invalid field %q for query", f)}
			}
			if f != conversion.FieldID {
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
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(conversion.FieldName, field.TypeString, value)
	}
	if cuo.mutation.NameCleared() {
		_spec.ClearField(conversion.FieldName, field.TypeString)
	}
	if value, ok := cuo.mutation.Int8ToString(); ok {
		_spec.SetField(conversion.FieldInt8ToString, field.TypeString, value)
	}
	if cuo.mutation.Int8ToStringCleared() {
		_spec.ClearField(conversion.FieldInt8ToString, field.TypeString)
	}
	if value, ok := cuo.mutation.Uint8ToString(); ok {
		_spec.SetField(conversion.FieldUint8ToString, field.TypeString, value)
	}
	if cuo.mutation.Uint8ToStringCleared() {
		_spec.ClearField(conversion.FieldUint8ToString, field.TypeString)
	}
	if value, ok := cuo.mutation.Int16ToString(); ok {
		_spec.SetField(conversion.FieldInt16ToString, field.TypeString, value)
	}
	if cuo.mutation.Int16ToStringCleared() {
		_spec.ClearField(conversion.FieldInt16ToString, field.TypeString)
	}
	if value, ok := cuo.mutation.Uint16ToString(); ok {
		_spec.SetField(conversion.FieldUint16ToString, field.TypeString, value)
	}
	if cuo.mutation.Uint16ToStringCleared() {
		_spec.ClearField(conversion.FieldUint16ToString, field.TypeString)
	}
	if value, ok := cuo.mutation.Int32ToString(); ok {
		_spec.SetField(conversion.FieldInt32ToString, field.TypeString, value)
	}
	if cuo.mutation.Int32ToStringCleared() {
		_spec.ClearField(conversion.FieldInt32ToString, field.TypeString)
	}
	if value, ok := cuo.mutation.Uint32ToString(); ok {
		_spec.SetField(conversion.FieldUint32ToString, field.TypeString, value)
	}
	if cuo.mutation.Uint32ToStringCleared() {
		_spec.ClearField(conversion.FieldUint32ToString, field.TypeString)
	}
	if value, ok := cuo.mutation.Int64ToString(); ok {
		_spec.SetField(conversion.FieldInt64ToString, field.TypeString, value)
	}
	if cuo.mutation.Int64ToStringCleared() {
		_spec.ClearField(conversion.FieldInt64ToString, field.TypeString)
	}
	if value, ok := cuo.mutation.Uint64ToString(); ok {
		_spec.SetField(conversion.FieldUint64ToString, field.TypeString, value)
	}
	if cuo.mutation.Uint64ToStringCleared() {
		_spec.ClearField(conversion.FieldUint64ToString, field.TypeString)
	}
	_node = &Conversion{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
