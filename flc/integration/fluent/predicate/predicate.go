// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package predicate

import (
	"github.com/usalko/fluent/dialect/sql"
)

// Api is the predicate function for api builders.
type Api func(*sql.Selector)

// Builder is the predicate function for builder builders.
type Builder func(*sql.Selector)

// Card is the predicate function for card builders.
type Card func(*sql.Selector)

// Comment is the predicate function for comment builders.
type Comment func(*sql.Selector)

// ExValueScan is the predicate function for ex_value_scan builders.
type ExValueScan func(*sql.Selector)

// ExValueScanOrErr calls the predicate only if the error is not nit.
func ExValueScanOrErr(p ExValueScan, err error) ExValueScan {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// FieldType is the predicate function for field_type builders.
type FieldType func(*sql.Selector)

// File is the predicate function for file builders.
type File func(*sql.Selector)

// FileType is the predicate function for file_type builders.
type FileType func(*sql.Selector)

// Goods is the predicate function for goods builders.
type Goods func(*sql.Selector)

// Group is the predicate function for group builders.
type Group func(*sql.Selector)

// GroupInfo is the predicate function for group_info builders.
type GroupInfo func(*sql.Selector)

// Item is the predicate function for item builders.
type Item func(*sql.Selector)

// License is the predicate function for license builders.
type License func(*sql.Selector)

// Node is the predicate function for node builders.
type Node func(*sql.Selector)

// PC is the predicate function for p_c builders.
type PC func(*sql.Selector)

// Pet is the predicate function for pet builders.
type Pet func(*sql.Selector)

// Spec is the predicate function for spec builders.
type Spec func(*sql.Selector)

// Task is the predicate function for fluenttask builders.
type Task func(*sql.Selector)

// User is the predicate function for user builders.
type User func(*sql.Selector)
