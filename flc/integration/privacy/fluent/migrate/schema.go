// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/dialect/sql/schema"
	"github.com/usalko/fluent/schema/field"
)

var (
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"planned", "in_progress", "closed"}, Default: "planned"},
		{Name: "uuid", Type: field.TypeUUID, Nullable: true},
		{Name: "user_tasks", Type: field.TypeInt, Nullable: true},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_users_tasks",
				Columns:    []*schema.Column{TasksColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "age", Type: field.TypeUint, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// TaskTeamsColumns holds the columns for the "task_teams" table.
	TaskTeamsColumns = []*schema.Column{
		{Name: "task_id", Type: field.TypeInt},
		{Name: "team_id", Type: field.TypeInt},
	}
	// TaskTeamsTable holds the schema information for the "task_teams" table.
	TaskTeamsTable = &schema.Table{
		Name:       "task_teams",
		Columns:    TaskTeamsColumns,
		PrimaryKey: []*schema.Column{TaskTeamsColumns[0], TaskTeamsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "task_teams_task_id",
				Columns:    []*schema.Column{TaskTeamsColumns[0]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "task_teams_team_id",
				Columns:    []*schema.Column{TaskTeamsColumns[1]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserTeamsColumns holds the columns for the "user_teams" table.
	UserTeamsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "team_id", Type: field.TypeInt},
	}
	// UserTeamsTable holds the schema information for the "user_teams" table.
	UserTeamsTable = &schema.Table{
		Name:       "user_teams",
		Columns:    UserTeamsColumns,
		PrimaryKey: []*schema.Column{UserTeamsColumns[0], UserTeamsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_teams_user_id",
				Columns:    []*schema.Column{UserTeamsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_teams_team_id",
				Columns:    []*schema.Column{UserTeamsColumns[1]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TasksTable,
		TeamsTable,
		UsersTable,
		TaskTeamsTable,
		UserTeamsTable,
	}
)

func init() {
	TasksTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.Annotation = &fluent_sql.Annotation{}
	UsersTable.Annotation.Checks = map[string]string{
		"backticks": "`name` IS NOT NULL",
	}
	TaskTeamsTable.ForeignKeys[0].RefTable = TasksTable
	TaskTeamsTable.ForeignKeys[1].RefTable = TeamsTable
	UserTeamsTable.ForeignKeys[0].RefTable = UsersTable
	UserTeamsTable.ForeignKeys[1].RefTable = TeamsTable
}
