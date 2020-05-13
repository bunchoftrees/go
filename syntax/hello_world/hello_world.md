# Hello World!

Standard introduction to Go with the obligatory "Hello World!" program, as seen below:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

In this simple application, we learn about the package declaration:

`package main` defines an executable program.

`import "fmt"` imports the formatted I/O ["fmt"](https://golang.org/pkg/fmt/) package.

Inside of the main function `func main()`, inside of curly brackets, we call the `.Println()` function with a string argument. 

To run this (without building), we can simply call `$ go run hello_world.go` which builds and executes using a temporary location, and then cleans itself up. Using `go build` will build an executable, which can then be run with `$ ./hello_world`. 

For easier iteration and development locally, it's quicker to run using `$ go run ...` and makes testing a bit faster when handling larger applications.

#### additional reading

* [cmd/go (go commands)](https://golang.org/cmd/go/)