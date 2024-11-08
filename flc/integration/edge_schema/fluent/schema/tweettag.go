// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// TweetTag holds the schema definition for the TweetTag entity.
type TweetTag struct {
	fluent.Schema
}

// Fields of the TweetTag.
func (TweetTag) Fields() []fluent.Field {
	return []fluent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("added_at").
			Default(time.Now),
		field.Int("tag_id"),
		field.Int("tweet_id"),
	}
}

// Edges of the TweetTag.
func (TweetTag) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("tag", Tag.Type).
			Unique().
			Required().
			Field("tag_id"),
		edge.To("tweet", Tweet.Type).
			Unique().
			Required().
			Field("tweet_id"),
	}
}
