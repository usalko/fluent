// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	fluent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(uuid.NewString),
	}
}

// Edges of the Pet.
func (Pet) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
		edge.To("cars", Car.Type),
		edge.To("friends", Pet.Type),
		edge.To("best_friend", Pet.Type).
			Unique(),
	}
}
