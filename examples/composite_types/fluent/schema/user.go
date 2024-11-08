// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"database/sql/driver"
	"fmt"

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
		field.String("address").
			GoType(&Address{}).
			SchemaType(map[string]string{
				dialect.Postgres: "address",
			}),
	}
}

type Address struct {
	Street, City string
}

var _ field.ValueScanner = (*Address)(nil)

// Scan implements the database/sql.Scanner interface.
func (a *Address) Scan(v interface{}) (err error) {
	switch v := v.(type) {
	case nil:
	case string:
		_, err = fmt.Sscanf(v, "(%q,%q)", &a.Street, &a.City)
	case []byte:
		_, err = fmt.Sscanf(string(v), "(%q,%q)", &a.Street, &a.City)
	}
	return
}

// Value implements the driver.Valuer interface.
func (a *Address) Value() (driver.Value, error) {
	return fmt.Sprintf("(%q,%q)", a.Street, a.City), nil
}
