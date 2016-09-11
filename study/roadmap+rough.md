Roadmap, Rough Notes:


ideas from -- EqJS:
    Values, Types, Operators
    Program Structure
    Functions
    Data Structures: Objects, Arrays
    Higher Order Functions
    How Objects Inherit/Work



Tactical: 
  - string manipulation
  - type casting, conversion
  - How to override a value in an array
  - How to partially override an object in an array
  - How to delete from an array
  - How to increment things
  - For loops, booleans, 
  - return values
  - closure/context

---------------
Idiomatic Basics:
  https://golang.org/doc/faq
        - when should I use pointer to an interface?
        - should I define methods on values or pointers?
        - Make vs new
        - atomic operations and mutees
        - where is my favourite helper function for testing

    ternary:
      if expr {
          n = trueVal
      } else {
          n = falseVal
      }


  Style:
    https://golang.org/doc/effective_go.html
    https://github.com/golang/go/wiki/CodeReviewComments
      - gofmt
      - goimports
---------------



Distirbuted Systems:


  Concurrency:
    http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/
    - Goroutines
    - channels
      https://www.goinggo.net/2014/02/the-nature-of-channels-in-go.html
    - Select
    - Patterns:
      - timing Out
      - Reocvering from failed goroutines
      - Fan-in/Fan-out
      - Workers
      - Request/Response


  Network Communication: 
    - http library
    - RPC (remote procedure call) Patterns: 

  Failure: RAFT, DISTRIBUTED SYSTEMS: 
    https://blog.gopheracademy.com/writing-a-distributed-systems-library/

---------------

Finer Points:

  Advantages/Disadvantages, Generics:
    http://stackoverflow.com/questions/2198529/what-are-the-advantages-and-disadvantages-of-go-programming-language
      - disadvantages: 
        - no support for generics
      - very strongly typed
      - compiles to machine code
      - not object oriented in traditional sense: composition


  "Go is Not Good"
  http://yager.io/programming/go.html
    Generics
    Haskell: "type classes"
    Rust: "traits"
    multiple return
    embedded systems

  FAQ:
    https://golang.org/doc/faq
      - guarantee that a type satisfies an interface

  Details, history: https://talks.golang.org/2012/splash.article

Automated Docs Generation:
  https://godoc.org/golang.org/x/tools/cmd/godoc
---------------
  https://golang.org/doc/faq

    - makes dependency analysis easy
    - no type system heirarchy
    - fully garbage-collected, support for concurent execution/communication
    - support for multicore machines
  _____________
  _____________
  _____________
  How to search:
    `site:golang.org 'name of function'`

  Review Packages that stand out:
    https://golang.org/pkg/
      http
      sort
      strconv
      strings
      reflect
      index, suffixarray (substring search in logarithmic time)
      testing
      fmt
      json
      text
      time
      path
      rpc
      net
      csv
      io
  _____________
  _____________
  _____________





  General Projects Searching...:
    https://godoc.org/







---------------







Redeploy personal site on AWS Lambda function
  - poker chip?
  - http://leebyron.com/

