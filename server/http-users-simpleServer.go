package main

import (
  "net/http"
  "fmt"
  "strings"
  "strconv"
  "encoding/json"
  "net/url"
)

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

  if request.Method == "GET" {
    var id = strings.Split(u.Path, "/")
    if (len(id) < 3) { //No Id
      js, err := json.Marshal(Users)
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusOK)
      w.Write(js)
      return
    } else if len(id) == 3 {
      requestedId, convError := strconv.Atoi(id[2])
      if convError != nil {
        http.Error(w, convError.Error(), http.StatusInternalServerError)
        return
      }
      for i := 0; i < len(Users); i++ {
        if (Users[i].Id == requestedId) {
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

      fmt.Println("ID ", requestedId, "Not Found")
      w.WriteHeader(http.StatusNotFound)
      return
    }
    return
  }

  if request.Method == "POST" {
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


  // Alternative Implementation?
  // http://stackoverflow.com/questions/18926303/iterate-through-a-struct-in-go
  if request.Method == "PATCH" {
    // http://www.alexedwards.net/blog/golang-response-snippets#json
    var id = strings.Split(u.Path, "/")

    if (len(id) < 3) { //No Id
      fmt.Println("Update without ID not implemented.")
      w.WriteHeader(http.StatusNotImplemented)
      return
    } else if len(id) == 3 {
      requestedId, convError := strconv.Atoi(id[2])
      if convError != nil {
        http.Error(w, convError.Error(), http.StatusInternalServerError)
        return
      }
      for i := 0; i < len(Users); i++ {
        if (Users[i].Id == requestedId) {
          decoder := json.NewDecoder(request.Body)
          var u UserSchema
          err = decoder.Decode(&u)
          if err != nil {
            panic(err)
          }
          newUser := User{requestedId, u}
          // update array
          leftSide := append(Users[:i], newUser)
          Users = append(leftSide, Users[i+1:]...)

          fmt.Println("Successful Update, ID:", requestedId)
          // send the deleted element
          w.WriteHeader(http.StatusAccepted)
          return
        }
      }

      //no match
      fmt.Println("Update Failed: ID", requestedId, "Not Found")
      w.WriteHeader(http.StatusNotFound)
      return
    }
    return
  }

  if request.Method == "DELETE" {
    var id = strings.Split(u.Path, "/")

    if (len(id) < 3) {//No Id
      fmt.Println("Delete without ID not implemented.")
      w.WriteHeader(http.StatusNotImplemented)
      return
    } else if len(id) == 3 {
      requestedId, convError := strconv.Atoi(id[2])
      if convError != nil {
        http.Error(w, convError.Error(), http.StatusInternalServerError)
        return
      }
      for i := 0; i < len(Users); i++ {
        if (Users[i].Id == requestedId) {
          js, err := json.Marshal(Users[i])
          if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
          }
          Users = append(Users[:i], Users[i+1:]...)
          fmt.Println("Successful Delete, ID:", requestedId)
          w.WriteHeader(http.StatusAccepted)
          w.Write(js)
          return
        }
      }
      fmt.Println("Delete Failed: ID ", requestedId, "Not Found")
      w.WriteHeader(http.StatusNotFound)
      return
    }
  }
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", handler)
  fmt.Println("Listening...")
  fmt.Println("Users...", Users)
  http.ListenAndServe(":8000", mux)
}
