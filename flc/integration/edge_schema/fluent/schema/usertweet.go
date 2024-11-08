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

// UserTweet holds the schema definition for the UserTweet entity.
type UserTweet struct {
	fluent.Schema
}

// Fields of the UserTweet.
func (UserTweet) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Int("user_id"),
		field.Int("tweet_id"),
	}
}

// Edges of the UserTweet.
func (UserTweet) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("tweet", Tweet.Type).
			Unique().
			Required().
			Field("tweet_id"),
	}
}

// Indexes of the UserTweet.
func (UserTweet) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("tweet_id").
			Unique(),
	}
}
