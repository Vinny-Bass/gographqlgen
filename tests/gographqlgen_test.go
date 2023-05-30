package graphqlgen_test

import (
	"strings"
	"testing"

	"github.com/Vinny-Bass/gographqlgen"
)

type Author struct {
	Name  string
	Books []Book
}

type Book struct {
	Title  string
	Author *Author
}

func TestSchemaGeneration(t *testing.T) {

	authorType := gographqlgen.ParseStruct(Author{})
	authorSchema := gographqlgen.GenerateSchema(authorType, false)

	expectedSchema := `
		type Author {
			Name: String
			Books: [Book]
		}
	`

	if strings.Join(strings.Fields(authorSchema), "") != strings.Join(strings.Fields(expectedSchema), "") {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedSchema, authorSchema)
	}

	expectedSchema = `
		type Book {
			Title: String
			Author: Author
		}
	`
	bookType := gographqlgen.ParseStruct(Book{})
	bookSchema := gographqlgen.GenerateSchema(bookType, true)

	if bookSchema != strings.Join(strings.Fields(expectedSchema), "") {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedSchema, bookSchema)
	}

}
