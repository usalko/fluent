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

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	fluent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Default("unknown"),
		field.Int("owner_id").
			Optional(),
	}
}

// Edges of the Pet.
func (Pet) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Field("owner_id").
			Unique(),
	}
}

// Annotations of the Pet.
func (Pet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_sql.Schema("db2"),
	}
}
