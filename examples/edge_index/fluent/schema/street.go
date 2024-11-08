// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// Street holds the schema definition for the Street entity.
type Street struct {
	fluent.Schema
}

// Fields of the Street.
func (Street) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
	}
}

// Edges of the Street.
func (Street) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("city", City.Type).
			Ref("streets").
			Unique(),
	}
}

// Indexes of the Street.
func (Street) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("name").
			Edges("city").
			Unique(),
	}
}
