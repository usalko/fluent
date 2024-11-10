// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/examples/migration/fluent/card"
	"github.com/usalko/fluent/examples/migration/fluent/payment"
	"github.com/usalko/fluent/examples/migration/fluent/user"
	"github.com/usalko/fluent/schema/field"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	mutation *CardMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (cc *CardCreate) SetType(s string) *CardCreate {
	cc.mutation.SetType(s)
	return cc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cc *CardCreate) SetNillableType(s *string) *CardCreate {
	if s != nil {
		cc.SetType(*s)
	}
	return cc
}

// SetNumberHash sets the "number_hash" field.
func (cc *CardCreate) SetNumberHash(s string) *CardCreate {
	cc.mutation.SetNumberHash(s)
	return cc
}

// SetCvvHash sets the "cvv_hash" field.
func (cc *CardCreate) SetCvvHash(s string) *CardCreate {
	cc.mutation.SetCvvHash(s)
	return cc
}

// SetExpiresAt sets the "expires_at" field.
func (cc *CardCreate) SetExpiresAt(time_time time.Time) *CardCreate {
	cc.mutation.SetExpiresAt(time_time)
	return cc
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (cc *CardCreate) SetNillableExpiresAt(time_time *time.Time) *CardCreate {
	if time_time != nil {
		cc.SetExpiresAt(*time_time)
	}
	return cc
}

// SetOwnerID sets the "owner_id" field.
func (cc *CardCreate) SetOwnerID(i int) *CardCreate {
	cc.mutation.SetOwnerID(i)
	return cc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (cc *CardCreate) SetNillableOwnerID(i *int) *CardCreate {
	if i != nil {
		cc.SetOwnerID(*i)
	}
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CardCreate) SetOwner(u *User) *CardCreate {
	return cc.SetOwnerID(u.ID)
}

// AddPaymentIDs adds the "payments" edge to the Payment entity by IDs.
func (cc *CardCreate) AddPaymentIDs(ids ...int) *CardCreate {
	cc.mutation.AddPaymentIDs(ids...)
	return cc
}

// AddPayments adds the "payments" edges to the Payment entity.
func (cc *CardCreate) AddPayments(p ...*Payment) *CardCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPaymentIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (cc *CardCreate) Mutation() *CardMutation {
	return cc.mutation
}

// Save creates the Card in the database.
func (cc *CardCreate) Save(ctx context.Context) (*Card, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CardCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CardCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CardCreate) defaults() {
	if _, ok := cc.mutation.GetType(); !ok {
		v := card.DefaultType
		cc.mutation.SetType(v)
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		v := card.DefaultOwnerID
		cc.mutation.SetOwnerID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CardCreate) check() error {
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`fluent: missing required field "Card.type"`)}
	}
	if _, ok := cc.mutation.NumberHash(); !ok {
		return &ValidationError{Name: "number_hash", err: errors.New(`fluent: missing required field "Card.number_hash"`)}
	}
	if _, ok := cc.mutation.CvvHash(); !ok {
		return &ValidationError{Name: "cvv_hash", err: errors.New(`fluent: missing required field "Card.cvv_hash"`)}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`fluent: missing required field "Card.owner_id"`)}
	}
	if len(cc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`fluent: missing required edge "Card.owner"`)}
	}
	return nil
}

func (cc *CardCreate) sqlSave(ctx context.Context) (*Card, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CardCreate) createSpec() (*Card, *sqlgraph.CreateSpec) {
	var (
		_node = &Card{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(card.Table, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.GetType(); ok {
		_spec.SetField(card.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := cc.mutation.NumberHash(); ok {
		_spec.SetField(card.FieldNumberHash, field.TypeString, value)
		_node.NumberHash = value
	}
	if value, ok := cc.mutation.CvvHash(); ok {
		_spec.SetField(card.FieldCvvHash, field.TypeString, value)
		_node.CvvHash = value
	}
	if value, ok := cc.mutation.ExpiresAt(); ok {
		_spec.SetField(card.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
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
	if nodes := cc.mutation.PaymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   card.PaymentsTable,
			Columns: []string{card.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CardCreateBulk is the builder for creating many Card entities in bulk.
type CardCreateBulk struct {
	config
	err      error
	builders []*CardCreate
}

// Save creates the Card entities in the database.
func (ccb *CardCreateBulk) Save(ctx context.Context) ([]*Card, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Card, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CardMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CardCreateBulk) SaveX(ctx context.Context) []*Card {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CardCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CardCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
