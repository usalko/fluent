// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
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
		field.Time("expired"),
		field.String("number"),
	}
}

// Edges of the Card.
func (Card) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("owner", User.Type).
			Ref("card").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
	}
}
