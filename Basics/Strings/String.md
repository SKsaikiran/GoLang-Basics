<h1 align="center">Strings</h1>

A string is a sequence of characters. For example, **"Hello"** is a string that includes characters: **H , e , l , l , o.**

We use double quotes to represent strings in Go. For example,
```
// using var
var name1 = "Go Programming"

// using shorthand notation
name2 := "Go Programming" 
```
Here, both ``name1`` and ``name2`` are strings with the value **"Go Programming"**.
## Example: Golang String
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

---------------------------------------------------
## Golang String using backtick (` `)
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

--------------------------------------------------------------------------
## Access Characters of String in Go
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
---------------------------------------------------------------------------
## Find the length of a string
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

-------------------------------------------------------------------------

## Join Two Strings Together
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
## Golang String Methods
In Go, the strings package provides various methods that can be used to perform different operations on strings.
|Functions|Descriptions|
|:-------|:-------|
|``Compare()``| Compares two strings  |
|``Contains()``| Checks if a substring is present inside a string  |
|``Replaces()``| Replaces a substring with another substring  |
|``ToLower() ``|  converts a string to lowercase |
|``ToUpper()``|  converts a string to uppercase |
|``Split()``|  splits a string into multiple substrings |

To use these methods, we must first import the strings package in our code.
```
import (
  "fmt"
  "strings"
)
```
----------------------------------------
## Compare Two Strings in Go
We use the ``Compare()`` of the ``strings`` package to compare two strings. For example,
```
// Program to compare string using Compare()

package main
import (
  "fmt"
  "strings"
)

func main() {

  // create three strings
  string1 := "Programiz"
  string2 := "Programiz Pro"
  string3 := "Programiz"

  // compare strings
  fmt.Println(strings.Compare(string1, string2))  // -1
  fmt.Println(strings.Compare(string2, string3))  // 1
  fmt.Println(strings.Compare(string1, string3))  // 0

}
```
Here, we have used
``strings.Compare(string1, string2)``
To compare two strings: ``string1`` and ``string2``. The function returns:
- -1 because string1 is smaller than string2
- 1 because string2 is greater than string3
- 0 because string1 and string3 are equal
```
**Note**: We have imported ``strings`` at the beginning of the program and used ``strings.Compare()`` not ``Compare()``.
```
---------------------------------------------------------------------
## Check if String contains Substring
To check if a substring is present inside a string, we use the ``Contains()`` method of the Go ``strings`` package.

Let's see an example,
```
// Program to illustrate the use of Contains()

package main
import (
  "fmt"
  "strings"
)

func main() {

  text := "Go Programming"
  substring1 := "Go"
  substring2 := "Golang"

  // check if Go is present in Go Programming
  result := strings.Contains(text, substring1)
  fmt.Println(result)

  // check if Golang is present in Go Programming
  result = strings.Contains(text, substring2)
  fmt.Println(result)
}
```
### Output
```
true
false
```
Here, we get the output

- ``true`` because the substring **"Go"** is present inside the string **"Go Programming"**
- ``false`` because the substring **"Golang"** is not present inside the string **"Go Programming"**
-----------------------------------------------------------------------
## Replace a String in Go
To replace a string, we use the ``Replace()`` method present inside the strings package. For example,
```
// Program using Replace() to replace strings

package main
import (
  "fmt"
  "strings"
)

func main() {
    
  text := "car"
  fmt.Println("Old String:", text)
  
  // replace r with t
  replacedText := strings.Replace(text, "r", "t", 1)

  fmt.Println("New String:", replacedText)
}
```
### Output
```
Old String: car
New String: cat
```
Notice the use of the ``Replace()`` method
``strings.Replace(text, "r", "t", 1)``
Here,

- ``text`` - string where we perform the replace operation
- ``"r"`` - old character that needs to be replaced
- ``"t"`` - new character that replaces the old character
- ``1`` - represents how many old characters to be replaced
```
Note : If we need to replace multiple characters, we can change the value of numbers from 1 to any other. For example,

// replace 2 r with 2 a
strings.Replace("Programiz", "r", "R", 2)

// Output: PRogRamiz
```
-----------------------------------------------------------------
## Change Case of Go String
The strings package provides

- ``ToUpper()`` - to change string to uppercase
- ``ToLower()`` - to change string to lowercase
We use the ``ToUpper()`` function to change the given string to uppercase. The function ``ToUpper()`` is provided by the package ``strings``. For example,
```
// Program to convert a string to uppercaseand lowercase

package main
import (
  "fmt"
  "strings"
)

func main() {

  text1 := "go is fun to learn"

  // convert to uppercase
  text1 = strings.ToUpper(text1)

  fmt.Println(text1)

  text2 := "I LOVE GOLANG"

  // convert to lowercase
  text2 = strings.ToLower(text2)
  fmt.Println(text2)
}
```
### Output
```
GO IS FUN TO LEARN
i love golang
```
-------------------------------------------------------------------------------
## Split Strings in Golang
In Go, we can split a string into multiple substrings using the ``Split()`` method. For example,
```
package main
import (
  "fmt"
  "strings"
)

func main() {
  var message = "I Love Golang"
  
  // split string from space " "
  splittedString := strings.Split(message, " ")

  fmt.Println(splittedString)
}

// Output: [I Love Golang]
```
*Notice the code,*
``strings.Split(message, " ")``
Here, we split the string at ``" "``. Hence, we get individual words as output.

The ``Split()`` method returns a slice of all the substrings. In our example, ``[I Love Golang]`` is a slice.

-------------------------------------------
<h1 align="center">Other String Operations</h1>
<details>
<summary>Compare Golang Strings using ==</summary>
<br>
In Go, we can also use the == operator to compare two strings. For example,
  
 ```
  // Program to compare two strings using == operator

package main
import "fmt"

func main() {

// create 2 strings
  string1 := "Programiz"
  string2 := "programiz"

  // compare two strings
  result := string1 == string2

  fmt.Println(result)
}

// Output: false
  ```
  
  The ``==`` operator returns

- ``true`` - if two strings are equal
- ``false`` - if two strings are not equal
  
  ```
  Note: The ``==`` operator is case sensitive. Hence, ``Programiz`` and ``programiz`` are not equal.
  ```
</details>

<details>
<summary>Create Strings from a Slice</summary>
<br>

  In Go, we can also create a string by joining all the elements of a string slice. For example,
  
  ```
  // Program to create a single string from slices of strings using Join()

package main
import (
  "fmt"
  "strings"
)

func main() {

  // create a string slice
  words := []string{"I", "love", "Golang"}

  // join each element of the slice
  message := strings.Join(words, " ")
  fmt.Println(message)
}

// Output: I love Golang
  ```
  
  Here, we have used the ``Join()`` method of the ``strings`` package to join each element of the slice.

</details>

<details>
<summary>Loop Through Strings in Golang</summary>
<br>
We use for range loop to iterate through each character of a Golang string. For example,
  
```
 // Program to iterate through a string

package main
import "fmt"
 
func main() {    
  text := "Golang"
 
  // for range loop to iterate through a string
  for _, character := range text  { 	 
  fmt.Printf("%c\n", character)
  }

} 
```  
### Output 
 ```
G
  
o
  
l
  
a
  
n
  
g 
```   
</details>
