// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("age"),
		field.String("name"),
	}
}

// Edges of the User.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("card", Card.Type).
			Unique(),
	}
}