// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package custom_type

import (
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/flc/integration/migrate/fluentv1/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CustomType {
	return predicate.CustomType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CustomType {
	return predicate.CustomType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CustomType {
	return predicate.CustomType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CustomType {
	return predicate.CustomType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CustomType {
	return predicate.CustomType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CustomType {
	return predicate.CustomType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CustomType {
	return predicate.CustomType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CustomType {
	return predicate.CustomType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CustomType {
	return predicate.CustomType(sql.FieldLTE(FieldID, id))
}

// Custom applies equality check predicate on the "custom" field. It's identical to CustomEQ.
func Custom(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldEQ(FieldCustom, v))
}

// CustomEQ applies the EQ predicate on the "custom" field.
func CustomEQ(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldEQ(FieldCustom, v))
}

// CustomNEQ applies the NEQ predicate on the "custom" field.
func CustomNEQ(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldNEQ(FieldCustom, v))
}

// CustomIn applies the In predicate on the "custom" field.
func CustomIn(vs ...string) predicate.CustomType {
	return predicate.CustomType(sql.FieldIn(FieldCustom, vs...))
}

// CustomNotIn applies the NotIn predicate on the "custom" field.
func CustomNotIn(vs ...string) predicate.CustomType {
	return predicate.CustomType(sql.FieldNotIn(FieldCustom, vs...))
}

// CustomGT applies the GT predicate on the "custom" field.
func CustomGT(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldGT(FieldCustom, v))
}

// CustomGTE applies the GTE predicate on the "custom" field.
func CustomGTE(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldGTE(FieldCustom, v))
}

// CustomLT applies the LT predicate on the "custom" field.
func CustomLT(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldLT(FieldCustom, v))
}

// CustomLTE applies the LTE predicate on the "custom" field.
func CustomLTE(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldLTE(FieldCustom, v))
}

// CustomContains applies the Contains predicate on the "custom" field.
func CustomContains(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldContains(FieldCustom, v))
}

// CustomHasPrefix applies the HasPrefix predicate on the "custom" field.
func CustomHasPrefix(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldHasPrefix(FieldCustom, v))
}

// CustomHasSuffix applies the HasSuffix predicate on the "custom" field.
func CustomHasSuffix(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldHasSuffix(FieldCustom, v))
}

// CustomIsNil applies the IsNil predicate on the "custom" field.
func CustomIsNil() predicate.CustomType {
	return predicate.CustomType(sql.FieldIsNull(FieldCustom))
}

// CustomNotNil applies the NotNil predicate on the "custom" field.
func CustomNotNil() predicate.CustomType {
	return predicate.CustomType(sql.FieldNotNull(FieldCustom))
}

// CustomEqualFold applies the EqualFold predicate on the "custom" field.
func CustomEqualFold(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldEqualFold(FieldCustom, v))
}

// CustomContainsFold applies the ContainsFold predicate on the "custom" field.
func CustomContainsFold(v string) predicate.CustomType {
	return predicate.CustomType(sql.FieldContainsFold(FieldCustom, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CustomType) predicate.CustomType {
	return predicate.CustomType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CustomType) predicate.CustomType {
	return predicate.CustomType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CustomType) predicate.CustomType {
	return predicate.CustomType(sql.NotPredicates(p))
}
