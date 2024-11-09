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

package category

import (
	"time"

	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/schema/schematype"
	"github.com/usalko/fluent/fluent_gql/internal/todouuid/fluent/predicate"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldID, id))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldText, v))
}

// Config applies equality check predicate on the "config" field. It's identical to ConfigEQ.
func Config(v *schematype.CategoryConfig) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldConfig, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v time.Duration) predicate.Category {
	vc := int64(v)
	return predicate.Category(sql.FieldEQ(FieldDuration, vc))
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v uint64) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldCount, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldText, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldStatus, vs...))
}

// ConfigEQ applies the EQ predicate on the "config" field.
func ConfigEQ(v *schematype.CategoryConfig) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldConfig, v))
}

// ConfigNEQ applies the NEQ predicate on the "config" field.
func ConfigNEQ(v *schematype.CategoryConfig) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldConfig, v))
}

// ConfigIn applies the In predicate on the "config" field.
func ConfigIn(vs ...*schematype.CategoryConfig) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldConfig, vs...))
}

// ConfigNotIn applies the NotIn predicate on the "config" field.
func ConfigNotIn(vs ...*schematype.CategoryConfig) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldConfig, vs...))
}

// ConfigGT applies the GT predicate on the "config" field.
func ConfigGT(v *schematype.CategoryConfig) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldConfig, v))
}

// ConfigGTE applies the GTE predicate on the "config" field.
func ConfigGTE(v *schematype.CategoryConfig) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldConfig, v))
}

// ConfigLT applies the LT predicate on the "config" field.
func ConfigLT(v *schematype.CategoryConfig) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldConfig, v))
}

// ConfigLTE applies the LTE predicate on the "config" field.
func ConfigLTE(v *schematype.CategoryConfig) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldConfig, v))
}

// ConfigIsNil applies the IsNil predicate on the "config" field.
func ConfigIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldConfig))
}

// ConfigNotNil applies the NotNil predicate on the "config" field.
func ConfigNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldConfig))
}

// TypesIsNil applies the IsNil predicate on the "types" field.
func TypesIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldTypes))
}

// TypesNotNil applies the NotNil predicate on the "types" field.
func TypesNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldTypes))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v time.Duration) predicate.Category {
	vc := int64(v)
	return predicate.Category(sql.FieldEQ(FieldDuration, vc))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v time.Duration) predicate.Category {
	vc := int64(v)
	return predicate.Category(sql.FieldNEQ(FieldDuration, vc))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...time.Duration) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Category(sql.FieldIn(FieldDuration, v...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...time.Duration) predicate.Category {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Category(sql.FieldNotIn(FieldDuration, v...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v time.Duration) predicate.Category {
	vc := int64(v)
	return predicate.Category(sql.FieldGT(FieldDuration, vc))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v time.Duration) predicate.Category {
	vc := int64(v)
	return predicate.Category(sql.FieldGTE(FieldDuration, vc))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v time.Duration) predicate.Category {
	vc := int64(v)
	return predicate.Category(sql.FieldLT(FieldDuration, vc))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v time.Duration) predicate.Category {
	vc := int64(v)
	return predicate.Category(sql.FieldLTE(FieldDuration, vc))
}

// DurationIsNil applies the IsNil predicate on the "duration" field.
func DurationIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldDuration))
}

// DurationNotNil applies the NotNil predicate on the "duration" field.
func DurationNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldDuration))
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v uint64) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldCount, v))
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v uint64) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldCount, v))
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...uint64) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldCount, vs...))
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...uint64) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldCount, vs...))
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v uint64) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldCount, v))
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v uint64) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldCount, v))
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v uint64) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldCount, v))
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v uint64) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldCount, v))
}

// CountIsNil applies the IsNil predicate on the "count" field.
func CountIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldCount))
}

// CountNotNil applies the NotNil predicate on the "count" field.
func CountNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldCount))
}

// StringsIsNil applies the IsNil predicate on the "strings" field.
func StringsIsNil() predicate.Category {
	return predicate.Category(sql.FieldIsNull(FieldStrings))
}

// StringsNotNil applies the NotNil predicate on the "strings" field.
func StringsNotNil() predicate.Category {
	return predicate.Category(sql.FieldNotNull(FieldStrings))
}

// HasTodos applies the HasEdge predicate on the "todos" edge.
func HasTodos() predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TodosTable, TodosColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTodosWith applies the HasEdge predicate on the "todos" edge with a given conditions (other predicates).
func HasTodosWith(preds ...predicate.Todo) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := newTodosStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubCategories applies the HasEdge predicate on the "sub_categories" edge.
func HasSubCategories() predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SubCategoriesTable, SubCategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubCategoriesWith applies the HasEdge predicate on the "sub_categories" edge with a given conditions (other predicates).
func HasSubCategoriesWith(preds ...predicate.Category) predicate.Category {
	return predicate.Category(func(s *sql.Selector) {
		step := newSubCategoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Category) predicate.Category {
	return predicate.Category(sql.NotPredicates(p))
}
