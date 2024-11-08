// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/usalko/fluent/flc/integration/edge_schema/fluent"
)

// The AttachedFileFunc type is an adapter to allow the use of ordinary
// function as AttachedFile mutator.
type AttachedFileFunc func(context.Context, *fluent.AttachedFileMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f AttachedFileFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.AttachedFileMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.AttachedFileMutation", m)
}

// The FileFunc type is an adapter to allow the use of ordinary
// function as File mutator.
type FileFunc func(context.Context, *fluent.FileMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f FileFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.FileMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.FileMutation", m)
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

// The GroupTagFunc type is an adapter to allow the use of ordinary
// function as GroupTag mutator.
type GroupTagFunc func(context.Context, *fluent.GroupTagMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupTagFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.GroupTagMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.GroupTagMutation", m)
}

// The ProcessFunc type is an adapter to allow the use of ordinary
// function as Process mutator.
type ProcessFunc func(context.Context, *fluent.ProcessMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f ProcessFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.ProcessMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.ProcessMutation", m)
}

// The RelationshipFunc type is an adapter to allow the use of ordinary
// function as Relationship mutator.
type RelationshipFunc func(context.Context, *fluent.RelationshipMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f RelationshipFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.RelationshipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.RelationshipMutation", m)
}

// The RelationshipInfoFunc type is an adapter to allow the use of ordinary
// function as RelationshipInfo mutator.
type RelationshipInfoFunc func(context.Context, *fluent.RelationshipInfoMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f RelationshipInfoFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.RelationshipInfoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.RelationshipInfoMutation", m)
}

// The RoleFunc type is an adapter to allow the use of ordinary
// function as Role mutator.
type RoleFunc func(context.Context, *fluent.RoleMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.RoleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.RoleMutation", m)
}

// The RoleUserFunc type is an adapter to allow the use of ordinary
// function as RoleUser mutator.
type RoleUserFunc func(context.Context, *fluent.RoleUserMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleUserFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.RoleUserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.RoleUserMutation", m)
}

// The TagFunc type is an adapter to allow the use of ordinary
// function as Tag mutator.
type TagFunc func(context.Context, *fluent.TagMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f TagFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.TagMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.TagMutation", m)
}

// The TweetFunc type is an adapter to allow the use of ordinary
// function as Tweet mutator.
type TweetFunc func(context.Context, *fluent.TweetMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f TweetFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.TweetMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.TweetMutation", m)
}

// The TweetLikeFunc type is an adapter to allow the use of ordinary
// function as TweetLike mutator.
type TweetLikeFunc func(context.Context, *fluent.TweetLikeMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f TweetLikeFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.TweetLikeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.TweetLikeMutation", m)
}

// The TweetTagFunc type is an adapter to allow the use of ordinary
// function as TweetTag mutator.
type TweetTagFunc func(context.Context, *fluent.TweetTagMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f TweetTagFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.TweetTagMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.TweetTagMutation", m)
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

// The UserGroupFunc type is an adapter to allow the use of ordinary
// function as UserGroup mutator.
type UserGroupFunc func(context.Context, *fluent.UserGroupMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f UserGroupFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.UserGroupMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.UserGroupMutation", m)
}

// The UserTweetFunc type is an adapter to allow the use of ordinary
// function as UserTweet mutator.
type UserTweetFunc func(context.Context, *fluent.UserTweetMutation) (fluent.Value, error)

// Mutate calls f(ctx, m).
func (f UserTweetFunc) Mutate(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
	if mv, ok := m.(*fluent.UserTweetMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *fluent.UserTweetMutation", m)
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
