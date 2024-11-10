// Copyright 2024-present Fluent
//
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
			),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Annotations(
				fluent_gql.OrderField("STATUS"),
			),
		field.Int("priority").
			Default(0).
			Annotations(
				fluent_gql.OrderField("PRIORITY"),
			),
		field.Text("text").
			NotEmpty().
			Annotations(
				fluent_gql.OrderField("TEXT"),
			),
		field.Bytes("blob").
			Optional(),
	}
}

// Edges returns todo edges.
func (Todo) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("children", Todo.Type).
			//nolint SA1019 we keep this as the example.
			Annotations(fluent_gql.Bind()).
			From("parent").
			//nolint SA1019 we keep this as the example.
			Annotations(fluent_gql.Bind()).
			Unique(),
		edge.From("category", Category.Type).
			Ref("todos").
			Unique(),
		edge.To("secret", VerySecret.Type).
			Unique(),
	}
}
