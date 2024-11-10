// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package blob

import (
	"github.com/google/uuid"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldEQ(FieldUUID, v))
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int) predicate.Blob {
	return predicate.Blob(sql.FieldEQ(FieldCount, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.Blob {
	return predicate.Blob(sql.FieldLTE(FieldUUID, v))
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int) predicate.Blob {
	return predicate.Blob(sql.FieldEQ(FieldCount, v))
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int) predicate.Blob {
	return predicate.Blob(sql.FieldNEQ(FieldCount, v))
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int) predicate.Blob {
	return predicate.Blob(sql.FieldIn(FieldCount, vs...))
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int) predicate.Blob {
	return predicate.Blob(sql.FieldNotIn(FieldCount, vs...))
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int) predicate.Blob {
	return predicate.Blob(sql.FieldGT(FieldCount, v))
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int) predicate.Blob {
	return predicate.Blob(sql.FieldGTE(FieldCount, v))
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int) predicate.Blob {
	return predicate.Blob(sql.FieldLT(FieldCount, v))
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int) predicate.Blob {
	return predicate.Blob(sql.FieldLTE(FieldCount, v))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Blob {
	return predicate.Blob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Blob) predicate.Blob {
	return predicate.Blob(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLinks applies the HasEdge predicate on the "links" edge.
func HasLinks() predicate.Blob {
	return predicate.Blob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LinksTable, LinksPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLinksWith applies the HasEdge predicate on the "links" edge with a given conditions (other predicates).
func HasLinksWith(preds ...predicate.Blob) predicate.Blob {
	return predicate.Blob(func(s *sql.Selector) {
		step := newLinksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBlobLinks applies the HasEdge predicate on the "blob_links" edge.
func HasBlobLinks() predicate.Blob {
	return predicate.Blob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BlobLinksTable, BlobLinksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlobLinksWith applies the HasEdge predicate on the "blob_links" edge with a given conditions (other predicates).
func HasBlobLinksWith(preds ...predicate.BlobLink) predicate.Blob {
	return predicate.Blob(func(s *sql.Selector) {
		step := newBlobLinksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Blob) predicate.Blob {
	return predicate.Blob(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Blob) predicate.Blob {
	return predicate.Blob(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Blob) predicate.Blob {
	return predicate.Blob(sql.NotPredicates(p))
}
