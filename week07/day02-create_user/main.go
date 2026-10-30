package main

import (
	"log"
	"net/http"
	"time"

	"myfirstgo/week07/day02-create_user/internal/httpapi"
	"myfirstgo/week07/day02-create_user/internal/notify"
	"myfirstgo/week07/day02-create_user/internal/storage"
	"myfirstgo/week07/day02-create_user/internal/user"
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
