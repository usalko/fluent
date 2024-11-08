// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	fluent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("value"),
	}
}

// Edges of the Tag.
func (Tag) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("tweets", Tweet.Type).
			Through("tweet_tags", TweetTag.Type),
		edge.To("groups", Group.Type).
			Through("group_tags", GroupTag.Type),
	}
}
