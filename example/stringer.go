package example

import "fmt"

const enumStringVarFoo = "Foo1Foo2Unknown"

var enumStringVarIndexFoo = [...]uint8{0, 4, 8, 15}

func (i Foo) StringVal() string {
	if i < 0 || i+1 >= Foo(len(enumStringVarIndexFoo)) {
		return fmt.Sprintf("%d", i)
	}
	return enumStringVarFoo[enumStringVarIndexFoo[i]:enumStringVarIndexFoo[i+1]]
}

func StringToFoo(s string) Foo {
	switch s {
	case "Foo1":
		return Foo_Foo1
	case "Foo2":
		return Foo_Foo2
	case "Unknown":
		return Foo_Foo_Unknown
	default:
		return 0
	}
}

const enumStringVarBar = "Bar1Bar2Unknown"

var enumStringVarIndexBar = [...]uint8{0, 4, 8, 15}

func (i Bar) StringVal() string {
	if i < 0 || i+1 >= Bar(len(enumStringVarIndexBar)) {
		return fmt.Sprintf("%d", i)
	}
	return enumStringVarBar[enumStringVarIndexBar[i]:enumStringVarIndexBar[i+1]]
}

func StringToBar(s string) Bar {
	switch s {
	case "Bar1":
		return Bar_Bar1
	case "Bar2":
		return Bar_Bar2
	case "Unknown":
		return Bar_Bar_Unknown
	default:
		return 0
	}
}
