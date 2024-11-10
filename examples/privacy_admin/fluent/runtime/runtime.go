// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/usalko/fluent/examples/privacy_admin/fluent/schema"
	"github.com/usalko/fluent/examples/privacy_admin/fluent/user"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	user.Policy = privacy.NewPolicies(schema.User{})
	user.Hooks[0] = func(next fluent.Mutator) fluent.Mutator {
		return fluent.MutateFunc(func(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}

const (
	Version = "v0.1.6" // Version of fluent codegen.
)
