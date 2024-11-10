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
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/fluent_gql"
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/schema/annotation"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	fluent.Schema
}

// Fields of the Group.
func (Group) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Default("Unknown"),
	}
}

// Edges of the Group.
func (Group) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("users", User.Type).
			Ref("groups").
			Annotations(fluent_gql.RelayConnection()),
	}
}

// Annotations returns Group annotations.
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_gql.RelayConnection(),
		fluent_gql.QueryField(),
		fluent_gql.Directives(
			annotation.HasPermissions([]string{"ADMIN", "MODERATOR"}),
		),
		fluent_gql.MultiOrder(),
		fluent_gql.Implements("NamedNode"),
	}
}
