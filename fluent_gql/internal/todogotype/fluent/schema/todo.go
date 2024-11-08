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
	todoschema "github.com/usalko/fluent/fluent_gql/internal/todo/fluent/schema"
	"github.com/usalko/fluent/fluent_gql/internal/todogotype/fluent/schema/bigintgql"
	"github.com/usalko/fluent/schema/field"
)

// Todo defines the todo type schema.
type Todo struct {
	fluent.Schema
}

// Mixin returns todo mixed-in schema.
func (Todo) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		todoschema.FilterFields(todoschema.Todo{}, "category_id"),
	}
}

// Fields returns todo fields.
func (Todo) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("id").
			Default("plain"),
		field.String("category_id").
			GoType(bigintgql.BigInt{}).
			Optional().
			Immutable(),
	}
}
