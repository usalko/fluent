// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.Float("age"),
		field.String("first_name"),
		field.String("last_name"),
		field.Strings("tags").
			Optional(),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// In case schema.ModeInspect is used without a dev-database, unnamed check constraints
		// should be normalized (i.e. identical to their definition in the database). In this
		// case, it is fluent_sql.Check("(`age` > 0)"). See: https://atlasgo.io/concepts/dev-database.
		fluent_sql.Check("age > 0"),

		// Named check constraints are compared by their name.
		// Thus, the definition does not need to be normalized.
		fluent_sql.Checks(map[string]string{
			"first_last_not_empty": "(`first_name` <> '' AND `last_name` <> '')",
		}),
	}
}

// Edges of the User.
func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("cards", Card.Type),
	}
}
