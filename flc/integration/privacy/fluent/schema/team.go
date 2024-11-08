// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/privacy"
	"github.com/usalko/fluent/flc/integration/privacy/rule"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/mixin"
)

// Team defines the schema of a team.
type Team struct {
	fluent.Schema
}

// Mixin list of schemas to the team.
func (Team) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		BaseMixin{},
	}
}

// Fields of the team.
func (Team) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the team.
func (Team) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("tasks", Task.Type).
			Ref("teams"),
		edge.From("users", User.Type).
			Ref("teams"),
	}
}

// Policy of the team.
func (Team) Policy() fluent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNotAdmin(),
			rule.DenyUpdateRule(),
		},
	}
}

// TeamMixin shared between task and user.
type TeamMixin struct {
	mixin.Schema
}

// Edges of the team-mixin.
func (TeamMixin) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("teams", Team.Type),
	}
}

// Policy of the team-mixin.
func (TeamMixin) Policy() fluent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterTeamRule(),
		},
	}
}
