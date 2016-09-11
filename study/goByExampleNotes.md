GoByExample.com

  go run
  go build
    - builds binaries you can then run

Variables:
  -explicitly declared, use by the compiler
    - type check correctness of function calls

  var
  var b, c int = 1,2
  infer initialized variables...

  zero-value: var e int //e is 0
  := shorthand for declaring and initializing

  declare:  type
  initialize: allocate memory

const:
  - types:
    - (unrelated) concepts from C
      - character: single character; for Go, in Unicode
      - character array: array of characters
      - string: managed array of characters
    - string: managed
    - 
  - can appear anywhere a var statement can
  - perform arithmetic with arbitrary precision
  - no type, unless given one
    - numbers can be given types by using in a context that requires one
  - use: typeof for type

  ```
    tst := "string"
    tst2 := 10
    tst3 := 1.2

    fmt.Println(reflect.TypeOf(tst))
    fmt.Println(reflect.TypeOf(tst2))
    fmt.Println(reflect.TypeOf(tst3))
  ```

for:
  - ;;;
  - range
  - channels
  - other data structures

switch:
  - example: separating multiple expressions:
    switch time.Now().Weekday() {
       case time.Saturday, time.Sunday:
           fmt.Println("it's the weekend")
       default:
           fmt.Println("it's a weekday")
       }

Arrays:
  typed by:
    - elements
    - number
    Create: var a [5]int // default zero valued
    Set: a[4] = 100
    Get: a[4]
    Length: len()
    - can compose -> two-dimnentional
      [2][3]int

Slices:
  - only elements
  - not number

  s := make([]string, 3) //empty
  set/get/len

  append()
  copy(to, from)

  :low
  high:
  declare:  t := []string{"g", "h", "i"}

  *Low Priority*: What about n-dimensional slices? 

  External:
    Golang Blog:
    https://blog.golang.org/go-slices-usage-and-internals

Maps:
  - key value pairs

  m := make(map[string]int) //empty, string keys, int values
  log all with print map
  fmt.Println("map:", m)
  set/get by key
  len for length
  delete(map, "key")

  _, isPresent := m["k2"] // discarded, boolean

  Declare and initialize: 
     n := map[string]int{"foo": 1, "bar": 2}

  External:
    Golang Blog: Maps In Action
    https://blog.golang.org/go-maps-in-action

    Golang Book: Arrays, Slices, and Map
    https://www.golang-book.com/books/intro/6


Range: (STOP: ) https://gobyexample.com/range
  - runs over a variety of data structures...
  ________
Ben Johnson's Blog:
  https://medium.com/@benbjohnson/go-walkthrough-encoding-binary-96dc5d4abb5d#.yib218yxm
  https://medium.com/@benbjohnson/go-walkthrough-encoding-json-package-9681d1d37a8f#.5252cojeg
  https://medium.com/@benbjohnson/go-walkthrough-io-package-8ac5e95a9fbd#.z4t49h3e4
  https://medium.com/@benbjohnson/go-walkthrough-strconv-7a24632a9e73#.unvhd72a9
  https://medium.com/@benbjohnson/go-walkthrough-encoding-package-bc5e912232d#.f7c6fdes9
  https://medium.com/@benbjohnson/go-walkthrough-bytes-strings-packages-499be9f4b5bd#.g4vsfkbyu