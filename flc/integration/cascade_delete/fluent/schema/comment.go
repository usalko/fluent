// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	fluent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("text"),
		field.Int("post_id"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("post", Post.Type).
			Ref("comments").
			Field("post_id").
			Required().
			Unique(),
	}
}
