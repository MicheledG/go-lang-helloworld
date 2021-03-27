/*
this helloworld file will be studied answering the following questions:
- how do we run the code in our project?
  "$> go run main.go", this is the command to be launched.
  "go" command allows
  - go build => takes N files, compiles the code in the files leaving the
  - go run => takes N files, compiles the code in the files, executes the obtained results
  - go fmt => formats code in provided files
  - go install => compiles and "installs" a package
  - go get => downloads the raw source code of someoen els's package
  - go test => runs any test associated with the current project
- what does 'package main' mean?
  Package == Project == Workspace  
  Package is a set of commong source code filse, each package's file has to state
  the name of the package
  Types of packages
  - executables
    - once compiled they are executable, they are useful to "do stuff"
	- a package is executable if the package has name "main", indeed "go build" recognize
	  it and creates a runnable file
	- must have a 'func' called main
  - reusable
    - not executed, but imported since they have useful stuff to be used in other packages
    - any other package name different from "main" is not executables
- what does 'import "fmt"' mean?
  - importing stuff defined in another package
  - 'fmt' it is part of the standard library built for go, it allows formatted I/O
- what's that 'func' thing?
  - 'func' is abbreviation for 'function', it declares a new function
- how is the main.go file organized?
  1. package declaration
  2. import other packages that we need
  3. delcare functions, tell Go to do things
*/
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
