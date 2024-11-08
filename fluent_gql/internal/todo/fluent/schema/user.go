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
	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/fluent_gql"
	"github.com/usalko/fluent/schema"
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
		field.String("name").
			Default("Anonymous"),
		field.UUID("username", uuid.UUID{}).
			Default(uuid.New),
		field.String("password").
			Sensitive().
			Optional(),
		field.JSON("required_metadata", map[string]any{}),
		field.JSON("metadata", map[string]any{}).
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("groups", Group.Type).
			Comment("The groups of the user").
			Annotations(
				fluent_gql.RelayConnection(),
				fluent_gql.OrderField("GROUPS_COUNT"),
			),
		edge.To("friends", User.Type).
			Through("friendships", Friendship.Type).
			Annotations(fluent_gql.RelayConnection()),
	}
}

// Annotations returns User annotations.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_gql.RelayConnection(),
		fluent_gql.QueryField(),
		fluent_gql.Mutations(
			fluent_gql.MutationCreate(),
			fluent_gql.MutationUpdate(),
		),
	}
}
