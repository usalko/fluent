// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect"
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
		field.Int("id").
			SchemaType(map[string]string{
				dialect.SQLite: "integer",
			}).
			Immutable(),
		field.Int("parent_id").
			Optional().
			Immutable(),
		field.Int("spouse_id").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("pets", Pet.Type),
		edge.To("children", User.Type).
			From("parent").
			Field("parent_id").
			Immutable().
			Comment("The parent edge and its field are immutable").
			Unique(),
		edge.To("spouse", User.Type).
			Field("spouse_id").
			Unique(),
		edge.To("card", Card.Type).
			Unique(),
		edge.To("metadata", Metadata.Type).
			Unique().
			StorageKey(edge.Column("id")),
		edge.From("info", Info.Type).
			Ref("user"),
		edge.To("rentals", Rental.Type),
	}
}
