// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	fluent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []fluent.Field {
	return []fluent.Field{
		field.UUID("id", uuid.Nil).
			Default(uuid.New),
		field.String("name"),
		field.Float("age"),
		field.Float("weight"),
		field.UUID("best_friend_id", uuid.Nil).
			Annotations(
				fluent_sql.Default(uuid.Nil.String()),
			),
		field.Int("owner_id").
			Default(0),
	}
}

// Edges of the Pet.
func (Pet) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("best_friend", Pet.Type).
			Unique().
			Required().
			Field("best_friend_id"),
		edge.To("owner", User.Type).
			Unique().
			Required().
			Field("owner_id"),
	}
}

// Indexes of the Pet.
func (Pet) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("name", "owner_id").
			Unique(),
	}
}
