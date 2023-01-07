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
### Access Characters of String in Go
We know that a string is a sequence of characters. Hence, we can access individual characters of a string.

Just like the **Go array**, we use index numbers to access string characters. For example,
```
// Program to access individual character of string

package main
import "fmt"
 
func main() {

  // create and initialize a string
  name := "Programiz"

  // access first character
  fmt.Printf("%c\n", name[0])  // P

  // access fourth character
  fmt.Printf("%c\n", name[3])  // g

  // access last character
  fmt.Printf("%c", name[8])  // z
}
```
Remember a string index starts from **0**, not **1**.

Hence,
- `` name[0] ``- returns the first character of the string
- `` name[3] ``- returns the fourth character
- `` name[8] ``- returns the ninth (last) character

### Find the length of a string
In Go, we use the `` len() `` function to find the length of a string. For example,
```
// Program to count the length of a string

package main
import "fmt"
 
func main() {
 
  // create string
  message := "Welcome to Programiz"
    
  // use len() function to count length
  stringLength := len(message)

  fmt.Println("Length of a string is:", len(stringLength))
 
}
```
#### Output
`` Length of a string is: 20 ``

Here,`` len() `` returns the number of characters present inside the string.

### Join Two Strings Together
In Go, we can use the ``+`` operator to join (concatenates) strings together. For example,
```
// Program to concatenate two strings

package main
import "fmt"

func main() {  
  message1 := "I love"
  message2 := "Go programming"
    
  // concatenation using + operator
  result := message1 + " " + message2

  fmt.Println(result)
}
```
### Output
`` I love Go programming ``

Here, we have used the ``+`` operator to join three strings: ``message1``,``" "``, and ``message2``.
### Golang String Methods
In Go, the strings package provides various methods that can be used to perform different operations on strings.
|Functions|Descriptions|
|:-------|:-------|
|``Compare()``| Compares two strings  |
|``Contains()``| Checks if a substring is present inside a string  |
|``Replaces()``| Replaces a substring with another substring  |
|``ToLower() ``|  converts a string to lowercase |
|``ToUpper()``|  converts a string to uppercase |
|``Split()``|  splits a string into multiple substrings |


