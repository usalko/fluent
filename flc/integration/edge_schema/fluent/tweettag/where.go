// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package tweettag

import (
	"time"

	"github.com/google/uuid"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/edge_schema/fluent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldLTE(FieldID, id))
}

// AddedAt applies equality check predicate on the "added_at" field. It's identical to AddedAtEQ.
func AddedAt(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldEQ(FieldAddedAt, v))
}

// TagID applies equality check predicate on the "tag_id" field. It's identical to TagIDEQ.
func TagID(v int) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldEQ(FieldTagID, v))
}

// TweetID applies equality check predicate on the "tweet_id" field. It's identical to TweetIDEQ.
func TweetID(v int) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldEQ(FieldTweetID, v))
}

// AddedAtEQ applies the EQ predicate on the "added_at" field.
func AddedAtEQ(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldEQ(FieldAddedAt, v))
}

// AddedAtNEQ applies the NEQ predicate on the "added_at" field.
func AddedAtNEQ(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldNEQ(FieldAddedAt, v))
}

// AddedAtIn applies the In predicate on the "added_at" field.
func AddedAtIn(vs ...time.Time) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldIn(FieldAddedAt, vs...))
}

// AddedAtNotIn applies the NotIn predicate on the "added_at" field.
func AddedAtNotIn(vs ...time.Time) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldNotIn(FieldAddedAt, vs...))
}

// AddedAtGT applies the GT predicate on the "added_at" field.
func AddedAtGT(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldGT(FieldAddedAt, v))
}

// AddedAtGTE applies the GTE predicate on the "added_at" field.
func AddedAtGTE(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldGTE(FieldAddedAt, v))
}

// AddedAtLT applies the LT predicate on the "added_at" field.
func AddedAtLT(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldLT(FieldAddedAt, v))
}

// AddedAtLTE applies the LTE predicate on the "added_at" field.
func AddedAtLTE(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldLTE(FieldAddedAt, v))
}

// TagIDEQ applies the EQ predicate on the "tag_id" field.
func TagIDEQ(v int) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldEQ(FieldTagID, v))
}

// TagIDNEQ applies the NEQ predicate on the "tag_id" field.
func TagIDNEQ(v int) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldNEQ(FieldTagID, v))
}

// TagIDIn applies the In predicate on the "tag_id" field.
func TagIDIn(vs ...int) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldIn(FieldTagID, vs...))
}

// TagIDNotIn applies the NotIn predicate on the "tag_id" field.
func TagIDNotIn(vs ...int) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldNotIn(FieldTagID, vs...))
}

// TweetIDEQ applies the EQ predicate on the "tweet_id" field.
func TweetIDEQ(v int) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldEQ(FieldTweetID, v))
}

// TweetIDNEQ applies the NEQ predicate on the "tweet_id" field.
func TweetIDNEQ(v int) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldNEQ(FieldTweetID, v))
}

// TweetIDIn applies the In predicate on the "tweet_id" field.
func TweetIDIn(vs ...int) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldIn(FieldTweetID, vs...))
}

// TweetIDNotIn applies the NotIn predicate on the "tweet_id" field.
func TweetIDNotIn(vs ...int) predicate.TweetTag {
	return predicate.TweetTag(sql.FieldNotIn(FieldTweetID, vs...))
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		step := newTagStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTweet applies the HasEdge predicate on the "tweet" edge.
func HasTweet() predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TweetTable, TweetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTweetWith applies the HasEdge predicate on the "tweet" edge with a given conditions (other predicates).
func HasTweetWith(preds ...predicate.Tweet) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		step := newTweetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TweetTag) predicate.TweetTag {
	return predicate.TweetTag(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TweetTag) predicate.TweetTag {
	return predicate.TweetTag(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TweetTag) predicate.TweetTag {
	return predicate.TweetTag(sql.NotPredicates(p))
}
