// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/edge_field/fluent/node"
	"github.com/usalko/fluent/schema/field"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	mutation *NodeMutation
	hooks    []Hook
}

// SetValue sets the "value" field.
func (nc *NodeCreate) SetValue(i int) *NodeCreate {
	nc.mutation.SetValue(i)
	return nc
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (nc *NodeCreate) SetNillableValue(i *int) *NodeCreate {
	if i != nil {
		nc.SetValue(*i)
	}
	return nc
}

// SetPrevID sets the "prev_id" field.
func (nc *NodeCreate) SetPrevID(i int) *NodeCreate {
	nc.mutation.SetPrevID(i)
	return nc
}

// SetNillablePrevID sets the "prev_id" field if the given value is not nil.
func (nc *NodeCreate) SetNillablePrevID(i *int) *NodeCreate {
	if i != nil {
		nc.SetPrevID(*i)
	}
	return nc
}

// SetPrev sets the "prev" edge to the Node entity.
func (nc *NodeCreate) SetPrev(n *Node) *NodeCreate {
	return nc.SetPrevID(n.ID)
}

// SetNextID sets the "next" edge to the Node entity by ID.
func (nc *NodeCreate) SetNextID(id int) *NodeCreate {
	nc.mutation.SetNextID(id)
	return nc
}

// SetNillableNextID sets the "next" edge to the Node entity by ID if the given value is not nil.
func (nc *NodeCreate) SetNillableNextID(id *int) *NodeCreate {
	if id != nil {
		nc = nc.SetNextID(*id)
	}
	return nc
}

// SetNext sets the "next" edge to the Node entity.
func (nc *NodeCreate) SetNext(n *Node) *NodeCreate {
	return nc.SetNextID(n.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (nc *NodeCreate) Mutation() *NodeMutation {
	return nc.mutation
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	nc.defaults()
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NodeCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NodeCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NodeCreate) defaults() {
	if _, ok := nc.mutation.Value(); !ok {
		v := node.DefaultValue
		nc.mutation.SetValue(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NodeCreate) check() error {
	if _, ok := nc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Node.value"`)}
	}
	return nil
}

func (nc *NodeCreate) sqlSave(ctx context.Context) (*Node, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NodeCreate) createSpec() (*Node, *sqlgraph.CreateSpec) {
	var (
		_node = &Node{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(node.Table, sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt))
	)
	if value, ok := nc.mutation.Value(); ok {
		_spec.SetField(node.FieldValue, field.TypeInt, value)
		_node.Value = value
	}
	if nodes := nc.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   node.PrevTable,
			Columns: []string{node.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PrevID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   node.NextTable,
			Columns: []string{node.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NodeCreateBulk is the builder for creating many Node entities in bulk.
type NodeCreateBulk struct {
	config
	err      error
	builders []*NodeCreate
}

// Save creates the Node entities in the database.
func (ncb *NodeCreateBulk) Save(ctx context.Context) ([]*Node, error) {
	if ncb.err != nil {
		return nil, ncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Node, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NodeMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NodeCreateBulk) SaveX(ctx context.Context) []*Node {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NodeCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NodeCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
