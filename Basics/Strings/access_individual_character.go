// Program to access individual character of string

package main
import "fmt"
 
func main() {

  // create and initialize a string
  name := "SUPman"

  // access first character
  fmt.Printf("%c\n", name[0])  // S

  // access fourth character
  fmt.Printf("%c\n", name[3])  // m

  // access last character
  fmt.Printf("%c", name[5])  // n
}
