// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// RoleUser holds the schema definition for the RoleUser entity.
type RoleUser struct {
	fluent.Schema
}

func (RoleUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "role_id"),
	}
}

// Fields of the RoleUser.
func (RoleUser) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("created_at").
			Default(time.Now),
		// Edge fields.
		field.Int("role_id"),
		field.Int("user_id"),
	}
}

// Edges of the RoleUser.
func (RoleUser) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("role", Role.Type).
			Unique().
			Required().
			Field("role_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}
