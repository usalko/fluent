// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	fluent.Schema
}

// Fields of the Post.
func (Post) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("text"),
		field.Int("author_id").
			Optional().
			Nillable(),
	}
}

// Edges of the Post.
func (Post) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("author", User.Type).
			Field("author_id").
			Unique(),
	}
}

// Indexes of the Post.
func (Post) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("author_id", "text"),
	}
}
