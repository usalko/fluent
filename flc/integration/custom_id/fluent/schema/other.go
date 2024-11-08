// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect"
	"github.com/usalko/fluent/flc/integration/custom_id/sid"
	"github.com/usalko/fluent/schema/field"
)

// Other holds the schema definition for the Other entity.
type Other struct {
	fluent.Schema
}

// Fields of the Other.
func (Other) Fields() []fluent.Field {
	return []fluent.Field{
		field.Other("id", sid.ID("")).
			SchemaType(map[string]string{
				dialect.MySQL:    "bigint",
				dialect.Postgres: "bigint",
				dialect.SQLite:   "integer",
			}).
			Default(sid.New).
			Immutable(),
	}
}
