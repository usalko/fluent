package tests

import (
	"testing"

	"github.com/usalko/fluent/flc/gen"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSnake1(t *testing.T) {
	name := "Gladys"
	if result := gen.SnakeCase(name); result != "gladys" {
		t.Fatalf(`%s != %s`, name, result)
	}
}

func TestSnake2(t *testing.T) {
	name := "IsItAString"
	if result := gen.SnakeCase(name); result != "is_it_astring" {
		t.Fatalf(`%s != %s`, name, result)
	}
}

func TestSnake3(t *testing.T) {
	name := "UUID"
	if result := gen.SnakeCase(name); result != "uuid" {
		t.Fatalf(`%s != %s`, name, result)
	}
}
