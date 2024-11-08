// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/__"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/p"
	schemadir "github.com/usalko/fluent/flc/integration/fluent/schema/dir"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/comment"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUniqueInt sets the "unique_int" field.
func (cu *CommentUpdate) SetUniqueInt(i int) *CommentUpdate {
	cu.mutation.ResetUniqueInt()
	cu.mutation.SetUniqueInt(i)
	return cu
}

// SetNillableUniqueInt sets the "unique_int" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableUniqueInt(i *int) *CommentUpdate {
	if i != nil {
		cu.SetUniqueInt(*i)
	}
	return cu
}

// AddUniqueInt adds i to the "unique_int" field.
func (cu *CommentUpdate) AddUniqueInt(i int) *CommentUpdate {
	cu.mutation.AddUniqueInt(i)
	return cu
}

// SetUniqueFloat sets the "unique_float" field.
func (cu *CommentUpdate) SetUniqueFloat(f float64) *CommentUpdate {
	cu.mutation.ResetUniqueFloat()
	cu.mutation.SetUniqueFloat(f)
	return cu
}

// SetNillableUniqueFloat sets the "unique_float" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableUniqueFloat(f *float64) *CommentUpdate {
	if f != nil {
		cu.SetUniqueFloat(*f)
	}
	return cu
}

// AddUniqueFloat adds f to the "unique_float" field.
func (cu *CommentUpdate) AddUniqueFloat(f float64) *CommentUpdate {
	cu.mutation.AddUniqueFloat(f)
	return cu
}

// SetNillableInt sets the "nillable_int" field.
func (cu *CommentUpdate) SetNillableInt(i int) *CommentUpdate {
	cu.mutation.ResetNillableInt()
	cu.mutation.SetNillableInt(i)
	return cu
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableNillableInt(i *int) *CommentUpdate {
	if i != nil {
		cu.SetNillableInt(*i)
	}
	return cu
}

// AddNillableInt adds i to the "nillable_int" field.
func (cu *CommentUpdate) AddNillableInt(i int) *CommentUpdate {
	cu.mutation.AddNillableInt(i)
	return cu
}

// ClearNillableInt clears the value of the "nillable_int" field.
func (cu *CommentUpdate) ClearNillableInt() *CommentUpdate {
	cu.mutation.ClearNillableInt()
	return cu
}

// SetTable sets the "table" field.
func (cu *CommentUpdate) SetTable(s string) *CommentUpdate {
	cu.mutation.SetTable(s)
	return cu
}

// SetNillableTable sets the "table" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableTable(s *string) *CommentUpdate {
	if s != nil {
		cu.SetTable(*s)
	}
	return cu
}

// ClearTable clears the value of the "table" field.
func (cu *CommentUpdate) ClearTable() *CommentUpdate {
	cu.mutation.ClearTable()
	return cu
}

// SetDir sets the "dir" field.
func (cu *CommentUpdate) SetDir(s schemadir.Dir) *CommentUpdate {
	cu.mutation.SetDir(s)
	return cu
}

// SetNillableDir sets the "dir" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableDir(s *schemadir.Dir) *CommentUpdate {
	if s != nil {
		cu.SetDir(*s)
	}
	return cu
}

// ClearDir clears the value of the "dir" field.
func (cu *CommentUpdate) ClearDir() *CommentUpdate {
	cu.mutation.ClearDir()
	return cu
}

// SetClient sets the "client" field.
func (cu *CommentUpdate) SetClient(s string) *CommentUpdate {
	cu.mutation.SetClient(s)
	return cu
}

// SetNillableClient sets the "client" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableClient(s *string) *CommentUpdate {
	if s != nil {
		cu.SetClient(*s)
	}
	return cu
}

// ClearClient clears the value of the "client" field.
func (cu *CommentUpdate) ClearClient() *CommentUpdate {
	cu.mutation.ClearClient()
	return cu
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.gremlinSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CommentUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := cu.gremlin().Query()
	if err := cu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	cu.mutation.done = true
	return res.ReadInt()
}

