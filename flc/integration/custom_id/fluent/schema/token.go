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

// Token holds the schema definition for the Token entity.
type Token struct {
	fluent.Schema
}

// Fields of the Token.
func (Token) Fields() []fluent.Field {
	return []fluent.Field{
		field.Other("id", sid.ID("")).
			SchemaType(map[string]string{
				dialect.MySQL:    "bigint",
				dialect.Postgres: "bigint",
				dialect.SQLite:   "integer",
			}).
			Default(sid.New).
			Immutable(),
		field.String("body").NotEmpty(),
	}
}

// Edges of the Token.
func (Token) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("account", Account.Type).Ref("token").Required().Unique(),
	}
}
