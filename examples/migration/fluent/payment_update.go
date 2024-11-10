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
	"github.com/usalko/fluent/examples/migration/fluent/card"
	"github.com/usalko/fluent/examples/migration/fluent/payment"
	"github.com/usalko/fluent/examples/migration/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCardID sets the "card_id" field.
func (pu *PaymentUpdate) SetCardID(i int) *PaymentUpdate {
	pu.mutation.SetCardID(i)
	return pu
}

// SetNillableCardID sets the "card_id" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableCardID(i *int) *PaymentUpdate {
	if i != nil {
		pu.SetCardID(*i)
	}
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PaymentUpdate) SetAmount(f float64) *PaymentUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(f)
	return pu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableAmount(f *float64) *PaymentUpdate {
	if f != nil {
		pu.SetAmount(*f)
	}
	return pu
}

// AddAmount adds f to the "amount" field.
func (pu *PaymentUpdate) AddAmount(f float64) *PaymentUpdate {
	pu.mutation.AddAmount(f)
	return pu
}

// SetCurrency sets the "currency" field.
func (pu *PaymentUpdate) SetCurrency(payment_currency payment.Currency) *PaymentUpdate {
	pu.mutation.SetCurrency(payment_currency)
	return pu
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableCurrency(payment_currency *payment.Currency) *PaymentUpdate {
	if payment_currency != nil {
		pu.SetCurrency(*payment_currency)
	}
	return pu
}

// SetTime sets the "time" field.
func (pu *PaymentUpdate) SetTime(time_time time.Time) *PaymentUpdate {
	pu.mutation.SetTime(time_time)
	return pu
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableTime(time_time *time.Time) *PaymentUpdate {
	if time_time != nil {
		pu.SetTime(*time_time)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *PaymentUpdate) SetDescription(s string) *PaymentUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableDescription(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// SetStatus sets the "status" field.
func (pu *PaymentUpdate) SetStatus(payment_status payment.Status) *PaymentUpdate {
	pu.mutation.SetStatus(payment_status)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableStatus(payment_status *payment.Status) *PaymentUpdate {
	if payment_status != nil {
		pu.SetStatus(*payment_status)
	}
	return pu
}

// SetCard sets the "card" edge to the Card entity.
func (pu *PaymentUpdate) SetCard(c *Card) *PaymentUpdate {
	return pu.SetCardID(c.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// ClearCard clears the "card" edge to the Card entity.
func (pu *PaymentUpdate) ClearCard() *PaymentUpdate {
	pu.mutation.ClearCard()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PaymentUpdate) check() error {
	if v, ok := pu.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`fluent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Currency(); ok {
		if err := payment.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf(`fluent: validator failed for field "Payment.currency": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := payment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`fluent: validator failed for field "Payment.status": %w`, err)}
		}
	}
	if pu.mutation.CardCleared() && len(pu.mutation.CardIDs()) > 0 {
		return errors.New(`fluent: clearing a required unique edge "Payment.card"`)
	}
	return nil
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.Currency(); ok {
		_spec.SetField(payment.FieldCurrency, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.Time(); ok {
		_spec.SetField(payment.FieldTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(payment.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(payment.FieldStatus, field.TypeEnum, value)
	}
	if pu.mutation.CardCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CardIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetCardID sets the "card_id" field.
func (puo *PaymentUpdateOne) SetCardID(i int) *PaymentUpdateOne {
	puo.mutation.SetCardID(i)
	return puo
}

// SetNillableCardID sets the "card_id" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableCardID(i *int) *PaymentUpdateOne {
	if i != nil {
		puo.SetCardID(*i)
	}
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PaymentUpdateOne) SetAmount(f float64) *PaymentUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(f)
	return puo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableAmount(f *float64) *PaymentUpdateOne {
	if f != nil {
		puo.SetAmount(*f)
	}
	return puo
}

// AddAmount adds f to the "amount" field.
func (puo *PaymentUpdateOne) AddAmount(f float64) *PaymentUpdateOne {
	puo.mutation.AddAmount(f)
	return puo
}

// SetCurrency sets the "currency" field.
func (puo *PaymentUpdateOne) SetCurrency(payment_currency payment.Currency) *PaymentUpdateOne {
	puo.mutation.SetCurrency(payment_currency)
	return puo
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableCurrency(payment_currency *payment.Currency) *PaymentUpdateOne {
	if payment_currency != nil {
		puo.SetCurrency(*payment_currency)
	}
	return puo
}

// SetTime sets the "time" field.
func (puo *PaymentUpdateOne) SetTime(time_time time.Time) *PaymentUpdateOne {
	puo.mutation.SetTime(time_time)
	return puo
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableTime(time_time *time.Time) *PaymentUpdateOne {
	if time_time != nil {
		puo.SetTime(*time_time)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *PaymentUpdateOne) SetDescription(s string) *PaymentUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableDescription(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// SetStatus sets the "status" field.
func (puo *PaymentUpdateOne) SetStatus(payment_status payment.Status) *PaymentUpdateOne {
	puo.mutation.SetStatus(payment_status)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableStatus(payment_status *payment.Status) *PaymentUpdateOne {
	if payment_status != nil {
		puo.SetStatus(*payment_status)
	}
	return puo
}

// SetCard sets the "card" edge to the Card entity.
func (puo *PaymentUpdateOne) SetCard(c *Card) *PaymentUpdateOne {
	return puo.SetCardID(c.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// ClearCard clears the "card" edge to the Card entity.
func (puo *PaymentUpdateOne) ClearCard() *PaymentUpdateOne {
	puo.mutation.ClearCard()
	return puo
}

// Where appends a list predicates to the PaymentUpdate builder.
func (puo *PaymentUpdateOne) Where(ps ...predicate.Payment) *PaymentUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PaymentUpdateOne) check() error {
	if v, ok := puo.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`fluent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Currency(); ok {
		if err := payment.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf(`fluent: validator failed for field "Payment.currency": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := payment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`fluent: validator failed for field "Payment.status": %w`, err)}
		}
	}
	if puo.mutation.CardCleared() && len(puo.mutation.CardIDs()) > 0 {
		return errors.New(`fluent: clearing a required unique edge "Payment.card"`)
	}
	return nil
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`fluent: missing "Payment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.Currency(); ok {
		_spec.SetField(payment.FieldCurrency, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.Time(); ok {
		_spec.SetField(payment.FieldTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(payment.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(payment.FieldStatus, field.TypeEnum, value)
	}
	if puo.mutation.CardCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CardIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