func (cu *CommentUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(comment.Label)
	for _, p := range cu.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := cu.mutation.UniqueInt(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, value)
	}
	if value, ok := cu.mutation.AddedUniqueInt(); ok {
		addValue := rv.Clone().Union(__.Values(comment.FieldUniqueInt), __.Constant(value)).Sum().Next()
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, addValue).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, fmt.Sprintf("+= %v", value))),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, __.Union(__.Values(comment.FieldUniqueInt), __.Constant(value)).Sum())
	}
	if value, ok := cu.mutation.UniqueFloat(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, value)
	}
	if value, ok := cu.mutation.AddedUniqueFloat(); ok {
		addValue := rv.Clone().Union(__.Values(comment.FieldUniqueFloat), __.Constant(value)).Sum().Next()
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, addValue).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, fmt.Sprintf("+= %v", value))),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, __.Union(__.Values(comment.FieldUniqueFloat), __.Constant(value)).Sum())
	}
	if value, ok := cu.mutation.NillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, value)
	}
	if value, ok := cu.mutation.AddedNillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, __.Union(__.Values(comment.FieldNillableInt), __.Constant(value)).Sum())
	}
	if value, ok := cu.mutation.Table(); ok {
		v.Property(dsl.Single, comment.FieldTable, value)
	}
	if value, ok := cu.mutation.Dir(); ok {
		v.Property(dsl.Single, comment.FieldDir, value)
	}
	if value, ok := cu.mutation.GetClient(); ok {
		v.Property(dsl.Single, comment.FieldClient, value)
	}
	var properties []any
	if cu.mutation.NillableIntCleared() {
		properties = append(properties, comment.FieldNillableInt)
	}
	if cu.mutation.TableCleared() {
		properties = append(properties, comment.FieldTable)
	}
	if cu.mutation.DirCleared() {
		properties = append(properties, comment.FieldDir)
	}
	if cu.mutation.ClientCleared() {
		properties = append(properties, comment.FieldClient)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
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

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetUniqueInt sets the "unique_int" field.
func (cuo *CommentUpdateOne) SetUniqueInt(i int) *CommentUpdateOne {
	cuo.mutation.ResetUniqueInt()
	cuo.mutation.SetUniqueInt(i)
	return cuo
}

// SetNillableUniqueInt sets the "unique_int" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUniqueInt(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetUniqueInt(*i)
	}
	return cuo
}

// AddUniqueInt adds i to the "unique_int" field.
func (cuo *CommentUpdateOne) AddUniqueInt(i int) *CommentUpdateOne {
	cuo.mutation.AddUniqueInt(i)
	return cuo
}

// SetUniqueFloat sets the "unique_float" field.
func (cuo *CommentUpdateOne) SetUniqueFloat(f float64) *CommentUpdateOne {
	cuo.mutation.ResetUniqueFloat()
	cuo.mutation.SetUniqueFloat(f)
	return cuo
}

// SetNillableUniqueFloat sets the "unique_float" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUniqueFloat(f *float64) *CommentUpdateOne {
	if f != nil {
		cuo.SetUniqueFloat(*f)
	}
	return cuo
}

// AddUniqueFloat adds f to the "unique_float" field.
func (cuo *CommentUpdateOne) AddUniqueFloat(f float64) *CommentUpdateOne {
	cuo.mutation.AddUniqueFloat(f)
	return cuo
}

// SetNillableInt sets the "nillable_int" field.
func (cuo *CommentUpdateOne) SetNillableInt(i int) *CommentUpdateOne {
	cuo.mutation.ResetNillableInt()
	cuo.mutation.SetNillableInt(i)
	return cuo
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableNillableInt(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetNillableInt(*i)
	}
	return cuo
}

// AddNillableInt adds i to the "nillable_int" field.
func (cuo *CommentUpdateOne) AddNillableInt(i int) *CommentUpdateOne {
	cuo.mutation.AddNillableInt(i)
	return cuo
}

// ClearNillableInt clears the value of the "nillable_int" field.
func (cuo *CommentUpdateOne) ClearNillableInt() *CommentUpdateOne {
	cuo.mutation.ClearNillableInt()
	return cuo
}

