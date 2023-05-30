package gographqlgen

import (
	"bytes"
	"fmt"
	"reflect"
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
func GenerateSchema(t Type) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "type %s {\n", t.Name)
	for _, field := range t.Fields {
		fmt.Fprintf(&buf, "  %s: %s\n", field.Name, field.Type)
	}
	fmt.Fprintf(&buf, "}\n")

	return buf.String()
}

// type Author struct {
// 	Name  string
// 	Books []Book
// }

// type Book struct {
// 	Title  string
// 	Author *Author
// 	name   string
// }

// func main() {
// 	// Parse the structs and generate schemas
// 	authorType := ParseStruct(Author{})
// 	authorSchema := GenerateSchema(authorType)
// 	fmt.Println(authorSchema)

// 	bookType := ParseStruct(Book{})
// 	bookSchema := GenerateSchema(bookType)
// 	fmt.Println(bookSchema)
// }
