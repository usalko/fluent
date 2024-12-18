// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	fluent.Schema
}

// Fields of the Tenant.
func (Tenant) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
	}
}

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
		field.Int("tenant_id"),
	}
}
