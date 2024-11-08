// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/privacy"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// Relationship holds the schema definition for the Relationship entity.
type Relationship struct {
	fluent.Schema
}

func (Relationship) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "relative_id"),
	}
}

// Fields of the Relationship.
func (Relationship) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("weight").
			Default(1),

		// Edge fields for the fields that compose this edge.
		// They also function as the primary key of the table.
		field.Int("user_id"),
		field.Int("relative_id"),

		// An edge field to external node that holds
		// additional information about this edge.
		field.Int("info_id").
			Optional(),
	}
}

// Edges of the Relationship.
func (Relationship) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("relative", User.Type).
			Required().
			Unique().
			Field("relative_id"),

		// An optional edge to an entity that holds
		// information about this relationship.
		edge.To("info", RelationshipInfo.Type).
			Field("info_id").
			Unique(),
	}
}

// Indexes of the Relationship.
func (Relationship) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("weight"),

		// A relationship-info can be connected to no more
		// than one relationship object (and edge schema).
		index.Edges("info").
			Unique(),
	}
}

// Policy defines the privacy policy of the Relationship.
func (Relationship) Policy() fluent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
