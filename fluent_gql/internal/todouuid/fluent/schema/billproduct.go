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
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/schema"
	"github.com/usalko/fluent/schema/field"
)

// BillProduct holds the schema definition for the BillProduct entity.
type BillProduct struct {
	fluent.Schema
}

// Mixin returns BillProduct mixed-in schema.
func (BillProduct) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		// Reuse the fields and edges from base example.
		schema.BillProduct{},
	}
}

// Fields returns BillProduct fields.
func (BillProduct) Fields() []fluent.Field {
	return []fluent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
	}
}
