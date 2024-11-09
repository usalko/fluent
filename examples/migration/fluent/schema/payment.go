// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	fluent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("card_id"),
		field.Float("amount").
			Positive(),
		field.Enum("currency").
			Values("USD", "EUR", "VND", "ILS"),
		field.Time("time"),
		field.String("description"),
		field.Enum("status").
			Values("pending", "completed", "failed"),
	}
}

// Edges of the Payment.
func (Payment) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("card", Card.Type).
			Ref("payments").
			Unique().
			Required().
			Field("card_id"),
	}
}

// Indexes of the Payment.
func (Payment) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("status", "time"),
	}
}

// Annotations of the Payment.
func (Payment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Named check constraints are compared by their name.
		// Thus, the definition does not need to be normalized.
		fluent_sql.Checks(map[string]string{
			"amount_positive": "(`amount` > 0)",
		}),
	}
}
