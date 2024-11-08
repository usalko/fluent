// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema/field"
)

// Conversion holds the schema definition for the Conversion entity.
type Conversion struct {
	fluent.Schema
}

// Fields of the Conversion.
func (Conversion) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Optional(),
		// Convert integer fields to string
		// Postgres uses the same type for int8 and int16
		// Postgres loses unsigned so we have assume value is signed
		field.String("int8_to_string").
			Optional().
			Annotations(fluent_sql.Annotation{Size: 6}),
		field.String("uint8_to_string").
			Optional().
			Annotations(fluent_sql.Annotation{Size: 6}),
		field.String("int16_to_string").
			Optional().
			Annotations(fluent_sql.Annotation{Size: 6}),
		field.String("uint16_to_string").
			Optional().
			Annotations(fluent_sql.Annotation{Size: 6}),
		field.String("int32_to_string").
			Optional().
			Annotations(fluent_sql.Annotation{Size: 12}),
		field.String("uint32_to_string").
			Optional().
			Annotations(fluent_sql.Annotation{Size: 12}),
		field.String("int64_to_string").
			Optional().
			Annotations(fluent_sql.Annotation{Size: 21}),
		field.String("uint64_to_string").
			Optional().
			Annotations(fluent_sql.Annotation{Size: 21}),
	}
}
