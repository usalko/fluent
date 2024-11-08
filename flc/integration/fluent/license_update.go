// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/fluent/license"
	"github.com/usalko/fluent/flc/integration/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// LicenseUpdate is the builder for updating License entities.
type LicenseUpdate struct {
	config
	hooks     []Hook
	mutation  *LicenseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LicenseUpdate builder.
func (lu *LicenseUpdate) Where(ps ...predicate.License) *LicenseUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUpdateTime sets the "update_time" field.
func (lu *LicenseUpdate) SetUpdateTime(t time.Time) *LicenseUpdate {
	lu.mutation.SetUpdateTime(t)
	return lu
}

// Mutation returns the LicenseMutation object of the builder.
func (lu *LicenseUpdate) Mutation() *LicenseMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LicenseUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LicenseUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LicenseUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LicenseUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LicenseUpdate) defaults() {
	if _, ok := lu.mutation.UpdateTime(); !ok {
		v := license.UpdateDefaultUpdateTime()
		lu.mutation.SetUpdateTime(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lu *LicenseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LicenseUpdate {
	lu.modifiers = append(lu.modifiers, modifiers...)
	return lu
}

func (lu *LicenseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(license.Table, license.Columns, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UpdateTime(); ok {
		_spec.SetField(license.FieldUpdateTime, field.TypeTime, value)
	}
	_spec.AddModifiers(lu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{license.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LicenseUpdateOne is the builder for updating a single License entity.
type LicenseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LicenseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (luo *LicenseUpdateOne) SetUpdateTime(t time.Time) *LicenseUpdateOne {
	luo.mutation.SetUpdateTime(t)
	return luo
}

// Mutation returns the LicenseMutation object of the builder.
func (luo *LicenseUpdateOne) Mutation() *LicenseMutation {
	return luo.mutation
}

// Where appends a list predicates to the LicenseUpdate builder.
func (luo *LicenseUpdateOne) Where(ps ...predicate.License) *LicenseUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LicenseUpdateOne) Select(field string, fields ...string) *LicenseUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated License entity.
func (luo *LicenseUpdateOne) Save(ctx context.Context) (*License, error) {
	luo.defaults()
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LicenseUpdateOne) SaveX(ctx context.Context) *License {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LicenseUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LicenseUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LicenseUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdateTime(); !ok {
		v := license.UpdateDefaultUpdateTime()
		luo.mutation.SetUpdateTime(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (luo *LicenseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LicenseUpdateOne {
	luo.modifiers = append(luo.modifiers, modifiers...)
	return luo
}

func (luo *LicenseUpdateOne) sqlSave(ctx context.Context) (_node *License, err error) {
	_spec := sqlgraph.NewUpdateSpec(license.Table, license.Columns, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "License.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, license.FieldID)
		for _, f := range fields {
			if !license.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != license.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.UpdateTime(); ok {
		_spec.SetField(license.FieldUpdateTime, field.TypeTime, value)
	}
	_spec.AddModifiers(luo.modifiers...)
	_node = &License{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{license.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}