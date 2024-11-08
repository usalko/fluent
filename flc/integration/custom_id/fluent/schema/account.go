// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect"
	"github.com/usalko/fluent/flc/integration/custom_id/sid"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	fluent.Schema
}

// Fields of the Account.
func (Account) Fields() []fluent.Field {
	return []fluent.Field{
		field.Other("id", sid.ID("")).
			SchemaType(map[string]string{
				dialect.MySQL:    "bigint",
				dialect.Postgres: "bigint",
				dialect.SQLite:   "integer",
			}).
			Default(sid.New).
			Immutable(),
		field.String("email").NotEmpty(),
	}
}

// Edges of the Account.
func (Account) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("token", Token.Type),
	}
}
