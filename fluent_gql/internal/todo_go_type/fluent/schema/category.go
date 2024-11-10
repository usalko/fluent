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
	"math/big"
	"math/rand"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/schema"
	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/schema/bigintgql"
	"github.com/usalko/fluent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	fluent.Schema
}

// Mixin returns category mixed-in schema.
func (Category) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		// Reuse the fields and edges from base example.
		schema.Category{},
	}
}

// Fields returns category fields.
func (Category) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("id").
			GoType(bigintgql.BigInt{}).
			Unique().
			Immutable().
			DefaultFunc(func() bigintgql.BigInt {
				//nolint:gosec
				return bigintgql.BigInt{Int: big.NewInt(int64(rand.Float64()))}
			}),
	}
}