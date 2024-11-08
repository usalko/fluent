// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// UserGroup holds the schema definition for the UserGroup entity.
type UserGroup struct {
	fluent.Schema
}

// Fields of the UserGroup.
func (UserGroup) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("joined_at").
			Default(time.Now),
		field.Int("user_id"),
		field.Int("group_id"),
	}
}

// Edges of the UserGroup.
func (UserGroup) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("group", Group.Type).
			Unique().
			Required().
			Field("group_id"),
	}
}
