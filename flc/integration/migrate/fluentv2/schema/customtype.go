// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect"
	"github.com/usalko/fluent/schema/field"
)

// CustomType holds the schema definition for the CustomType entity.
type CustomType struct {
	fluent.Schema
}

// Fields of the CustomType.
func (CustomType) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("custom").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "customtype",
			}),
		field.Time("tz0").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "timestamp(0)",
				dialect.Postgres: "timestamptz(0)",
			}),
		field.Time("tz3").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "timestamp(3)",
				dialect.Postgres: "timestamptz(3)",
			}),
	}
}
