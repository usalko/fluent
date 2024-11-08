// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/mixin"
)

// User holds the schema for the user entity.
type User struct {
	fluent.Schema
}

func (User) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		UserMixin{},
	}
}

// Fields of the user.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("age"),
		field.String("name").
			StructTag(`json:"first_name" graphql:"first_name"`),
		field.String("last").
			Default("unknown").
			StructTag(`graphql:"last_name"`),
		field.String("nickname").
			Optional().
			Unique(),
		field.String("address").
			Optional().
			DefaultFunc(func() string { return "static" }),
		field.String("phone").
			Optional().
			Unique(),
		field.String("password").
			Optional().
			Sensitive(),
		field.Enum("role").
			Values("user", "admin", "free-user", "test user").
			Default("user"),
		field.Enum("employment").
			Values("Full-Time", "Part-Time", "Contract").
			Default("Full-Time"),
		field.String("SSOCert").
			Optional(),
		// Some users store the associations
		// count as a separate field.
		field.Int("files_count").
			Optional(),
	}
}

// Edges of the user.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("card", Card.Type).Comment("Cards associated with this user. O2O edge").Unique(),
		edge.To("pets", Pet.Type),
		edge.To("files", File.Type),
		edge.To("groups", Group.Type),
		edge.To("friends", User.Type),
		edge.To("following", User.Type).From("followers"),
		edge.To("team", Pet.Type).Unique(),
		edge.To("spouse", User.Type).Unique(),
		edge.To("parent", User.Type).Unique().From("children"),
	}
}

// UserMixin composes create/update time mixin.
type UserMixin struct {
	mixin.Schema
}

// Fields of the time mixin.
func (UserMixin) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("optional_int").
			Optional().
			Positive(),
	}
}
