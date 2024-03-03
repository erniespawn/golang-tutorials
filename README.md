# golang-tutorials

// https://www.youtube.com/watch?v=YzLrWHZa-Kc&t=3848s

-
- 
- time
- strings functions 
-
- while loop
- Variables 
- convert float to string


## for loop

```go





```



## Printf data types
```go
package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolen
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses base on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n",
		"tea", 1, 'A', 3.14, true, 2, 5)
}

tea 1 A 3.140000 true 2 5

```


## runes
```go
package main

import (
	"fmt"
	"unicode/utf8"
)

var p1 = fmt.Println

func main() {
	rStr := "abcdef"
	p1("Rune count :", utf8.RuneCountInString(rStr))
	// for loop
	for i, runeVal := range rStr {
		// Dont forget the "\n" at the end
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}


Rune count : 6
0 : U+0061 'a' : a
1 : U+0062 'b' : b
2 : U+0063 'c' : c
3 : U+0064 'd' : d
4 : U+0065 'e' : e
5 : U+0066 'f' : f

```


## Strings functions
```go
package main

import (
	"fmt"
	"strings"
)

var p1 = fmt.Println

func main() {
	// Showing how to replace a word
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	p1(sV2)

	// Showing how to split
	p1("Split :", strings.Split("a-b-c-d", "-"))

	// Return bool if has Prefix
	p1("Prefix :", strings.HasPrefix("tacocat", "taco"))

	// Return bool if has Suffix
	p1("Suffix :", strings.HasSuffix("tacocat", "cat"))

	// Replace into Lower
	p1("Lower :", strings.ToLower("Can YOU Fix This?"))

	// Trim Spaces  \n\t
	sV3 := "\nSome stings of words   this\n"
	p1("Trim spaces :", strings.TrimSpace(sV3))

	// Replace characters
	st1 := "Listen to this music"
	p1("Replace :", strings.Replace(st1, "music", "noice", -1))
}

Another word
Split : [a b c d]
Prefix : true
Suffix : true
Lower : can you fix this?
Trim spaces : Some stings of words   this
Replace : Listen to this noice

```




## Data Types
```go
package main

import (
	"fmt"
	"reflect"
)

var p1 = fmt.Println

func main() {
	p1(reflect.TypeOf(25))
	p1(reflect.TypeOf(3.12))
	p1(reflect.TypeOf(true))
	p1(reflect.TypeOf("Hello"))
	p1(reflect.TypeOf('😀'))
}

int
float64
bool
string
int32

```

## User input
```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var p1 = fmt.Println

func main() {
	p1("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		p1("Hello", name)
	} else {
		log.Fatal(err)
	}
}
```