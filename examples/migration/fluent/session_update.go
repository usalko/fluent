// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/examples/migration/fluent/predicate"
	"github.com/usalko/fluent/examples/migration/fluent/session"
	"github.com/usalko/fluent/examples/migration/fluent/session_device"
	"github.com/usalko/fluent/schema/field"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetActive sets the "active" field.
func (su *SessionUpdate) SetActive(b bool) *SessionUpdate {
	su.mutation.SetActive(b)
	return su
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (su *SessionUpdate) SetNillableActive(b *bool) *SessionUpdate {
	if b != nil {
		su.SetActive(*b)
	}
	return su
}

// SetIssuedAt sets the "issued_at" field.
func (su *SessionUpdate) SetIssuedAt(time_time time.Time) *SessionUpdate {
	su.mutation.SetIssuedAt(time_time)
	return su
}

// SetNillableIssuedAt sets the "issued_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableIssuedAt(time_time *time.Time) *SessionUpdate {
	if time_time != nil {
		su.SetIssuedAt(*time_time)
	}
	return su
}

// SetExpiresAt sets the "expires_at" field.
func (su *SessionUpdate) SetExpiresAt(time_time time.Time) *SessionUpdate {
	su.mutation.SetExpiresAt(time_time)
	return su
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableExpiresAt(time_time *time.Time) *SessionUpdate {
	if time_time != nil {
		su.SetExpiresAt(*time_time)
	}
	return su
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (su *SessionUpdate) ClearExpiresAt() *SessionUpdate {
	su.mutation.ClearExpiresAt()
	return su
}

// SetToken sets the "token" field.
func (su *SessionUpdate) SetToken(s string) *SessionUpdate {
	su.mutation.SetToken(s)
	return su
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (su *SessionUpdate) SetNillableToken(s *string) *SessionUpdate {
	if s != nil {
		su.SetToken(*s)
	}
	return su
}

// ClearToken clears the value of the "token" field.
func (su *SessionUpdate) ClearToken() *SessionUpdate {
	su.mutation.ClearToken()
	return su
}

// SetMethod sets the "method" field.
func (su *SessionUpdate) SetMethod(msi map[string]interface{}) *SessionUpdate {
	su.mutation.SetMethod(msi)
	return su
}

// ClearMethod clears the value of the "method" field.
func (su *SessionUpdate) ClearMethod() *SessionUpdate {
	su.mutation.ClearMethod()
	return su
}

// SetDeviceID sets the "device_id" field.
func (su *SessionUpdate) SetDeviceID(uuid_uuid uuid.UUID) *SessionUpdate {
	su.mutation.SetDeviceID(uuid_uuid)
	return su
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableDeviceID(uuid_uuid *uuid.UUID) *SessionUpdate {
	if uuid_uuid != nil {
		su.SetDeviceID(*uuid_uuid)
	}
	return su
}

// ClearDeviceID clears the value of the "device_id" field.
func (su *SessionUpdate) ClearDeviceID() *SessionUpdate {
	su.mutation.ClearDeviceID()
	return su
}

// SetDevice sets the "device" edge to the SessionDevice entity.
func (su *SessionUpdate) SetDevice(s *SessionDevice) *SessionUpdate {
	return su.SetDeviceID(s.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// ClearDevice clears the "device" edge to the SessionDevice entity.
func (su *SessionUpdate) ClearDevice() *SessionUpdate {
	su.mutation.ClearDevice()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Active(); ok {
		_spec.SetField(session.FieldActive, field.TypeBool, value)
	}
	if value, ok := su.mutation.IssuedAt(); ok {
		_spec.SetField(session.FieldIssuedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeTime, value)
	}
	if su.mutation.ExpiresAtCleared() {
		_spec.ClearField(session.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := su.mutation.Token(); ok {
		_spec.SetField(session.FieldToken, field.TypeString, value)
	}
	if su.mutation.TokenCleared() {
		_spec.ClearField(session.FieldToken, field.TypeString)
	}
	if value, ok := su.mutation.Method(); ok {
		_spec.SetField(session.FieldMethod, field.TypeJSON, value)
	}
	if su.mutation.MethodCleared() {
		_spec.ClearField(session.FieldMethod, field.TypeJSON)
	}
	if su.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session_device.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session_device.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetActive sets the "active" field.
func (suo *SessionUpdateOne) SetActive(b bool) *SessionUpdateOne {
	suo.mutation.SetActive(b)
	return suo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableActive(b *bool) *SessionUpdateOne {
	if b != nil {
		suo.SetActive(*b)
	}
	return suo
}

// SetIssuedAt sets the "issued_at" field.
func (suo *SessionUpdateOne) SetIssuedAt(time_time time.Time) *SessionUpdateOne {
	suo.mutation.SetIssuedAt(time_time)
	return suo
}

// SetNillableIssuedAt sets the "issued_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableIssuedAt(time_time *time.Time) *SessionUpdateOne {
	if time_time != nil {
		suo.SetIssuedAt(*time_time)
	}
	return suo
}

// SetExpiresAt sets the "expires_at" field.
func (suo *SessionUpdateOne) SetExpiresAt(time_time time.Time) *SessionUpdateOne {
	suo.mutation.SetExpiresAt(time_time)
	return suo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableExpiresAt(time_time *time.Time) *SessionUpdateOne {
	if time_time != nil {
		suo.SetExpiresAt(*time_time)
	}
	return suo
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (suo *SessionUpdateOne) ClearExpiresAt() *SessionUpdateOne {
	suo.mutation.ClearExpiresAt()
	return suo
}

// SetToken sets the "token" field.
func (suo *SessionUpdateOne) SetToken(s string) *SessionUpdateOne {
	suo.mutation.SetToken(s)
	return suo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableToken(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetToken(*s)
	}
	return suo
}

// ClearToken clears the value of the "token" field.
func (suo *SessionUpdateOne) ClearToken() *SessionUpdateOne {
	suo.mutation.ClearToken()
	return suo
}

// SetMethod sets the "method" field.
func (suo *SessionUpdateOne) SetMethod(msi map[string]interface{}) *SessionUpdateOne {
	suo.mutation.SetMethod(msi)
	return suo
}

// ClearMethod clears the value of the "method" field.
func (suo *SessionUpdateOne) ClearMethod() *SessionUpdateOne {
	suo.mutation.ClearMethod()
	return suo
}

// SetDeviceID sets the "device_id" field.
func (suo *SessionUpdateOne) SetDeviceID(uuid_uuid uuid.UUID) *SessionUpdateOne {
	suo.mutation.SetDeviceID(uuid_uuid)
	return suo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableDeviceID(uuid_uuid *uuid.UUID) *SessionUpdateOne {
	if uuid_uuid != nil {
		suo.SetDeviceID(*uuid_uuid)
	}
	return suo
}

// ClearDeviceID clears the value of the "device_id" field.
func (suo *SessionUpdateOne) ClearDeviceID() *SessionUpdateOne {
	suo.mutation.ClearDeviceID()
	return suo
}

// SetDevice sets the "device" edge to the SessionDevice entity.
func (suo *SessionUpdateOne) SetDevice(s *SessionDevice) *SessionUpdateOne {
	return suo.SetDeviceID(s.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// ClearDevice clears the "device" edge to the SessionDevice entity.
func (suo *SessionUpdateOne) ClearDevice() *SessionUpdateOne {
	suo.mutation.ClearDevice()
	return suo
}

// Where appends a list predicates to the SessionUpdate builder.
func (suo *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`fluent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Active(); ok {
		_spec.SetField(session.FieldActive, field.TypeBool, value)
	}
	if value, ok := suo.mutation.IssuedAt(); ok {
		_spec.SetField(session.FieldIssuedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeTime, value)
	}
	if suo.mutation.ExpiresAtCleared() {
		_spec.ClearField(session.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := suo.mutation.Token(); ok {
		_spec.SetField(session.FieldToken, field.TypeString, value)
	}
	if suo.mutation.TokenCleared() {
		_spec.ClearField(session.FieldToken, field.TypeString)
	}
	if value, ok := suo.mutation.Method(); ok {
		_spec.SetField(session.FieldMethod, field.TypeJSON, value)
	}
	if suo.mutation.MethodCleared() {
		_spec.ClearField(session.FieldMethod, field.TypeJSON)
	}
	if suo.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session_device.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session_device.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
