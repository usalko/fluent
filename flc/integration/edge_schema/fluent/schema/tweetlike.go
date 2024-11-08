// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/privacy"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// TweetLike holds the schema definition for the TweetLike entity.
type TweetLike struct {
	fluent.Schema
}

func (TweetLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "tweet_id"),
	}
}

// Fields of the TweetLike.
func (TweetLike) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("liked_at").
			Default(time.Now),
		field.Int("user_id"),
		field.Int("tweet_id"),
	}
}

// Edges of the TweetLike.
func (TweetLike) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("tweet", Tweet.Type).
			Unique().
			Required().
			Field("tweet_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}

// Policy defines the privacy policy of the TweetLike.
func (TweetLike) Policy() fluent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
