// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package groupinfo

import (
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/fluent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldLTE(FieldID, id))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldEQ(FieldDesc, v))
}

// MaxUsers applies equality check predicate on the "max_users" field. It's identical to MaxUsersEQ.
func MaxUsers(v int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldEQ(FieldMaxUsers, v))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldHasSuffix(FieldDesc, v))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldContainsFold(FieldDesc, v))
}

// MaxUsersEQ applies the EQ predicate on the "max_users" field.
func MaxUsersEQ(v int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldEQ(FieldMaxUsers, v))
}

// MaxUsersNEQ applies the NEQ predicate on the "max_users" field.
func MaxUsersNEQ(v int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldNEQ(FieldMaxUsers, v))
}

// MaxUsersIn applies the In predicate on the "max_users" field.
func MaxUsersIn(vs ...int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldIn(FieldMaxUsers, vs...))
}

// MaxUsersNotIn applies the NotIn predicate on the "max_users" field.
func MaxUsersNotIn(vs ...int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldNotIn(FieldMaxUsers, vs...))
}

// MaxUsersGT applies the GT predicate on the "max_users" field.
func MaxUsersGT(v int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldGT(FieldMaxUsers, v))
}

// MaxUsersGTE applies the GTE predicate on the "max_users" field.
func MaxUsersGTE(v int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldGTE(FieldMaxUsers, v))
}

// MaxUsersLT applies the LT predicate on the "max_users" field.
func MaxUsersLT(v int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldLT(FieldMaxUsers, v))
}

// MaxUsersLTE applies the LTE predicate on the "max_users" field.
func MaxUsersLTE(v int) predicate.GroupInfo {
	return predicate.GroupInfo(sql.FieldLTE(FieldMaxUsers, v))
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, GroupsTable, GroupsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		step := newGroupsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfo(sql.NotPredicates(p))
}
