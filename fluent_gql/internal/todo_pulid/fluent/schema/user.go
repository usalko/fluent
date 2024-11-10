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
	"github.com/usalko/fluent/fluent_gql/internal/todo_pulid/fluent/schema/pulid"
	"github.com/usalko/fluent/schema/edge"
)

// User defines the user type schema.
type User struct {
	fluent.Schema
}

// Mixin returns user mixed-in schema.
func (User) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		pulid.MixinWithPrefix("UR"),
		schema.FilterEdges(schema.User{}, "friends", "friendships"),
	}
}

// Edges of the User.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("friends", User.Type).
			Through("friendships", Friendship.Type),
	}
}