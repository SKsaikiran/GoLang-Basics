A string is a sequence of characters. For example, **"Hello"** is a string that includes characters: **H , e , l , l , o.**

We use double quotes to represent strings in Go. For example,
```
// using var
var name1 = "Go Programming"

// using shorthand notation
name2 := "Go Programming" 
```
Here, both ``name1`` and ``name2`` are strings with the value **"Go Programming"**.
### Example: Golang String
```
// Program to create a string in Golang

package main
import "fmt"

func main() {

  // create string using var
  var message1 = "Hello,"

  // create string using shorthand notation
  message2 := "Welcome to Programiz"

  fmt.Println(message1)
  fmt.Println(message2)
}
```
### Output
```
Hello,
Welcome to Programiz
```
### Golang String using backtick (` `)
  In **Go**, we can also represent strings using the tick mark notation. For example,
```
Program to represent a string with a backtick

package main
import "fmt"

func main() {

  //  represent string with `  `    
  message := `I love Go Programming`
    
  fmt.Println(message)
}
```
### Output
`` I love Go Programming ``
