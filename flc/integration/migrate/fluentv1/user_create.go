// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluentv1

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/migrate/fluentv1/car"
	"github.com/usalko/fluent/flc/integration/migrate/fluentv1/user"
	"github.com/usalko/fluent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetAge sets the "age" field.
func (uc *UserCreate) SetAge(i int32) *UserCreate {
	uc.mutation.SetAge(i)
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetDescription sets the "description" field.
func (uc *UserCreate) SetDescription(s string) *UserCreate {
	uc.mutation.SetDescription(s)
	return uc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uc *UserCreate) SetNillableDescription(s *string) *UserCreate {
	if s != nil {
		uc.SetDescription(*s)
	}
	return uc
}

// SetNickname sets the "nickname" field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetAddress sets the "address" field.
func (uc *UserCreate) SetAddress(s string) *UserCreate {
	uc.mutation.SetAddress(s)
	return uc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uc *UserCreate) SetNillableAddress(s *string) *UserCreate {
	if s != nil {
		uc.SetAddress(*s)
	}
	return uc
}

// SetRenamed sets the "renamed" field.
func (uc *UserCreate) SetRenamed(s string) *UserCreate {
	uc.mutation.SetRenamed(s)
	return uc
}

// SetNillableRenamed sets the "renamed" field if the given value is not nil.
func (uc *UserCreate) SetNillableRenamed(s *string) *UserCreate {
	if s != nil {
		uc.SetRenamed(*s)
	}
	return uc
}

// SetOldToken sets the "old_token" field.
func (uc *UserCreate) SetOldToken(s string) *UserCreate {
	uc.mutation.SetOldToken(s)
	return uc
}

// SetNillableOldToken sets the "old_token" field if the given value is not nil.
func (uc *UserCreate) SetNillableOldToken(s *string) *UserCreate {
	if s != nil {
		uc.SetOldToken(*s)
	}
	return uc
}

// SetBlob sets the "blob" field.
func (uc *UserCreate) SetBlob(b []byte) *UserCreate {
	uc.mutation.SetBlob(b)
	return uc
}

// SetState sets the "state" field.
func (uc *UserCreate) SetState(us user.State) *UserCreate {
	uc.mutation.SetState(us)
	return uc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (uc *UserCreate) SetNillableState(us *user.State) *UserCreate {
	if us != nil {
		uc.SetState(*us)
	}
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(s string) *UserCreate {
	uc.mutation.SetStatus(s)
	return uc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(s *string) *UserCreate {
	if s != nil {
		uc.SetStatus(*s)
	}
	return uc
}

// SetWorkplace sets the "workplace" field.
func (uc *UserCreate) SetWorkplace(s string) *UserCreate {
	uc.mutation.SetWorkplace(s)
	return uc
}

// SetNillableWorkplace sets the "workplace" field if the given value is not nil.
func (uc *UserCreate) SetNillableWorkplace(s *string) *UserCreate {
	if s != nil {
		uc.SetWorkplace(*s)
	}
	return uc
}

// SetDropOptional sets the "drop_optional" field.
func (uc *UserCreate) SetDropOptional(s string) *UserCreate {
	uc.mutation.SetDropOptional(s)
	return uc
}

// SetNillableDropOptional sets the "drop_optional" field if the given value is not nil.
func (uc *UserCreate) SetNillableDropOptional(s *string) *UserCreate {
	if s != nil {
		uc.SetDropOptional(*s)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// SetParentID sets the "parent" edge to the User entity by ID.
func (uc *UserCreate) SetParentID(id int) *UserCreate {
	uc.mutation.SetParentID(id)
	return uc
}

// SetNillableParentID sets the "parent" edge to the User entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableParentID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetParentID(*id)
	}
	return uc
}

// SetParent sets the "parent" edge to the User entity.
func (uc *UserCreate) SetParent(u *User) *UserCreate {
	return uc.SetParentID(u.ID)
}

// AddChildIDs adds the "children" edge to the User entity by IDs.
func (uc *UserCreate) AddChildIDs(ids ...int) *UserCreate {
	uc.mutation.AddChildIDs(ids...)
	return uc
}

// AddChildren adds the "children" edges to the User entity.
func (uc *UserCreate) AddChildren(u ...*User) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddChildIDs(ids...)
}

// SetSpouseID sets the "spouse" edge to the User entity by ID.
func (uc *UserCreate) SetSpouseID(id int) *UserCreate {
	uc.mutation.SetSpouseID(id)
	return uc
}

// SetNillableSpouseID sets the "spouse" edge to the User entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableSpouseID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetSpouseID(*id)
	}
	return uc
}

// SetSpouse sets the "spouse" edge to the User entity.
func (uc *UserCreate) SetSpouse(u *User) *UserCreate {
	return uc.SetSpouseID(u.ID)
}

// SetCarID sets the "car" edge to the Car entity by ID.
func (uc *UserCreate) SetCarID(id int) *UserCreate {
	uc.mutation.SetCarID(id)
	return uc
}

// SetNillableCarID sets the "car" edge to the Car entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableCarID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetCarID(*id)
	}
	return uc
}

// SetCar sets the "car" edge to the Car entity.
func (uc *UserCreate) SetCar(c *Car) *UserCreate {
	return uc.SetCarID(c.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.OldToken(); !ok {
		v := user.DefaultOldToken()
		uc.mutation.SetOldToken(v)
	}
	if _, ok := uc.mutation.State(); !ok {
		v := user.DefaultState
		uc.mutation.SetState(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`fluentv1: missing required field "User.age"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`fluentv1: missing required field "User.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`fluentv1: validator failed for field "User.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`fluentv1: missing required field "User.nickname"`)}
	}
	if _, ok := uc.mutation.OldToken(); !ok {
		return &ValidationError{Name: "old_token", err: errors.New(`fluentv1: missing required field "User.old_token"`)}
	}
	if v, ok := uc.mutation.Blob(); ok {
		if err := user.BlobValidator(v); err != nil {
			return &ValidationError{Name: "blob", err: fmt.Errorf(`fluentv1: validator failed for field "User.blob": %w`, err)}
		}
	}
	if v, ok := uc.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`fluentv1: validator failed for field "User.state": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Workplace(); ok {
		if err := user.WorkplaceValidator(v); err != nil {
			return &ValidationError{Name: "workplace", err: fmt.Errorf(`fluentv1: validator failed for field "User.workplace": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt32, value)
		_node.Age = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := uc.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := uc.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := uc.mutation.Renamed(); ok {
		_spec.SetField(user.FieldRenamed, field.TypeString, value)
		_node.Renamed = value
	}
	if value, ok := uc.mutation.OldToken(); ok {
		_spec.SetField(user.FieldOldToken, field.TypeString, value)
		_node.OldToken = value
	}
	if value, ok := uc.mutation.Blob(); ok {
		_spec.SetField(user.FieldBlob, field.TypeBytes, value)
		_node.Blob = value
	}
	if value, ok := uc.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeEnum, value)
		_node.State = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := uc.mutation.Workplace(); ok {
		_spec.SetField(user.FieldWorkplace, field.TypeString, value)
		_node.Workplace = value
	}
	if value, ok := uc.mutation.DropOptional(); ok {
		_spec.SetField(user.FieldDropOptional, field.TypeString, value)
		_node.DropOptional = value
	}
	if nodes := uc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_spouse = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
