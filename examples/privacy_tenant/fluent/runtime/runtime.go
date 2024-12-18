// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/usalko/fluent/examples/privacy_tenant/fluent/group"
	"github.com/usalko/fluent/examples/privacy_tenant/fluent/schema"
	"github.com/usalko/fluent/examples/privacy_tenant/fluent/tenant"
	"github.com/usalko/fluent/examples/privacy_tenant/fluent/user"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	groupMixin := schema.Group{}.Mixin()
	group.Policy = privacy.NewPolicies(groupMixin[0], groupMixin[1], schema.Group{})
	group.Hooks[0] = func(next fluent.Mutator) fluent.Mutator {
		return fluent.MutateFunc(func(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
			if err := group.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.DefaultName holds the default value on creation for the name field.
	group.DefaultName = groupDescName.Default.(string)
	tenantMixin := schema.Tenant{}.Mixin()
	tenant.Policy = privacy.NewPolicies(tenantMixin[0], schema.Tenant{})
	tenant.Hooks[0] = func(next fluent.Mutator) fluent.Mutator {
		return fluent.MutateFunc(func(ctx context.Context, m fluent.Mutation) (fluent.Value, error) {
			if err := tenant.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	tenantFields := schema.Tenant{}.Fields()
	_ = tenantFields
	// tenantDescName is the schema descriptor for name field.
	tenantDescName := tenantFields[0].Descriptor()
	// tenant.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tenant.NameValidator = tenantDescName.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	user.Policy = privacy.NewPolicies(userMixin[0], userMixin[1], schema.User{})
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
	Version = "v0.1.7" // Version of fluent codegen.
)
