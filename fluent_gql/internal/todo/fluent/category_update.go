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
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/dialect/sql/sqljson"
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/category"
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/predicate"
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/schema/schematype"
	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/todo"
	"github.com/usalko/fluent/schema/field"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks     []Hook
	mutation  *CategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetText sets the "text" field.
func (cu *CategoryUpdate) SetText(s string) *CategoryUpdate {
	cu.mutation.SetText(s)
	return cu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableText(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetText(*s)
	}
	return cu
}

// SetStatus sets the "status" field.
func (cu *CategoryUpdate) SetStatus(c category.Status) *CategoryUpdate {
	cu.mutation.SetStatus(c)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableStatus(c *category.Status) *CategoryUpdate {
	if c != nil {
		cu.SetStatus(*c)
	}
	return cu
}

// SetConfig sets the "config" field.
func (cu *CategoryUpdate) SetConfig(sc *schematype.CategoryConfig) *CategoryUpdate {
	cu.mutation.SetConfig(sc)
	return cu
}

// ClearConfig clears the value of the "config" field.
func (cu *CategoryUpdate) ClearConfig() *CategoryUpdate {
	cu.mutation.ClearConfig()
	return cu
}

// SetTypes sets the "types" field.
func (cu *CategoryUpdate) SetTypes(st *schematype.CategoryTypes) *CategoryUpdate {
	cu.mutation.SetTypes(st)
	return cu
}

// ClearTypes clears the value of the "types" field.
func (cu *CategoryUpdate) ClearTypes() *CategoryUpdate {
	cu.mutation.ClearTypes()
	return cu
}

// SetDuration sets the "duration" field.
func (cu *CategoryUpdate) SetDuration(t time.Duration) *CategoryUpdate {
	cu.mutation.ResetDuration()
	cu.mutation.SetDuration(t)
	return cu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableDuration(t *time.Duration) *CategoryUpdate {
	if t != nil {
		cu.SetDuration(*t)
	}
	return cu
}

// AddDuration adds t to the "duration" field.
func (cu *CategoryUpdate) AddDuration(t time.Duration) *CategoryUpdate {
	cu.mutation.AddDuration(t)
	return cu
}

// ClearDuration clears the value of the "duration" field.
func (cu *CategoryUpdate) ClearDuration() *CategoryUpdate {
	cu.mutation.ClearDuration()
	return cu
}

// SetCount sets the "count" field.
func (cu *CategoryUpdate) SetCount(u uint64) *CategoryUpdate {
	cu.mutation.ResetCount()
	cu.mutation.SetCount(u)
	return cu
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableCount(u *uint64) *CategoryUpdate {
	if u != nil {
		cu.SetCount(*u)
	}
	return cu
}

// AddCount adds u to the "count" field.
func (cu *CategoryUpdate) AddCount(u int64) *CategoryUpdate {
	cu.mutation.AddCount(u)
	return cu
}

// ClearCount clears the value of the "count" field.
func (cu *CategoryUpdate) ClearCount() *CategoryUpdate {
	cu.mutation.ClearCount()
	return cu
}

// SetStrings sets the "strings" field.
func (cu *CategoryUpdate) SetStrings(s []string) *CategoryUpdate {
	cu.mutation.SetStrings(s)
	return cu
}

// AppendStrings appends s to the "strings" field.
func (cu *CategoryUpdate) AppendStrings(s []string) *CategoryUpdate {
	cu.mutation.AppendStrings(s)
	return cu
}

// ClearStrings clears the value of the "strings" field.
func (cu *CategoryUpdate) ClearStrings() *CategoryUpdate {
	cu.mutation.ClearStrings()
	return cu
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (cu *CategoryUpdate) AddTodoIDs(ids ...int) *CategoryUpdate {
	cu.mutation.AddTodoIDs(ids...)
	return cu
}

// AddTodos adds the "todos" edges to the Todo entity.
func (cu *CategoryUpdate) AddTodos(t ...*Todo) *CategoryUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddTodoIDs(ids...)
}

// AddSubCategoryIDs adds the "sub_categories" edge to the Category entity by IDs.
func (cu *CategoryUpdate) AddSubCategoryIDs(ids ...int) *CategoryUpdate {
	cu.mutation.AddSubCategoryIDs(ids...)
	return cu
}

// AddSubCategories adds the "sub_categories" edges to the Category entity.
func (cu *CategoryUpdate) AddSubCategories(c ...*Category) *CategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddSubCategoryIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearTodos clears all "todos" edges to the Todo entity.
func (cu *CategoryUpdate) ClearTodos() *CategoryUpdate {
	cu.mutation.ClearTodos()
	return cu
}

// RemoveTodoIDs removes the "todos" edge to Todo entities by IDs.
func (cu *CategoryUpdate) RemoveTodoIDs(ids ...int) *CategoryUpdate {
	cu.mutation.RemoveTodoIDs(ids...)
	return cu
}

// RemoveTodos removes "todos" edges to Todo entities.
func (cu *CategoryUpdate) RemoveTodos(t ...*Todo) *CategoryUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveTodoIDs(ids...)
}

// ClearSubCategories clears all "sub_categories" edges to the Category entity.
func (cu *CategoryUpdate) ClearSubCategories() *CategoryUpdate {
	cu.mutation.ClearSubCategories()
	return cu
}

// RemoveSubCategoryIDs removes the "sub_categories" edge to Category entities by IDs.
func (cu *CategoryUpdate) RemoveSubCategoryIDs(ids ...int) *CategoryUpdate {
	cu.mutation.RemoveSubCategoryIDs(ids...)
	return cu
}

// RemoveSubCategories removes "sub_categories" edges to Category entities.
func (cu *CategoryUpdate) RemoveSubCategories(c ...*Category) *CategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveSubCategoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CategoryUpdate) check() error {
	if v, ok := cu.mutation.Text(); ok {
		if err := category.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`fluent: validator failed for field "Category.text": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Status(); ok {
		if err := category.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`fluent: validator failed for field "Category.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CategoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CategoryUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Text(); ok {
		_spec.SetField(category.FieldText, field.TypeString, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(category.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.Config(); ok {
		_spec.SetField(category.FieldConfig, field.TypeOther, value)
	}
	if cu.mutation.ConfigCleared() {
		_spec.ClearField(category.FieldConfig, field.TypeOther)
	}
	if value, ok := cu.mutation.Types(); ok {
		_spec.SetField(category.FieldTypes, field.TypeJSON, value)
	}
	if cu.mutation.TypesCleared() {
		_spec.ClearField(category.FieldTypes, field.TypeJSON)
	}
	if value, ok := cu.mutation.Duration(); ok {
		_spec.SetField(category.FieldDuration, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedDuration(); ok {
		_spec.AddField(category.FieldDuration, field.TypeInt64, value)
	}
	if cu.mutation.DurationCleared() {
		_spec.ClearField(category.FieldDuration, field.TypeInt64)
	}
	if value, ok := cu.mutation.Count(); ok {
		_spec.SetField(category.FieldCount, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedCount(); ok {
		_spec.AddField(category.FieldCount, field.TypeUint64, value)
	}
	if cu.mutation.CountCleared() {
		_spec.ClearField(category.FieldCount, field.TypeUint64)
	}
	if value, ok := cu.mutation.Strings(); ok {
		_spec.SetField(category.FieldStrings, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedStrings(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, category.FieldStrings, value)
		})
	}
	if cu.mutation.StringsCleared() {
		_spec.ClearField(category.FieldStrings, field.TypeJSON)
	}
	if cu.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedTodosIDs(); len(nodes) > 0 && !cu.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.SubCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.SubCategoriesTable,
			Columns: category.SubCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedSubCategoriesIDs(); len(nodes) > 0 && !cu.mutation.SubCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.SubCategoriesTable,
			Columns: category.SubCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SubCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.SubCategoriesTable,
			Columns: category.SubCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetText sets the "text" field.
func (cuo *CategoryUpdateOne) SetText(s string) *CategoryUpdateOne {
	cuo.mutation.SetText(s)
	return cuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableText(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetText(*s)
	}
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CategoryUpdateOne) SetStatus(c category.Status) *CategoryUpdateOne {
	cuo.mutation.SetStatus(c)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableStatus(c *category.Status) *CategoryUpdateOne {
	if c != nil {
		cuo.SetStatus(*c)
	}
	return cuo
}

// SetConfig sets the "config" field.
func (cuo *CategoryUpdateOne) SetConfig(sc *schematype.CategoryConfig) *CategoryUpdateOne {
	cuo.mutation.SetConfig(sc)
	return cuo
}

// ClearConfig clears the value of the "config" field.
func (cuo *CategoryUpdateOne) ClearConfig() *CategoryUpdateOne {
	cuo.mutation.ClearConfig()
	return cuo
}

// SetTypes sets the "types" field.
func (cuo *CategoryUpdateOne) SetTypes(st *schematype.CategoryTypes) *CategoryUpdateOne {
	cuo.mutation.SetTypes(st)
	return cuo
}

// ClearTypes clears the value of the "types" field.
func (cuo *CategoryUpdateOne) ClearTypes() *CategoryUpdateOne {
	cuo.mutation.ClearTypes()
	return cuo
}

// SetDuration sets the "duration" field.
func (cuo *CategoryUpdateOne) SetDuration(t time.Duration) *CategoryUpdateOne {
	cuo.mutation.ResetDuration()
	cuo.mutation.SetDuration(t)
	return cuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableDuration(t *time.Duration) *CategoryUpdateOne {
	if t != nil {
		cuo.SetDuration(*t)
	}
	return cuo
}

// AddDuration adds t to the "duration" field.
func (cuo *CategoryUpdateOne) AddDuration(t time.Duration) *CategoryUpdateOne {
	cuo.mutation.AddDuration(t)
	return cuo
}

// ClearDuration clears the value of the "duration" field.
func (cuo *CategoryUpdateOne) ClearDuration() *CategoryUpdateOne {
	cuo.mutation.ClearDuration()
	return cuo
}

// SetCount sets the "count" field.
func (cuo *CategoryUpdateOne) SetCount(u uint64) *CategoryUpdateOne {
	cuo.mutation.ResetCount()
	cuo.mutation.SetCount(u)
	return cuo
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableCount(u *uint64) *CategoryUpdateOne {
	if u != nil {
		cuo.SetCount(*u)
	}
	return cuo
}

// AddCount adds u to the "count" field.
func (cuo *CategoryUpdateOne) AddCount(u int64) *CategoryUpdateOne {
	cuo.mutation.AddCount(u)
	return cuo
}

// ClearCount clears the value of the "count" field.
func (cuo *CategoryUpdateOne) ClearCount() *CategoryUpdateOne {
	cuo.mutation.ClearCount()
	return cuo
}

// SetStrings sets the "strings" field.
func (cuo *CategoryUpdateOne) SetStrings(s []string) *CategoryUpdateOne {
	cuo.mutation.SetStrings(s)
	return cuo
}

// AppendStrings appends s to the "strings" field.
func (cuo *CategoryUpdateOne) AppendStrings(s []string) *CategoryUpdateOne {
	cuo.mutation.AppendStrings(s)
	return cuo
}

// ClearStrings clears the value of the "strings" field.
func (cuo *CategoryUpdateOne) ClearStrings() *CategoryUpdateOne {
	cuo.mutation.ClearStrings()
	return cuo
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (cuo *CategoryUpdateOne) AddTodoIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.AddTodoIDs(ids...)
	return cuo
}

// AddTodos adds the "todos" edges to the Todo entity.
func (cuo *CategoryUpdateOne) AddTodos(t ...*Todo) *CategoryUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddTodoIDs(ids...)
}

// AddSubCategoryIDs adds the "sub_categories" edge to the Category entity by IDs.
func (cuo *CategoryUpdateOne) AddSubCategoryIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.AddSubCategoryIDs(ids...)
	return cuo
}

// AddSubCategories adds the "sub_categories" edges to the Category entity.
func (cuo *CategoryUpdateOne) AddSubCategories(c ...*Category) *CategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddSubCategoryIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearTodos clears all "todos" edges to the Todo entity.
func (cuo *CategoryUpdateOne) ClearTodos() *CategoryUpdateOne {
	cuo.mutation.ClearTodos()
	return cuo
}

// RemoveTodoIDs removes the "todos" edge to Todo entities by IDs.
func (cuo *CategoryUpdateOne) RemoveTodoIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.RemoveTodoIDs(ids...)
	return cuo
}

// RemoveTodos removes "todos" edges to Todo entities.
func (cuo *CategoryUpdateOne) RemoveTodos(t ...*Todo) *CategoryUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveTodoIDs(ids...)
}

// ClearSubCategories clears all "sub_categories" edges to the Category entity.
func (cuo *CategoryUpdateOne) ClearSubCategories() *CategoryUpdateOne {
	cuo.mutation.ClearSubCategories()
	return cuo
}

// RemoveSubCategoryIDs removes the "sub_categories" edge to Category entities by IDs.
func (cuo *CategoryUpdateOne) RemoveSubCategoryIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.RemoveSubCategoryIDs(ids...)
	return cuo
}

// RemoveSubCategories removes "sub_categories" edges to Category entities.
func (cuo *CategoryUpdateOne) RemoveSubCategories(c ...*Category) *CategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveSubCategoryIDs(ids...)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cuo *CategoryUpdateOne) Where(ps ...predicate.Category) *CategoryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CategoryUpdateOne) check() error {
	if v, ok := cuo.mutation.Text(); ok {
		if err := category.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`fluent: validator failed for field "Category.text": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Status(); ok {
		if err := category.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`fluent: validator failed for field "Category.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CategoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CategoryUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`fluent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Text(); ok {
		_spec.SetField(category.FieldText, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(category.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.Config(); ok {
		_spec.SetField(category.FieldConfig, field.TypeOther, value)
	}
	if cuo.mutation.ConfigCleared() {
		_spec.ClearField(category.FieldConfig, field.TypeOther)
	}
	if value, ok := cuo.mutation.Types(); ok {
		_spec.SetField(category.FieldTypes, field.TypeJSON, value)
	}
	if cuo.mutation.TypesCleared() {
		_spec.ClearField(category.FieldTypes, field.TypeJSON)
	}
	if value, ok := cuo.mutation.Duration(); ok {
		_spec.SetField(category.FieldDuration, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedDuration(); ok {
		_spec.AddField(category.FieldDuration, field.TypeInt64, value)
	}
	if cuo.mutation.DurationCleared() {
		_spec.ClearField(category.FieldDuration, field.TypeInt64)
	}
	if value, ok := cuo.mutation.Count(); ok {
		_spec.SetField(category.FieldCount, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedCount(); ok {
		_spec.AddField(category.FieldCount, field.TypeUint64, value)
	}
	if cuo.mutation.CountCleared() {
		_spec.ClearField(category.FieldCount, field.TypeUint64)
	}
	if value, ok := cuo.mutation.Strings(); ok {
		_spec.SetField(category.FieldStrings, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedStrings(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, category.FieldStrings, value)
		})
	}
	if cuo.mutation.StringsCleared() {
		_spec.ClearField(category.FieldStrings, field.TypeJSON)
	}
	if cuo.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedTodosIDs(); len(nodes) > 0 && !cuo.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.SubCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.SubCategoriesTable,
			Columns: category.SubCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedSubCategoriesIDs(); len(nodes) > 0 && !cuo.mutation.SubCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.SubCategoriesTable,
			Columns: category.SubCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SubCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.SubCategoriesTable,
			Columns: category.SubCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
