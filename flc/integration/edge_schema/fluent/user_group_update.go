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
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/group"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/user"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/user_group"
	"github.com/usalko/fluent/schema/field"
)

// UserGroupUpdate is the builder for updating UserGroup entities.
type UserGroupUpdate struct {
	config
	hooks    []Hook
	mutation *UserGroupMutation
}

// Where appends a list predicates to the UserGroupUpdate builder.
func (ugu *UserGroupUpdate) Where(ps ...predicate.UserGroup) *UserGroupUpdate {
	ugu.mutation.Where(ps...)
	return ugu
}

// SetJoinedAt sets the "joined_at" field.
func (ugu *UserGroupUpdate) SetJoinedAt(tt time.Time) *UserGroupUpdate {
	ugu.mutation.SetJoinedAt(tt)
	return ugu
}

// SetNillableJoinedAt sets the "joined_at" field if the given value is not nil.
func (ugu *UserGroupUpdate) SetNillableJoinedAt(tt *time.Time) *UserGroupUpdate {
	if tt != nil {
		ugu.SetJoinedAt(*tt)
	}
	return ugu
}

// SetUserID sets the "user_id" field.
func (ugu *UserGroupUpdate) SetUserID(i int) *UserGroupUpdate {
	ugu.mutation.SetUserID(i)
	return ugu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ugu *UserGroupUpdate) SetNillableUserID(i *int) *UserGroupUpdate {
	if i != nil {
		ugu.SetUserID(*i)
	}
	return ugu
}

