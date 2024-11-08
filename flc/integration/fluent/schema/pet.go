// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"

	"github.com/google/uuid"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	fluent.Schema
}

// Annotations of the Pet.
func (Pet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_sql.Annotation{Table: "pet"},
	}
}

// Fields of the Pet.
func (Pet) Fields() []fluent.Field {
	return []fluent.Field{
		field.Float("age").
			Default(0),
		field.String("name"),
		field.UUID("uuid", uuid.UUID{}).
			Optional(),
		field.String("nickname").
			Optional(),
		field.Bool("trained").
			Default(false),
		field.Time("optional_time").
			Optional(),
	}
}

// Edges of the Dog.
func (Pet) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("team", User.Type).
			Ref("team").
			Unique(),
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}

func (Pet) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("name").
			Edges("owner"),
		index.Fields("nickname").
			Unique(),
	}
}
