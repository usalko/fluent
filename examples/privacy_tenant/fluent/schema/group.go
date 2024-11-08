// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/examples/privacy_tenant/fluent/privacy"
	"github.com/usalko/fluent/examples/privacy_tenant/rule"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	fluent.Schema
}

// Mixin of the User schema.
func (Group) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the User.
func (Group) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Default("Unknown"),
	}
}

// Edges of the Group.
func (Group) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("users", User.Type).
			Ref("groups"),
	}
}

// Policy defines the privacy policy of the Group.
func (Group) Policy() fluent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Limit DenyMismatchedTenants only for
			// Create operations
			privacy.OnMutationOperation(
				rule.DenyMismatchedTenants(),
				fluent.OpCreate,
			),
		},
	}
}
