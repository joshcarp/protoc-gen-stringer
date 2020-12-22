package main

import (
	"bytes"
	"fmt"
	enumstring "github.com/joshcarp/protoc-gen-stringer/options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func main() {
	protogen.Options{}.Run(GenerateFiles)
}

// includeImport returns true if the proto identified by the import path should be included in the
// generation.
func includeImport(path string) bool {
	return path != "google/protobuf/descriptor.proto"
}

// GenerateFile generates the contents of a index.sysl file.
func GenerateFiles(gen *protogen.Plugin) error {
	var buf bytes.Buffer
	g := gen.NewGeneratedFile("stringer.go", gen.Files[0].GoImportPath)
	p := func(format string, a ...interface{}) (int, error) {
		return buf.Write([]byte(fmt.Sprintf(format, a...)))
	}
	for _, file := range gen.Files {
		if includeImport(*file.Proto.Name) {
			p("package %s\n\n", file.GoPackageName)
			p("import \"fmt\"\n")
			for _, e := range file.Enums {
				enumStringName := "enumStringVar" + e.GoIdent.GoName
				enumIndexName := "enumStringVarIndex" + e.GoIdent.GoName
				p(`const %s ="`, enumStringName)
				for _, val := range e.Values {
					p("%s", enum(e, val))
				}
				p("\"\n")
				p(`var %s = [...]uint8 {`, enumIndexName)
				lastindex := 0
				for i, val := range e.Values {
					p("%d", lastindex)
					lastindex = len(enum(e, val)) + lastindex
					if i != len(e.Values)-1 {
						p(", ")
					}
				}
				p(", %d}\n", lastindex)

				p(`func (i %s) StringVal() string {
	if i < 0 || i+1 >= %s(len(%s)) {
		return fmt.Sprintf("%%d", i)
	}
	return %s[%s[i]:%s[i+1]]
}
`, e.GoIdent.GoName, e.GoIdent.GoName, enumIndexName, enumStringName, enumIndexName, enumIndexName)
			}
		}
	}
	g.P(buf.String())
	return nil
}

func enum(e *protogen.Enum, val *protogen.EnumValue) string {
	return proto.GetExtension(e.Desc.Values().ByNumber(val.Desc.Number()).Options(), enumstring.E_StringVal).(string)
}
