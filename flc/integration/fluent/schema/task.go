// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"

	"github.com/usalko/fluent/flc/integration/fluent/schema/task"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	fluent.Schema
}

// Fields of the Task.
func (Task) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("priority").
			GoType(task.Priority(0)).
			Default(int(task.PriorityMid)),
		field.JSON("priorities", map[string]task.Priority{}).
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Nillable(),
		field.String("name").
			Optional().
			Deprecated(),
		field.String("owner").
			Optional(),
		field.Int("order").
			Optional(),
		field.Int("order_option").
			Optional(),
		field.String("op").
			MaxLen(45).
			Default(""),
	}
}

// Indexes of the Task.
func (Task) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("name", "owner").
			Unique(),
	}
}
