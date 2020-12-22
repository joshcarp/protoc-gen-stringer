package foobar

import (
	"fmt"
	"testing"
)

func TestHello_StringVal(t *testing.T) {
	fmt.Println(Hello_Foo.StringVal())
	fmt.Println(Hello_Bar.StringVal())

}