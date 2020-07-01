# Learning Go

All the source code in Go is organized into single folder. The location is specified by the environmental variable `GOPATH`. To check which directly is being used we can do.

```bash
echo $GOPATH
```

All the source code goes into the `src` folder. Other folders are `bin` and `pkg`. `bin` folder contains executable binary and `pkg` is used for compiled library files.

## 1.0 Hello world

All go programs start with a package directive. For a single executable function the package has to be main.

```go
package main
```

The execution of all go programs starts with the main function.

```go
func main() {
	fmt.Println("Hello World!")
}
```

But before we can use the `fmt` package, we need to import it.

```go
import (
	"fmt"
)
```

### hello.go

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
```

To run this you can simply do

```bash
go run hello.go
```

To build a binary

```bash
go build hello.go
```

This will create a new file called `hello` which can be simply run typing `./hello`.

## 1.1 Variables

You can create new variables by using the `var` keyword.

```go
var name string
name = "Kumar" // this will assign the name variable to "Kumar"
```

```go
var name int
name = "Kumar" // This will throw error since the variable was declared as string
```

The variable can be assigned when it is declared and the type is infered by the Go compiler.

```go
var name = "Kumar"
```

or

```go
name := "Kumar"
```
