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
		field.Int("value"),
		field.Int("parent_id").
			Optional(),
	}
}

// Edges of the Node.
func (Node) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("children", Node.Type).
			From("parent").
			Field("parent_id").
			Unique(),
	}
}
