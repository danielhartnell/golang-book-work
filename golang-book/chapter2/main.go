// The package declaration
// Every go program starts with a package declaration
package main

// Import allows us to include code from other packages
// fmt is shorthand for "format"
// This package implements formatting for input and output
import "fmt"

// This is a comment

/*
 This is another comment
 With multiple lines!
*/

// Function declaration
// A function has inputs, outputs and statements
// Statements are executed in order
// A function has a name, parameters, an optional return type and a body
func main() {
  // We are calling the Println function and passing in our string
  // This is a statement
  fmt.Println("Hello, my name is Daniel")
}
