// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/examples/migration/fluent/pet"
	"github.com/usalko/fluent/examples/migration/fluent/user"
	"github.com/usalko/fluent/schema/field"
	"github.com/google/uuid"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PetCreate) SetName(s string) *PetCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetAge sets the "age" field.
func (pc *PetCreate) SetAge(f float64) *PetCreate {
	pc.mutation.SetAge(f)
	return pc
}

// SetWeight sets the "weight" field.
func (pc *PetCreate) SetWeight(f float64) *PetCreate {
	pc.mutation.SetWeight(f)
	return pc
}

// SetBestFriendID sets the "best_friend_id" field.
func (pc *PetCreate) SetBestFriendID(u uuid.UUID) *PetCreate {
	pc.mutation.SetBestFriendID(u)
	return pc
}

// SetOwnerID sets the "owner_id" field.
func (pc *PetCreate) SetOwnerID(i int) *PetCreate {
	pc.mutation.SetOwnerID(i)
	return pc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (pc *PetCreate) SetNillableOwnerID(i *int) *PetCreate {
	if i != nil {
		pc.SetOwnerID(*i)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PetCreate) SetID(u uuid.UUID) *PetCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PetCreate) SetNillableID(u *uuid.UUID) *PetCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetBestFriend sets the "best_friend" edge to the Pet entity.
func (pc *PetCreate) SetBestFriend(p *Pet) *PetCreate {
	return pc.SetBestFriendID(p.ID)
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PetCreate) SetOwner(u *User) *PetCreate {
	return pc.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (pc *PetCreate) Mutation() *PetMutation {
	return pc.mutation
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PetCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PetCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PetCreate) defaults() {
	if _, ok := pc.mutation.OwnerID(); !ok {
		v := pet.DefaultOwnerID
		pc.mutation.SetOwnerID(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := pet.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PetCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Pet.name"`)}
	}
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Pet.age"`)}
	}
	if _, ok := pc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Pet.weight"`)}
	}
	if _, ok := pc.mutation.BestFriendID(); !ok {
		return &ValidationError{Name: "best_friend_id", err: errors.New(`ent: missing required field "Pet.best_friend_id"`)}
	}
	if _, ok := pc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Pet.owner_id"`)}
	}
	if len(pc.mutation.BestFriendIDs()) == 0 {
		return &ValidationError{Name: "best_friend", err: errors.New(`ent: missing required edge "Pet.best_friend"`)}
	}
	if len(pc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Pet.owner"`)}
	}
	return nil
}

func (pc *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PetCreate) createSpec() (*Pet, *sqlgraph.CreateSpec) {
	var (
		_node = &Pet{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(pet.Table, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(pet.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.SetField(pet.FieldAge, field.TypeFloat64, value)
		_node.Age = value
	}
	if value, ok := pc.mutation.Weight(); ok {
		_spec.SetField(pet.FieldWeight, field.TypeFloat64, value)
		_node.Weight = value
	}
	if nodes := pc.mutation.BestFriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   pet.BestFriendTable,
			Columns: []string{pet.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BestFriendID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PetCreateBulk is the builder for creating many Pet entities in bulk.
type PetCreateBulk struct {
	config
	err      error
	builders []*PetCreate
}

// Save creates the Pet entities in the database.
func (pcb *PetCreateBulk) Save(ctx context.Context) ([]*Pet, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pet, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PetMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PetCreateBulk) SaveX(ctx context.Context) []*Pet {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PetCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PetCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
