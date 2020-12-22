package foobar

import "fmt"

const enumStringVarHello = "Unspecifiedba123123rb123arba123rba123rba1234rbac3trbawetrwrrbactrfasdfrbawctrwertr"

var enumStringVarIndexHello = [...]uint8{0, 11, 20, 26, 32, 38, 45, 51, 60, 71, 82}

func (i Hello) StringVal() string {
	if i < 0 || i+1 >= Hello(len(enumStringVarIndexHello)) {
		return fmt.Sprintf("%d", i)
	}
	return enumStringVarHello[enumStringVarIndexHello[i]:enumStringVarIndexHello[i+1]]
}

const enumStringVarBar = "BarBarBarbar1233BarBarBarba123123rBarBarBarb123arBarBarBarba123rBarBarBarba123rBarBarBarba1234rBarBarBarbac3trBarBarBarbawetrwrrBarBarBarbactrfasdfrBarBarBarbawctrwertr"

var enumStringVarIndexBar = [...]uint8{0, 16, 34, 49, 64, 79, 95, 110, 128, 148, 168}

func (i Bar) StringVal() string {
	if i < 0 || i+1 >= Bar(len(enumStringVarIndexBar)) {
		return fmt.Sprintf("%d", i)
	}
	return enumStringVarBar[enumStringVarIndexBar[i]:enumStringVarIndexBar[i+1]]
}
