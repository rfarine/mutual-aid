package handler

import (
  "fmt"
  "net/http"
  "time"
  "encoding/json"
)

type Response struct {
  Date string `json:"date"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
  currentTime := time.Now().Format(time.RFC850)
  fmt.Fprintf(w, currentTime)
  date := Response{Date: currentTime}
  json.NewEncoder(w).Encode(date)
}
