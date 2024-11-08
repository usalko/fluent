// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/fluent_gql"
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/schema/customstruct"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Todo defines the todo type schema.
type Todo struct {
	fluent.Schema
}

// Fields returns todo fields.
func (Todo) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				fluent_gql.OrderField("CREATED_AT"),
				fluent_gql.Skip(fluent_gql.SkipMutationCreateInput),
			),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
				"Pending", "PENDING",
			).
			Annotations(
				fluent_gql.OrderField("STATUS"),
			),
		field.Int("priority").
			Default(0).
			Annotations(
				fluent_gql.OrderField("PRIORITY_ORDER"),
				fluent_gql.MapsTo("priorityOrder"),
			),
		field.Text("text").
			NotEmpty().
			Annotations(
				fluent_gql.OrderField("TEXT"),
			),
		field.Bytes("blob").
			Annotations(
				fluent_gql.Skip(),
			).
			Optional(),
		field.Int("category_id").
			Optional().
			Immutable().
			Annotations(
				fluent_gql.MapsTo("categoryID", "category_id", "categoryX"),
			),
		field.JSON("init", map[string]any{}).
			Optional().
			Annotations(fluent_gql.Type("Map")),
		field.JSON("custom", []customstruct.Custom{}).
			Annotations(
				fluent_gql.Skip(fluent_gql.SkipMutationCreateInput),
				fluent_gql.Skip(fluent_gql.SkipMutationUpdateInput),
			).
			Optional(),
		field.JSON("customp", []*customstruct.Custom{}).
			Annotations(
				fluent_gql.Skip(fluent_gql.SkipMutationCreateInput),
				fluent_gql.Skip(fluent_gql.SkipMutationUpdateInput),
			).
			Optional(),
		field.Int("value").
			Default(0),
	}
}

// Edges returns todo edges.
func (Todo) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("children", Todo.Type).
			Annotations(
				fluent_gql.RelayConnection(),
				// For non-unique edges, the order field can be only on edge count.
				// The convention is "UPPER(<edge-name>)_COUNT".
				fluent_gql.OrderField("CHILDREN_COUNT"),
			).
			From("parent").
			Annotations(
				// For unique edges, the order field can be on the edge field that is defined
				// as fluent_gql.OrderField. The convention is "UPPER(<edge-name>)_<gql-order-field>".
				fluent_gql.OrderField("PARENT_STATUS"),
			).
			Unique(),
		edge.From("category", Category.Type).
			Ref("todos").
			Field("category_id").
			Unique().
			Immutable().
			Annotations(
				fluent_gql.OrderField("CATEGORY_TEXT"),
			),
		edge.To("secret", VerySecret.Type).
			Unique(),
	}
}

// Annotations returns Todo annotations.
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_gql.RelayConnection(),
		fluent_gql.QueryField().Description("This is the todo item"),
		fluent_gql.Mutations(fluent_gql.MutationCreate(), fluent_gql.MutationUpdate()),
		fluent_gql.MultiOrder(),
	}
}
