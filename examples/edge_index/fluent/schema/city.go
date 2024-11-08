// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// City holds the schema definition for the City entity.
type City struct {
	fluent.Schema
}

// Fields of the City.
func (City) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
	}
}

// Edges of the City.
func (City) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("streets", Street.Type),
	}
}
