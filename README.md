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

## 1.2 Arrays

An array can be created by using the following syntax:

```go
age := [3]int{}
```

This will create array with 3 integer. All the elements will be initialized as 0 as that is the zero value for integer type.

We can se the value by using the index

```go
age[2] = 3
```

This will set the third (since Go is zero indexed) element to 3.

> There is no way to change the size of the array in Go. To use a variable sized data structure we can use slice.

## 1.3 Slices

To create a new slice

```go
age := []int{1, 2, 3}
```

To add elements to array

```go
age = append(age, 4, 5) // append will return a new slice
fmt.Println(age) // prints [1 2 3 4 5]
```

To remove element at index 3 from slice:

```go
append(age[:2], age[3:]...) // This will return [1 2 4 5]
```

## 1.2 Conditionals

If statements in Go does not require to be within parenthesis.

```go
age := 5
if age > 3 {
    fmt.Println("age is greater than 3!")
} else {
    fmt.Println("age is less than 3")
}
```
