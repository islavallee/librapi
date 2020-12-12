package main

import (
	"github.com/islavallee/librapi/pkg/http"
)

func main() {
	s := http.NewServer(http.NewHandler())
	s.Serve(8080)
}
