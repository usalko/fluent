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
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/user"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetSetID sets the "set_id" field.
func (fu *FileUpdate) SetSetID(i int) *FileUpdate {
	fu.mutation.ResetSetID()
	fu.mutation.SetSetID(i)
	return fu
}

// SetNillableSetID sets the "set_id" field if the given value is not nil.
func (fu *FileUpdate) SetNillableSetID(i *int) *FileUpdate {
	if i != nil {
		fu.SetSetID(*i)
	}
	return fu
}

// AddSetID adds i to the "set_id" field.
func (fu *FileUpdate) AddSetID(i int) *FileUpdate {
	fu.mutation.AddSetID(i)
	return fu
}

// ClearSetID clears the value of the "set_id" field.
func (fu *FileUpdate) ClearSetID() *FileUpdate {
	fu.mutation.ClearSetID()
	return fu
}

// SetSize sets the "size" field.
func (fu *FileUpdate) SetSize(i int) *FileUpdate {
	fu.mutation.ResetSize()
	fu.mutation.SetSize(i)
	return fu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fu *FileUpdate) SetNillableSize(i *int) *FileUpdate {
	if i != nil {
		fu.SetSize(*i)
	}
	return fu
}

// AddSize adds i to the "size" field.
func (fu *FileUpdate) AddSize(i int) *FileUpdate {
	fu.mutation.AddSize(i)
	return fu
}

// SetName sets the "name" field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fu *FileUpdate) SetNillableName(s *string) *FileUpdate {
	if s != nil {
		fu.SetName(*s)
	}
	return fu
}

