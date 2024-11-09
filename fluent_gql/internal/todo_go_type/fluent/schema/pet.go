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
	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/schema/uintgql"
	"github.com/usalko/fluent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	fluent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []fluent.Field {
	return []fluent.Field{
		field.Uint64("id").
			GoType(uintgql.Uint64(0)),
		field.String("name").
			Default("Anonymous"),
	}
}
