package handler

import (
  "net/http"
  "os"
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
  id := r.PathParameters["id"]

  res, err := client.Query(f.Get(f.RefClass(f.Collection("person"), id)))

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  var person Person

  if err := res.At(f.ObjKey("data")).Get(&person); err != nil {
    panic(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(person)
}
