package main

import (
	"log"
	"net/http"
	"time"

	"myfirstgo/week07/day01-http/internal/httpapi"
	"myfirstgo/week07/day01-http/internal/notify"
	"myfirstgo/week07/day01-http/internal/storage"
	"myfirstgo/week07/day01-http/internal/user"
)

func main() {

	repo := storage.NewMemoryRepo()
	n := notify.Console{}
	svc := user.NewService(repo, n)

	_, err := svc.Register("Anna", "anna@mail.ru")
	if err != nil {
		log.Fatalf("register Anna: %v", err)
	}
	_, err = svc.Register("Boris", "boris@mail.ru")
	if err != nil {
		log.Fatalf("register Boris: %v", err)
	}

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           httpapi.NewHandler(svc).Routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
