// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	fluent.Schema
}

// Fields of the Card.
func (Card) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("number").
			Optional(),
		field.Int("owner_id").
			Optional(),
	}
}

// Edges of the Card.
func (Card) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("owner", User.Type).
			Ref("card").
			Field("owner_id").
			Unique(),
	}
}

// Indexes of the Card.
func (Card) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("number", "owner_id"),
	}
}
