// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/examples/triggers/fluent/predicate"
	"github.com/usalko/fluent/examples/triggers/fluent/user_audit_log"
	"github.com/usalko/fluent/schema/field"
)

// UserAuditLogUpdate is the builder for updating UserAuditLog entities.
type UserAuditLogUpdate struct {
	config
	hooks    []Hook
	mutation *UserAuditLogMutation
}

// Where appends a list predicates to the UserAuditLogUpdate builder.
func (ualu *UserAuditLogUpdate) Where(ps ...predicate.UserAuditLog) *UserAuditLogUpdate {
	ualu.mutation.Where(ps...)
	return ualu
}

// SetOperationType sets the "operation_type" field.
func (ualu *UserAuditLogUpdate) SetOperationType(s string) *UserAuditLogUpdate {
	ualu.mutation.SetOperationType(s)
	return ualu
}

// SetNillableOperationType sets the "operation_type" field if the given value is not nil.
func (ualu *UserAuditLogUpdate) SetNillableOperationType(s *string) *UserAuditLogUpdate {
	if s != nil {
		ualu.SetOperationType(*s)
	}
	return ualu
}

// SetOperationTime sets the "operation_time" field.
func (ualu *UserAuditLogUpdate) SetOperationTime(s string) *UserAuditLogUpdate {
	ualu.mutation.SetOperationTime(s)
	return ualu
}

// SetNillableOperationTime sets the "operation_time" field if the given value is not nil.
func (ualu *UserAuditLogUpdate) SetNillableOperationTime(s *string) *UserAuditLogUpdate {
	if s != nil {
		ualu.SetOperationTime(*s)
	}
	return ualu
}

// SetOldValue sets the "old_value" field.
func (ualu *UserAuditLogUpdate) SetOldValue(s string) *UserAuditLogUpdate {
	ualu.mutation.SetOldValue(s)
	return ualu
}

// SetNillableOldValue sets the "old_value" field if the given value is not nil.
func (ualu *UserAuditLogUpdate) SetNillableOldValue(s *string) *UserAuditLogUpdate {
	if s != nil {
		ualu.SetOldValue(*s)
	}
	return ualu
}

// ClearOldValue clears the value of the "old_value" field.
func (ualu *UserAuditLogUpdate) ClearOldValue() *UserAuditLogUpdate {
	ualu.mutation.ClearOldValue()
	return ualu
}

// SetNewValue sets the "new_value" field.
func (ualu *UserAuditLogUpdate) SetNewValue(s string) *UserAuditLogUpdate {
	ualu.mutation.SetNewValue(s)
	return ualu
}

// SetNillableNewValue sets the "new_value" field if the given value is not nil.
func (ualu *UserAuditLogUpdate) SetNillableNewValue(s *string) *UserAuditLogUpdate {
	if s != nil {
		ualu.SetNewValue(*s)
	}
	return ualu
}

// ClearNewValue clears the value of the "new_value" field.
func (ualu *UserAuditLogUpdate) ClearNewValue() *UserAuditLogUpdate {
	ualu.mutation.ClearNewValue()
	return ualu
}

// Mutation returns the UserAuditLogMutation object of the builder.
func (ualu *UserAuditLogUpdate) Mutation() *UserAuditLogMutation {
	return ualu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ualu *UserAuditLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ualu.sqlSave, ualu.mutation, ualu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ualu *UserAuditLogUpdate) SaveX(ctx context.Context) int {
	affected, err := ualu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ualu *UserAuditLogUpdate) Exec(ctx context.Context) error {
	_, err := ualu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ualu *UserAuditLogUpdate) ExecX(ctx context.Context) {
	if err := ualu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ualu *UserAuditLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user_audit_log.Table, user_audit_log.Columns, sqlgraph.NewFieldSpec(user_audit_log.FieldID, field.TypeInt))
	if ps := ualu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ualu.mutation.OperationType(); ok {
		_spec.SetField(user_audit_log.FieldOperationType, field.TypeString, value)
	}
	if value, ok := ualu.mutation.OperationTime(); ok {
		_spec.SetField(user_audit_log.FieldOperationTime, field.TypeString, value)
	}
	if value, ok := ualu.mutation.OldValue(); ok {
		_spec.SetField(user_audit_log.FieldOldValue, field.TypeString, value)
	}
	if ualu.mutation.OldValueCleared() {
		_spec.ClearField(user_audit_log.FieldOldValue, field.TypeString)
	}
	if value, ok := ualu.mutation.NewValue(); ok {
		_spec.SetField(user_audit_log.FieldNewValue, field.TypeString, value)
	}
	if ualu.mutation.NewValueCleared() {
		_spec.ClearField(user_audit_log.FieldNewValue, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ualu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user_audit_log.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ualu.mutation.done = true
	return n, nil
}

// UserAuditLogUpdateOne is the builder for updating a single UserAuditLog entity.
type UserAuditLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserAuditLogMutation
}

// SetOperationType sets the "operation_type" field.
func (ualuo *UserAuditLogUpdateOne) SetOperationType(s string) *UserAuditLogUpdateOne {
	ualuo.mutation.SetOperationType(s)
	return ualuo
}