// SetTable sets the "table" field.
func (cuo *CommentUpdateOne) SetTable(s string) *CommentUpdateOne {
	cuo.mutation.SetTable(s)
	return cuo
}

// SetNillableTable sets the "table" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableTable(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetTable(*s)
	}
	return cuo
}

// ClearTable clears the value of the "table" field.
func (cuo *CommentUpdateOne) ClearTable() *CommentUpdateOne {
	cuo.mutation.ClearTable()
	return cuo
}

// SetDir sets the "dir" field.
func (cuo *CommentUpdateOne) SetDir(s schemadir.Dir) *CommentUpdateOne {
	cuo.mutation.SetDir(s)
	return cuo
}

// SetNillableDir sets the "dir" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableDir(s *schemadir.Dir) *CommentUpdateOne {
	if s != nil {
		cuo.SetDir(*s)
	}
	return cuo
}

// ClearDir clears the value of the "dir" field.
func (cuo *CommentUpdateOne) ClearDir() *CommentUpdateOne {
	cuo.mutation.ClearDir()
	return cuo
}

// SetClient sets the "client" field.
func (cuo *CommentUpdateOne) SetClient(s string) *CommentUpdateOne {
	cuo.mutation.SetClient(s)
	return cuo
}

// SetNillableClient sets the "client" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableClient(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetClient(*s)
	}
	return cuo
}

// ClearClient clears the value of the "client" field.
func (cuo *CommentUpdateOne) ClearClient() *CommentUpdateOne {
	cuo.mutation.ClearClient()
	return cuo
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cuo *CommentUpdateOne) Where(ps ...predicate.Comment) *CommentUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	return withHooks(ctx, cuo.gremlinSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CommentUpdateOne) gremlinSave(ctx context.Context) (*Comment, error) {
	res := &gremlin.Response{}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	query, bindings := cuo.gremlin(id).Query()
	if err := cuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	cuo.mutation.done = true
	c := &Comment{config: cuo.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cuo *CommentUpdateOne) gremlin(id string) *dsl.Traversal {
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
	if value, ok := cuo.mutation.UniqueInt(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, value)
	}
	if value, ok := cuo.mutation.AddedUniqueInt(); ok {
		addValue := rv.Clone().Union(__.Values(comment.FieldUniqueInt), __.Constant(value)).Sum().Next()
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, addValue).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, fmt.Sprintf("+= %v", value))),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, __.Union(__.Values(comment.FieldUniqueInt), __.Constant(value)).Sum())
	}
	if value, ok := cuo.mutation.UniqueFloat(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, value)
	}
	if value, ok := cuo.mutation.AddedUniqueFloat(); ok {
		addValue := rv.Clone().Union(__.Values(comment.FieldUniqueFloat), __.Constant(value)).Sum().Next()
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, addValue).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, fmt.Sprintf("+= %v", value))),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, __.Union(__.Values(comment.FieldUniqueFloat), __.Constant(value)).Sum())
	}
	if value, ok := cuo.mutation.NillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, value)
	}
	if value, ok := cuo.mutation.AddedNillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, __.Union(__.Values(comment.FieldNillableInt), __.Constant(value)).Sum())
	}
	if value, ok := cuo.mutation.Table(); ok {
		v.Property(dsl.Single, comment.FieldTable, value)
	}
	if value, ok := cuo.mutation.Dir(); ok {
		v.Property(dsl.Single, comment.FieldDir, value)
	}
	if value, ok := cuo.mutation.GetClient(); ok {
		v.Property(dsl.Single, comment.FieldClient, value)
	}
	var properties []any
	if cuo.mutation.NillableIntCleared() {
		properties = append(properties, comment.FieldNillableInt)
	}
	if cuo.mutation.TableCleared() {
		properties = append(properties, comment.FieldTable)
	}
	if cuo.mutation.DirCleared() {
		properties = append(properties, comment.FieldDir)
	}
	if cuo.mutation.ClientCleared() {
		properties = append(properties, comment.FieldClient)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if len(cuo.fields) > 0 {
		fields := make([]any, 0, len(cuo.fields)+1)
		fields = append(fields, true)
		for _, f := range cuo.fields {
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