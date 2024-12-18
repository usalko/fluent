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
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Friendship holds the edge schema definition of the Friendship relationship.
type Friendship struct {
	fluent.Schema
}

// Fields of the Friendship.
func (Friendship) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Int("user_id"),
		field.Int("friend_id"),
	}
}

// Edges of the Friendship.
func (Friendship) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("friend", User.Type).
			Required().
			Unique().
			Field("friend_id"),
	}
}

// Annotations of the Friendship.
func (Friendship) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_gql.RelayConnection(),
		fluent_gql.Mutations(
			fluent_gql.MutationUpdate(),
		),
	}
}