// SetGroupID sets the "group_id" field.
func (ugu *UserGroupUpdate) SetGroupID(i int) *UserGroupUpdate {
	ugu.mutation.SetGroupID(i)
	return ugu
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (ugu *UserGroupUpdate) SetNillableGroupID(i *int) *UserGroupUpdate {
	if i != nil {
		ugu.SetGroupID(*i)
	}
	return ugu
}

// SetUser sets the "user" edge to the User entity.
func (ugu *UserGroupUpdate) SetUser(u *User) *UserGroupUpdate {
	return ugu.SetUserID(u.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (ugu *UserGroupUpdate) SetGroup(g *Group) *UserGroupUpdate {
	return ugu.SetGroupID(g.ID)
}

// Mutation returns the UserGroupMutation object of the builder.
func (ugu *UserGroupUpdate) Mutation() *UserGroupMutation {
	return ugu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ugu *UserGroupUpdate) ClearUser() *UserGroupUpdate {
	ugu.mutation.ClearUser()
	return ugu
}

// ClearGroup clears the "group" edge to the Group entity.
func (ugu *UserGroupUpdate) ClearGroup() *UserGroupUpdate {
	ugu.mutation.ClearGroup()
	return ugu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ugu *UserGroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ugu.sqlSave, ugu.mutation, ugu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ugu *UserGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := ugu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ugu *UserGroupUpdate) Exec(ctx context.Context) error {
	_, err := ugu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugu *UserGroupUpdate) ExecX(ctx context.Context) {
	if err := ugu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ugu *UserGroupUpdate) check() error {
	if ugu.mutation.UserCleared() && len(ugu.mutation.UserIDs()) > 0 {
		return errors.New(`fluent: clearing a required unique edge "UserGroup.user"`)
	}
	if ugu.mutation.GroupCleared() && len(ugu.mutation.GroupIDs()) > 0 {
		return errors.New(`fluent: clearing a required unique edge "UserGroup.group"`)
	}
	return nil
}

func (ugu *UserGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ugu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user_group.Table, user_group.Columns, sqlgraph.NewFieldSpec(user_group.FieldID, field.TypeInt))
	if ps := ugu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ugu.mutation.JoinedAt(); ok {
		_spec.SetField(user_group.FieldJoinedAt, field.TypeTime, value)
	}
	if ugu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user_group.UserTable,
			Columns: []string{user_group.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ugu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user_group.UserTable,
			Columns: []string{user_group.UserColumn},
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
	if ugu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user_group.GroupTable,
			Columns: []string{user_group.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ugu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user_group.GroupTable,
			Columns: []string{user_group.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ugu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user_group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ugu.mutation.done = true
	return n, nil
}

// UserGroupUpdateOne is the builder for updating a single UserGroup entity.
type UserGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserGroupMutation
}

// SetJoinedAt sets the "joined_at" field.
func (uguo *UserGroupUpdateOne) SetJoinedAt(tt time.Time) *UserGroupUpdateOne {
	uguo.mutation.SetJoinedAt(tt)
	return uguo
}

// SetNillableJoinedAt sets the "joined_at" field if the given value is not nil.
func (uguo *UserGroupUpdateOne) SetNillableJoinedAt(tt *time.Time) *UserGroupUpdateOne {
	if tt != nil {
		uguo.SetJoinedAt(*tt)
	}
	return uguo
}

// SetUserID sets the "user_id" field.
func (uguo *UserGroupUpdateOne) SetUserID(i int) *UserGroupUpdateOne {
	uguo.mutation.SetUserID(i)
	return uguo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uguo *UserGroupUpdateOne) SetNillableUserID(i *int) *UserGroupUpdateOne {
	if i != nil {
		uguo.SetUserID(*i)
	}
	return uguo
}

// SetGroupID sets the "group_id" field.
func (uguo *UserGroupUpdateOne) SetGroupID(i int) *UserGroupUpdateOne {
	uguo.mutation.SetGroupID(i)
	return uguo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (uguo *UserGroupUpdateOne) SetNillableGroupID(i *int) *UserGroupUpdateOne {
	if i != nil {
		uguo.SetGroupID(*i)
	}
	return uguo
}

// SetUser sets the "user" edge to the User entity.
func (uguo *UserGroupUpdateOne) SetUser(u *User) *UserGroupUpdateOne {
	return uguo.SetUserID(u.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (uguo *UserGroupUpdateOne) SetGroup(g *Group) *UserGroupUpdateOne {
	return uguo.SetGroupID(g.ID)
}

// Mutation returns the UserGroupMutation object of the builder.
func (uguo *UserGroupUpdateOne) Mutation() *UserGroupMutation {
	return uguo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uguo *UserGroupUpdateOne) ClearUser() *UserGroupUpdateOne {
	uguo.mutation.ClearUser()
	return uguo
}

// ClearGroup clears the "group" edge to the Group entity.
func (uguo *UserGroupUpdateOne) ClearGroup() *UserGroupUpdateOne {
	uguo.mutation.ClearGroup()
	return uguo
}

// Where appends a list predicates to the UserGroupUpdate builder.
func (uguo *UserGroupUpdateOne) Where(ps ...predicate.UserGroup) *UserGroupUpdateOne {
	uguo.mutation.Where(ps...)
	return uguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uguo *UserGroupUpdateOne) Select(field string, fields ...string) *UserGroupUpdateOne {
	uguo.fields = append([]string{field}, fields...)
	return uguo
}

// Save executes the query and returns the updated UserGroup entity.
func (uguo *UserGroupUpdateOne) Save(ctx context.Context) (*UserGroup, error) {
	return withHooks(ctx, uguo.sqlSave, uguo.mutation, uguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uguo *UserGroupUpdateOne) SaveX(ctx context.Context) *UserGroup {
	node, err := uguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uguo *UserGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := uguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uguo *UserGroupUpdateOne) ExecX(ctx context.Context) {
	if err := uguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uguo *UserGroupUpdateOne) check() error {
	if uguo.mutation.UserCleared() && len(uguo.mutation.UserIDs()) > 0 {
		return errors.New(`fluent: clearing a required unique edge "UserGroup.user"`)
	}
	if uguo.mutation.GroupCleared() && len(uguo.mutation.GroupIDs()) > 0 {
		return errors.New(`fluent: clearing a required unique edge "UserGroup.group"`)
	}
	return nil
}

func (uguo *UserGroupUpdateOne) sqlSave(ctx context.Context) (_node *UserGroup, err error) {
	if err := uguo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user_group.Table, user_group.Columns, sqlgraph.NewFieldSpec(user_group.FieldID, field.TypeInt))
	id, ok := uguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`fluent: missing "UserGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user_group.FieldID)
		for _, f := range fields {
			if !user_group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != user_group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uguo.mutation.JoinedAt(); ok {
		_spec.SetField(user_group.FieldJoinedAt, field.TypeTime, value)
	}
	if uguo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user_group.UserTable,
			Columns: []string{user_group.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uguo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user_group.UserTable,
			Columns: []string{user_group.UserColumn},
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
	if uguo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user_group.GroupTable,
			Columns: []string{user_group.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uguo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user_group.GroupTable,
			Columns: []string{user_group.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserGroup{config: uguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user_group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uguo.mutation.done = true
	return _node, nil
}
