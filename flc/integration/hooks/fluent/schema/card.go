// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"time"

	gen "github.com/usalko/fluent/flc/integration/hooks/fluent"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/card"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/hook"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/mixin"
)

// RejectUpdate rejects ~all update operations
// that are not on a specific entity.
type RejectUpdate struct {
	mixin.Schema
}

func (RejectUpdate) Hooks() []fluent.Hook {
	return []fluent.Hook{
		hook.If(
			hook.Reject(fluent.OpUpdate),
			// Accept only updates that contains the "expired_at" field.
			hook.Not(hook.HasFields(card.FieldExpiredAt)),
		),
	}
}

// Card holds the schema definition for the CreditCard entity.
type Card struct {
	fluent.Schema
}

func (Card) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		RejectUpdate{},
	}
}

func (Card) Hooks() []fluent.Hook {
	return []fluent.Hook{
		hook.On(
			func(next fluent.Mutator) fluent.Mutator {
				return fluent.MutateFunc(func(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
					num, ok := m.Field(card.FieldNumber)
					if !ok {
						return nil, fmt.Errorf("missing card number value")
					}
					// Validator in hooks.
					if len(num.(string)) < 4 {
						return nil, fmt.Errorf("card number is too short")
					}
					return next.Mutate(ctx, m)
				})
			},
			fluent.OpCreate,
		),
		func(next fluent.Mutator) fluent.Mutator {
			return hook.CardFunc(func(ctx context.Context, m *gen.CardMutation) (fluent.Value, error) {
				m.SetInHook("value was set in hook")
				if _, ok := m.Name(); !ok {
					m.SetName("unknown")
				}
				return next.Mutate(ctx, m)
			})
		},
	}
}

func (Card) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("number").
			Immutable().
			Default("unknown").
			NotEmpty(),
		field.String("name").
			Optional().
			Comment("Exact name written on card"),
		field.Time("created_at").
			Default(time.Now),
		field.String("in_hook").
			Comment("InHook is a mandatory field that is set by the hook."),
		field.Time("expired_at").
			Optional(),
	}
}

func (Card) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("owner", User.Type).
			Ref("cards").
			Unique(),
	}
}

func (Card) Interceptors() []fluent.Interceptor {
	return []fluent.Interceptor{
		fluent.InterceptFunc(func(next fluent.Querier) fluent.Querier {
			return fluent.QuerierFunc(func(ctx context.Context, q fluent.Query) (fluent.Value, error) {
				return next.Query(ctx, q)
			})
		}),
	}
}
