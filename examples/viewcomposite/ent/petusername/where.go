// Code generated by ent, DO NOT EDIT.

package petusername

import (
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/examples/viewcomposite/fluent/predicate"
)

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PetUserName {
	return predicate.PetUserName(sql.FieldContainsFold(FieldName, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PetUserName) predicate.PetUserName {
	return predicate.PetUserName(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PetUserName) predicate.PetUserName {
	return predicate.PetUserName(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PetUserName) predicate.PetUserName {
	return predicate.PetUserName(sql.NotPredicates(p))
}
