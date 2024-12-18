// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package generated_code

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/dialect/sql/sqljson"
	"github.com/usalko/fluent/flc/integration/another_target/generated_code/predicate"
	"github.com/usalko/fluent/flc/integration/another_target/generated_code/schema"
	"github.com/usalko/fluent/flc/integration/another_target/generated_code/user"
	"github.com/usalko/fluent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetT sets the "t" field.
func (uu *UserUpdate) SetT(schema_t *schema.T) *UserUpdate {
	uu.mutation.SetT(schema_t)
	return uu
}

// ClearT clears the value of the "t" field.
func (uu *UserUpdate) ClearT() *UserUpdate {
	uu.mutation.ClearT()
	return uu
}

// SetURL sets the "url" field.
func (uu *UserUpdate) SetURL(url_url *url.URL) *UserUpdate {
	uu.mutation.SetURL(url_url)
	return uu
}

// ClearURL clears the value of the "url" field.
func (uu *UserUpdate) ClearURL() *UserUpdate {
	uu.mutation.ClearURL()
	return uu
}

// SetURLs sets the "URLs" field.
func (uu *UserUpdate) SetURLs(url_url []*url.URL) *UserUpdate {
	uu.mutation.SetURLs(url_url)
	return uu
}

// AppendURLs appends url_url to the "URLs" field.
func (uu *UserUpdate) AppendURLs(url_url []*url.URL) *UserUpdate {
	uu.mutation.AppendURLs(url_url)
	return uu
}

// ClearURLs clears the value of the "URLs" field.
func (uu *UserUpdate) ClearURLs() *UserUpdate {
	uu.mutation.ClearURLs()
	return uu
}

// SetRaw sets the "raw" field.
func (uu *UserUpdate) SetRaw(json_rawmessage json.RawMessage) *UserUpdate {
	uu.mutation.SetRaw(json_rawmessage)
	return uu
}

// AppendRaw appends json_rawmessage to the "raw" field.
func (uu *UserUpdate) AppendRaw(json_rawmessage json.RawMessage) *UserUpdate {
	uu.mutation.AppendRaw(json_rawmessage)
	return uu
}

// ClearRaw clears the value of the "raw" field.
func (uu *UserUpdate) ClearRaw() *UserUpdate {
	uu.mutation.ClearRaw()
	return uu
}

// SetDirs sets the "dirs" field.
func (uu *UserUpdate) SetDirs(http_dir []http.Dir) *UserUpdate {
	uu.mutation.SetDirs(http_dir)
	return uu
}

// AppendDirs appends http_dir to the "dirs" field.
func (uu *UserUpdate) AppendDirs(http_dir []http.Dir) *UserUpdate {
	uu.mutation.AppendDirs(http_dir)
	return uu
}

// SetInts sets the "ints" field.
func (uu *UserUpdate) SetInts(i []int) *UserUpdate {
	uu.mutation.SetInts(i)
	return uu
}

// AppendInts appends i to the "ints" field.
func (uu *UserUpdate) AppendInts(i []int) *UserUpdate {
	uu.mutation.AppendInts(i)
	return uu
}

// ClearInts clears the value of the "ints" field.
func (uu *UserUpdate) ClearInts() *UserUpdate {
	uu.mutation.ClearInts()
	return uu
}

// SetFloats sets the "floats" field.
func (uu *UserUpdate) SetFloats(f []float64) *UserUpdate {
	uu.mutation.SetFloats(f)
	return uu
}

// AppendFloats appends f to the "floats" field.
func (uu *UserUpdate) AppendFloats(f []float64) *UserUpdate {
	uu.mutation.AppendFloats(f)
	return uu
}

// ClearFloats clears the value of the "floats" field.
func (uu *UserUpdate) ClearFloats() *UserUpdate {
	uu.mutation.ClearFloats()
	return uu
}

