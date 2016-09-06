slices: built on fixed size arrays
  - flexible, extensible
array: foundation. size is part of type.
  - commonly hold storage for a slice


 A slice is not an array. A slice describes a piece of an array.
  - length
  - pointer to an element of an array

illustration:

```
type sliceHeader struct {
    Length        int
    ZerothElement *byte
}

slice := sliceHeader{
    Length:        50,
    ZerothElement: &buffer[100],
}
```

It is Idiomatic to use a pointer receiver for a method that modifies a slice.

func PtrSubtractOneFromLength(slicePtr *[]byte) {
    slice := *slicePtr
    *slicePtr = slice[0 : len(slice)-1]
}

func main() {
    fmt.Println("Before: len(slice) =", len(slice))
    PtrSubtractOneFromLength(&slice)
    fmt.Println("After:  len(slice) =", len(slice))
}