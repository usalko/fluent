// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/task"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/team"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/user"
	"github.com/usalko/fluent/schema/field"
	"github.com/google/uuid"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TaskUpdate) SetTitle(s string) *TaskUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableTitle(s *string) *TaskUpdate {
	if s != nil {
		tu.SetTitle(*s)
	}
	return tu
}

// SetDescription sets the "description" field.
func (tu *TaskUpdate) SetDescription(s string) *TaskUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDescription(s *string) *TaskUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TaskUpdate) ClearDescription() *TaskUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetStatus sets the "status" field.
func (tu *TaskUpdate) SetStatus(t task.Status) *TaskUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStatus(t *task.Status) *TaskUpdate {
	if t != nil {
		tu.SetStatus(*t)
	}
	return tu
}

// SetUUID sets the "uuid" field.
func (tu *TaskUpdate) SetUUID(u uuid.UUID) *TaskUpdate {
	tu.mutation.SetUUID(u)
	return tu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableUUID(u *uuid.UUID) *TaskUpdate {
	if u != nil {
		tu.SetUUID(*u)
	}
	return tu
}

// ClearUUID clears the value of the "uuid" field.
func (tu *TaskUpdate) ClearUUID() *TaskUpdate {
	tu.mutation.ClearUUID()
	return tu
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (tu *TaskUpdate) AddTeamIDs(ids ...int) *TaskUpdate {
	tu.mutation.AddTeamIDs(ids...)
	return tu
}

// AddTeams adds the "teams" edges to the Team entity.
func (tu *TaskUpdate) AddTeams(t ...*Team) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTeamIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (tu *TaskUpdate) SetOwnerID(id int) *TaskUpdate {
	tu.mutation.SetOwnerID(id)
	return tu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (tu *TaskUpdate) SetNillableOwnerID(id *int) *TaskUpdate {
	if id != nil {
		tu = tu.SetOwnerID(*id)
	}
	return tu
}

// SetOwner sets the "owner" edge to the User entity.
func (tu *TaskUpdate) SetOwner(u *User) *TaskUpdate {
	return tu.SetOwnerID(u.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (tu *TaskUpdate) ClearTeams() *TaskUpdate {
	tu.mutation.ClearTeams()
	return tu
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (tu *TaskUpdate) RemoveTeamIDs(ids ...int) *TaskUpdate {
	tu.mutation.RemoveTeamIDs(ids...)
	return tu
}

// RemoveTeams removes "teams" edges to Team entities.
func (tu *TaskUpdate) RemoveTeams(t ...*Team) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTeamIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (tu *TaskUpdate) ClearOwner() *TaskUpdate {
	tu.mutation.ClearOwner()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Task.title": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(task.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.UUID(); ok {
		_spec.SetField(task.FieldUUID, field.TypeUUID, value)
	}
	if tu.mutation.UUIDCleared() {
		_spec.ClearField(task.FieldUUID, field.TypeUUID)
	}
	if tu.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !tu.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.OwnerTable,
			Columns: []string{task.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.OwnerTable,
			Columns: []string{task.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetTitle sets the "title" field.
func (tuo *TaskUpdateOne) SetTitle(s string) *TaskUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTitle(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetTitle(*s)
	}
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TaskUpdateOne) SetDescription(s string) *TaskUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDescription(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TaskUpdateOne) ClearDescription() *TaskUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TaskUpdateOne) SetStatus(t task.Status) *TaskUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStatus(t *task.Status) *TaskUpdateOne {
	if t != nil {
		tuo.SetStatus(*t)
	}
	return tuo
}

// SetUUID sets the "uuid" field.
func (tuo *TaskUpdateOne) SetUUID(u uuid.UUID) *TaskUpdateOne {
	tuo.mutation.SetUUID(u)
	return tuo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableUUID(u *uuid.UUID) *TaskUpdateOne {
	if u != nil {
		tuo.SetUUID(*u)
	}
	return tuo
}

// ClearUUID clears the value of the "uuid" field.
func (tuo *TaskUpdateOne) ClearUUID() *TaskUpdateOne {
	tuo.mutation.ClearUUID()
	return tuo
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (tuo *TaskUpdateOne) AddTeamIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.AddTeamIDs(ids...)
	return tuo
}

// AddTeams adds the "teams" edges to the Team entity.
func (tuo *TaskUpdateOne) AddTeams(t ...*Team) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTeamIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (tuo *TaskUpdateOne) SetOwnerID(id int) *TaskUpdateOne {
	tuo.mutation.SetOwnerID(id)
	return tuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOwnerID(id *int) *TaskUpdateOne {
	if id != nil {
		tuo = tuo.SetOwnerID(*id)
	}
	return tuo
}

// SetOwner sets the "owner" edge to the User entity.
func (tuo *TaskUpdateOne) SetOwner(u *User) *TaskUpdateOne {
	return tuo.SetOwnerID(u.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (tuo *TaskUpdateOne) ClearTeams() *TaskUpdateOne {
	tuo.mutation.ClearTeams()
	return tuo
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (tuo *TaskUpdateOne) RemoveTeamIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.RemoveTeamIDs(ids...)
	return tuo
}

// RemoveTeams removes "teams" edges to Team entities.
func (tuo *TaskUpdateOne) RemoveTeams(t ...*Team) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTeamIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (tuo *TaskUpdateOne) ClearOwner() *TaskUpdateOne {
	tuo.mutation.ClearOwner()
	return tuo
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Task.title": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(task.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.UUID(); ok {
		_spec.SetField(task.FieldUUID, field.TypeUUID, value)
	}
	if tuo.mutation.UUIDCleared() {
		_spec.ClearField(task.FieldUUID, field.TypeUUID)
	}
	if tuo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !tuo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.OwnerTable,
			Columns: []string{task.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.OwnerTable,
			Columns: []string{task.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
