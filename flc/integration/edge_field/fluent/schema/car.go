// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	fluent.Schema
}

// Fields of the Car.
func (Car) Fields() []fluent.Field {
	return []fluent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("number").
			Optional(),
	}
}

// Edges of the Car.
func (Car) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("rentals", Rental.Type),
	}
}
