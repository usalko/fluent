// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/google/uuid"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// Rental holds the schema definition for the Rental entity.
type Rental struct {
	fluent.Schema
}

// Fields of the Rental.
func (Rental) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("date").
			Default(time.Now),
		field.Int("user_id").
			Immutable(),
		field.UUID("car_id", uuid.UUID{}).
			Immutable(),
	}
}

// Edges of the Rental.
func (Rental) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("user", User.Type).
			Ref("rentals").
			Field("user_id").
			Unique().
			Required().
			Immutable(),
		edge.From("car", Car.Type).
			Ref("rentals").
			Field("car_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the Rental.
func (Rental) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("car_id", "user_id").
			Unique(),
	}
}
