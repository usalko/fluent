// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package failure

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
)

type User struct {
	fluent.Schema
}

func (User) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("panic", User{}.Type),
	}
}
