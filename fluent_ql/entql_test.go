// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package fluent_ql_test

import (
	"strconv"
	"testing"

	"github.com/usalko/fluent/fluent_ql"

	"github.com/stretchr/testify/assert"
)

func TestPString(t *testing.T) {
	tests := []struct {
		P fluent_ql.P
		S string
	}{
		{
			P: fluent_ql.And(
				fluent_ql.FieldEQ("name", "a8m"),
				fluent_ql.FieldIn("org", "fb", "fluent"),
			),
			S: `name == "a8m" && org in ["fb","fluent"]`,
		},
		{
			P: fluent_ql.Or(
				fluent_ql.Not(fluent_ql.FieldEQ("name", "mashraki")),
				fluent_ql.FieldIn("org", "fb", "fluent"),
			),
			S: `!(name == "mashraki") || org in ["fb","fluent"]`,
		},
		{
			P: fluent_ql.HasEdgeWith(
				"groups",
				fluent_ql.HasEdgeWith(
					"admins",
					fluent_ql.Not(fluent_ql.FieldEQ("name", "a8m")),
				),
			),
			S: `has_edge(groups, has_edge(admins, !(name == "a8m")))`,
		},
		{
			P: fluent_ql.And(
				fluent_ql.FieldGT("age", 30),
				fluent_ql.FieldContains("workplace", "fb"),
			),
			S: `age > 30 && contains(workplace, "fb")`,
		},
		{
			P: fluent_ql.Not(fluent_ql.FieldLT("score", 32.23)),
			S: `!(score < 32.23)`,
		},
		{
			P: fluent_ql.And(
				fluent_ql.FieldNil("active"),
				fluent_ql.FieldNotNil("name"),
			),
			S: `active == nil && name != nil`,
		},
		{
			P: fluent_ql.Or(
				fluent_ql.FieldNotIn("id", 1, 2, 3),
				fluent_ql.FieldHasSuffix("name", "admin"),
			),
			S: `id not in [1,2,3] || has_suffix(name, "admin")`,
		},
		{
			P: fluent_ql.EQ(fluent_ql.F("current"), fluent_ql.F("total")).Negate(),
			S: `!(current == total)`,
		},
	}
	for i := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := tests[i].P.String()
			assert.Equal(t, tests[i].S, s)
		})
	}
}
