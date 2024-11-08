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

package todo

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/usalko/fluent/fluent_gql/internal/todo"
	"github.com/usalko/fluent/fluent_gql/internal/todouuid/fluent"
)

// Resolver is the resolver root.
type Resolver struct{ client *fluent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *fluent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client},
		Directives: DirectiveRoot{
			HasPermissions: todo.HasPermission(),
		},
	})
}
