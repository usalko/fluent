// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	fluent.Schema
}

// Fields of the File.
func (File) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
		field.Bool("deleted").
			Default(false),
		field.Int("parent_id").
			Optional(),
	}
}

// Edges of the File.
func (File) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("children", File.Type).
			From("parent").
			Unique().
			Field("parent_id"),
	}
}
