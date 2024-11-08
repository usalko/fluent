// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// FileType holds the schema definition for the FileType entity.
type FileType struct {
	fluent.Schema
}

// Fields of the FileType.
func (FileType) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Unique(),
		field.Enum("type").
			NamedValues(
				"PNG", "png",
				"SVG", "svg",
				"JPG", "jpg",
			).
			Default("png"),
		field.Enum("state").
			NamedValues(
				"On", "ON",
				"Off", "OFF",
			).
			Default("ON"),
	}
}

// Edges of the FileType.
func (FileType) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("files", File.Type),
	}
}
