// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema/field"
)

// Zoo holds the schema definition for the Zoo entity.
type Zoo struct {
	fluent.Schema
}

// Fields of the Zoo.
func (Zoo) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("id").
			Annotations(
				fluent_sql.DefaultExprs(map[string]string{
					dialect.MySQL:    "floor(rand() * ~(1<<31))",
					dialect.SQLite:   "abs(random())",
					dialect.Postgres: "floor(random() * ~(1<<31))",
				}),
			),
	}
}
