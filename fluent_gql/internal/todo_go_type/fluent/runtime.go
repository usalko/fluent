// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by flc, DO NOT EDIT.

package fluent

import (
	"time"

	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/category"
	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/friendship"
	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/group"
	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/pet"
	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/schema"
	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/schema/bigintgql"
	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/todo"
	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/user"
	"github.com/usalko/fluent/fluent_gql/internal/todo_go_type/fluent/very_secret"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryMixin := schema.Category{}.Mixin()
	categoryMixinFields0 := categoryMixin[0].Fields()
	_ = categoryMixinFields0
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescText is the schema descriptor for text field.
	categoryDescText := categoryMixinFields0[1].Descriptor()
	// category.TextValidator is a validator for the "text" field. It is called by the builders before save.
	category.TextValidator = categoryDescText.Validators[0].(func(string) error)
	// categoryDescID is the schema descriptor for id field.
	categoryDescID := categoryFields[0].Descriptor()
	// category.DefaultID holds the default value on creation for the id field.
	category.DefaultID = categoryDescID.Default.(func() bigintgql.BigInt)
	friendshipFields := schema.Friendship{}.Fields()
	_ = friendshipFields
	// friendshipDescCreatedAt is the schema descriptor for created_at field.
	friendshipDescCreatedAt := friendshipFields[0].Descriptor()
	// friendship.DefaultCreatedAt holds the default value on creation for the created_at field.
	friendship.DefaultCreatedAt = friendshipDescCreatedAt.Default.(func() time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[1].Descriptor()
	// group.DefaultName holds the default value on creation for the name field.
	group.DefaultName = groupDescName.Default.(string)
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescName is the schema descriptor for name field.
	petDescName := petFields[1].Descriptor()
	// pet.DefaultName holds the default value on creation for the name field.
	pet.DefaultName = petDescName.Default.(string)
	todoMixin := schema.Todo{}.Mixin()
	todoMixinFields0 := todoMixin[0].Fields()
	_ = todoMixinFields0
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoMixinFields0[0].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescPriority is the schema descriptor for priority field.
	todoDescPriority := todoMixinFields0[2].Descriptor()
	// todo.DefaultPriority holds the default value on creation for the priority field.
	todo.DefaultPriority = todoDescPriority.Default.(int)
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoMixinFields0[3].Descriptor()
	// todo.TextValidator is a validator for the "text" field. It is called by the builders before save.
	todo.TextValidator = todoDescText.Validators[0].(func(string) error)
	// todoDescValue is the schema descriptor for value field.
	todoDescValue := todoMixinFields0[8].Descriptor()
	// todo.DefaultValue holds the default value on creation for the value field.
	todo.DefaultValue = todoDescValue.Default.(int)
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoFields[0].Descriptor()
	// todo.DefaultID holds the default value on creation for the id field.
	todo.DefaultID = todoDescID.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	very_secretFields := schema.VerySecret{}.Fields()
	_ = very_secretFields
	// very_secretDescID is the schema descriptor for id field.
	very_secretDescID := very_secretFields[0].Descriptor()
	// very_secret.DefaultID holds the default value on creation for the id field.
	very_secret.DefaultID = very_secretDescID.Default.(string)
}
