package handler

import (
  "fmt"
  "net/http"
  "os"
  f "github.com/fauna/faunadb-go/v3/faunadb"
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
