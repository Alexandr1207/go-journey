package main

import (
	"log"
	"net/http"
	"time"

	"myfirstgo/week07/day03-middleware/internal/httpapi"
	"myfirstgo/week07/day03-middleware/internal/notify"
	"myfirstgo/week07/day03-middleware/internal/storage"
	"myfirstgo/week07/day03-middleware/internal/user"
)

func main() {

	repo := storage.NewMemoryRepo()
	n := notify.Email{From: "noreply@site.com"}
	svc := user.NewService(repo, n)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           httpapi.NewHandler(svc).Routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
