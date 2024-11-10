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
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// OneToMany holds the schema definition for the OneToMany entity.
type OneToMany struct {
	fluent.Schema
}

// Fields of the OneToMany.
func (OneToMany) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				fluent_gql.OrderField("NAME"),
			),
		field.String("field2").
			Optional(),
		field.Int("parent_id").
			Optional().
			Annotations(
				fluent_gql.Skip(fluent_gql.SkipAll),
			),
	}
}

// Edges of the OneToMany.
func (OneToMany) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("children", OneToMany.Type).
			From("parent").
			Field("parent_id").
			Unique(),
	}
}

func (OneToMany) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_gql.QueryField("oneToMany"),
		fluent_gql.RelayConnection(),
	}
}
