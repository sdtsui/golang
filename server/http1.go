package main

import (
  "io"
  "net/http"
  "strconv"
  "fmt"
)

var count = 0

func hello(w http.ResponseWriter, r *http.Request) {
  count++
  fmt.Println("Count:", count);
  io.WriteString(w, strconv.Itoa(count))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", hello)
  http.ListenAndServe(":8000", mux)
}

// https://golang.org/pkg/net/http/#ServeMux
//   mux := http.NewServeMux()
//   mux.HandleFunc("/", hello)

