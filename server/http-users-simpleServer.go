package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  "net/url"
)


type User struct {
  Username string
  Password string
  Email string
}

// http://www.alexedwards.net/blog/golang-response-snippets#json
// type Users struct { }

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

var Users = make([]User, 0)

var usersCount = 1


func handler(w http.ResponseWriter, request *http.Request) {
  decoder := json.NewDecoder(request.Body)
  fmt.Println("Method:", request.Method)

  if request.Method == "GET" { //TODO: write all users

  }

  if request.Method == "POST" { //TODO: create a new user, append

  }

  if request.Method == "PATCH" {//TODO: update an existing user in the slice

  }

  if request.Method == "DELETE" { //TODO: remove a user from the slice

  }

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


  Users = append(Users, firstUser)
  fmt.Println("FirstUsers:", firstUser, Users)

  // fmt.Println("Current Users...", Users)

  http.ListenAndServe(":8000", mux)
}

// https://golang.org/pkg/net/http/#ServeMux
//   mux := http.NewServeMux()
//   mux.HandleFunc("/", hello)

// a note on lowerCase key names:
// http://stackoverflow.com/questions/11693865/lower-case-key-names-with-json-marshal-in-go