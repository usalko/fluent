// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
		field.String("public_info"),
		field.String("private_info"),
	}
}

// CleanUser represents a user without its PII field.
type CleanUser struct {
	fluent.View
}

// Annotations of the CleanUser.
func (CleanUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_sql.ViewFor(dialect.Postgres, func(s *sql.Selector) {
			s.Select("id", "name", "public_info").From(sql.Table("users"))
		}),
	}
}

// Fields of the CleanUser.
func (CleanUser) Fields() []fluent.Field {
	return []fluent.Field{
		// Note, unlike real schemas (tables, defined with fluent.Schema),
		// the "id" field should be defined manually if needed.
		field.Int("id"),
		field.String("name"),
		field.String("public_info"),
	}
}
