package example

import "fmt"

func ExampleFoo_StringVal(){
	fmt.Println(Foo_Foo2.StringVal())
	// Output:
	// Foo2
}

func ExampleStringToFoo(){
	fmt.Println(StringToFoo("Unknown"))
	fmt.Println(StringToFoo("90mds90fumasd0fuasdf"))
	// Output:
	// Foo_Unknown true
	// Foo1 false
}