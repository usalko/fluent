// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// An example for an ent/schema that uses struct embedding to define default table schema (in the database).

// Group holds the schema definition for the Group entity.
type Group struct {
	base
}

// Fields of the Group.
func (Group) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the Group.
func (Group) Edges() []fluent.Edge {
	return []fluent.Edge{
		// The edge schema (join table) is defaults to the edge owner schema (Group).
		edge.To("users", User.Type),
	}
}
