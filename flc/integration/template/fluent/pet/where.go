// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package pet

import (
	"time"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/template/fluent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Pet {
	return predicate.Pet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Pet {
	return predicate.Pet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Pet {
	return predicate.Pet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Pet {
	return predicate.Pet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Pet {
	return predicate.Pet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Pet {
	return predicate.Pet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Pet {
	return predicate.Pet(sql.FieldLTE(FieldID, id))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldAge, v))
}

// LicensedAt applies equality check predicate on the "licensed_at" field. It's identical to LicensedAtEQ.
func LicensedAt(v time.Time) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldLicensedAt, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Pet {
	return predicate.Pet(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Pet {
	return predicate.Pet(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Pet {
	return predicate.Pet(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Pet {
	return predicate.Pet(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Pet {
	return predicate.Pet(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Pet {
	return predicate.Pet(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Pet {
	return predicate.Pet(sql.FieldLTE(FieldAge, v))
}

// LicensedAtEQ applies the EQ predicate on the "licensed_at" field.
func LicensedAtEQ(v time.Time) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldLicensedAt, v))
}

// LicensedAtNEQ applies the NEQ predicate on the "licensed_at" field.
func LicensedAtNEQ(v time.Time) predicate.Pet {
	return predicate.Pet(sql.FieldNEQ(FieldLicensedAt, v))
}

// LicensedAtIn applies the In predicate on the "licensed_at" field.
func LicensedAtIn(vs ...time.Time) predicate.Pet {
	return predicate.Pet(sql.FieldIn(FieldLicensedAt, vs...))
}

// LicensedAtNotIn applies the NotIn predicate on the "licensed_at" field.
func LicensedAtNotIn(vs ...time.Time) predicate.Pet {
	return predicate.Pet(sql.FieldNotIn(FieldLicensedAt, vs...))
}

// LicensedAtGT applies the GT predicate on the "licensed_at" field.
func LicensedAtGT(v time.Time) predicate.Pet {
	return predicate.Pet(sql.FieldGT(FieldLicensedAt, v))
}

// LicensedAtGTE applies the GTE predicate on the "licensed_at" field.
func LicensedAtGTE(v time.Time) predicate.Pet {
	return predicate.Pet(sql.FieldGTE(FieldLicensedAt, v))
}

// LicensedAtLT applies the LT predicate on the "licensed_at" field.
func LicensedAtLT(v time.Time) predicate.Pet {
	return predicate.Pet(sql.FieldLT(FieldLicensedAt, v))
}

// LicensedAtLTE applies the LTE predicate on the "licensed_at" field.
func LicensedAtLTE(v time.Time) predicate.Pet {
	return predicate.Pet(sql.FieldLTE(FieldLicensedAt, v))
}

// LicensedAtIsNil applies the IsNil predicate on the "licensed_at" field.
func LicensedAtIsNil() predicate.Pet {
	return predicate.Pet(sql.FieldIsNull(FieldLicensedAt))
}

// LicensedAtNotNil applies the NotNil predicate on the "licensed_at" field.
func LicensedAtNotNil() predicate.Pet {
	return predicate.Pet(sql.FieldNotNull(FieldLicensedAt))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Pet) predicate.Pet {
	return predicate.Pet(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Pet) predicate.Pet {
	return predicate.Pet(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Pet) predicate.Pet {
	return predicate.Pet(sql.NotPredicates(p))
}