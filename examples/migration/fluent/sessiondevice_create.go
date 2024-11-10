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
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/examples/migration/fluent/session"
	"github.com/usalko/fluent/examples/migration/fluent/sessiondevice"
	"github.com/usalko/fluent/schema/field"
)

// SessionDeviceCreate is the builder for creating a SessionDevice entity.
type SessionDeviceCreate struct {
	config
	mutation *SessionDeviceMutation
	hooks    []Hook
}

// SetIPAddress sets the "ip_address" field.
func (sdc *SessionDeviceCreate) SetIPAddress(s string) *SessionDeviceCreate {
	sdc.mutation.SetIPAddress(s)
	return sdc
}

// SetUserAgent sets the "user_agent" field.
func (sdc *SessionDeviceCreate) SetUserAgent(s string) *SessionDeviceCreate {
	sdc.mutation.SetUserAgent(s)
	return sdc
}

// SetLocation sets the "location" field.
func (sdc *SessionDeviceCreate) SetLocation(s string) *SessionDeviceCreate {
	sdc.mutation.SetLocation(s)
	return sdc
}

// SetCreatedAt sets the "created_at" field.
func (sdc *SessionDeviceCreate) SetCreatedAt(t time.Time) *SessionDeviceCreate {
	sdc.mutation.SetCreatedAt(t)
	return sdc
}

// SetUpdatedAt sets the "updated_at" field.
func (sdc *SessionDeviceCreate) SetUpdatedAt(t time.Time) *SessionDeviceCreate {
	sdc.mutation.SetUpdatedAt(t)
	return sdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sdc *SessionDeviceCreate) SetNillableUpdatedAt(t *time.Time) *SessionDeviceCreate {
	if t != nil {
		sdc.SetUpdatedAt(*t)
	}
	return sdc
}

// SetID sets the "id" field.
func (sdc *SessionDeviceCreate) SetID(u uuid.UUID) *SessionDeviceCreate {
	sdc.mutation.SetID(u)
	return sdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sdc *SessionDeviceCreate) SetNillableID(u *uuid.UUID) *SessionDeviceCreate {
	if u != nil {
		sdc.SetID(*u)
	}
	return sdc
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (sdc *SessionDeviceCreate) AddSessionIDs(ids ...uuid.UUID) *SessionDeviceCreate {
	sdc.mutation.AddSessionIDs(ids...)
	return sdc
}

// AddSessions adds the "sessions" edges to the Session entity.
func (sdc *SessionDeviceCreate) AddSessions(s ...*Session) *SessionDeviceCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdc.AddSessionIDs(ids...)
}

// Mutation returns the SessionDeviceMutation object of the builder.
func (sdc *SessionDeviceCreate) Mutation() *SessionDeviceMutation {
	return sdc.mutation
}

// Save creates the SessionDevice in the database.
func (sdc *SessionDeviceCreate) Save(ctx context.Context) (*SessionDevice, error) {
	sdc.defaults()
	return withHooks(ctx, sdc.sqlSave, sdc.mutation, sdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *SessionDeviceCreate) SaveX(ctx context.Context) *SessionDevice {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdc *SessionDeviceCreate) Exec(ctx context.Context) error {
	_, err := sdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdc *SessionDeviceCreate) ExecX(ctx context.Context) {
	if err := sdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdc *SessionDeviceCreate) defaults() {
	if _, ok := sdc.mutation.ID(); !ok {
		v := sessiondevice.DefaultID()
		sdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdc *SessionDeviceCreate) check() error {
	if _, ok := sdc.mutation.IPAddress(); !ok {
		return &ValidationError{Name: "ip_address", err: errors.New(`fluent: missing required field "SessionDevice.ip_address"`)}
	}
	if v, ok := sdc.mutation.IPAddress(); ok {
		if err := sessiondevice.IPAddressValidator(v); err != nil {
			return &ValidationError{Name: "ip_address", err: fmt.Errorf(`fluent: validator failed for field "SessionDevice.ip_address": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.UserAgent(); !ok {
		return &ValidationError{Name: "user_agent", err: errors.New(`fluent: missing required field "SessionDevice.user_agent"`)}
	}
	if v, ok := sdc.mutation.UserAgent(); ok {
		if err := sessiondevice.UserAgentValidator(v); err != nil {
			return &ValidationError{Name: "user_agent", err: fmt.Errorf(`fluent: validator failed for field "SessionDevice.user_agent": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`fluent: missing required field "SessionDevice.location"`)}
	}
	if v, ok := sdc.mutation.Location(); ok {
		if err := sessiondevice.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`fluent: validator failed for field "SessionDevice.location": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`fluent: missing required field "SessionDevice.created_at"`)}
	}
	return nil
}

func (sdc *SessionDeviceCreate) sqlSave(ctx context.Context) (*SessionDevice, error) {
	if err := sdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
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
	sdc.mutation.id = &_node.ID
	sdc.mutation.done = true
	return _node, nil
}

func (sdc *SessionDeviceCreate) createSpec() (*SessionDevice, *sqlgraph.CreateSpec) {
	var (
		_node = &SessionDevice{config: sdc.config}
		_spec = sqlgraph.NewCreateSpec(sessiondevice.Table, sqlgraph.NewFieldSpec(sessiondevice.FieldID, field.TypeUUID))
	)
	if id, ok := sdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sdc.mutation.IPAddress(); ok {
		_spec.SetField(sessiondevice.FieldIPAddress, field.TypeString, value)
		_node.IPAddress = value
	}
	if value, ok := sdc.mutation.UserAgent(); ok {
		_spec.SetField(sessiondevice.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = value
	}
	if value, ok := sdc.mutation.Location(); ok {
		_spec.SetField(sessiondevice.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := sdc.mutation.CreatedAt(); ok {
		_spec.SetField(sessiondevice.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sdc.mutation.UpdatedAt(); ok {
		_spec.SetField(sessiondevice.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := sdc.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sessiondevice.SessionsTable,
			Columns: []string{sessiondevice.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SessionDeviceCreateBulk is the builder for creating many SessionDevice entities in bulk.
type SessionDeviceCreateBulk struct {
	config
	err      error
	builders []*SessionDeviceCreate
}

// Save creates the SessionDevice entities in the database.
func (sdcb *SessionDeviceCreateBulk) Save(ctx context.Context) ([]*SessionDevice, error) {
	if sdcb.err != nil {
		return nil, sdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*SessionDevice, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionDeviceMutation)
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
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *SessionDeviceCreateBulk) SaveX(ctx context.Context) []*SessionDevice {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdcb *SessionDeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := sdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdcb *SessionDeviceCreateBulk) ExecX(ctx context.Context) {
	if err := sdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
