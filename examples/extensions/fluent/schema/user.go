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
		field.Bytes("location").
			// Ideally, we would use a custom GoType
			// to represent the "geometry" type.
			SchemaType(map[string]string{
				dialect.Postgres: "GEOMETRY(Point, 4326)",
			}),
	}
}
