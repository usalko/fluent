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
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/schema"
	"github.com/usalko/fluent/schema/field"

	"github.com/google/uuid"
)

// Group defines the group type schema.
type Group struct {
	fluent.Schema
}

// Mixin returns group mixed-in schema.
func (Group) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		schema.Group{},
	}
}

// Fields returns group fields.
func (Group) Fields() []fluent.Field {
	return []fluent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
	}
}