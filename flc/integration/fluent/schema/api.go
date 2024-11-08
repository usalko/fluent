// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import "github.com/usalko/fluent"

// Api represents an example schema with wrong usage of acronym (e.g., API).
// nolint:stylecheck
type Api struct {
	fluent.Schema
}
