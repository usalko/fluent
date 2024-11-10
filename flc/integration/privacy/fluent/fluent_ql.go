// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"github.com/usalko/fluent/flc/integration/privacy/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/task"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/team"
	"github.com/usalko/fluent/flc/integration/privacy/fluent/user"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/fluent_ql"
	"github.com/usalko/fluent/schema/field"
)

// schemaGraph holds a representation of fluent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 3)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
		Type: "Task",
		Fields: map[string]*sqlgraph.FieldSpec{
			task.FieldTitle:       {Type: field.TypeString, Column: task.FieldTitle},
			task.FieldDescription: {Type: field.TypeString, Column: task.FieldDescription},
			task.FieldStatus:      {Type: field.TypeEnum, Column: task.FieldStatus},
			task.FieldUUID:        {Type: field.TypeUUID, Column: task.FieldUUID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   team.Table,
			Columns: team.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: team.FieldID,
			},
		},
		Type: "Team",
		Fields: map[string]*sqlgraph.FieldSpec{
			team.FieldName: {Type: field.TypeString, Column: team.FieldName},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldName: {Type: field.TypeString, Column: user.FieldName},
			user.FieldAge:  {Type: field.TypeUint, Column: user.FieldAge},
		},
	}
	graph.MustAddE(
		"teams",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
		},
		"Task",
		"Team",
	)
	graph.MustAddE(
		"owner",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.OwnerTable,
			Columns: []string{task.OwnerColumn},
			Bidi:    false,
		},
		"Task",
		"User",
	)
	graph.MustAddE(
		"tasks",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   team.TasksTable,
			Columns: team.TasksPrimaryKey,
			Bidi:    false,
		},
		"Team",
		"Task",
	)
	graph.MustAddE(
		"users",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
		},
		"Team",
		"User",
	)
	graph.MustAddE(
		"teams",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
		},
		"User",
		"Team",
	)
	graph.MustAddE(
		"tasks",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
		},
		"User",
		"Task",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (tq *TaskQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TaskQuery builder.
func (tq *TaskQuery) Filter() *TaskFilter {
	return &TaskFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TaskMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an fluent_ql.Where implementation to apply filters on the TaskMutation builder.
func (m *TaskMutation) Filter() *TaskFilter {
	return &TaskFilter{config: m.config, predicateAdder: m}
}

// TaskFilter provides a generic filtering capability at runtime for TaskQuery.
type TaskFilter struct {
	predicateAdder
	config
}

// Where applies the fluent_ql predicate on the query filter.
func (f *TaskFilter) Where(p fluent_ql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the fluent_ql int predicate on the id field.
func (f *TaskFilter) WhereID(p fluent_ql.IntP) {
	f.Where(p.Field(task.FieldID))
}

// WhereTitle applies the fluent_ql string predicate on the title field.
func (f *TaskFilter) WhereTitle(p fluent_ql.StringP) {
	f.Where(p.Field(task.FieldTitle))
}

// WhereDescription applies the fluent_ql string predicate on the description field.
func (f *TaskFilter) WhereDescription(p fluent_ql.StringP) {
	f.Where(p.Field(task.FieldDescription))
}

// WhereStatus applies the fluent_ql string predicate on the status field.
func (f *TaskFilter) WhereStatus(p fluent_ql.StringP) {
	f.Where(p.Field(task.FieldStatus))
}

// WhereUUID applies the fluent_ql [16]byte predicate on the uuid field.
func (f *TaskFilter) WhereUUID(p fluent_ql.ValueP) {
	f.Where(p.Field(task.FieldUUID))
}

// WhereHasTeams applies a predicate to check if query has an edge teams.
func (f *TaskFilter) WhereHasTeams() {
	f.Where(fluent_ql.HasEdge("teams"))
}

// WhereHasTeamsWith applies a predicate to check if query has an edge teams with a given conditions (other predicates).
func (f *TaskFilter) WhereHasTeamsWith(preds ...predicate.Team) {
	f.Where(fluent_ql.HasEdgeWith("teams", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasOwner applies a predicate to check if query has an edge owner.
func (f *TaskFilter) WhereHasOwner() {
	f.Where(fluent_ql.HasEdge("owner"))
}

// WhereHasOwnerWith applies a predicate to check if query has an edge owner with a given conditions (other predicates).
func (f *TaskFilter) WhereHasOwnerWith(preds ...predicate.User) {
	f.Where(fluent_ql.HasEdgeWith("owner", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (tq *TeamQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TeamQuery builder.
func (tq *TeamQuery) Filter() *TeamFilter {
	return &TeamFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TeamMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an fluent_ql.Where implementation to apply filters on the TeamMutation builder.
func (m *TeamMutation) Filter() *TeamFilter {
	return &TeamFilter{config: m.config, predicateAdder: m}
}

// TeamFilter provides a generic filtering capability at runtime for TeamQuery.
type TeamFilter struct {
	predicateAdder
	config
}

// Where applies the fluent_ql predicate on the query filter.
func (f *TeamFilter) Where(p fluent_ql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the fluent_ql int predicate on the id field.
func (f *TeamFilter) WhereID(p fluent_ql.IntP) {
	f.Where(p.Field(team.FieldID))
}

// WhereName applies the fluent_ql string predicate on the name field.
func (f *TeamFilter) WhereName(p fluent_ql.StringP) {
	f.Where(p.Field(team.FieldName))
}

// WhereHasTasks applies a predicate to check if query has an edge tasks.
func (f *TeamFilter) WhereHasTasks() {
	f.Where(fluent_ql.HasEdge("tasks"))
}

// WhereHasTasksWith applies a predicate to check if query has an edge tasks with a given conditions (other predicates).
func (f *TeamFilter) WhereHasTasksWith(preds ...predicate.Task) {
	f.Where(fluent_ql.HasEdgeWith("tasks", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasUsers applies a predicate to check if query has an edge users.
func (f *TeamFilter) WhereHasUsers() {
	f.Where(fluent_ql.HasEdge("users"))
}

// WhereHasUsersWith applies a predicate to check if query has an edge users with a given conditions (other predicates).
func (f *TeamFilter) WhereHasUsersWith(preds ...predicate.User) {
	f.Where(fluent_ql.HasEdgeWith("users", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an fluent_ql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the fluent_ql predicate on the query filter.
func (f *UserFilter) Where(p fluent_ql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the fluent_ql int predicate on the id field.
func (f *UserFilter) WhereID(p fluent_ql.IntP) {
	f.Where(p.Field(user.FieldID))
}

// WhereName applies the fluent_ql string predicate on the name field.
func (f *UserFilter) WhereName(p fluent_ql.StringP) {
	f.Where(p.Field(user.FieldName))
}

// WhereAge applies the fluent_ql uint predicate on the age field.
func (f *UserFilter) WhereAge(p fluent_ql.UintP) {
	f.Where(p.Field(user.FieldAge))
}

// WhereHasTeams applies a predicate to check if query has an edge teams.
func (f *UserFilter) WhereHasTeams() {
	f.Where(fluent_ql.HasEdge("teams"))
}

// WhereHasTeamsWith applies a predicate to check if query has an edge teams with a given conditions (other predicates).
func (f *UserFilter) WhereHasTeamsWith(preds ...predicate.Team) {
	f.Where(fluent_ql.HasEdgeWith("teams", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasTasks applies a predicate to check if query has an edge tasks.
func (f *UserFilter) WhereHasTasks() {
	f.Where(fluent_ql.HasEdge("tasks"))
}

// WhereHasTasksWith applies a predicate to check if query has an edge tasks with a given conditions (other predicates).
func (f *UserFilter) WhereHasTasksWith(preds ...predicate.Task) {
	f.Where(fluent_ql.HasEdgeWith("tasks", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}