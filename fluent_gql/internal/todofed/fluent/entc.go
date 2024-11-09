// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/usalko/fluent/flc"
	"github.com/usalko/fluent/flc/gen"
	"github.com/usalko/fluent/fluent_gql"
)

func main() {
	ex, err := fluent_gql.NewExtension(
		fluent_gql.WithConfigPath("./gqlgen.yml"),
	)
	if err != nil {
		log.Fatalf("creating fluent_gql extension: %v", err)
	}
	err = flc.Generate("./fluent/schema", &gen.Config{
		Header: `
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
		`,
	}, flc.Extensions(ex), flc.TemplateDir("./fluent/template"))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
