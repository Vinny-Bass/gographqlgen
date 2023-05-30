package gographqlgen

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Parses Golang struct into intermediate representation.
func ParseStruct(s interface{}) Type {
	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fields := make([]Field, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fields[i] = Field{
			Name: field.Name,
			Type: getFieldType(field.Type),
		}
	}

	return Type{
		Name:   t.Name(),
		Fields: fields,
	}
}

// Converts intermediate representation into GraphQL schema.
func GenerateSchema(t Type, trimSpace bool) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "type %s {\n", t.Name)
	for _, field := range t.Fields {
		fmt.Fprintf(&buf, "  %s: %s\n", field.Name, field.Type)
	}
	fmt.Fprintf(&buf, "}\n")

	if !trimSpace {
		return buf.String()
	}

	return strings.Join(strings.Fields(buf.String()), "")
}
