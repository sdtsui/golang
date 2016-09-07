package main

import (
  "net/http"
  "fmt"
  "strings"
  "strconv"
  "encoding/json"
  "net/url"
  // "io"
)

//NEXT:
// receive JSON
// search a slice, return an element
// delete from a slice
// update an element in a slice
// http://www.alexedwards.net/blog/golang-response-snippets#json
type UserSchema struct {
  Username string
  Password string
  Email string
}
type User struct {
  Id int `json:"id"`
  User UserSchema `json:"user"`
}
var firstUserSchema = UserSchema{
  Username: "Daniel",
  Password: "12345",
  Email: "daniel@blacklight.io",
}
var secondUserSchema = UserSchema{
  Username: "Daniel2",
  Password: "54321",
  Email: "dan@blacklight.io",
}
var Users []User
var usersCount = 0

func init() {
  fmt.Println("Initializing server...")
  var NewUser1 = User{usersCount, firstUserSchema}
  Users = append(Users, NewUser1)
  usersCount++

  var NewUser2 = User{usersCount, secondUserSchema}
  Users = append(Users, NewUser2)
  usersCount++

  fmt.Println("Initial Users :", Users)
}

func handler(w http.ResponseWriter, request *http.Request) {
  u, err := url.Parse(request.RequestURI)
  if err != nil {
      panic(err)
  }
  fmt.Println("Host :", u.Host)
  fmt.Println("Path :", u.Path)

  if request.Method == "GET" { //TODO: write all users

    // var id = strings.Split(u.Path, "/")[2]
    // n, err := strconv.Atoi(id)
    var id = strings.Split(u.Path, "/")
    if (len(id) < 3) { //doesn't have ID
      // fmt.Println("SENDING USERS: ", Users)
      js, err := json.Marshal(Users)
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusOK)
      w.Write(js)

      return
    } else if len(id) == 3 { //has ID
      requestedId, convError := strconv.Atoi(id[2])
      if convError != nil {
        http.Error(w, convError.Error(), http.StatusInternalServerError)
        return
      }
      for i := 0; i < len(Users); i++ {
        if (Users[i].Id == requestedId) {
          //match
          js, err := json.Marshal(Users[i])
          if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
          }
          w.Header().Set("Content-Type", "application/json")
          w.WriteHeader(http.StatusOK)
          w.Write(js)
          return
        }
      }

      //no match
      fmt.Println("ID ", requestedId, "Not Found")
      w.WriteHeader(http.StatusNotFound)
      return
    }


    return
  }

  if request.Method == "POST" { //TODO: create a new user, append
    // insert a new element
    decoder := json.NewDecoder(request.Body)
    var u UserSchema
    err = decoder.Decode(&u)
    if err != nil {
      panic(err)
    }
    newUser := User{usersCount, u}
    usersCount++
    Users = append(Users, newUser)
    w.WriteHeader(http.StatusCreated)

    return
  }

  if request.Method == "PATCH" {//TODO: update an existing user in the slice
    // update an element
    return
  }

  if request.Method == "DELETE" { //TODO: remove a user from the slice
    // remove an element
    return
  }
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", handler)
  fmt.Println("Listening...")
  fmt.Println("Users...", Users)
  http.ListenAndServe(":8000", mux)
}
