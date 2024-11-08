// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package fakent fis a fake generated Ent package.
package fakent

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/flc/load/testdata/cycle"
)

type Hook = fluent.Hook

var _ = &cycle.Used{}
