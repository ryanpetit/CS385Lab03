// Found @: https://hackernoon.com/how-to-create-a-web-server-in-go-a064277287c9
package main

import (
  "net/http"
   "os"
)

func returnHostName(w http.ResponseWriter, r *http.Request) {
  message, _ := os.Hostname()
  message = message + "\n"
  w.Write([]byte(message))
}

func main() {
  http.HandleFunc("/", returnHostName)
  if err := http.ListenAndServe(":80", nil); err != nil {
    panic(err)
  }
}
