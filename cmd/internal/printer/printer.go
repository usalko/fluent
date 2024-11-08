// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package printer

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/usalko/fluent/flc/gen"

	"github.com/olekukonko/tablewriter"
)

// A Config controls the output of Fprint.
type Config struct {
	io.Writer
}

// Print prints a table description of the graph to the given writer.
func (p Config) Print(g *gen.Graph) {
	for _, n := range g.Nodes {
		p.node(n)
	}
}

// Fprint executes "pretty-printer" on the given writer.
func Fprint(w io.Writer, g *gen.Graph) {
	Config{Writer: w}.Print(g)
}

// node returns description of a type. The format of the description is:
//
//	Type:
//			<Fields Table>
//
//			<Edges Table>
func (p Config) node(t *gen.Type) {
	var (
		b      strings.Builder
		id     []*gen.Field
		table  = tablewriter.NewWriter(&b)
		header = []string{"Field", "Type", "Unique", "Optional", "Nillable", "Default", "UpdateDefault", "Immutable", "StructTag", "Validators", "Comment"}
	)
	b.WriteString(t.Name + ":\n")
	table.SetAutoFormatHeaders(false)
	table.SetHeader(header)
	if t.ID != nil {
		id = append(id, t.ID)
	}
	for _, f := range append(id, t.Fields...) {
		v := reflect.ValueOf(*f)
		row := make([]string, len(header))
		for i := 0; i < len(row)-1; i++ {
			field := v.FieldByNameFunc(func(name string) bool {
				// The first field is mapped from "Name" to "Field".
				return name == "Name" && i == 0 || name == header[i]
			})
			row[i] = fmt.Sprint(field.Interface())
		}
		row[len(row)-1] = f.Comment()
		table.Append(row)
	}
	table.Render()
	table = tablewriter.NewWriter(&b)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{"Edge", "Type", "Inverse", "BackRef", "Relation", "Unique", "Optional", "Comment"})
	for _, e := range t.Edges {
		table.Append([]string{
			e.Name,
			e.Type.Name,
			strconv.FormatBool(e.IsInverse()),
			e.Inverse,
			e.Rel.Type.String(),
			strconv.FormatBool(e.Unique),
			strconv.FormatBool(e.Optional),
			e.Comment(),
		})
	}
	if table.NumLines() > 0 {
		table.Render()
	}
	io.WriteString(p, strings.ReplaceAll(b.String(), "\n", "\n\t")+"\n")
}
