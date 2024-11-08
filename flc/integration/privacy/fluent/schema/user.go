// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/privacy"
	"github.com/usalko/fluent/flc/integration/privacy/rule"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// User defines the schema of a user.
type User struct {
	fluent.Schema
}

// Mixin list of schemas to the user.
func (User) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		BaseMixin{},
		TeamMixin{},
	}
}

// Fields of the user.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Immutable().
			NotEmpty().
			Unique(),
		field.Uint("age").
			Optional(),
	}
}

// Edges of the user.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("tasks", Task.Type),
	}
}

// Policy of the user.
func (User) Policy() fluent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowUserCreateIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

// Annotations of the user.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_sql.Checks(map[string]string{
			"backticks": "`name` IS NOT NULL",
		}),
	}
}
