// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect"
	"github.com/usalko/fluent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("postal_code").
			SchemaType(map[string]string{
				// Set the database column type to "us_postal_code" only in PostgreSQL.
				// In case this schema is used with other databases, it falls back to the
				// default type (e.g., "varchar").
				dialect.Postgres: "us_postal_code",
			}),
	}
}
