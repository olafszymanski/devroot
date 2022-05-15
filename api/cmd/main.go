package main

import (
	"net/http"

	"github.com/olafszymanski/devroot/api/api"
)

func main() {
	r := api.NewRouter()

	http.ListenAndServe(":3000", r)
}
