package gographqlgen

import (
	"fmt"
	"reflect"
	"strings"
)

type Field struct {
	Name string
	Type string
}

type Type struct {
	Name   string
	Fields []Field
}

func uppercaseFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	firstLetter := strings.ToUpper(string(s[0]))
	if len(s) > 1 {
		return firstLetter + s[1:]
	}
	return firstLetter
}

// Function to get name of the type
func getFieldType(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		return uppercaseFirst(getFieldType(t.Elem()))
	} else if t.Kind() == reflect.Slice {
		return fmt.Sprintf("[%s]", getFieldType(t.Elem()))
	}
	return uppercaseFirst(t.Name())
}
