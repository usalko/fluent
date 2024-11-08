// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	fluent.Schema
}

// Fields of the Group.
func (Group) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
	}
}

// Annotations of the Group.
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_sql.Annotation{
			Table:     "versioned_groups",
			Charset:   "ascii",
			Collation: "ascii_general_ci",
		},
	}
}
