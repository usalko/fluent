// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/google/uuid"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("id").
			StorageKey("oid"),
		field.Int32("age"),
		field.String("name").
			MaxLen(10),
		field.Text("description").
			Optional(),
		field.String("nickname").
			Unique(),
		field.String("address").
			Optional(),
		field.String("renamed").
			Optional(),
		field.String("old_token").
			DefaultFunc(uuid.NewString),
		field.Bytes("blob").
			Optional().
			MaxLen(255),
		field.Enum("state").
			Optional().
			Values("logged_in", "logged_out").
			Default("logged_in"),
		field.String("status").
			Optional(),
		field.String("workplace").
			MaxLen(30).
			Optional(),
		field.String("drop_optional").
			Optional(),
	}
}

func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("children", User.Type).
			From("parent").
			Unique(),
		edge.To("spouse", User.Type).
			Unique(),
		edge.To("car", Car.Type).
			Unique(),
	}
}

func (User) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("description").
			Annotations(fluent_sql.Prefix(50)),
		index.Fields("name", "address").
			Unique(),
	}
}

type Car struct {
	fluent.Schema
}

// Annotations of the Car.
func (Car) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_sql.Annotation{Table: "Car"},
	}
}

func (Car) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("owner", User.Type).
			Ref("car").
			Unique(),
	}
}
