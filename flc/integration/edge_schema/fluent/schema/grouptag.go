// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// GroupTag holds the schema definition for the GroupTag entity.
type GroupTag struct {
	fluent.Schema
}

// Fields of the GroupTag.
func (GroupTag) Fields() []fluent.Field {
	return []fluent.Field{
		// An edge schema with the builtin ID
		// field, but without any other field.
		field.Int("tag_id"),
		field.Int("group_id"),
	}
}

// Edges of the GroupTag.
func (GroupTag) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("tag", Tag.Type).
			Unique().
			Required().
			Field("tag_id"),
		edge.To("group", Group.Type).
			Unique().
			Required().
			Field("group_id"),
	}
}
