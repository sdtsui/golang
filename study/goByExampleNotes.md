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

    gotcha:
      Copying to free up the underlying array:
        func CopyDigits(filename string) []byte {
            b, _ := ioutil.ReadFile(filename)
            b = digitRegexp.Find(b)
            c := make([]byte, len(b))
            copy(c, b)
            return c
        }

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

    declare reference: var m  map[string]int
    initialize: var m = make(map[string]int)
      ||
    map literal:   commits :=map[string]int{ "rsc": 3711}
    get: m[key]
    set: m[key] = 66
    delete: delete(m, "route")

    for key, value := range m {
      // can use _ to discard
    }

    SLICE IS:
      - ptr
      - len
      - cap
        - make([]int, LEN, CAP)
        - cap(s) returns its capacity!
    - len/cap allows manipulation to be as efficient as manipulating array indicies

    Golang Book: Arrays, Slices, and Map
    https://www.golang-book.com/books/intro/6
      - exploiting zero values in maps and slices:
        - if things exist 
          `if visited[n]`
        - to append to nil slices
          likes[l] = append(likes[l], p)
      - key types: comparables
      - NOT COMPARABLE: slices, maps, functions (can't be compared using ==, can't be map keys)
        - bool
        - numeric
        - string
        - pointer
        - channel
        - interface types
        - structs
          ```
          type Key struct {
              Path, Country string
          }
          hits := make(map[Key]int)
          ```
        - arrays that contain only those types

    *maps are not safe for concurrent use*
      mediate accesses with:
        sync.RWMutex
          - counter.Lock(), counter.RLock()
          - do some stuff, read/write
          - counter.Unlock(), counter.RUnlock()

    *map interation order is randomized*
      - must maintain a separate slice to store order if iteration is needed



Range: (STOP: ) https://gobyexample.com/range
  - runs over a variety of data structures...
  for k, v := range SOMEMAP {
    //key, value, 
    // can use _
  }


Functions & Multiple Return Values:

  - typing of multiple consecutive values
    ```
      func plusPlus(a, b, c int) int {
          return a + b + c
      }
    ```
  - returning variable number of arguments:
    return x, y

  - variadic functions:
    - sum(nums ...int) {

    }

  - closures: just like JS
  - recursion: just like JS

Pointers: 
  - Can only change values at memory addresses
    ival = 0 does nothing outside of scope
    *iptr = 0 changes the value at memory address. 
      - can print memory address

Struct: 
  - Typed collection of fields
    - mutable
    - pointer sugar in the language:
      ```
      To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
      ```
  Methods:
    - defined on structs
    - value or receiver types:
      - IF the method needs to modify the receiver, the receiver must be a pointer.
      `func (r *rect) incrementArea(){}`
      `func (r rect) Area() {}

  Interfaces: mechanism for grouping and naming related sets of methods
    - define an interface with method names and return types
    - implement all the methods: "named collection of method signatures"
    - 


    External: Introduction to using interfaces: http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
      - Interfaces are hard. 
      - It's important to know how to design interfaces.
      - The type system is important...and key to using interfaces effectively.
      - Set of Methods, but alo a type.
        - value is of *Two Worlds*:
          - {type, value}
            - method table for value's underlying type
            - data held by that value


        --- breakpoint ---
        DOG/CAT EXAMPLE:
          If you change that one signature, and you try to run the same program exactly as-is (http://play.golang.org/p/TvR758rfre), you will see the following error:

          prog.go:40: cannot use Cat literal (type Cat) as type Animal in array element:
              Cat does not implement Animal (Speak method requires pointer receiver)
          This error message is a bit confusing at first, to be honest. What it’s saying is not that the interface Animal demands that you define your method as a pointer receiver, but that you have tried to convert a Cat struct into an Animal interface value, but only *Cat satisfies that interface. You can fix this bug by passing in a *Cat pointer to the Animal slice instead of a Cat value, by using new(Cat) instead of Cat{} (you could also say &Cat{}, I simply prefer the look of new(Cat)):

          animals := []Animal{Dog{}, new(Cat), Llama{}, JavaProgrammer{}}



        External[Punt]: http://research.swtch.com/interfaces

Errors:  https://gobyexample.com/errors


  ________
Ben Johnson's Blog:
  https://medium.com/@benbjohnson/go-walkthrough-encoding-binary-96dc5d4abb5d#.yib218yxm
  https://medium.com/@benbjohnson/go-walkthrough-encoding-json-package-9681d1d37a8f#.5252cojeg
  https://medium.com/@benbjohnson/go-walkthrough-io-package-8ac5e95a9fbd#.z4t49h3e4
  https://medium.com/@benbjohnson/go-walkthrough-strconv-7a24632a9e73#.unvhd72a9
  https://medium.com/@benbjohnson/go-walkthrough-encoding-package-bc5e912232d#.f7c6fdes9
  https://medium.com/@benbjohnson/go-walkthrough-bytes-strings-packages-499be9f4b5bd#.g4vsfkbyu