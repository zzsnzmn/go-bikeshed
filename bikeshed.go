package main

import (
      "fmt"
      "net/http"
)

type Bikeshed struct{}

func (b Bikeshed) ServeHTTP(
      w http.ResponseWriter,
      r *http.Request) {
      fmt.Fprint(w, "the bikeshed should clearly be "+r.FormValue("color"))
}

func main() {
     /**
     assuming go is installed run
     > go run bikeshed.go
     
     or build it
     > go build bikeshed.go
     > ./bikeshed

     hit the link: localhost:4000?color=mauve
     */
      var b Bikeshed
      http.ListenAndServe("localhost:4000", b)
}