// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Mixin of the User schema.
func (User) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Default("Unknown"),
		field.Strings("foods").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("groups", Group.Type),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() fluent.Policy {
	// Privacy policy defined in the BaseMixin and TenantMixin.
	return nil
}
