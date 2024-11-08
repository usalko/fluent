// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	fluent.Schema
}

// Fields of the Role.
func (Role) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Role.
func (Role) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("user", User.Type).
			Ref("roles").
			Through("roles_users", RoleUser.Type),
	}
}
