// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/usalko/fluent/flc/integration/hooks/fluent/card"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/pet"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/schema"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cardMixin := schema.Card{}.Mixin()
	cardMixinHooks0 := cardMixin[0].Hooks()
	cardHooks := schema.Card{}.Hooks()
	card.Hooks[0] = cardMixinHooks0[0]
	card.Hooks[1] = cardHooks[0]
	card.Hooks[2] = cardHooks[1]
	cardInters := schema.Card{}.Interceptors()
	card.Interceptors[0] = cardInters[0]
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescNumber is the schema descriptor for number field.
	cardDescNumber := cardFields[0].Descriptor()
	// card.DefaultNumber holds the default value on creation for the number field.
	card.DefaultNumber = cardDescNumber.Default.(string)
	// card.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	card.NumberValidator = cardDescNumber.Validators[0].(func(string) error)
	// cardDescCreatedAt is the schema descriptor for created_at field.
	cardDescCreatedAt := cardFields[2].Descriptor()
	// card.DefaultCreatedAt holds the default value on creation for the created_at field.
	card.DefaultCreatedAt = cardDescCreatedAt.Default.(func() time.Time)
	petMixin := schema.Pet{}.Mixin()
	petMixinHooks0 := petMixin[0].Hooks()
	pet.Hooks[0] = petMixinHooks0[0]
	petMixinInters0 := petMixin[0].Interceptors()
	pet.Interceptors[0] = petMixinInters0[0]
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	user.Hooks[1] = userHooks[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescVersion is the schema descriptor for version field.
	userDescVersion := userMixinFields0[0].Descriptor()
	// user.DefaultVersion holds the default value on creation for the version field.
	user.DefaultVersion = userDescVersion.Default.(int)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[3].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
}

const (
	Version = "v0.1.5" // Version of fluent codegen.
)
