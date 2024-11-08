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

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	fluent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
	}
}

// PetUserName represents a user/pet name returned from the pet_user_names view.
type PetUserName struct {
	fluent.View
}

// Annotations of the PetUserName.
func (PetUserName) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// The definition below returns all names in the system
		// by unifying the "users" table and the "pets" table.
		fluent_sql.ViewFor(dialect.Postgres, func(s *sql.Selector) {
			s.SelectDistinct("name").
				From(
					s.New().Select("name").From(sql.Table("users")).
						Union(
							s.New().Select("name").From(sql.Table("pets")),
						).
						As("all_names"),
				)
		}),
		// Alternatively, you can use raw definitions to define the view.
		// But note, this definition is skipped if the ViewFor annotation
		// is defined for the dialect we generated migration to (Postgres).
		fluent_sql.View(`SELECT DISTINCT name
FROM (
    SELECT users.name
    FROM users
    UNION
    SELECT pets.name
    FROM pets
) AS all_names;
`),
	}
}

// Fields of the PetUserName.
func (PetUserName) Fields() []fluent.Field {
	return []fluent.Field{
		// Skip adding the "id" field as
		// it is not needed for this view.
		field.String("name"),
	}
}
