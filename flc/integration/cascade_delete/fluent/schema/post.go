// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	fluent.Schema
}

// Fields of the Post.
func (Post) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("text").
			Default("What's on your mind?"),
		field.Int("author_id").
			Optional(),
	}
}

// Edges of the Post.
func (Post) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("author", User.Type).
			Ref("posts").
			Field("author_id").
			Unique(),
		edge.To("comments", Comment.Type).
			Annotations(fluent_sql.Annotation{
				OnDelete: fluent_sql.Cascade,
			}),
	}
}
