// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/mixin"
)

var (
	incrementalDisabled = false
)

type Mixin struct {
	mixin.Schema
}

// Annotations of the Mixin schema.
func (Mixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_sql.Annotation{Charset: "utf8mb4"},
	}
}

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Fields of the User schema.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("id").
			StorageKey("user_id").
			Annotations(fluent_sql.Annotation{
				Incremental: &incrementalDisabled,
			}),
		field.String("name").
			Optional().
			Annotations(fluent_sql.Annotation{
				Size: 128,
			}).Comment(`Name of the user.
Comment line1
Comment line2`),
		field.String("label").
			Optional(),
	}
}

// Mixin of the User schema.
func (User) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		Mixin{},
	}
}

// Annotations of the User schema.
func (User) Annotations() []schema.Annotation {
	incremental := false
	return []schema.Annotation{
		fluent_sql.Annotation{
			Table:       "Users",
			Incremental: &incremental,
		},
	}
}
