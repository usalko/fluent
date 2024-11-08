// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	fluent.Schema
}

// Mixin of the User.
func (User) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		VersionMixin{},
	}
}

// Fields of the User.
func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.Enum("status").
			Values("online", "offline"),
	}
}

// VersionMixin provides an optimistic concurrency
// control mechanism using a "version" field.
type VersionMixin struct {
	mixin.Schema
}

// Fields of the VersionMixin.
func (VersionMixin) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int64("version").
			DefaultFunc(func() int64 {
				return time.Now().UnixNano()
			}).
			Comment("Unix time of when the latest update occurred"),
	}
}
