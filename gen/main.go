package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Model struct {
	structDefs []StructDef
}

type StructDef struct {
	name         string
	fields       []StructField
	alias        string
	pointerAlias string
}

func (def StructDef) hasDwSize() bool {
	for _, f := range def.fields {
		if f.name == "dwSize" {
			return true
		}
	}
	return false
}

type StructField struct {
	typeName  string
	pointers  int
	name      string
	arraySize string
}

func (f StructField) isArray() bool {
	return f.arraySize != ""
}

const USE_FILE = true

func main() {
	write := fmt.Print
	writeln := fmt.Println
	writef := fmt.Printf

	if USE_FILE {
		file, err := os.Create("../structs.go")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		write = func(a ...interface{}) (int, error) {
			return fmt.Fprint(file, a...)
		}
		writeln = func(a ...interface{}) (int, error) {
			return fmt.Fprintln(file, a...)
		}
		writef = func(format string, a ...interface{}) (int, error) {
			return fmt.Fprintf(file, format, a...)
		}
	}

	preprocessModel()

	write(`package di8

/*
#include "include.h"

` + generateSizeOfDefinitions() + `*/
import "C"
import "unsafe"`)
	for _, def := range model.structDefs {
		writeln()
		writef("type %v struct {\n", trimDI(def.name))

		for _, field := range def.fields {
			if field.name == "dwSize" {
				continue
			}

			writef("\t%v %v\n",
				cToGoVarName(field.name),
				cToGoTypeName(field.typeName, field.arraySize),
			)
		}

		writeln("}\n")
	}
}

func preprocessModel() {
	return
	for _, def := range model.structDefs {
		for i, field := range def.fields {
			if strings.HasPrefix(field.typeName, "LP") {
				if field.pointers > 0 {
					panic("more than one indirection")
				}
				def.fields[i].pointers++
				def.fields[i].typeName = strings.TrimPrefix(
					def.fields[i].typeName, "LP")
			}
		}
	}
}

func generateSizeOfDefinitions() string {
	sizeofs := ""
	for _, def := range model.structDefs {
		if def.name != def.alias {
			panic(def.name + " != " + def.alias)
		}

		if def.hasDwSize() {
			sizeofs += fmt.Sprintf("#define sizeof_%v sizeof(%v)\n",
				def.name, def.name)
		}
	}
	return sizeofs
}

func cToGoVarName(cName string) string {
	if strings.HasPrefix(cName, "cb") {
		return strings.TrimPrefix(cName, "cb") + "Size"
	}
	if strings.HasPrefix(cName, "dwc") {
		return strings.TrimPrefix(cName, "dwc") + "Size"
	}
	if cName == "cAxes" {
		return "AxesCount"
	}
	if strings.HasPrefix(cName, "guid") {
		return toFirstUpper(cName)
	}

	_, goName := splitCName(cName)
	if strings.HasPrefix(goName, "Di") && unicode.IsUpper(rune(goName[2])) {
		goName = strings.TrimPrefix(goName, "Di")
	}
	return toFirstUpper(goName)
}

func trimDI(s string) string {
	return strings.TrimPrefix(s, "DI")
}

var ctogo = map[string]string{
	"DWORD":     "uint32",
	"CHAR":      "byte",
	"LPDWORD":   "*uint32",
	"WORD":      "uint16",
	"LPVOID":    "unsafe.Pointer",
	"LONG":      "int32",
	"LPLONG":    "*int32",
	"BYTE":      "byte",
	"WCHAR":     "uint16",
	"UINT_PTR":  "uintptr",
	"HWND":      "unsafe.Pointer",
	"HINSTANCE": "unsafe.Pointer",
	"TCHAR":     "byte",
	"D3DCOLOR":  "uint32",
	"IUnknown":  "unsafe.Pointer",
	"LPTSTR":    "string",
}

func cToGoTypeName(cName, arraySize string) string {
	name := cToGoTypeNameWithoutArray(cName)
	if len(arraySize) == 0 {
		return name
	}
	if arraySize == "MAX_PATH" {
		arraySize = "260"
	}
	return "[" + arraySize + "]" + name
}

func cToGoTypeNameWithoutArray(cName string) string {
	if goName, ok := ctogo[cName]; ok {
		return goName
	}
	if strings.HasPrefix(cName, "DI") {
		return strings.TrimPrefix(cName, "DI")
	}
	if strings.HasPrefix(cName, "LPDI") {
		return "*" + strings.TrimPrefix(cName, "LPDI")
	}
	if strings.HasPrefix(cName, "LPCDI") {
		return "*" + strings.TrimPrefix(cName, "LPCDI")
	}
	return cName
}

func toFirstUpper(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func splitCName(s string) (hungarian, name string) {
	for i, r := range s {
		if unicode.IsUpper(r) {
			return s[:i], s[i:]
		}
	}
	return "", s
}

/* TODO and NOTE s

CPOINT is deprecated; it is used in DIPROPCPOINTS
DIPROPCPOINTS  is also deprecated

*/
