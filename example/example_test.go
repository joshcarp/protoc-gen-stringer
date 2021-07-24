package example

import "fmt"

func ExampleFoo_StringVal(){
	fmt.Println(Foo_Foo2.StringVal())
	// Output:
	// Foo2
}

func ExampleStringToFoo(){
	fmt.Println(StringToFoo("Unknown"))
	// Output:
	// Foo_Unknown
}