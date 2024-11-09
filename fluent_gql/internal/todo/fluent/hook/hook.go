// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by flc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent"
)

// The BillProductFunc type is an adapter to allow the use of ordinary
// function as BillProduct mutator.
type BillProductFunc func(context.Context, *fluent.BillProductMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f BillProductFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.BillProductMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.BillProductMutation", m)
}

// The CategoryFunc type is an adapter to allow the use of ordinary
// function as Category mutator.
type CategoryFunc func(context.Context, *fluent.CategoryMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f CategoryFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.CategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.CategoryMutation", m)
}

// The FriendshipFunc type is an adapter to allow the use of ordinary
// function as Friendship mutator.
type FriendshipFunc func(context.Context, *fluent.FriendshipMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f FriendshipFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.FriendshipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.FriendshipMutation", m)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *fluent.GroupMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.GroupMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.GroupMutation", m)
}

// The OneToManyFunc type is an adapter to allow the use of ordinary
// function as OneToMany mutator.
type OneToManyFunc func(context.Context, *fluent.OneToManyMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f OneToManyFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.OneToManyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.OneToManyMutation", m)
}

// The ProjectFunc type is an adapter to allow the use of ordinary
// function as Project mutator.
type ProjectFunc func(context.Context, *fluent.ProjectMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.ProjectMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.ProjectMutation", m)
}

// The TodoFunc type is an adapter to allow the use of ordinary
// function as Todo mutator.
type TodoFunc func(context.Context, *fluent.TodoMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f TodoFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.TodoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.TodoMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *fluent.UserMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.UserMutation", m)
}

// The VerySecretFunc type is an adapter to allow the use of ordinary
// function as VerySecret mutator.
type VerySecretFunc func(context.Context, *fluent.VerySecretMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f VerySecretFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.VerySecretMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.VerySecretMutation", m)
}

// The WorkspaceFunc type is an adapter to allow the use of ordinary
// function as Workspace mutator.
type WorkspaceFunc func(context.Context, *fluent.WorkspaceMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f WorkspaceFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.WorkspaceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.WorkspaceMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, fluent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m fluent.Mutation) bool {
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
	return func(ctx context.Context, m fluent.Mutation) bool {
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
	return func(ctx context.Context, m fluent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op fluent.Op) Condition {
	return func(_ context.Context, m fluent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m fluent.Mutation) bool {
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
	return func(_ context.Context, m fluent.Mutation) bool {
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
	return func(_ context.Context, m fluent.Mutation) bool {
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
func If(hk fluent.Hook, cond Condition) fluent.Hook {
	return func(next fluent.Mutator) fluent.Mutator {
		return fluent.MutateFunc(func(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, fluent.Delete|fluent.Create)
func On(hk fluent.Hook, op fluent.Op) fluent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, fluent.Update|fluent.UpdateOne)
func Unless(hk fluent.Hook, op fluent.Op) fluent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) fluent.Hook {
	return func(fluent.Mutator) fluent.Mutator {
		return fluent.MutateFunc(func(context.Context, fluent.Mutation) (fluent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []fluent.Hook {
//		return []fluent.Hook{
//			Reject(fluent.Delete|fluent.Update),
//		}
//	}
func Reject(op fluent.Op) fluent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []fluent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...fluent.Hook) Chain {
	return Chain{append([]fluent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() fluent.Hook {
	return func(mutator fluent.Mutator) fluent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...fluent.Hook) Chain {
	newHooks := make([]fluent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
