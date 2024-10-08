// Program to compare string using Compare()

package main
import (
  "fmt"
  "strings"
)

func main() {

  // create three strings
  string1 := "GOlang"
  string2 := "GOlang Pro"
  string3 := "GOlang"

  // compare strings
  fmt.Println(strings.Compare(string1, string2))  // -1
  fmt.Println(strings.Compare(string2, string3))  // 1
  fmt.Println(strings.Compare(string1, string3))  // 0

}