// SetStrings sets the "strings" field.
func (uu *UserUpdate) SetStrings(s []string) *UserUpdate {
	uu.mutation.SetStrings(s)
	return uu
}

// AppendStrings appends s to the "strings" field.
func (uu *UserUpdate) AppendStrings(s []string) *UserUpdate {
	uu.mutation.AppendStrings(s)
	return uu
}

// ClearStrings clears the value of the "strings" field.
func (uu *UserUpdate) ClearStrings() *UserUpdate {
	uu.mutation.ClearStrings()
	return uu
}

// SetIntsValidate sets the "ints_validate" field.
func (uu *UserUpdate) SetIntsValidate(i []int) *UserUpdate {
	uu.mutation.SetIntsValidate(i)
	return uu
}

// AppendIntsValidate appends i to the "ints_validate" field.
func (uu *UserUpdate) AppendIntsValidate(i []int) *UserUpdate {
	uu.mutation.AppendIntsValidate(i)
	return uu
}

// ClearIntsValidate clears the value of the "ints_validate" field.
func (uu *UserUpdate) ClearIntsValidate() *UserUpdate {
	uu.mutation.ClearIntsValidate()
	return uu
}

// SetFloatsValidate sets the "floats_validate" field.
func (uu *UserUpdate) SetFloatsValidate(f []float64) *UserUpdate {
	uu.mutation.SetFloatsValidate(f)
	return uu
}

// AppendFloatsValidate appends f to the "floats_validate" field.
func (uu *UserUpdate) AppendFloatsValidate(f []float64) *UserUpdate {
	uu.mutation.AppendFloatsValidate(f)
	return uu
}

// ClearFloatsValidate clears the value of the "floats_validate" field.
func (uu *UserUpdate) ClearFloatsValidate() *UserUpdate {
	uu.mutation.ClearFloatsValidate()
	return uu
}

// SetStringsValidate sets the "strings_validate" field.
func (uu *UserUpdate) SetStringsValidate(s []string) *UserUpdate {
	uu.mutation.SetStringsValidate(s)
	return uu
}

// AppendStringsValidate appends s to the "strings_validate" field.
func (uu *UserUpdate) AppendStringsValidate(s []string) *UserUpdate {
	uu.mutation.AppendStringsValidate(s)
	return uu
}

// ClearStringsValidate clears the value of the "strings_validate" field.
func (uu *UserUpdate) ClearStringsValidate() *UserUpdate {
	uu.mutation.ClearStringsValidate()
	return uu
}

// SetAddr sets the "addr" field.
func (uu *UserUpdate) SetAddr(schema_addr schema.Addr) *UserUpdate {
	uu.mutation.SetAddr(schema_addr)
	return uu
}

// SetNillableAddr sets the "addr" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAddr(schema_addr *schema.Addr) *UserUpdate {
	if schema_addr != nil {
		uu.SetAddr(*schema_addr)
	}
	return uu
}

// ClearAddr clears the value of the "addr" field.
func (uu *UserUpdate) ClearAddr() *UserUpdate {
	uu.mutation.ClearAddr()
	return uu
}

// SetUnknown sets the "unknown" field.
func (uu *UserUpdate) SetUnknown(a any) *UserUpdate {
	uu.mutation.SetUnknown(a)
	return uu
}

