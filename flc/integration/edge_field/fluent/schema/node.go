// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Node holds the schema definition for the Node entity.
type Node struct {
	fluent.Schema
}

// Fields of the Node.
func (Node) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("value").
			Default(0),
		field.Int("prev_id").
			Optional(),
	}
}

// Edges of the Node.
func (Node) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("next", Node.Type).
			Unique().
			From("prev").
			Field("prev_id").
			Unique(),
	}
}