// SetNillableOperationType sets the "operation_type" field if the given value is not nil.
func (ualuo *UserAuditLogUpdateOne) SetNillableOperationType(s *string) *UserAuditLogUpdateOne {
	if s != nil {
		ualuo.SetOperationType(*s)
	}
	return ualuo
}

// SetOperationTime sets the "operation_time" field.
func (ualuo *UserAuditLogUpdateOne) SetOperationTime(s string) *UserAuditLogUpdateOne {
	ualuo.mutation.SetOperationTime(s)
	return ualuo
}

// SetNillableOperationTime sets the "operation_time" field if the given value is not nil.
func (ualuo *UserAuditLogUpdateOne) SetNillableOperationTime(s *string) *UserAuditLogUpdateOne {
	if s != nil {
		ualuo.SetOperationTime(*s)
	}
	return ualuo
}

// SetOldValue sets the "old_value" field.
func (ualuo *UserAuditLogUpdateOne) SetOldValue(s string) *UserAuditLogUpdateOne {
	ualuo.mutation.SetOldValue(s)
	return ualuo
}

// SetNillableOldValue sets the "old_value" field if the given value is not nil.
func (ualuo *UserAuditLogUpdateOne) SetNillableOldValue(s *string) *UserAuditLogUpdateOne {
	if s != nil {
		ualuo.SetOldValue(*s)
	}
	return ualuo
}

// ClearOldValue clears the value of the "old_value" field.
func (ualuo *UserAuditLogUpdateOne) ClearOldValue() *UserAuditLogUpdateOne {
	ualuo.mutation.ClearOldValue()
	return ualuo
}

// SetNewValue sets the "new_value" field.
func (ualuo *UserAuditLogUpdateOne) SetNewValue(s string) *UserAuditLogUpdateOne {
	ualuo.mutation.SetNewValue(s)
	return ualuo
}

// SetNillableNewValue sets the "new_value" field if the given value is not nil.
func (ualuo *UserAuditLogUpdateOne) SetNillableNewValue(s *string) *UserAuditLogUpdateOne {
	if s != nil {
		ualuo.SetNewValue(*s)
	}
	return ualuo
}

// ClearNewValue clears the value of the "new_value" field.
func (ualuo *UserAuditLogUpdateOne) ClearNewValue() *UserAuditLogUpdateOne {
	ualuo.mutation.ClearNewValue()
	return ualuo
}

// Mutation returns the UserAuditLogMutation object of the builder.
func (ualuo *UserAuditLogUpdateOne) Mutation() *UserAuditLogMutation {
	return ualuo.mutation
}

// Where appends a list predicates to the UserAuditLogUpdate builder.
func (ualuo *UserAuditLogUpdateOne) Where(ps ...predicate.UserAuditLog) *UserAuditLogUpdateOne {
	ualuo.mutation.Where(ps...)
	return ualuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ualuo *UserAuditLogUpdateOne) Select(field string, fields ...string) *UserAuditLogUpdateOne {
	ualuo.fields = append([]string{field}, fields...)
	return ualuo
}

// Save executes the query and returns the updated UserAuditLog entity.
func (ualuo *UserAuditLogUpdateOne) Save(ctx context.Context) (*UserAuditLog, error) {
	return withHooks(ctx, ualuo.sqlSave, ualuo.mutation, ualuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ualuo *UserAuditLogUpdateOne) SaveX(ctx context.Context) *UserAuditLog {
	node, err := ualuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ualuo *UserAuditLogUpdateOne) Exec(ctx context.Context) error {
	_, err := ualuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ualuo *UserAuditLogUpdateOne) ExecX(ctx context.Context) {
	if err := ualuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ualuo *UserAuditLogUpdateOne) sqlSave(ctx context.Context) (_node *UserAuditLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(user_audit_log.Table, user_audit_log.Columns, sqlgraph.NewFieldSpec(user_audit_log.FieldID, field.TypeInt))
	id, ok := ualuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`fluent: missing "UserAuditLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ualuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user_audit_log.FieldID)
		for _, f := range fields {
			if !user_audit_log.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != user_audit_log.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ualuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ualuo.mutation.OperationType(); ok {
		_spec.SetField(user_audit_log.FieldOperationType, field.TypeString, value)
	}
	if value, ok := ualuo.mutation.OperationTime(); ok {
		_spec.SetField(user_audit_log.FieldOperationTime, field.TypeString, value)
	}
	if value, ok := ualuo.mutation.OldValue(); ok {
		_spec.SetField(user_audit_log.FieldOldValue, field.TypeString, value)
	}
	if ualuo.mutation.OldValueCleared() {
		_spec.ClearField(user_audit_log.FieldOldValue, field.TypeString)
	}
	if value, ok := ualuo.mutation.NewValue(); ok {
		_spec.SetField(user_audit_log.FieldNewValue, field.TypeString, value)
	}
	if ualuo.mutation.NewValueCleared() {
		_spec.ClearField(user_audit_log.FieldNewValue, field.TypeString)
	}
	_node = &UserAuditLog{config: ualuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ualuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user_audit_log.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ualuo.mutation.done = true
	return _node, nil
}
