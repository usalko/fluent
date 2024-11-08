// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/mixin"
)

type IDMixin struct {
	mixin.Schema
}

func (IDMixin) Fields() []fluent.Field {
	return []fluent.Field{
		field.Float("before_id").
			Optional().
			Positive(),
		field.Int("id").
			Positive().
			Immutable(),
		field.Float("after_id").
			Optional().
			Negative(),
	}
}

// Car holds the schema definition for the Car entity.
type Car struct {
	fluent.Schema
}

// Mixin of the Car.
func (Car) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		IDMixin{},
	}
}

// Fields of the Car.
func (Car) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("model"),
	}
}

// Edges of the Car.
func (Car) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("owner", Pet.Type).
			Ref("cars").
			Unique(),
	}
}
