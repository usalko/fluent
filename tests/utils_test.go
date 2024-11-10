package tests

import (
	"testing"

	"github.com/usalko/fluent/flc/gen"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSnakeToSnakeCaseSpecialFromCamel(t *testing.T) {
	name := "Gladys"
	if result := gen.SnakeCase(name); result != "gladys" {
		t.Fatalf(`%s != %s`, name, result)
	}
}
