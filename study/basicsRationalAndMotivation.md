Basics, Rationale, Motivation:
  http://thenewstack.io/a-closer-look-at-golang-from-an-architects-perspective/
    - type embedding
      - composition over inheritance
      - promote code reuse
    - concurrency is first-class
    - great for:
      - distributed
      - concurrent
      - cloud computing environments

  structs similar to classes
  Composition of Inheritance:  Does not support inheritance, but composition of types.

  ```
  type People interface {
    SayHello()
    ToString()
  }
  peopleArr := [...]People{alex, john,jithesh}
  ```

  GOPATH:
  ```
    go get gopkg.in/mgo.v2

  import (
   "gopkg.in/mgo.v2"
   )
  ```
  _____________
  _____________
  _____________
