package main

import (
	"bytes"
	"fmt"
	enumstring "github.com/joshcarp/protoc-gen-stringer/options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"log"
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
	p := func(format string, a ...interface{}) {
		_, _ = buf.Write([]byte(fmt.Sprintf(format, a...)))
	}
	for _, filename := range gen.Request.GetFileToGenerate() {
		file := gen.FilesByPath[filename]
		if includeImport(*file.Proto.Name) {
			p("package %s\n\n", file.GoPackageName)
			p("import \"fmt\"\n")
			for _, e := range file.Enums {
				enumStringName := "enumStringVar" + e.GoIdent.GoName
				enumIndexName := "enumStringVarIndex" + e.GoIdent.GoName
				p(`const %s ="`, enumStringName)
				uniqueStrings := make(map[string]bool)
				for _, val := range e.Values {
					if uniqueStrings[enum(e, val)] {
						log.Fatal("error: duplicate enum string ", enum(e, val))
					}
					uniqueStrings[enum(e, val)] = true
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
			generateStringToEnum(p, e)
			}
		}
	}
	g.P(buf.String())
	return nil
}

func enum(e *protogen.Enum, val *protogen.EnumValue) string {
	return proto.GetExtension(e.Desc.Values().ByNumber(val.Desc.Number()).Options(), enumstring.E_StringVal).(string)
}

func generateStringToEnum(p func (string, ...interface{}), e *protogen.Enum){
	p("\nfunc StringTo%s(s string) %s {\n", e.GoIdent.GoName, e.GoIdent.GoName)
	p("switch s {\n")

	for _, val := range e.Values{
		p("case \"%s\":\n", enum(e, val))
		p("return %s\n", val.GoIdent.GoName)
	}
	p("default:\n")
	p("return 0\n")
	p("}\n")
	p("}\n")
}
