// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/examples/privacy_tenant/fluent/privacy"
	"github.com/usalko/fluent/examples/privacy_tenant/rule"
	"github.com/usalko/fluent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	fluent.Schema
}

// Mixin of the Tenant schema.
func (Tenant) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			NotEmpty(),
	}
}

// Policy defines the privacy policy of the User.
func (Tenant) Policy() fluent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// For Tenant type, we only allow admin users to mutate
			// the tenant information and deny otherwise.
			rule.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
