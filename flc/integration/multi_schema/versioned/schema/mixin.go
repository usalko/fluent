// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/mixin"
)

// This file contains two example of how to define the base schema.
// 1. Use Mixin and use it in all schemas that reside in "db1".
// 2. Create a "base" schema and use struct embedding to in all schemas that reside in "db1".

// Example 1:

// Mixin holds the default configuration for most schemas in this package.
type Mixin struct {
	mixin.Schema
}

// Annotations of the Mixin.
func (Mixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_sql.Schema("db1"),
	}
}

// Example 2:

// Base schema.
type base struct {
	fluent.Schema
}

// Annotations of the base schema.
func (base) Annotations() []schema.Annotation {
	return []schema.Annotation{
		fluent_sql.Schema("db1"),
	}
}
