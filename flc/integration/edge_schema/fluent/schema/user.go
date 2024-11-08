// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/privacy"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Default("Unknown"),
	}
}

// Edges of the User.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("groups", Group.Type).
			Through("joined_groups", UserGroup.Type),
		edge.To("friends", User.Type).
			Through("friendships", Friendship.Type),
		edge.To("relatives", User.Type).
			Through("relationship", Relationship.Type),
		edge.To("liked_tweets", Tweet.Type).
			Through("likes", TweetLike.Type),
		edge.To("tweets", Tweet.Type).
			Through("user_tweets", UserTweet.Type),
		edge.To("roles", Role.Type).
			Through("roles_users", RoleUser.Type),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() fluent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
