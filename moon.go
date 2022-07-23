package moon

import (
	"fmt"
	"net/http"
)

type Moon struct {
}

func (m *Moon) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		switch r.URL.Path {
		case "/":
			_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
			if err != nil {
				return
			}
		}
	}
}

func New() *Moon {
	return new(Moon)
}
