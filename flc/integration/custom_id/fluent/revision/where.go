// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package revision

import (
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Revision {
	return predicate.Revision(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Revision {
	return predicate.Revision(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Revision {
	return predicate.Revision(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Revision {
	return predicate.Revision(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Revision {
	return predicate.Revision(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Revision {
	return predicate.Revision(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Revision {
	return predicate.Revision(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Revision {
	return predicate.Revision(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Revision {
	return predicate.Revision(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Revision {
	return predicate.Revision(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Revision {
	return predicate.Revision(sql.FieldContainsFold(FieldID, id))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Revision) predicate.Revision {
	return predicate.Revision(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Revision) predicate.Revision {
	return predicate.Revision(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Revision) predicate.Revision {
	return predicate.Revision(sql.NotPredicates(p))
}