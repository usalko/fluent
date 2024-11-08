// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package cycle

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/flc/load/testdata/cycle/fakent"
	"github.com/usalko/fluent/schema/field"
)

// User holds the user schema.
type User struct {
	fluent.Schema
}

func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.JSON("used", &Used{}),
		field.Enum("e").
			GoType(Enum(0)),
	}
}

type (
	Used        struct{}
	NotUsed     struct{}
	notExported struct{}
	Enum        int
)

func (Enum) Values() []string { return nil }

// The cause for cycle.
var _ fakent.Hook = nil
