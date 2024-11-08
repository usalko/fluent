// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	schemadir "github.com/usalko/fluent/flc/integration/fluent/schema/dir"
	"github.com/usalko/fluent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	fluent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("unique_int").
			Unique(),
		field.Float("unique_float").
			Unique(),
		field.Int("nillable_int").
			Optional().
			Nillable(),
		field.String("table").
			Optional(),
		field.JSON("dir", schemadir.Dir("")).
			Optional(),
		field.String("client").
			Optional(),
	}
}
