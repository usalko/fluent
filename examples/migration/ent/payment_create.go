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
	"github.com/usalko/fluent/schema/field"
)

// PaymentCreate is the builder for creating a Payment entity.
type PaymentCreate struct {
	config
	mutation *PaymentMutation
	hooks    []Hook
}

// SetCardID sets the "card_id" field.
func (pc *PaymentCreate) SetCardID(i int) *PaymentCreate {
	pc.mutation.SetCardID(i)
	return pc
}

// SetAmount sets the "amount" field.
func (pc *PaymentCreate) SetAmount(f float64) *PaymentCreate {
	pc.mutation.SetAmount(f)
	return pc
}

// SetCurrency sets the "currency" field.
func (pc *PaymentCreate) SetCurrency(pa payment.Currency) *PaymentCreate {
	pc.mutation.SetCurrency(pa)
	return pc
}

// SetTime sets the "time" field.
func (pc *PaymentCreate) SetTime(t time.Time) *PaymentCreate {
	pc.mutation.SetTime(t)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PaymentCreate) SetDescription(s string) *PaymentCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetStatus sets the "status" field.
func (pc *PaymentCreate) SetStatus(pa payment.Status) *PaymentCreate {
	pc.mutation.SetStatus(pa)
	return pc
}

// SetCard sets the "card" edge to the Card entity.
func (pc *PaymentCreate) SetCard(c *Card) *PaymentCreate {
	return pc.SetCardID(c.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (pc *PaymentCreate) Mutation() *PaymentMutation {
	return pc.mutation
}

// Save creates the Payment in the database.
func (pc *PaymentCreate) Save(ctx context.Context) (*Payment, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaymentCreate) SaveX(ctx context.Context) *Payment {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PaymentCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PaymentCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PaymentCreate) check() error {
	if _, ok := pc.mutation.CardID(); !ok {
		return &ValidationError{Name: "card_id", err: errors.New(`ent: missing required field "Payment.card_id"`)}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Payment.amount"`)}
	}
	if v, ok := pc.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Currency(); !ok {
		return &ValidationError{Name: "currency", err: errors.New(`ent: missing required field "Payment.currency"`)}
	}
	if v, ok := pc.mutation.Currency(); ok {
		if err := payment.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf(`ent: validator failed for field "Payment.currency": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "Payment.time"`)}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Payment.description"`)}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Payment.status"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := payment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Payment.status": %w`, err)}
		}
	}
	if len(pc.mutation.CardIDs()) == 0 {
		return &ValidationError{Name: "card", err: errors.New(`ent: missing required edge "Payment.card"`)}
	}
	return nil
}

func (pc *PaymentCreate) sqlSave(ctx context.Context) (*Payment, error) {
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
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PaymentCreate) createSpec() (*Payment, *sqlgraph.CreateSpec) {
	var (
		_node = &Payment{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(payment.Table, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := pc.mutation.Currency(); ok {
		_spec.SetField(payment.FieldCurrency, field.TypeEnum, value)
		_node.Currency = value
	}
	if value, ok := pc.mutation.Time(); ok {
		_spec.SetField(payment.FieldTime, field.TypeTime, value)
		_node.Time = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(payment.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(payment.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := pc.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.CardTable,
			Columns: []string{payment.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CardID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PaymentCreateBulk is the builder for creating many Payment entities in bulk.
type PaymentCreateBulk struct {
	config
	err      error
	builders []*PaymentCreate
}

// Save creates the Payment entities in the database.
func (pcb *PaymentCreateBulk) Save(ctx context.Context) ([]*Payment, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Payment, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentMutation)
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PaymentCreateBulk) SaveX(ctx context.Context) []*Payment {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PaymentCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PaymentCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
