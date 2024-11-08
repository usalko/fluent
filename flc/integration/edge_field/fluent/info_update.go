// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/dialect/sql/sqljson"
	"github.com/usalko/fluent/flc/integration/edge_field/fluent/info"
	"github.com/usalko/fluent/flc/integration/edge_field/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/edge_field/fluent/user"
	"github.com/usalko/fluent/schema/field"
)

// InfoUpdate is the builder for updating Info entities.
type InfoUpdate struct {
	config
	hooks    []Hook
	mutation *InfoMutation
}

// Where appends a list predicates to the InfoUpdate builder.
func (iu *InfoUpdate) Where(ps ...predicate.Info) *InfoUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetContent sets the "content" field.
func (iu *InfoUpdate) SetContent(jm json.RawMessage) *InfoUpdate {
	iu.mutation.SetContent(jm)
	return iu
}

// AppendContent appends jm to the "content" field.
func (iu *InfoUpdate) AppendContent(jm json.RawMessage) *InfoUpdate {
	iu.mutation.AppendContent(jm)
	return iu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (iu *InfoUpdate) SetUserID(id int) *InfoUpdate {
	iu.mutation.SetUserID(id)
	return iu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (iu *InfoUpdate) SetNillableUserID(id *int) *InfoUpdate {
	if id != nil {
		iu = iu.SetUserID(*id)
	}
	return iu
}

// SetUser sets the "user" edge to the User entity.
func (iu *InfoUpdate) SetUser(u *User) *InfoUpdate {
	return iu.SetUserID(u.ID)
}

// Mutation returns the InfoMutation object of the builder.
func (iu *InfoUpdate) Mutation() *InfoMutation {
	return iu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (iu *InfoUpdate) ClearUser() *InfoUpdate {
	iu.mutation.ClearUser()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InfoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InfoUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InfoUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InfoUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *InfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(info.Table, info.Columns, sqlgraph.NewFieldSpec(info.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Content(); ok {
		_spec.SetField(info.FieldContent, field.TypeJSON, value)
	}
	if value, ok := iu.mutation.AppendedContent(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, info.FieldContent, value)
		})
	}
	if iu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   info.UserTable,
			Columns: []string{info.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   info.UserTable,
			Columns: []string{info.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{info.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// InfoUpdateOne is the builder for updating a single Info entity.
type InfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InfoMutation
}

// SetContent sets the "content" field.
func (iuo *InfoUpdateOne) SetContent(jm json.RawMessage) *InfoUpdateOne {
	iuo.mutation.SetContent(jm)
	return iuo
}

// AppendContent appends jm to the "content" field.
func (iuo *InfoUpdateOne) AppendContent(jm json.RawMessage) *InfoUpdateOne {
	iuo.mutation.AppendContent(jm)
	return iuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (iuo *InfoUpdateOne) SetUserID(id int) *InfoUpdateOne {
	iuo.mutation.SetUserID(id)
	return iuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (iuo *InfoUpdateOne) SetNillableUserID(id *int) *InfoUpdateOne {
	if id != nil {
		iuo = iuo.SetUserID(*id)
	}
	return iuo
}

// SetUser sets the "user" edge to the User entity.
func (iuo *InfoUpdateOne) SetUser(u *User) *InfoUpdateOne {
	return iuo.SetUserID(u.ID)
}

// Mutation returns the InfoMutation object of the builder.
func (iuo *InfoUpdateOne) Mutation() *InfoMutation {
	return iuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (iuo *InfoUpdateOne) ClearUser() *InfoUpdateOne {
	iuo.mutation.ClearUser()
	return iuo
}

// Where appends a list predicates to the InfoUpdate builder.
func (iuo *InfoUpdateOne) Where(ps ...predicate.Info) *InfoUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InfoUpdateOne) Select(field string, fields ...string) *InfoUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Info entity.
func (iuo *InfoUpdateOne) Save(ctx context.Context) (*Info, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InfoUpdateOne) SaveX(ctx context.Context) *Info {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InfoUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InfoUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *InfoUpdateOne) sqlSave(ctx context.Context) (_node *Info, err error) {
	_spec := sqlgraph.NewUpdateSpec(info.Table, info.Columns, sqlgraph.NewFieldSpec(info.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Info.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, info.FieldID)
		for _, f := range fields {
			if !info.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != info.FieldID {
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
	if value, ok := iuo.mutation.Content(); ok {
		_spec.SetField(info.FieldContent, field.TypeJSON, value)
	}
	if value, ok := iuo.mutation.AppendedContent(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, info.FieldContent, value)
		})
	}
	if iuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   info.UserTable,
			Columns: []string{info.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   info.UserTable,
			Columns: []string{info.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Info{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{info.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}