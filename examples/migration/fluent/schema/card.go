// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	fluent.Schema
}

// Fields of the Card.
func (Card) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("type").
			Default("unknown"),
		field.String("number_hash"),
		field.String("cvv_hash"),
		field.Time("expires_at").
			Optional(),
		field.Int("owner_id").
			Default(0),
	}
}

// Edges of the Card.
func (Card) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("owner", User.Type).
			Ref("cards").
			Unique().
			Required().
			Field("owner_id"),
		edge.To("payments", Payment.Type),
	}
}

// Annotations of the Card.
func (Card) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Named check constraints are compared by their name.
		// Thus, the definition does not need to be normalized.
		fluent_sql.Checks(map[string]string{
			"number_hash_length": "(LENGTH(`number_hash`) = 16)",
		}),
	}
}
