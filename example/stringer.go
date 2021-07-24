package example

import "fmt"

const enumStringVarExample = "Unspecifiedba123123rb123arba123r1ba123r2ba1234rbac3trbawetrwrrbactrfasdfrbawctrwertr"

var enumStringVarIndexExample = [...]uint8{0, 11, 20, 26, 33, 40, 47, 53, 62, 73, 84}

func (i Example) StringVal() string {
	if i < 0 || i+1 >= Example(len(enumStringVarIndexExample)) {
		return fmt.Sprintf("%d", i)
	}
	return enumStringVarExample[enumStringVarIndexExample[i]:enumStringVarIndexExample[i+1]]
}

func StringToExample(s string) Example {
	switch s {
	case "ba123123r":
		return Example_exampleBar2
	case "b123ar":
		return Example_exampleBar3
	case "ba123r1":
		return Example_exampleBar4
	case "ba123r2":
		return Example_exampleBar5
	case "ba1234r":
		return Example_exampleBar6
	case "bac3tr":
		return Example_exampleBar7
	case "bawetrwrr":
		return Example_exampleBar8
	case "bactrfasdfr":
		return Example_exampleBar9
	case "bawctrwertr":
		return Example_exampleBar10
	default:
		return 0
	}
}

const enumStringVarBar = "BarBarBarbar1233BarBarBarba123123rBarBarBarb123arBarBarBarba123r1BarBarBarba123r2BarBarBarba1234rBarBarBarbac3trBarBarBarbawetrwrrBarBarBarbactrfasdfrBarBarBarbawctrwertr"

var enumStringVarIndexBar = [...]uint8{0, 16, 34, 49, 65, 81, 97, 112, 130, 150, 170}

func (i Bar) StringVal() string {
	if i < 0 || i+1 >= Bar(len(enumStringVarIndexBar)) {
		return fmt.Sprintf("%d", i)
	}
	return enumStringVarBar[enumStringVarIndexBar[i]:enumStringVarIndexBar[i+1]]
}

func StringToBar(s string) Bar {
	switch s {
	case "BarBarBarba123123r":
		return Bar_BarBar2
	case "BarBarBarb123ar":
		return Bar_BarBar3
	case "BarBarBarba123r1":
		return Bar_BarBar4
	case "BarBarBarba123r2":
		return Bar_BarBar5
	case "BarBarBarba1234r":
		return Bar_BarBar6
	case "BarBarBarbac3tr":
		return Bar_BarBar7
	case "BarBarBarbawetrwrr":
		return Bar_BarBar8
	case "BarBarBarbactrfasdfr":
		return Bar_BarBar9
	case "BarBarBarbawctrwertr":
		return Bar_BarBar10
	default:
		return 0
	}
}
