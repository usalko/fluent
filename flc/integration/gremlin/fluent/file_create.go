// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/__"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/p"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/file"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/file_type"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/user"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
}

// SetSetID sets the "set_id" field.
func (fc *FileCreate) SetSetID(i int) *FileCreate {
	fc.mutation.SetSetID(i)
	return fc
}

// SetNillableSetID sets the "set_id" field if the given value is not nil.
func (fc *FileCreate) SetNillableSetID(i *int) *FileCreate {
	if i != nil {
		fc.SetSetID(*i)
	}
	return fc
}

// SetSize sets the "size" field.
func (fc *FileCreate) SetSize(i int) *FileCreate {
	fc.mutation.SetSize(i)
	return fc
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fc *FileCreate) SetNillableSize(i *int) *FileCreate {
	if i != nil {
		fc.SetSize(*i)
	}
	return fc
}

// SetName sets the "name" field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetUser sets the "user" field.
func (fc *FileCreate) SetUser(s string) *FileCreate {
	fc.mutation.SetUser(s)
	return fc
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (fc *FileCreate) SetNillableUser(s *string) *FileCreate {
	if s != nil {
		fc.SetUser(*s)
	}
	return fc
}

// SetGroup sets the "group" field.
func (fc *FileCreate) SetGroup(s string) *FileCreate {
	fc.mutation.SetGroup(s)
	return fc
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (fc *FileCreate) SetNillableGroup(s *string) *FileCreate {
	if s != nil {
		fc.SetGroup(*s)
	}
	return fc
}

// SetOp sets the "op" field.
func (fc *FileCreate) SetOp(b bool) *FileCreate {
	fc.mutation.SetOpField(b)
	return fc
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (fc *FileCreate) SetNillableOp(b *bool) *FileCreate {
	if b != nil {
		fc.SetOp(*b)
	}
	return fc
}

// SetFieldID sets the "field_id" field.
func (fc *FileCreate) SetFieldID(i int) *FileCreate {
	fc.mutation.SetFieldID(i)
	return fc
}

// SetNillableFieldID sets the "field_id" field if the given value is not nil.
func (fc *FileCreate) SetNillableFieldID(i *int) *FileCreate {
	if i != nil {
		fc.SetFieldID(*i)
	}
	return fc
}

// SetCreateTime sets the "create_time" field.
func (fc *FileCreate) SetCreateTime(t time.Time) *FileCreate {
	fc.mutation.SetCreateTime(t)
	return fc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (fc *FileCreate) SetNillableCreateTime(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetCreateTime(*t)
	}
	return fc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fc *FileCreate) SetOwnerID(id string) *FileCreate {
	fc.mutation.SetOwnerID(id)
	return fc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableOwnerID(id *string) *FileCreate {
	if id != nil {
		fc = fc.SetOwnerID(*id)
	}
	return fc
}

// SetOwner sets the "owner" edge to the User entity.
func (fc *FileCreate) SetOwner(u *User) *FileCreate {
	return fc.SetOwnerID(u.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (fc *FileCreate) SetTypeID(id string) *FileCreate {
	fc.mutation.SetTypeID(id)
	return fc
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableTypeID(id *string) *FileCreate {
	if id != nil {
		fc = fc.SetTypeID(*id)
	}
	return fc
}

// SetType sets the "type" edge to the FileType entity.
func (fc *FileCreate) SetType(f *FileType) *FileCreate {
	return fc.SetTypeID(f.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (fc *FileCreate) AddFieldIDs(ids ...string) *FileCreate {
	fc.mutation.AddFieldIDs(ids...)
	return fc
}

// AddField adds the "field" edges to the FieldType entity.
func (fc *FileCreate) AddField(f ...*FieldType) *FileCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	fc.defaults()
	return withHooks(ctx, fc.gremlinSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FileCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FileCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() {
	if _, ok := fc.mutation.Size(); !ok {
		v := file.DefaultSize
		fc.mutation.SetSize(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if v, ok := fc.mutation.SetID(); ok {
		if err := file.SetIDValidator(v); err != nil {
			return &ValidationError{Name: "set_id", err: fmt.Errorf(`fluent: validator failed for field "File.set_id": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`fluent: missing required field "File.size"`)}
	}
	if v, ok := fc.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`fluent: validator failed for field "File.size": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`fluent: missing required field "File.name"`)}
	}
	return nil
}

func (fc *FileCreate) gremlinSave(ctx context.Context) (*File, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := fc.gremlin().Query()
	if err := fc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &File{config: fc.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	fc.mutation.id = &rnode.ID
	fc.mutation.done = true
	return rnode, nil
}

func (fc *FileCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(file.Label)
	if value, ok := fc.mutation.SetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, value)
	}
	if value, ok := fc.mutation.Size(); ok {
		v.Property(dsl.Single, file.FieldSize, value)
	}
	if value, ok := fc.mutation.Name(); ok {
		v.Property(dsl.Single, file.FieldName, value)
	}
	if value, ok := fc.mutation.User(); ok {
		v.Property(dsl.Single, file.FieldUser, value)
	}
	if value, ok := fc.mutation.Group(); ok {
		v.Property(dsl.Single, file.FieldGroup, value)
	}
	if value, ok := fc.mutation.GetOp(); ok {
		v.Property(dsl.Single, file.FieldOp, value)
	}
	if value, ok := fc.mutation.FieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, value)
	}
	if value, ok := fc.mutation.CreateTime(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(file.Label, file.FieldCreateTime, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(file.Label, file.FieldCreateTime, value)),
		})
		v.Property(dsl.Single, file.FieldCreateTime, value)
	}
	for _, id := range fc.mutation.OwnerIDs() {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range fc.mutation.TypeIDs() {
		v.AddE(file_type.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range fc.mutation.FieldIDs() {
		v.AddE(file.FieldLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(file.FieldLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(file.Label, file.FieldLabel, id)),
		})
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	err      error
	builders []*FileCreate
}
