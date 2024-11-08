// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
	}
}

// UserAuditLog holds the schema definition for the UserAuditLog entity.
type UserAuditLog struct {
	fluent.Schema
}

// Fields of the UserAuditLog.
func (UserAuditLog) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("operation_type"),
		field.String("operation_time"),
		field.String("old_value").
			Optional(),
		field.String("new_value").
			Optional(),
	}
}