// ClearUnknown clears the value of the "unknown" field.
func (uu *UserUpdate) ClearUnknown() *UserUpdate {
	uu.mutation.ClearUnknown()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.IntsValidate(); ok {
		if err := user.IntsValidateValidator(v); err != nil {
			return &ValidationError{Name: "ints_validate", err: fmt.Errorf(`generated_code: validator failed for field "User.ints_validate": %w`, err)}
		}
	}
	if v, ok := uu.mutation.FloatsValidate(); ok {
		if err := user.FloatsValidateValidator(v); err != nil {
			return &ValidationError{Name: "floats_validate", err: fmt.Errorf(`generated_code: validator failed for field "User.floats_validate": %w`, err)}
		}
	}
	if v, ok := uu.mutation.StringsValidate(); ok {
		if err := user.StringsValidateValidator(v); err != nil {
			return &ValidationError{Name: "strings_validate", err: fmt.Errorf(`generated_code: validator failed for field "User.strings_validate": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uu *UserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdate {
	uu.modifiers = append(uu.modifiers, modifiers...)
	return uu
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.T(); ok {
		_spec.SetField(user.FieldT, field.TypeJSON, value)
	}
	if uu.mutation.TCleared() {
		_spec.ClearField(user.FieldT, field.TypeJSON)
	}
	if value, ok := uu.mutation.URL(); ok {
		_spec.SetField(user.FieldURL, field.TypeJSON, value)
	}
	if uu.mutation.URLCleared() {
		_spec.ClearField(user.FieldURL, field.TypeJSON)
	}
	if value, ok := uu.mutation.URLs(); ok {
		_spec.SetField(user.FieldURLs, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedURLs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldURLs, value)
		})
	}
	if uu.mutation.URLsCleared() {
		_spec.ClearField(user.FieldURLs, field.TypeJSON)
	}
	if value, ok := uu.mutation.Raw(); ok {
		_spec.SetField(user.FieldRaw, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedRaw(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldRaw, value)
		})
	}
	if uu.mutation.RawCleared() {
		_spec.ClearField(user.FieldRaw, field.TypeJSON)
	}
	if value, ok := uu.mutation.Dirs(); ok {
		_spec.SetField(user.FieldDirs, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedDirs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldDirs, value)
		})
	}
	if value, ok := uu.mutation.Ints(); ok {
		_spec.SetField(user.FieldInts, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedInts(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldInts, value)
		})
	}
	if uu.mutation.IntsCleared() {
		_spec.ClearField(user.FieldInts, field.TypeJSON)
	}
	if value, ok := uu.mutation.Floats(); ok {
		_spec.SetField(user.FieldFloats, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedFloats(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldFloats, value)
		})
	}
	if uu.mutation.FloatsCleared() {
		_spec.ClearField(user.FieldFloats, field.TypeJSON)
	}
	if value, ok := uu.mutation.Strings(); ok {
		_spec.SetField(user.FieldStrings, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedStrings(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStrings, value)
		})
	}
	if uu.mutation.StringsCleared() {
		_spec.ClearField(user.FieldStrings, field.TypeJSON)
	}
	if value, ok := uu.mutation.IntsValidate(); ok {
		_spec.SetField(user.FieldIntsValidate, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedIntsValidate(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldIntsValidate, value)
		})
	}
	if uu.mutation.IntsValidateCleared() {
		_spec.ClearField(user.FieldIntsValidate, field.TypeJSON)
	}
	if value, ok := uu.mutation.FloatsValidate(); ok {
		_spec.SetField(user.FieldFloatsValidate, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedFloatsValidate(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldFloatsValidate, value)
		})
	}
	if uu.mutation.FloatsValidateCleared() {
		_spec.ClearField(user.FieldFloatsValidate, field.TypeJSON)
	}
	if value, ok := uu.mutation.StringsValidate(); ok {
		_spec.SetField(user.FieldStringsValidate, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedStringsValidate(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStringsValidate, value)
		})
	}
	if uu.mutation.StringsValidateCleared() {
		_spec.ClearField(user.FieldStringsValidate, field.TypeJSON)
	}
	if value, ok := uu.mutation.Addr(); ok {
		_spec.SetField(user.FieldAddr, field.TypeJSON, value)
	}
	if uu.mutation.AddrCleared() {
		_spec.ClearField(user.FieldAddr, field.TypeJSON)
	}
	if value, ok := uu.mutation.Unknown(); ok {
		_spec.SetField(user.FieldUnknown, field.TypeJSON, value)
	}
	if uu.mutation.UnknownCleared() {
		_spec.ClearField(user.FieldUnknown, field.TypeJSON)
	}
	_spec.AddModifiers(uu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetT sets the "t" field.
func (uuo *UserUpdateOne) SetT(schema_t *schema.T) *UserUpdateOne {
	uuo.mutation.SetT(schema_t)
	return uuo
}

// ClearT clears the value of the "t" field.
func (uuo *UserUpdateOne) ClearT() *UserUpdateOne {
	uuo.mutation.ClearT()
	return uuo
}

// SetURL sets the "url" field.
func (uuo *UserUpdateOne) SetURL(url_url *url.URL) *UserUpdateOne {
	uuo.mutation.SetURL(url_url)
	return uuo
}

// ClearURL clears the value of the "url" field.
func (uuo *UserUpdateOne) ClearURL() *UserUpdateOne {
	uuo.mutation.ClearURL()
	return uuo
}

// SetURLs sets the "URLs" field.
func (uuo *UserUpdateOne) SetURLs(url_url []*url.URL) *UserUpdateOne {
	uuo.mutation.SetURLs(url_url)
	return uuo
}

// AppendURLs appends url_url to the "URLs" field.
func (uuo *UserUpdateOne) AppendURLs(url_url []*url.URL) *UserUpdateOne {
	uuo.mutation.AppendURLs(url_url)
	return uuo
}

// ClearURLs clears the value of the "URLs" field.
func (uuo *UserUpdateOne) ClearURLs() *UserUpdateOne {
	uuo.mutation.ClearURLs()
	return uuo
}

// SetRaw sets the "raw" field.
func (uuo *UserUpdateOne) SetRaw(json_rawmessage json.RawMessage) *UserUpdateOne {
	uuo.mutation.SetRaw(json_rawmessage)
	return uuo
}

// AppendRaw appends json_rawmessage to the "raw" field.
func (uuo *UserUpdateOne) AppendRaw(json_rawmessage json.RawMessage) *UserUpdateOne {
	uuo.mutation.AppendRaw(json_rawmessage)
	return uuo
}

// ClearRaw clears the value of the "raw" field.
func (uuo *UserUpdateOne) ClearRaw() *UserUpdateOne {
	uuo.mutation.ClearRaw()
	return uuo
}

// SetDirs sets the "dirs" field.
func (uuo *UserUpdateOne) SetDirs(http_dir []http.Dir) *UserUpdateOne {
	uuo.mutation.SetDirs(http_dir)
	return uuo
}

// AppendDirs appends http_dir to the "dirs" field.
func (uuo *UserUpdateOne) AppendDirs(http_dir []http.Dir) *UserUpdateOne {
	uuo.mutation.AppendDirs(http_dir)
	return uuo
}

// SetInts sets the "ints" field.
func (uuo *UserUpdateOne) SetInts(i []int) *UserUpdateOne {
	uuo.mutation.SetInts(i)
	return uuo
}

// AppendInts appends i to the "ints" field.
func (uuo *UserUpdateOne) AppendInts(i []int) *UserUpdateOne {
	uuo.mutation.AppendInts(i)
	return uuo
}

// ClearInts clears the value of the "ints" field.
func (uuo *UserUpdateOne) ClearInts() *UserUpdateOne {
	uuo.mutation.ClearInts()
	return uuo
}

// SetFloats sets the "floats" field.
func (uuo *UserUpdateOne) SetFloats(f []float64) *UserUpdateOne {
	uuo.mutation.SetFloats(f)
	return uuo
}

// AppendFloats appends f to the "floats" field.
func (uuo *UserUpdateOne) AppendFloats(f []float64) *UserUpdateOne {
	uuo.mutation.AppendFloats(f)
	return uuo
}

// ClearFloats clears the value of the "floats" field.
func (uuo *UserUpdateOne) ClearFloats() *UserUpdateOne {
	uuo.mutation.ClearFloats()
	return uuo
}

// SetStrings sets the "strings" field.
func (uuo *UserUpdateOne) SetStrings(s []string) *UserUpdateOne {
	uuo.mutation.SetStrings(s)
	return uuo
}

// AppendStrings appends s to the "strings" field.
func (uuo *UserUpdateOne) AppendStrings(s []string) *UserUpdateOne {
	uuo.mutation.AppendStrings(s)
	return uuo
}

// ClearStrings clears the value of the "strings" field.
func (uuo *UserUpdateOne) ClearStrings() *UserUpdateOne {
	uuo.mutation.ClearStrings()
	return uuo
}

// SetIntsValidate sets the "ints_validate" field.
func (uuo *UserUpdateOne) SetIntsValidate(i []int) *UserUpdateOne {
	uuo.mutation.SetIntsValidate(i)
	return uuo
}

// AppendIntsValidate appends i to the "ints_validate" field.
func (uuo *UserUpdateOne) AppendIntsValidate(i []int) *UserUpdateOne {
	uuo.mutation.AppendIntsValidate(i)
	return uuo
}

// ClearIntsValidate clears the value of the "ints_validate" field.
func (uuo *UserUpdateOne) ClearIntsValidate() *UserUpdateOne {
	uuo.mutation.ClearIntsValidate()
	return uuo
}

// SetFloatsValidate sets the "floats_validate" field.
func (uuo *UserUpdateOne) SetFloatsValidate(f []float64) *UserUpdateOne {
	uuo.mutation.SetFloatsValidate(f)
	return uuo
}

// AppendFloatsValidate appends f to the "floats_validate" field.
func (uuo *UserUpdateOne) AppendFloatsValidate(f []float64) *UserUpdateOne {
	uuo.mutation.AppendFloatsValidate(f)
	return uuo
}

// ClearFloatsValidate clears the value of the "floats_validate" field.
func (uuo *UserUpdateOne) ClearFloatsValidate() *UserUpdateOne {
	uuo.mutation.ClearFloatsValidate()
	return uuo
}

// SetStringsValidate sets the "strings_validate" field.
func (uuo *UserUpdateOne) SetStringsValidate(s []string) *UserUpdateOne {
	uuo.mutation.SetStringsValidate(s)
	return uuo
}

// AppendStringsValidate appends s to the "strings_validate" field.
func (uuo *UserUpdateOne) AppendStringsValidate(s []string) *UserUpdateOne {
	uuo.mutation.AppendStringsValidate(s)
	return uuo
}

// ClearStringsValidate clears the value of the "strings_validate" field.
func (uuo *UserUpdateOne) ClearStringsValidate() *UserUpdateOne {
	uuo.mutation.ClearStringsValidate()
	return uuo
}

// SetAddr sets the "addr" field.
func (uuo *UserUpdateOne) SetAddr(schema_addr schema.Addr) *UserUpdateOne {
	uuo.mutation.SetAddr(schema_addr)
	return uuo
}

// SetNillableAddr sets the "addr" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddr(schema_addr *schema.Addr) *UserUpdateOne {
	if schema_addr != nil {
		uuo.SetAddr(*schema_addr)
	}
	return uuo
}

// ClearAddr clears the value of the "addr" field.
func (uuo *UserUpdateOne) ClearAddr() *UserUpdateOne {
	uuo.mutation.ClearAddr()
	return uuo
}

// SetUnknown sets the "unknown" field.
func (uuo *UserUpdateOne) SetUnknown(a any) *UserUpdateOne {
	uuo.mutation.SetUnknown(a)
	return uuo
}

// ClearUnknown clears the value of the "unknown" field.
func (uuo *UserUpdateOne) ClearUnknown() *UserUpdateOne {
	uuo.mutation.ClearUnknown()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.IntsValidate(); ok {
		if err := user.IntsValidateValidator(v); err != nil {
			return &ValidationError{Name: "ints_validate", err: fmt.Errorf(`generated_code: validator failed for field "User.ints_validate": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.FloatsValidate(); ok {
		if err := user.FloatsValidateValidator(v); err != nil {
			return &ValidationError{Name: "floats_validate", err: fmt.Errorf(`generated_code: validator failed for field "User.floats_validate": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.StringsValidate(); ok {
		if err := user.StringsValidateValidator(v); err != nil {
			return &ValidationError{Name: "strings_validate", err: fmt.Errorf(`generated_code: validator failed for field "User.strings_validate": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uuo *UserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdateOne {
	uuo.modifiers = append(uuo.modifiers, modifiers...)
	return uuo
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated_code: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated_code: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.T(); ok {
		_spec.SetField(user.FieldT, field.TypeJSON, value)
	}
	if uuo.mutation.TCleared() {
		_spec.ClearField(user.FieldT, field.TypeJSON)
	}
	if value, ok := uuo.mutation.URL(); ok {
		_spec.SetField(user.FieldURL, field.TypeJSON, value)
	}
	if uuo.mutation.URLCleared() {
		_spec.ClearField(user.FieldURL, field.TypeJSON)
	}
	if value, ok := uuo.mutation.URLs(); ok {
		_spec.SetField(user.FieldURLs, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedURLs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldURLs, value)
		})
	}
	if uuo.mutation.URLsCleared() {
		_spec.ClearField(user.FieldURLs, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Raw(); ok {
		_spec.SetField(user.FieldRaw, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedRaw(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldRaw, value)
		})
	}
	if uuo.mutation.RawCleared() {
		_spec.ClearField(user.FieldRaw, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Dirs(); ok {
		_spec.SetField(user.FieldDirs, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedDirs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldDirs, value)
		})
	}
	if value, ok := uuo.mutation.Ints(); ok {
		_spec.SetField(user.FieldInts, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedInts(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldInts, value)
		})
	}
	if uuo.mutation.IntsCleared() {
		_spec.ClearField(user.FieldInts, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Floats(); ok {
		_spec.SetField(user.FieldFloats, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedFloats(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldFloats, value)
		})
	}
	if uuo.mutation.FloatsCleared() {
		_spec.ClearField(user.FieldFloats, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Strings(); ok {
		_spec.SetField(user.FieldStrings, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedStrings(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStrings, value)
		})
	}
	if uuo.mutation.StringsCleared() {
		_spec.ClearField(user.FieldStrings, field.TypeJSON)
	}
	if value, ok := uuo.mutation.IntsValidate(); ok {
		_spec.SetField(user.FieldIntsValidate, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedIntsValidate(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldIntsValidate, value)
		})
	}
	if uuo.mutation.IntsValidateCleared() {
		_spec.ClearField(user.FieldIntsValidate, field.TypeJSON)
	}
	if value, ok := uuo.mutation.FloatsValidate(); ok {
		_spec.SetField(user.FieldFloatsValidate, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedFloatsValidate(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldFloatsValidate, value)
		})
	}
	if uuo.mutation.FloatsValidateCleared() {
		_spec.ClearField(user.FieldFloatsValidate, field.TypeJSON)
	}
	if value, ok := uuo.mutation.StringsValidate(); ok {
		_spec.SetField(user.FieldStringsValidate, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedStringsValidate(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStringsValidate, value)
		})
	}
	if uuo.mutation.StringsValidateCleared() {
		_spec.ClearField(user.FieldStringsValidate, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Addr(); ok {
		_spec.SetField(user.FieldAddr, field.TypeJSON, value)
	}
	if uuo.mutation.AddrCleared() {
		_spec.ClearField(user.FieldAddr, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Unknown(); ok {
		_spec.SetField(user.FieldUnknown, field.TypeJSON, value)
	}
	if uuo.mutation.UnknownCleared() {
		_spec.ClearField(user.FieldUnknown, field.TypeJSON)
	}
	_spec.AddModifiers(uuo.modifiers...)
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
