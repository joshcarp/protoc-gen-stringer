package protostringer_test

import (
	"fmt"
	"testing"

	"github.com/joshcarp/protoc-gen-stringer/protostringer"
	"github.com/joshcarp/protoc-gen-stringer/tests/double"
)

func TestStringer(t *testing.T) {
	e := double.Example_exampleBar2
	e2 := double.Example_exampleBarEmpty
	fmt.Println(protostringer.String(e))
	fmt.Println(protostringer.String(e2))
	t.FailNow()
}

// func TestEnum(t *testing.T) {
// 	var e double.Example
// 	protostringer.Enum("exampleBarEmpty", &e)
// 	fmt.Println(e)
// 	t.FailNow()
// }
