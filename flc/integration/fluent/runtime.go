// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"math/big"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/flc/integration/fluent/card"
	"github.com/usalko/fluent/flc/integration/fluent/exvaluescan"
	"github.com/usalko/fluent/flc/integration/fluent/fieldtype"
	"github.com/usalko/fluent/flc/integration/fluent/file"
	"github.com/usalko/fluent/flc/integration/fluent/group"
	"github.com/usalko/fluent/flc/integration/fluent/groupinfo"
	"github.com/usalko/fluent/flc/integration/fluent/item"
	"github.com/usalko/fluent/flc/integration/fluent/license"
	"github.com/usalko/fluent/flc/integration/fluent/node"
	"github.com/usalko/fluent/flc/integration/fluent/pet"
	"github.com/usalko/fluent/flc/integration/fluent/schema"
	"github.com/usalko/fluent/flc/integration/fluent/schema/task"
	enttask "github.com/usalko/fluent/flc/integration/fluent/task"
	"github.com/usalko/fluent/flc/integration/fluent/user"

	"github.com/usalko/fluent/schema/field"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cardMixin := schema.Card{}.Mixin()
	cardMixinFields0 := cardMixin[0].Fields()
	_ = cardMixinFields0
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescCreateTime is the schema descriptor for create_time field.
	cardDescCreateTime := cardMixinFields0[0].Descriptor()
	// card.DefaultCreateTime holds the default value on creation for the create_time field.
	card.DefaultCreateTime = cardDescCreateTime.Default.(func() time.Time)
	// cardDescUpdateTime is the schema descriptor for update_time field.
	cardDescUpdateTime := cardMixinFields0[1].Descriptor()
	// card.DefaultUpdateTime holds the default value on creation for the update_time field.
	card.DefaultUpdateTime = cardDescUpdateTime.Default.(func() time.Time)
	// card.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	card.UpdateDefaultUpdateTime = cardDescUpdateTime.UpdateDefault.(func() time.Time)
	// cardDescBalance is the schema descriptor for balance field.
	cardDescBalance := cardFields[0].Descriptor()
	// card.DefaultBalance holds the default value on creation for the balance field.
	card.DefaultBalance = cardDescBalance.Default.(float64)
	// cardDescNumber is the schema descriptor for number field.
	cardDescNumber := cardFields[1].Descriptor()
	// card.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	card.NumberValidator = cardDescNumber.Validators[0].(func(string) error)
	// cardDescName is the schema descriptor for name field.
	cardDescName := cardFields[2].Descriptor()
	// card.NameValidator is a validator for the "name" field. It is called by the builders before save.
	card.NameValidator = cardDescName.Validators[0].(func(string) error)
	exvaluescanFields := schema.ExValueScan{}.Fields()
	_ = exvaluescanFields
	// exvaluescanDescBinary is the schema descriptor for binary field.
	exvaluescanDescBinary := exvaluescanFields[0].Descriptor()
	exvaluescan.ValueScanner.Binary = exvaluescanDescBinary.ValueScanner.(field.TypeValueScanner[*url.URL])
	// exvaluescanDescBinaryBytes is the schema descriptor for binary_bytes field.
	exvaluescanDescBinaryBytes := exvaluescanFields[1].Descriptor()
	exvaluescan.ValueScanner.BinaryBytes = exvaluescanDescBinaryBytes.ValueScanner.(field.TypeValueScanner[*url.URL])
	// exvaluescanDescBinaryOptional is the schema descriptor for binary_optional field.
	exvaluescanDescBinaryOptional := exvaluescanFields[2].Descriptor()
	exvaluescan.ValueScanner.BinaryOptional = exvaluescanDescBinaryOptional.ValueScanner.(field.TypeValueScanner[*url.URL])
	// exvaluescanDescText is the schema descriptor for text field.
	exvaluescanDescText := exvaluescanFields[3].Descriptor()
	exvaluescan.ValueScanner.Text = exvaluescanDescText.ValueScanner.(field.TypeValueScanner[*big.Int])
	// exvaluescanDescTextOptional is the schema descriptor for text_optional field.
	exvaluescanDescTextOptional := exvaluescanFields[4].Descriptor()
	exvaluescan.ValueScanner.TextOptional = exvaluescanDescTextOptional.ValueScanner.(field.TypeValueScanner[*big.Int])
	// exvaluescanDescBase64 is the schema descriptor for base64 field.
	exvaluescanDescBase64 := exvaluescanFields[5].Descriptor()
	exvaluescan.ValueScanner.Base64 = exvaluescanDescBase64.ValueScanner.(field.TypeValueScanner[string])
	// exvaluescanDescCustom is the schema descriptor for custom field.
	exvaluescanDescCustom := exvaluescanFields[6].Descriptor()
	exvaluescan.ValueScanner.Custom = exvaluescanDescCustom.ValueScanner.(field.TypeValueScanner[string])
	// exvaluescanDescCustomOptional is the schema descriptor for custom_optional field.
	exvaluescanDescCustomOptional := exvaluescanFields[7].Descriptor()
	exvaluescan.ValueScanner.CustomOptional = exvaluescanDescCustomOptional.ValueScanner.(field.TypeValueScanner[string])
	fieldtypeFields := schema.FieldType{}.Fields()
	_ = fieldtypeFields
	// fieldtypeDescInt64 is the schema descriptor for int64 field.
	fieldtypeDescInt64 := fieldtypeFields[4].Descriptor()
	// fieldtype.UpdateDefaultInt64 holds the default value on update for the int64 field.
	fieldtype.UpdateDefaultInt64 = fieldtypeDescInt64.UpdateDefault.(func() int64)
	// fieldtypeDescValidateOptionalInt32 is the schema descriptor for validate_optional_int32 field.
	fieldtypeDescValidateOptionalInt32 := fieldtypeFields[15].Descriptor()
	// fieldtype.ValidateOptionalInt32Validator is a validator for the "validate_optional_int32" field. It is called by the builders before save.
	fieldtype.ValidateOptionalInt32Validator = fieldtypeDescValidateOptionalInt32.Validators[0].(func(int32) error)
	// fieldtypeDescLinkOther is the schema descriptor for link_other field.
	fieldtypeDescLinkOther := fieldtypeFields[27].Descriptor()
	// fieldtype.DefaultLinkOther holds the default value on creation for the link_other field.
	fieldtype.DefaultLinkOther = fieldtypeDescLinkOther.Default.(*schema.Link)
	// fieldtypeDescLinkOtherFunc is the schema descriptor for link_other_func field.
	fieldtypeDescLinkOtherFunc := fieldtypeFields[28].Descriptor()
	// fieldtype.DefaultLinkOtherFunc holds the default value on creation for the link_other_func field.
	fieldtype.DefaultLinkOtherFunc = fieldtypeDescLinkOtherFunc.Default.(func() *schema.Link)
	// fieldtypeDescMAC is the schema descriptor for mac field.
	fieldtypeDescMAC := fieldtypeFields[29].Descriptor()
	// fieldtype.MACValidator is a validator for the "mac" field. It is called by the builders before save.
	fieldtype.MACValidator = fieldtypeDescMAC.Validators[0].(func(string) error)
	// fieldtypeDescDuration is the schema descriptor for duration field.
	fieldtypeDescDuration := fieldtypeFields[33].Descriptor()
	// fieldtype.UpdateDefaultDuration holds the default value on update for the duration field.
	fieldtype.UpdateDefaultDuration = fieldtypeDescDuration.UpdateDefault.(func() time.Duration)
	// fieldtypeDescDir is the schema descriptor for dir field.
	fieldtypeDescDir := fieldtypeFields[34].Descriptor()
	// fieldtype.DefaultDir holds the default value on creation for the dir field.
	fieldtype.DefaultDir = fieldtypeDescDir.Default.(func() http.Dir)
	// fieldtypeDescNdir is the schema descriptor for ndir field.
	fieldtypeDescNdir := fieldtypeFields[35].Descriptor()
	// fieldtype.NdirValidator is a validator for the "ndir" field. It is called by the builders before save.
	fieldtype.NdirValidator = fieldtypeDescNdir.Validators[0].(func(string) error)
	// fieldtypeDescStr is the schema descriptor for str field.
	fieldtypeDescStr := fieldtypeFields[36].Descriptor()
	// fieldtype.DefaultStr holds the default value on creation for the str field.
	fieldtype.DefaultStr = fieldtypeDescStr.Default.(func() sql.NullString)
	// fieldtypeDescNullStr is the schema descriptor for null_str field.
	fieldtypeDescNullStr := fieldtypeFields[37].Descriptor()
	// fieldtype.DefaultNullStr holds the default value on creation for the null_str field.
	fieldtype.DefaultNullStr = fieldtypeDescNullStr.Default.(func() *sql.NullString)
	// fieldtypeDescLink is the schema descriptor for link field.
	fieldtypeDescLink := fieldtypeFields[38].Descriptor()
	// fieldtype.LinkValidator is a validator for the "link" field. It is called by the builders before save.
	fieldtype.LinkValidator = fieldtypeDescLink.Validators[0].(func(string) error)
	// fieldtypeDescDeletedAt is the schema descriptor for deleted_at field.
	fieldtypeDescDeletedAt := fieldtypeFields[43].Descriptor()
	// fieldtype.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	fieldtype.DefaultDeletedAt = fieldtypeDescDeletedAt.Default.(func() *sql.NullTime)
	// fieldtype.UpdateDefaultDeletedAt holds the default value on update for the deleted_at field.
	fieldtype.UpdateDefaultDeletedAt = fieldtypeDescDeletedAt.UpdateDefault.(func() *sql.NullTime)
	// fieldtypeDescRawData is the schema descriptor for raw_data field.
	fieldtypeDescRawData := fieldtypeFields[44].Descriptor()
	// fieldtype.RawDataValidator is a validator for the "raw_data" field. It is called by the builders before save.
	fieldtype.RawDataValidator = func() func([]byte) error {
		validators := fieldtypeDescRawData.Validators
		fns := [...]func([]byte) error{
			validators[0].(func([]byte) error),
			validators[1].(func([]byte) error),
		}
		return func(raw_data []byte) error {
			for _, fn := range fns {
				if err := fn(raw_data); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// fieldtypeDescIP is the schema descriptor for ip field.
	fieldtypeDescIP := fieldtypeFields[46].Descriptor()
	// fieldtype.DefaultIP holds the default value on creation for the ip field.
	fieldtype.DefaultIP = fieldtypeDescIP.Default.(func() net.IP)
	// fieldtype.IPValidator is a validator for the "ip" field. It is called by the builders before save.
	fieldtype.IPValidator = fieldtypeDescIP.Validators[0].(func([]byte) error)
	// fieldtypeDescPair is the schema descriptor for pair field.
	fieldtypeDescPair := fieldtypeFields[59].Descriptor()
	// fieldtype.DefaultPair holds the default value on creation for the pair field.
	fieldtype.DefaultPair = fieldtypeDescPair.Default.(func() schema.Pair)
	// fieldtypeDescVstring is the schema descriptor for vstring field.
	fieldtypeDescVstring := fieldtypeFields[61].Descriptor()
	// fieldtype.DefaultVstring holds the default value on creation for the vstring field.
	fieldtype.DefaultVstring = fieldtypeDescVstring.Default.(func() schema.VString)
	// fieldtypeDescTriple is the schema descriptor for triple field.
	fieldtypeDescTriple := fieldtypeFields[62].Descriptor()
	// fieldtype.DefaultTriple holds the default value on creation for the triple field.
	fieldtype.DefaultTriple = fieldtypeDescTriple.Default.(func() schema.Triple)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescSetID is the schema descriptor for set_id field.
	fileDescSetID := fileFields[0].Descriptor()
	// file.SetIDValidator is a validator for the "set_id" field. It is called by the builders before save.
	file.SetIDValidator = fileDescSetID.Validators[0].(func(int) error)
	// fileDescSize is the schema descriptor for size field.
	fileDescSize := fileFields[1].Descriptor()
	// file.DefaultSize holds the default value on creation for the size field.
	file.DefaultSize = fileDescSize.Default.(int)
	// file.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	file.SizeValidator = fileDescSize.Validators[0].(func(int) error)
	filetypeFields := schema.FileType{}.Fields()
	_ = filetypeFields
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescActive is the schema descriptor for active field.
	groupDescActive := groupFields[0].Descriptor()
	// group.DefaultActive holds the default value on creation for the active field.
	group.DefaultActive = groupDescActive.Default.(bool)
	// groupDescType is the schema descriptor for type field.
	groupDescType := groupFields[2].Descriptor()
	// group.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	group.TypeValidator = func() func(string) error {
		validators := groupDescType.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_type string) error {
			for _, fn := range fns {
				if err := fn(_type); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// groupDescMaxUsers is the schema descriptor for max_users field.
	groupDescMaxUsers := groupFields[3].Descriptor()
	// group.DefaultMaxUsers holds the default value on creation for the max_users field.
	group.DefaultMaxUsers = groupDescMaxUsers.Default.(int)
	// group.MaxUsersValidator is a validator for the "max_users" field. It is called by the builders before save.
	group.MaxUsersValidator = groupDescMaxUsers.Validators[0].(func(int) error)
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[4].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = func() func(string) error {
		validators := groupDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	groupinfoFields := schema.GroupInfo{}.Fields()
	_ = groupinfoFields
	// groupinfoDescMaxUsers is the schema descriptor for max_users field.
	groupinfoDescMaxUsers := groupinfoFields[1].Descriptor()
	// groupinfo.DefaultMaxUsers holds the default value on creation for the max_users field.
	groupinfo.DefaultMaxUsers = groupinfoDescMaxUsers.Default.(int)
	itemFields := schema.Item{}.Fields()
	_ = itemFields
	// itemDescText is the schema descriptor for text field.
	itemDescText := itemFields[1].Descriptor()
	// item.TextValidator is a validator for the "text" field. It is called by the builders before save.
	item.TextValidator = itemDescText.Validators[0].(func(string) error)
	// itemDescID is the schema descriptor for id field.
	itemDescID := itemFields[0].Descriptor()
	// item.DefaultID holds the default value on creation for the id field.
	item.DefaultID = itemDescID.Default.(func() string)
	// item.IDValidator is a validator for the "id" field. It is called by the builders before save.
	item.IDValidator = itemDescID.Validators[0].(func(string) error)
	licenseMixin := schema.License{}.Mixin()
	licenseMixinFields0 := licenseMixin[0].Fields()
	_ = licenseMixinFields0
	licenseFields := schema.License{}.Fields()
	_ = licenseFields
	// licenseDescCreateTime is the schema descriptor for create_time field.
	licenseDescCreateTime := licenseMixinFields0[0].Descriptor()
	// license.DefaultCreateTime holds the default value on creation for the create_time field.
	license.DefaultCreateTime = licenseDescCreateTime.Default.(func() time.Time)
	// licenseDescUpdateTime is the schema descriptor for update_time field.
	licenseDescUpdateTime := licenseMixinFields0[1].Descriptor()
	// license.DefaultUpdateTime holds the default value on creation for the update_time field.
	license.DefaultUpdateTime = licenseDescUpdateTime.Default.(func() time.Time)
	// license.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	license.UpdateDefaultUpdateTime = licenseDescUpdateTime.UpdateDefault.(func() time.Time)
	nodeFields := schema.Node{}.Fields()
	_ = nodeFields
	// nodeDescUpdatedAt is the schema descriptor for updated_at field.
	nodeDescUpdatedAt := nodeFields[1].Descriptor()
	// node.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	node.UpdateDefaultUpdatedAt = nodeDescUpdatedAt.UpdateDefault.(func() time.Time)
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescAge is the schema descriptor for age field.
	petDescAge := petFields[0].Descriptor()
	// pet.DefaultAge holds the default value on creation for the age field.
	pet.DefaultAge = petDescAge.Default.(float64)
	// petDescTrained is the schema descriptor for trained field.
	petDescTrained := petFields[4].Descriptor()
	// pet.DefaultTrained holds the default value on creation for the trained field.
	pet.DefaultTrained = petDescTrained.Default.(bool)
	enttaskFields := schema.Task{}.Fields()
	_ = enttaskFields
	// enttaskDescPriority is the schema descriptor for priority field.
	enttaskDescPriority := enttaskFields[0].Descriptor()
	// enttask.DefaultPriority holds the default value on creation for the priority field.
	enttask.DefaultPriority = task.Priority(enttaskDescPriority.Default.(int))
	// enttaskDescCreatedAt is the schema descriptor for created_at field.
	enttaskDescCreatedAt := enttaskFields[2].Descriptor()
	// enttask.DefaultCreatedAt holds the default value on creation for the created_at field.
	enttask.DefaultCreatedAt = enttaskDescCreatedAt.Default.(func() time.Time)
	// enttaskDescOp is the schema descriptor for op field.
	enttaskDescOp := enttaskFields[7].Descriptor()
	// enttask.DefaultOp holds the default value on creation for the op field.
	enttask.DefaultOp = enttaskDescOp.Default.(string)
	// enttask.OpValidator is a validator for the "op" field. It is called by the builders before save.
	enttask.OpValidator = enttaskDescOp.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescOptionalInt is the schema descriptor for optional_int field.
	userDescOptionalInt := userMixinFields0[0].Descriptor()
	// user.OptionalIntValidator is a validator for the "optional_int" field. It is called by the builders before save.
	user.OptionalIntValidator = userDescOptionalInt.Validators[0].(func(int) error)
	// userDescLast is the schema descriptor for last field.
	userDescLast := userFields[2].Descriptor()
	// user.DefaultLast holds the default value on creation for the last field.
	user.DefaultLast = userDescLast.Default.(string)
	// userDescAddress is the schema descriptor for address field.
	userDescAddress := userFields[4].Descriptor()
	// user.DefaultAddress holds the default value on creation for the address field.
	user.DefaultAddress = userDescAddress.Default.(func() string)
}