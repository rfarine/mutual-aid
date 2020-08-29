package handler

import (
  "fmt"
  "net/http"
  "os"
  "github.com/gin-gonic/gin"
  f "gopkg.in/fauna/faunadb-go.v2/faunadb"
)

type Person struct {
  ID string `fauna:"id"`
  FirstName string `fauna:"firstName"`
  LastName string `fauna:"lastName"`
  Org string `fauna:"org"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
  apiKey := os.Getenv("FAUNADB_SECRET")
  client := f.NewFaunaClient(apiKey)

  res, err := client.Query(f.Get(f.RefClass(f.Collection("person"), "1")))

  if err != nil {
    panic(err)
  }

  var person Person

  if err := res.At(f.ObjKey("data")).Get(&person); err != nil {
    panic(err)
  }

  fmt.Println(person)
}
