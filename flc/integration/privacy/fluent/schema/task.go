// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/privacy"
	"github.com/usalko/fluent/flc/integration/privacy/rule"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Task defines the schema of a task.
type Task struct {
	fluent.Schema
}

// Mixin list of schemas to the task.
func (Task) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		BaseMixin{},
		TeamMixin{},
	}
}

// Fields of the task.
func (Task) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("title").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.Enum("status").
			Values("planned", "in_progress", "closed").
			Default("planned"),
		field.UUID("uuid", uuid.UUID{}).
			Optional(),
	}
}

// Edges of the task.
func (Task) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("owner", User.Type).
			Ref("tasks").
			Unique(),
	}
}

// Hooks for the task.
func (Task) Hooks() []fluent.Hook {
	return []fluent.Hook{
		rule.LogTaskMutationHook(),
	}
}

// Policy of the task.
func (Task) Policy() fluent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowTaskCreateIfOwner(),
			rule.DenyIfStatusChangedByOther(),
			rule.AllowIfViewerInTheSameTeam(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
