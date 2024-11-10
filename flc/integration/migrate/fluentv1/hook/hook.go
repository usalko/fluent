// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/usalko/fluent/flc/integration/migrate/fluentv1"
)

// The CarFunc type is an adapter to allow the use of ordinary
// function as Car mutator.
type CarFunc func(context.Context, *fluentv1.CarMutation) (fluentv1.Value, error)

// Mutate calls f(ctx, m).
func (f CarFunc) Mutate(ctx context.Context, m fluentv1.Mutation) (fluentv1.Value, error) {
	if mv, ok := m.(*fluentv1.CarMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluentv1.CarMutation", m)
}

// The ConversionFunc type is an adapter to allow the use of ordinary
// function as Conversion mutator.
type ConversionFunc func(context.Context, *fluentv1.ConversionMutation) (fluentv1.Value, error)

// Mutate calls f(ctx, m).
func (f ConversionFunc) Mutate(ctx context.Context, m fluentv1.Mutation) (fluentv1.Value, error) {
	if mv, ok := m.(*fluentv1.ConversionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluentv1.ConversionMutation", m)
}

// The CustomTypeFunc type is an adapter to allow the use of ordinary
// function as CustomType mutator.
type CustomTypeFunc func(context.Context, *fluentv1.CustomTypeMutation) (fluentv1.Value, error)

// Mutate calls f(ctx, m).
func (f CustomTypeFunc) Mutate(ctx context.Context, m fluentv1.Mutation) (fluentv1.Value, error) {
	if mv, ok := m.(*fluentv1.CustomTypeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluentv1.CustomTypeMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *fluentv1.UserMutation) (fluentv1.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m fluentv1.Mutation) (fluentv1.Value, error) {
	if mv, ok := m.(*fluentv1.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluentv1.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, fluentv1.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m fluentv1.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m fluentv1.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m fluentv1.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op fluentv1.Op) Condition {
	return func(_ context.Context, m fluentv1.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m fluentv1.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m fluentv1.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m fluentv1.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk fluentv1.Hook, cond Condition) fluentv1.Hook {
	return func(next fluentv1.Mutator) fluentv1.Mutator {
		return fluentv1.MutateFunc(func(ctx context.Context, m fluentv1.Mutation) (fluentv1.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, fluentv1.Delete|fluentv1.Create)
func On(hk fluentv1.Hook, op fluentv1.Op) fluentv1.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, fluentv1.Update|fluentv1.UpdateOne)
func Unless(hk fluentv1.Hook, op fluentv1.Op) fluentv1.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) fluentv1.Hook {
	return func(fluentv1.Mutator) fluentv1.Mutator {
		return fluentv1.MutateFunc(func(context.Context, fluentv1.Mutation) (fluentv1.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []fluentv1.Hook {
//		return []fluentv1.Hook{
//			Reject(fluentv1.Delete|fluentv1.Update),
//		}
//	}
func Reject(op fluentv1.Op) fluentv1.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []fluentv1.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...fluentv1.Hook) Chain {
	return Chain{append([]fluentv1.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() fluentv1.Hook {
	return func(mutator fluentv1.Mutator) fluentv1.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...fluentv1.Hook) Chain {
	newHooks := make([]fluentv1.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
