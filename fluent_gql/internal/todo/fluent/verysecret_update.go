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

	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/predicate"
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/verysecret"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/schema/field"
)

// VerySecretUpdate is the builder for updating VerySecret entities.
type VerySecretUpdate struct {
	config
	hooks     []Hook
	mutation  *VerySecretMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the VerySecretUpdate builder.
func (vsu *VerySecretUpdate) Where(ps ...predicate.VerySecret) *VerySecretUpdate {
	vsu.mutation.Where(ps...)
	return vsu
}

// SetPassword sets the "password" field.
func (vsu *VerySecretUpdate) SetPassword(s string) *VerySecretUpdate {
	vsu.mutation.SetPassword(s)
	return vsu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (vsu *VerySecretUpdate) SetNillablePassword(s *string) *VerySecretUpdate {
	if s != nil {
		vsu.SetPassword(*s)
	}
	return vsu
}

// Mutation returns the VerySecretMutation object of the builder.
func (vsu *VerySecretUpdate) Mutation() *VerySecretMutation {
	return vsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vsu *VerySecretUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, vsu.sqlSave, vsu.mutation, vsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vsu *VerySecretUpdate) SaveX(ctx context.Context) int {
	affected, err := vsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vsu *VerySecretUpdate) Exec(ctx context.Context) error {
	_, err := vsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vsu *VerySecretUpdate) ExecX(ctx context.Context) {
	if err := vsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (vsu *VerySecretUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *VerySecretUpdate {
	vsu.modifiers = append(vsu.modifiers, modifiers...)
	return vsu
}

func (vsu *VerySecretUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(verysecret.Table, verysecret.Columns, sqlgraph.NewFieldSpec(verysecret.FieldID, field.TypeInt))
	if ps := vsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vsu.mutation.Password(); ok {
		_spec.SetField(verysecret.FieldPassword, field.TypeString, value)
	}
	_spec.AddModifiers(vsu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, vsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{verysecret.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vsu.mutation.done = true
	return n, nil
}

// VerySecretUpdateOne is the builder for updating a single VerySecret entity.
type VerySecretUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *VerySecretMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetPassword sets the "password" field.
func (vsuo *VerySecretUpdateOne) SetPassword(s string) *VerySecretUpdateOne {
	vsuo.mutation.SetPassword(s)
	return vsuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (vsuo *VerySecretUpdateOne) SetNillablePassword(s *string) *VerySecretUpdateOne {
	if s != nil {
		vsuo.SetPassword(*s)
	}
	return vsuo
}

// Mutation returns the VerySecretMutation object of the builder.
func (vsuo *VerySecretUpdateOne) Mutation() *VerySecretMutation {
	return vsuo.mutation
}

// Where appends a list predicates to the VerySecretUpdate builder.
func (vsuo *VerySecretUpdateOne) Where(ps ...predicate.VerySecret) *VerySecretUpdateOne {
	vsuo.mutation.Where(ps...)
	return vsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vsuo *VerySecretUpdateOne) Select(field string, fields ...string) *VerySecretUpdateOne {
	vsuo.fields = append([]string{field}, fields...)
	return vsuo
}

// Save executes the query and returns the updated VerySecret entity.
func (vsuo *VerySecretUpdateOne) Save(ctx context.Context) (*VerySecret, error) {
	return withHooks(ctx, vsuo.sqlSave, vsuo.mutation, vsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vsuo *VerySecretUpdateOne) SaveX(ctx context.Context) *VerySecret {
	node, err := vsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vsuo *VerySecretUpdateOne) Exec(ctx context.Context) error {
	_, err := vsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vsuo *VerySecretUpdateOne) ExecX(ctx context.Context) {
	if err := vsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (vsuo *VerySecretUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *VerySecretUpdateOne {
	vsuo.modifiers = append(vsuo.modifiers, modifiers...)
	return vsuo
}

func (vsuo *VerySecretUpdateOne) sqlSave(ctx context.Context) (_node *VerySecret, err error) {
	_spec := sqlgraph.NewUpdateSpec(verysecret.Table, verysecret.Columns, sqlgraph.NewFieldSpec(verysecret.FieldID, field.TypeInt))
	id, ok := vsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VerySecret.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, verysecret.FieldID)
		for _, f := range fields {
			if !verysecret.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != verysecret.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vsuo.mutation.Password(); ok {
		_spec.SetField(verysecret.FieldPassword, field.TypeString, value)
	}
	_spec.AddModifiers(vsuo.modifiers...)
	_node = &VerySecret{config: vsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{verysecret.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vsuo.mutation.done = true
	return _node, nil
}
