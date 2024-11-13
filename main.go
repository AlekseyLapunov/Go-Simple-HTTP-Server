package main

import (
	"net/http"
	"github.com/AlekseyLapunov/Go-Simple-HTTP-Server/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}