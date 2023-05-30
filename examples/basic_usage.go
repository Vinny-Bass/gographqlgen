package main

import (
	"fmt"

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

func main() {
	author := Author{}
	book := Book{}

	authorType := gographqlgen.ParseStruct(author)
	schema := gographqlgen.GenerateSchema(authorType, false)

	bookType := gographqlgen.ParseStruct(book)
	bookSchema := gographqlgen.GenerateSchema(bookType, false)

	fmt.Println(schema)
	fmt.Println(bookSchema)
}
