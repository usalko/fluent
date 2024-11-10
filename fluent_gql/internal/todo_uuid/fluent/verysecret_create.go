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

	"github.com/google/uuid"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/fluent_gql/internal/todo_uuid/fluent/verysecret"
	"github.com/usalko/fluent/schema/field"
)

// VerySecretCreate is the builder for creating a VerySecret entity.
type VerySecretCreate struct {
	config
	mutation *VerySecretMutation
	hooks    []Hook
}

// SetPassword sets the "password" field.
func (vsc *VerySecretCreate) SetPassword(s string) *VerySecretCreate {
	vsc.mutation.SetPassword(s)
	return vsc
}

// SetID sets the "id" field.
func (vsc *VerySecretCreate) SetID(u uuid.UUID) *VerySecretCreate {
	vsc.mutation.SetID(u)
	return vsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vsc *VerySecretCreate) SetNillableID(u *uuid.UUID) *VerySecretCreate {
	if u != nil {
		vsc.SetID(*u)
	}
	return vsc
}

// Mutation returns the VerySecretMutation object of the builder.
func (vsc *VerySecretCreate) Mutation() *VerySecretMutation {
	return vsc.mutation
}

// Save creates the VerySecret in the database.
func (vsc *VerySecretCreate) Save(ctx context.Context) (*VerySecret, error) {
	vsc.defaults()
	return withHooks(ctx, vsc.sqlSave, vsc.mutation, vsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vsc *VerySecretCreate) SaveX(ctx context.Context) *VerySecret {
	v, err := vsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vsc *VerySecretCreate) Exec(ctx context.Context) error {
	_, err := vsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vsc *VerySecretCreate) ExecX(ctx context.Context) {
	if err := vsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vsc *VerySecretCreate) defaults() {
	if _, ok := vsc.mutation.ID(); !ok {
		v := verysecret.DefaultID()
		vsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vsc *VerySecretCreate) check() error {
	if _, ok := vsc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`fluent: missing required field "VerySecret.password"`)}
	}
	return nil
}

func (vsc *VerySecretCreate) sqlSave(ctx context.Context) (*VerySecret, error) {
	if err := vsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vsc.driver, _spec); err != nil {
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
	vsc.mutation.id = &_node.ID
	vsc.mutation.done = true
	return _node, nil
}

func (vsc *VerySecretCreate) createSpec() (*VerySecret, *sqlgraph.CreateSpec) {
	var (
		_node = &VerySecret{config: vsc.config}
		_spec = sqlgraph.NewCreateSpec(verysecret.Table, sqlgraph.NewFieldSpec(verysecret.FieldID, field.TypeUUID))
	)
	if id, ok := vsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := vsc.mutation.Password(); ok {
		_spec.SetField(verysecret.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	return _node, _spec
}

// VerySecretCreateBulk is the builder for creating many VerySecret entities in bulk.
type VerySecretCreateBulk struct {
	config
	err      error
	builders []*VerySecretCreate
}

// Save creates the VerySecret entities in the database.
func (vscb *VerySecretCreateBulk) Save(ctx context.Context) ([]*VerySecret, error) {
	if vscb.err != nil {
		return nil, vscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vscb.builders))
	nodes := make([]*VerySecret, len(vscb.builders))
	mutators := make([]Mutator, len(vscb.builders))
	for i := range vscb.builders {
		func(i int, root context.Context) {
			builder := vscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VerySecretMutation)
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
					_, err = mutators[i+1].Mutate(root, vscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vscb *VerySecretCreateBulk) SaveX(ctx context.Context) []*VerySecret {
	v, err := vscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vscb *VerySecretCreateBulk) Exec(ctx context.Context) error {
	_, err := vscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vscb *VerySecretCreateBulk) ExecX(ctx context.Context) {
	if err := vscb.Exec(ctx); err != nil {
		panic(err)
	}
}
