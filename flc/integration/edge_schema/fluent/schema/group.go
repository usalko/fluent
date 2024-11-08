// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	fluent.Schema
}

// Fields of the Group.
func (Group) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Default("Unknown"),
	}
}

// Edges of the Group.
func (Group) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("users", User.Type).
			Ref("groups").
			Through("joined_users", UserGroup.Type),
		edge.From("tags", Tag.Type).
			Ref("groups").
			Through("group_tags", GroupTag.Type),
	}
}
