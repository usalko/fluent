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
	"github.com/usalko/fluent/dialect"
	"github.com/usalko/fluent/fluent_gql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"

	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/schema/schematype"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	fluent.Schema
}

// Fields of the Category.
func (Category) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("id").
			Annotations(
				// Setting the OrderField explicitly on the "ID"
				// field, adds it to the generated GraphQL schema.
				fluent_gql.OrderField("ID"),
			),
		field.Text("text").
			NotEmpty().
			Annotations(
				fluent_gql.OrderField("TEXT"),
			),
		field.Enum("status").
			NamedValues(
				"Enabled", "ENABLED",
				"Disabled", "DISABLED",
			).
			Annotations(
				fluent_gql.Type("CategoryStatus"),
				fluent_gql.OrderField("STATUS"),
			),
		field.Other("config", &schematype.CategoryConfig{}).
			SchemaType(map[string]string{
				dialect.SQLite: "json",
			}).
			Optional(),
		field.JSON("types", &schematype.CategoryTypes{}).
			Optional(),
		field.Int64("duration").
			GoType(time.Duration(0)).
			Optional().
			Annotations(
				fluent_gql.OrderField("DURATION"),
				fluent_gql.Type("Duration"),
			),
		field.Uint64("count").
			Optional().
			Annotations(
				fluent_gql.OrderField("COUNT"),
				fluent_gql.Type("Uint64"),
			),
		field.Strings("strings").
			Optional().
			Deprecated("use `string` instead"),
	}
}

// Edges of the Category.
func (Category) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("todos", Todo.Type).
			Annotations(
				fluent_gql.RelayConnection(),
				fluent_gql.OrderField("TODOS_COUNT"),
			),
		edge.To("sub_categories", Category.Type).
			Annotations(fluent_gql.RelayConnection()),
	}
}

// Annotations returns Todo annotations.
func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_gql.QueryField(),
		fluent_gql.RelayConnection(),
		fluent_gql.Mutations(fluent_gql.MutationCreate(), fluent_gql.MutationUpdate()),
		fluent_gql.MultiOrder(),
	}
}
