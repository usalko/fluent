// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
	"github.com/usalko/fluent/schema/mixin"
)

// BaseMixin holds the schema definition for the BaseMixin entity.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the Mixin.
func (BaseMixin) Fields() []fluent.Field {
	return []fluent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("some_field"),
	}
}

// MixinID holds the schema definition for the MixinID entity.
type MixinID struct {
	fluent.Schema
}

// Fields of the MixinID.
func (MixinID) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("mixin_field"),
	}
}

// Indexes of the MixinID
func (MixinID) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("id"),
		index.Fields("id", "some_field"),
		index.Fields("id", "mixin_field"),
		index.Fields("id", "mixin_field", "some_field"),
	}
}

// Mixin of MixinID
func (MixinID) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		BaseMixin{},
	}
}
