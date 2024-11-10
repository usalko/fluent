// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/usalko/fluent/flc"
	"github.com/usalko/fluent/flc/gen"
)

func main() {
	err := flc.Generate("./schema", &gen.Config{
		Header: `
			// This source code is licensed under the Apache 2.0 license found
			// in the LICENSE file in the root directory of this source tree.

			// Code generated by ent, DO NOT EDIT.
		`,
		Features: []gen.Feature{
			gen.FeatureEntQL,
			gen.FeatureUpsert,
			gen.FeaturePrivacy,
			gen.FeatureSnapshot,
		},
	})
	if err != nil {
		log.Fatalf("running fluent codegen: %v", err)
	}
}
