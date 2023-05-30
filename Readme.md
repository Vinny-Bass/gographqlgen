# GoGraphQLGen

GoGraphQLGen is a library that transforms Golang structs into GraphQL schemas. It generates a string representation of your GraphQL schema based on your Golang structs.

## Installation
To use GoGraphQLGen, you need to install the package using go get:

``` bash
go get github.com/Vinny-Bass/gographqlgen
``` 
## Usage
Here's a basic usage example:

```go
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
	gen := graphqlgen.NewGraphqlGen()

	gen.ParseStruct(Author{})
	gen.ParseStruct(Book{})

	schema := gen.GenerateSchema()

	fmt.Println(schema)
}
```

## Testing
To run tests, navigate to the project root and use the go test command:

```bash
go test ./...
```

## Contributing
If you want to contribute, please feel free to fork the repository, make changes, and submit a pull request. I appreciate your help!

## License
GoGraphQLGen is licensed under the MIT License.