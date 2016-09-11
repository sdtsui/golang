package main

import (
  "net/http"
  "fmt"
  "strings"
  "strconv"
  "encoding/json"
  "net/url"
)

type UserDetail struct {
  Username string
  Password string
  Email string
}
type User struct {
  Id int `json:"id"`
  User UserDetail `json:"user"`
}
var firstUserDetail = UserDetail{
  Username: "Daniel",
  Password: "12345",
  Email: "daniel@blacklight.io",
}
var secondUserDetail = UserDetail{
  Username: "Daniel2",
  Password: "67890",
  Email: "dan@blacklight.io",
}
var Users []User
var usersCount = 0

func addUser(count int, detail UserDetail) {
  var newUser = User{count, detail}
  Users = append(Users, newUser)
  usersCount++
}

func init() {
  fmt.Println("Initializing server...")
  addUser(usersCount, firstUserDetail)
  addUser(usersCount, secondUserDetail)
  fmt.Println("Initial Users :", Users)
}

func handler(w http.ResponseWriter, request *http.Request) {
  u, err := url.Parse(request.RequestURI)
  if err != nil {
      panic(err)
  }
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
    var u UserDetail
    err = decoder.Decode(&u)
    if err != nil {
      panic(err)
    }
    addUser(usersCount, u)
    w.WriteHeader(http.StatusCreated)
    return
  }


  if request.Method == "PATCH" {
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
          var u UserDetail
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
