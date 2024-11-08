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
// Code generated by entc, DO NOT EDIT.

package friendship

import (
	"time"

	"github.com/usalko/fluent/fluent_gql/internal/todogotype/fluent/predicate"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Friendship {
	return predicate.Friendship(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Friendship {
	return predicate.Friendship(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Friendship {
	return predicate.Friendship(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Friendship {
	return predicate.Friendship(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Friendship {
	return predicate.Friendship(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Friendship {
	return predicate.Friendship(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Friendship {
	return predicate.Friendship(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Friendship {
	return predicate.Friendship(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Friendship {
	return predicate.Friendship(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldCreatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldUserID, v))
}

// FriendID applies equality check predicate on the "friend_id" field. It's identical to FriendIDEQ.
func FriendID(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldFriendID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldLTE(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Friendship {
	return predicate.Friendship(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Friendship {
	return predicate.Friendship(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldContainsFold(FieldUserID, v))
}

// FriendIDEQ applies the EQ predicate on the "friend_id" field.
func FriendIDEQ(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldFriendID, v))
}

// FriendIDNEQ applies the NEQ predicate on the "friend_id" field.
func FriendIDNEQ(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldNEQ(FieldFriendID, v))
}

// FriendIDIn applies the In predicate on the "friend_id" field.
func FriendIDIn(vs ...string) predicate.Friendship {
	return predicate.Friendship(sql.FieldIn(FieldFriendID, vs...))
}

// FriendIDNotIn applies the NotIn predicate on the "friend_id" field.
func FriendIDNotIn(vs ...string) predicate.Friendship {
	return predicate.Friendship(sql.FieldNotIn(FieldFriendID, vs...))
}

// FriendIDGT applies the GT predicate on the "friend_id" field.
func FriendIDGT(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldGT(FieldFriendID, v))
}

// FriendIDGTE applies the GTE predicate on the "friend_id" field.
func FriendIDGTE(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldGTE(FieldFriendID, v))
}

// FriendIDLT applies the LT predicate on the "friend_id" field.
func FriendIDLT(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldLT(FieldFriendID, v))
}

// FriendIDLTE applies the LTE predicate on the "friend_id" field.
func FriendIDLTE(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldLTE(FieldFriendID, v))
}

// FriendIDContains applies the Contains predicate on the "friend_id" field.
func FriendIDContains(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldContains(FieldFriendID, v))
}

// FriendIDHasPrefix applies the HasPrefix predicate on the "friend_id" field.
func FriendIDHasPrefix(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldHasPrefix(FieldFriendID, v))
}

// FriendIDHasSuffix applies the HasSuffix predicate on the "friend_id" field.
func FriendIDHasSuffix(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldHasSuffix(FieldFriendID, v))
}

// FriendIDEqualFold applies the EqualFold predicate on the "friend_id" field.
func FriendIDEqualFold(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldEqualFold(FieldFriendID, v))
}

// FriendIDContainsFold applies the ContainsFold predicate on the "friend_id" field.
func FriendIDContainsFold(v string) predicate.Friendship {
	return predicate.Friendship(sql.FieldContainsFold(FieldFriendID, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFriend applies the HasEdge predicate on the "friend" edge.
func HasFriend() predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FriendTable, FriendColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFriendWith applies the HasEdge predicate on the "friend" edge with a given conditions (other predicates).
func HasFriendWith(preds ...predicate.User) predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := newFriendStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Friendship) predicate.Friendship {
	return predicate.Friendship(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Friendship) predicate.Friendship {
	return predicate.Friendship(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Friendship) predicate.Friendship {
	return predicate.Friendship(sql.NotPredicates(p))
}
