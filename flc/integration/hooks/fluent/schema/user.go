// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/flc/integration/hooks/fluent/user"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/hook"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Mixin of the User.
func (User) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		VersionMixin{},
	}
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
		field.Uint("worth").
			Optional(),
		field.String("password").
			Optional().
			Sensitive(),
		field.Bool("active").
			Default(true),
	}
}

// Edges of the User.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("cards", Card.Type),
		edge.To("pets", Pet.Type),
		edge.To("friends", User.Type),
		edge.To("best_friend", User.Type).
			Unique(),
	}
}

// Hooks of the User.
func (User) Hooks() []fluent.Hook {
	return []fluent.Hook{
		hook.If(
			hook.FixedError(errors.New("password cannot be edited on update-many")),
			hook.And(
				hook.HasOp(fluent.OpUpdate),
				hook.Or(
					hook.HasFields(user.FieldPassword),
					hook.HasClearedFields(user.FieldPassword),
				),
			),
		),
	}
}

type VersionMixin struct {
	mixin.Schema
}

func (VersionMixin) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("version").
			Default(0),
	}
}

func (VersionMixin) Hooks() []fluent.Hook {
	return []fluent.Hook{
		hook.On(VersionHook(), fluent.OpUpdateOne),
	}
}

func VersionHook() fluent.Hook {
	type OldSetVersion interface {
		SetVersion(int)
		Version() (int, bool)
		OldVersion(context.Context) (int, error)
	}
	return func(next fluent.Mutator) fluent.Mutator {
		// A hook that validates the "version" field is incremented by 1 on each update.
		// Note that this is just a dummy example, and it doesn't promise consistency in
		// the database.
		return fluent.MutateFunc(func(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
			ver, ok := m.(OldSetVersion)
			if !ok {
				return next.Mutate(ctx, m)
			}
			oldV, err := ver.OldVersion(ctx)
			if err != nil {
				return nil, err
			}
			curV, exists := ver.Version()
			if !exists {
				return nil, fmt.Errorf("version field is required in update mutation")
			}
			if curV != oldV+1 {
				return nil, fmt.Errorf("version field must be incremented by 1")
			}
			// Add an SQL predicate that validates the "version" column is equal
			// to "oldV" (it wasn't changed during the mutation by other process).
			return next.Mutate(ctx, m)
		})
	}
}
