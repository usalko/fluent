// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// An example for an fluent/schema that mixed in an fluent/mixin to define default table schema (in the database).

// Friendship holds the edge schema definition of the Friendship relationship.
type Friendship struct {
	fluent.Schema
}

// Mixin defines the schemas that mixed into this schema.
func (Friendship) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		Mixin{},
	}
}

// Fields of the Friendship.
func (Friendship) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("weight").
			Default(1),
		field.Time("created_at").
			Default(time.Now),
		field.Int("user_id").
			Immutable(),
		field.Int("friend_id").
			Immutable(),
	}
}

// Edges of the Friendship.
func (Friendship) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Immutable().
			Field("user_id"),
		edge.To("friend", User.Type).
			Unique().
			Required().
			Immutable().
			Field("friend_id"),
	}
}

// Indexes of the Friendship.
func (Friendship) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("created_at"),
	}
}
