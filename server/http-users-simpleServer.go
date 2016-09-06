package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  "net/url"
)

//NEXT:
// Marshall slice, send JSON.
// receive JSON
// search a slice, return an element
// 
// delete from a slice
// update an element in a slice
// http://www.alexedwards.net/blog/golang-response-snippets#json
type User struct {
  Username string
  Password string
  Email string
}
type SavedUser struct {
  id int
  user User
}
var firstUser = User{
  Username: "Daniel",
  Password: "12345",
  Email: "daniel@blacklight.io",
}
var firstSavedUser = SavedUser{
  id: 0,
  user: firstUser,
}
// var Users = make([]User, 0)
var Users []User
var usersCount = 0

func init() {
  fmt.Println("Initializing server...")
  usersCount++
  Users = append(Users, firstUser)
  fmt.Println("Initial Users :", Users)
}

func handler(w http.ResponseWriter, request *http.Request) {
  decoder := json.NewDecoder(request.Body)
  fmt.Println("Method:", request.Method)

  if request.Method == "GET" { //TODO: write all users
    js, err := json.Marshal(Users)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Println("SENDING JS: ", js)
    w.Write(js)

    return
  }

  if request.Method == "POST" { //TODO: create a new user, append

  }

  if request.Method == "PATCH" {//TODO: update an existing user in the slice

  }

  if request.Method == "DELETE" { //TODO: remove a user from the slice

  }

  fmt.Println("DZ____________________")

  u, err := url.Parse(request.RequestURI)
  if err != nil {
      panic(err)
  }
  fmt.Println("Host :", u.Host)
  fmt.Println("Url :", u.Path)


  var t User
  err = decoder.Decode(&t)
  if err != nil {
    panic(err)
  }
  fmt.Println("Test Data : ", t.Username)
  fmt.Println("Test Data : ", t.Password)
  fmt.Println("Test Data : ", t.Email)

  // io.WriteString(w, "Hello world!")
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/users", handler)
  fmt.Println("Listening...")
  http.ListenAndServe(":8000", mux)
}
