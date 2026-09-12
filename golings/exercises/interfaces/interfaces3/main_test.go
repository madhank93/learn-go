// interfaces3
// Recover the concrete type held by an interface value with a type switch
// (or a comma-ok type assertion).

package main_test

import (
	"fmt"
	"testing"
)

// describe reports v as "int: N" for ints, "string: S" for strings,
// or "unknown" for anything else.
func describe(v any) string {
	// FIXME: switch on the dynamic type (int, string) and format each; default to "unknown".
	switch v.(type) {
	case int:
		return fmt.Sprintf("int: %v", v)
	case string:
		return fmt.Sprintf("string: %v", v)
	default:
		return "unknown"
	}
}

func TestDescribe(t *testing.T) {
	if got := describe(42); got != "int: 42" {
		t.Errorf("int case: got %q", got)
	}
	if got := describe("hi"); got != "string: hi" {
		t.Errorf("string case: got %q", got)
	}
	if got := describe(3.14); got != "unknown" {
		t.Errorf("default case: got %q", got)
	}
}
