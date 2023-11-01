package main

import (
	"fmt"
	"net/http"

	"github.com/DarkCodeMaster/Moon"
)

func main() {
	m := moon.New()

	m.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	m.POST("/test", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	m.Run(":8080")
}
