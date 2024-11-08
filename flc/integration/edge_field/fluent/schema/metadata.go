// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	fluent.Schema
}

func (Metadata) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("id"),
		field.Int("age").
			Default(0),
		field.Int("parent_id").
			Optional(),
	}
}
func (Metadata) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("user", User.Type).
			Ref("metadata").
			Unique(),
		edge.To("parent", Metadata.Type).
			Field("parent_id").
			Unique().
			From("children"),
	}
}