// SetUser sets the "user" field.
func (fu *FileUpdate) SetUser(s string) *FileUpdate {
	fu.mutation.SetUser(s)
	return fu
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (fu *FileUpdate) SetNillableUser(s *string) *FileUpdate {
	if s != nil {
		fu.SetUser(*s)
	}
	return fu
}

// ClearUser clears the value of the "user" field.
func (fu *FileUpdate) ClearUser() *FileUpdate {
	fu.mutation.ClearUser()
	return fu
}

// SetGroup sets the "group" field.
func (fu *FileUpdate) SetGroup(s string) *FileUpdate {
	fu.mutation.SetGroup(s)
	return fu
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (fu *FileUpdate) SetNillableGroup(s *string) *FileUpdate {
	if s != nil {
		fu.SetGroup(*s)
	}
	return fu
}

// ClearGroup clears the value of the "group" field.
func (fu *FileUpdate) ClearGroup() *FileUpdate {
	fu.mutation.ClearGroup()
	return fu
}

// SetOp sets the "op" field.
func (fu *FileUpdate) SetOp(b bool) *FileUpdate {
	fu.mutation.SetOpField(b)
	return fu
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (fu *FileUpdate) SetNillableOp(b *bool) *FileUpdate {
	if b != nil {
		fu.SetOp(*b)
	}
	return fu
}

// ClearOp clears the value of the "op" field.
func (fu *FileUpdate) ClearOp() *FileUpdate {
	fu.mutation.ClearOp()
	return fu
}

// SetFieldID sets the "field_id" field.
func (fu *FileUpdate) SetFieldID(i int) *FileUpdate {
	fu.mutation.ResetFieldID()
	fu.mutation.SetFieldID(i)
	return fu
}

// SetNillableFieldID sets the "field_id" field if the given value is not nil.
func (fu *FileUpdate) SetNillableFieldID(i *int) *FileUpdate {
	if i != nil {
		fu.SetFieldID(*i)
	}
	return fu
}

// AddFieldID adds i to the "field_id" field.
func (fu *FileUpdate) AddFieldID(i int) *FileUpdate {
	fu.mutation.AddFieldID(i)
	return fu
}

// ClearFieldID clears the value of the "field_id" field.
func (fu *FileUpdate) ClearFieldID() *FileUpdate {
	fu.mutation.ClearFieldID()
	return fu
}

// SetCreateTime sets the "create_time" field.
func (fu *FileUpdate) SetCreateTime(time_time time.Time) *FileUpdate {
	fu.mutation.SetCreateTime(time_time)
	return fu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (fu *FileUpdate) SetNillableCreateTime(time_time *time.Time) *FileUpdate {
	if time_time != nil {
		fu.SetCreateTime(*time_time)
	}
	return fu
}

// ClearCreateTime clears the value of the "create_time" field.
func (fu *FileUpdate) ClearCreateTime() *FileUpdate {
	fu.mutation.ClearCreateTime()
	return fu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fu *FileUpdate) SetOwnerID(id string) *FileUpdate {
	fu.mutation.SetOwnerID(id)
	return fu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fu *FileUpdate) SetNillableOwnerID(id *string) *FileUpdate {
	if id != nil {
		fu = fu.SetOwnerID(*id)
	}
	return fu
}

// SetOwner sets the "owner" edge to the User entity.
func (fu *FileUpdate) SetOwner(u *User) *FileUpdate {
	return fu.SetOwnerID(u.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (fu *FileUpdate) SetTypeID(id string) *FileUpdate {
	fu.mutation.SetTypeID(id)
	return fu
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (fu *FileUpdate) SetNillableTypeID(id *string) *FileUpdate {
	if id != nil {
		fu = fu.SetTypeID(*id)
	}
	return fu
}

// SetType sets the "type" edge to the FileType entity.
func (fu *FileUpdate) SetType(f *FileType) *FileUpdate {
	return fu.SetTypeID(f.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (fu *FileUpdate) AddFieldIDs(ids ...string) *FileUpdate {
	fu.mutation.AddFieldIDs(ids...)
	return fu
}

// AddField adds the "field" edges to the FieldType entity.
func (fu *FileUpdate) AddField(f ...*FieldType) *FileUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (fu *FileUpdate) ClearOwner() *FileUpdate {
	fu.mutation.ClearOwner()
	return fu
}

// ClearType clears the "type" edge to the FileType entity.
func (fu *FileUpdate) ClearType() *FileUpdate {
	fu.mutation.ClearType()
	return fu
}

// ClearFieldEdge clears all "field" edges to the FieldType entity.
func (fu *FileUpdate) ClearFieldEdge() *FileUpdate {
	fu.mutation.ClearFieldEdge()
	return fu
}

// RemoveFieldIDs removes the "field" edge to FieldType entities by IDs.
func (fu *FileUpdate) RemoveFieldIDs(ids ...string) *FileUpdate {
	fu.mutation.RemoveFieldIDs(ids...)
	return fu
}

// RemoveField removes "field" edges to FieldType entities.
func (fu *FileUpdate) RemoveField(f ...*FieldType) *FileUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveFieldIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.gremlinSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FileUpdate) check() error {
	if v, ok := fu.mutation.SetID(); ok {
		if err := file.SetIDValidator(v); err != nil {
			return &ValidationError{Name: "set_id", err: fmt.Errorf(`fluent: validator failed for field "File.set_id": %w`, err)}
		}
	}
	if v, ok := fu.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`fluent: validator failed for field "File.size": %w`, err)}
		}
	}
	return nil
}

func (fu *FileUpdate) gremlinSave(ctx context.Context) (int, error) {
	if err := fu.check(); err != nil {
		return 0, err
	}
	res := &gremlin.Response{}
	query, bindings := fu.gremlin().Query()
	if err := fu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	fu.mutation.done = true
	return res.ReadInt()
}

func (fu *FileUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(file.Label)
	for _, p := range fu.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := fu.mutation.SetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, value)
	}
	if value, ok := fu.mutation.AddedSetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, __.Union(__.Values(file.FieldSetID), __.Constant(value)).Sum())
	}
	if value, ok := fu.mutation.Size(); ok {
		v.Property(dsl.Single, file.FieldSize, value)
	}
	if value, ok := fu.mutation.AddedSize(); ok {
		v.Property(dsl.Single, file.FieldSize, __.Union(__.Values(file.FieldSize), __.Constant(value)).Sum())
	}
	if value, ok := fu.mutation.Name(); ok {
		v.Property(dsl.Single, file.FieldName, value)
	}
	if value, ok := fu.mutation.User(); ok {
		v.Property(dsl.Single, file.FieldUser, value)
	}
	if value, ok := fu.mutation.Group(); ok {
		v.Property(dsl.Single, file.FieldGroup, value)
	}
	if value, ok := fu.mutation.GetOp(); ok {
		v.Property(dsl.Single, file.FieldOp, value)
	}
	if value, ok := fu.mutation.FieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, value)
	}
	if value, ok := fu.mutation.AddedFieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, __.Union(__.Values(file.FieldFieldID), __.Constant(value)).Sum())
	}
	if value, ok := fu.mutation.CreateTime(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(file.Label, file.FieldCreateTime, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(file.Label, file.FieldCreateTime, value)),
		})
		v.Property(dsl.Single, file.FieldCreateTime, value)
	}
	var properties []any
	if fu.mutation.SetIDCleared() {
		properties = append(properties, file.FieldSetID)
	}
	if fu.mutation.UserCleared() {
		properties = append(properties, file.FieldUser)
	}
	if fu.mutation.GroupCleared() {
		properties = append(properties, file.FieldGroup)
	}
	if fu.mutation.OpCleared() {
		properties = append(properties, file.FieldOp)
	}
	if fu.mutation.FieldIDCleared() {
		properties = append(properties, file.FieldFieldID)
	}
	if fu.mutation.CreateTimeCleared() {
		properties = append(properties, file.FieldCreateTime)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if fu.mutation.OwnerCleared() {
		tr := rv.Clone().InE(user.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range fu.mutation.OwnerIDs() {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	if fu.mutation.TypeCleared() {
		tr := rv.Clone().InE(file_type.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range fu.mutation.TypeIDs() {
		v.AddE(file_type.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range fu.mutation.RemovedFieldIDs() {
		tr := rv.Clone().OutE(file.FieldLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range fu.mutation.FieldIDs() {
		v.AddE(file.FieldLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(file.FieldLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(file.Label, file.FieldLabel, id)),
		})
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetSetID sets the "set_id" field.
func (fuo *FileUpdateOne) SetSetID(i int) *FileUpdateOne {
	fuo.mutation.ResetSetID()
	fuo.mutation.SetSetID(i)
	return fuo
}

// SetNillableSetID sets the "set_id" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableSetID(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetSetID(*i)
	}
	return fuo
}

// AddSetID adds i to the "set_id" field.
func (fuo *FileUpdateOne) AddSetID(i int) *FileUpdateOne {
	fuo.mutation.AddSetID(i)
	return fuo
}

// ClearSetID clears the value of the "set_id" field.
func (fuo *FileUpdateOne) ClearSetID() *FileUpdateOne {
	fuo.mutation.ClearSetID()
	return fuo
}

// SetSize sets the "size" field.
func (fuo *FileUpdateOne) SetSize(i int) *FileUpdateOne {
	fuo.mutation.ResetSize()
	fuo.mutation.SetSize(i)
	return fuo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableSize(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetSize(*i)
	}
	return fuo
}

// AddSize adds i to the "size" field.
func (fuo *FileUpdateOne) AddSize(i int) *FileUpdateOne {
	fuo.mutation.AddSize(i)
	return fuo
}

// SetName sets the "name" field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableName(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetName(*s)
	}
	return fuo
}

// SetUser sets the "user" field.
func (fuo *FileUpdateOne) SetUser(s string) *FileUpdateOne {
	fuo.mutation.SetUser(s)
	return fuo
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableUser(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetUser(*s)
	}
	return fuo
}

// ClearUser clears the value of the "user" field.
func (fuo *FileUpdateOne) ClearUser() *FileUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// SetGroup sets the "group" field.
func (fuo *FileUpdateOne) SetGroup(s string) *FileUpdateOne {
	fuo.mutation.SetGroup(s)
	return fuo
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableGroup(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetGroup(*s)
	}
	return fuo
}

// ClearGroup clears the value of the "group" field.
func (fuo *FileUpdateOne) ClearGroup() *FileUpdateOne {
	fuo.mutation.ClearGroup()
	return fuo
}

// SetOp sets the "op" field.
func (fuo *FileUpdateOne) SetOp(b bool) *FileUpdateOne {
	fuo.mutation.SetOpField(b)
	return fuo
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableOp(b *bool) *FileUpdateOne {
	if b != nil {
		fuo.SetOp(*b)
	}
	return fuo
}

// ClearOp clears the value of the "op" field.
func (fuo *FileUpdateOne) ClearOp() *FileUpdateOne {
	fuo.mutation.ClearOp()
	return fuo
}

// SetFieldID sets the "field_id" field.
func (fuo *FileUpdateOne) SetFieldID(i int) *FileUpdateOne {
	fuo.mutation.ResetFieldID()
	fuo.mutation.SetFieldID(i)
	return fuo
}

// SetNillableFieldID sets the "field_id" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableFieldID(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetFieldID(*i)
	}
	return fuo
}

// AddFieldID adds i to the "field_id" field.
func (fuo *FileUpdateOne) AddFieldID(i int) *FileUpdateOne {
	fuo.mutation.AddFieldID(i)
	return fuo
}

// ClearFieldID clears the value of the "field_id" field.
func (fuo *FileUpdateOne) ClearFieldID() *FileUpdateOne {
	fuo.mutation.ClearFieldID()
	return fuo
}

// SetCreateTime sets the "create_time" field.
func (fuo *FileUpdateOne) SetCreateTime(time_time time.Time) *FileUpdateOne {
	fuo.mutation.SetCreateTime(time_time)
	return fuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableCreateTime(time_time *time.Time) *FileUpdateOne {
	if time_time != nil {
		fuo.SetCreateTime(*time_time)
	}
	return fuo
}

// ClearCreateTime clears the value of the "create_time" field.
func (fuo *FileUpdateOne) ClearCreateTime() *FileUpdateOne {
	fuo.mutation.ClearCreateTime()
	return fuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fuo *FileUpdateOne) SetOwnerID(id string) *FileUpdateOne {
	fuo.mutation.SetOwnerID(id)
	return fuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableOwnerID(id *string) *FileUpdateOne {
	if id != nil {
		fuo = fuo.SetOwnerID(*id)
	}
	return fuo
}

// SetOwner sets the "owner" edge to the User entity.
func (fuo *FileUpdateOne) SetOwner(u *User) *FileUpdateOne {
	return fuo.SetOwnerID(u.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (fuo *FileUpdateOne) SetTypeID(id string) *FileUpdateOne {
	fuo.mutation.SetTypeID(id)
	return fuo
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableTypeID(id *string) *FileUpdateOne {
	if id != nil {
		fuo = fuo.SetTypeID(*id)
	}
	return fuo
}

// SetType sets the "type" edge to the FileType entity.
func (fuo *FileUpdateOne) SetType(f *FileType) *FileUpdateOne {
	return fuo.SetTypeID(f.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (fuo *FileUpdateOne) AddFieldIDs(ids ...string) *FileUpdateOne {
	fuo.mutation.AddFieldIDs(ids...)
	return fuo
}

// AddField adds the "field" edges to the FieldType entity.
func (fuo *FileUpdateOne) AddField(f ...*FieldType) *FileUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (fuo *FileUpdateOne) ClearOwner() *FileUpdateOne {
	fuo.mutation.ClearOwner()
	return fuo
}

// ClearType clears the "type" edge to the FileType entity.
func (fuo *FileUpdateOne) ClearType() *FileUpdateOne {
	fuo.mutation.ClearType()
	return fuo
}

// ClearFieldEdge clears all "field" edges to the FieldType entity.
func (fuo *FileUpdateOne) ClearFieldEdge() *FileUpdateOne {
	fuo.mutation.ClearFieldEdge()
	return fuo
}

// RemoveFieldIDs removes the "field" edge to FieldType entities by IDs.
func (fuo *FileUpdateOne) RemoveFieldIDs(ids ...string) *FileUpdateOne {
	fuo.mutation.RemoveFieldIDs(ids...)
	return fuo
}

// RemoveField removes "field" edges to FieldType entities.
func (fuo *FileUpdateOne) RemoveField(f ...*FieldType) *FileUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveFieldIDs(ids...)
}

// Where appends a list predicates to the FileUpdate builder.
func (fuo *FileUpdateOne) Where(ps ...predicate.File) *FileUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	return withHooks(ctx, fuo.gremlinSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FileUpdateOne) check() error {
	if v, ok := fuo.mutation.SetID(); ok {
		if err := file.SetIDValidator(v); err != nil {
			return &ValidationError{Name: "set_id", err: fmt.Errorf(`fluent: validator failed for field "File.set_id": %w`, err)}
		}
	}
	if v, ok := fuo.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`fluent: validator failed for field "File.size": %w`, err)}
		}
	}
	return nil
}

func (fuo *FileUpdateOne) gremlinSave(ctx context.Context) (*File, error) {
	if err := fuo.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`fluent: missing "File.id" for update`)}
	}
	query, bindings := fuo.gremlin(id).Query()
	if err := fuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	fuo.mutation.done = true
	f := &File{config: fuo.config}
	if err := f.FromResponse(res); err != nil {
		return nil, err
	}
	return f, nil
}

func (fuo *FileUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := fuo.mutation.SetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, value)
	}
	if value, ok := fuo.mutation.AddedSetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, __.Union(__.Values(file.FieldSetID), __.Constant(value)).Sum())
	}
	if value, ok := fuo.mutation.Size(); ok {
		v.Property(dsl.Single, file.FieldSize, value)
	}
	if value, ok := fuo.mutation.AddedSize(); ok {
		v.Property(dsl.Single, file.FieldSize, __.Union(__.Values(file.FieldSize), __.Constant(value)).Sum())
	}
	if value, ok := fuo.mutation.Name(); ok {
		v.Property(dsl.Single, file.FieldName, value)
	}
	if value, ok := fuo.mutation.User(); ok {
		v.Property(dsl.Single, file.FieldUser, value)
	}
	if value, ok := fuo.mutation.Group(); ok {
		v.Property(dsl.Single, file.FieldGroup, value)
	}
	if value, ok := fuo.mutation.GetOp(); ok {
		v.Property(dsl.Single, file.FieldOp, value)
	}
	if value, ok := fuo.mutation.FieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, value)
	}
	if value, ok := fuo.mutation.AddedFieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, __.Union(__.Values(file.FieldFieldID), __.Constant(value)).Sum())
	}
	if value, ok := fuo.mutation.CreateTime(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(file.Label, file.FieldCreateTime, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(file.Label, file.FieldCreateTime, value)),
		})
		v.Property(dsl.Single, file.FieldCreateTime, value)
	}
	var properties []any
	if fuo.mutation.SetIDCleared() {
		properties = append(properties, file.FieldSetID)
	}
	if fuo.mutation.UserCleared() {
		properties = append(properties, file.FieldUser)
	}
	if fuo.mutation.GroupCleared() {
		properties = append(properties, file.FieldGroup)
	}
	if fuo.mutation.OpCleared() {
		properties = append(properties, file.FieldOp)
	}
	if fuo.mutation.FieldIDCleared() {
		properties = append(properties, file.FieldFieldID)
	}
	if fuo.mutation.CreateTimeCleared() {
		properties = append(properties, file.FieldCreateTime)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if fuo.mutation.OwnerCleared() {
		tr := rv.Clone().InE(user.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range fuo.mutation.OwnerIDs() {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	if fuo.mutation.TypeCleared() {
		tr := rv.Clone().InE(file_type.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range fuo.mutation.TypeIDs() {
		v.AddE(file_type.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range fuo.mutation.RemovedFieldIDs() {
		tr := rv.Clone().OutE(file.FieldLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range fuo.mutation.FieldIDs() {
		v.AddE(file.FieldLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(file.FieldLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(file.Label, file.FieldLabel, id)),
		})
	}
	if len(fuo.fields) > 0 {
		fields := make([]any, 0, len(fuo.fields)+1)
		fields = append(fields, true)
		for _, f := range fuo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
