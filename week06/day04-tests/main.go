package main

import (
	"errors"
	"fmt"
	"myfirstgo/week06/day04-tests/internal/notify"
	"myfirstgo/week06/day04-tests/internal/storage"
	"myfirstgo/week06/day04-tests/internal/user"
	"sync"
)

func main() {
	repo := storage.NewMemoryRepo()
	svc := user.NewService(repo, notify.Console{})

	u, err := svc.Register("Anna", "anna@x.com") // To: anna@x.com , msg: Welcome, Anna!
	fmt.Println(u, err)                          // {1 Anna anna@x.com} <nil>

	_, err = svc.Register("Anna2", "anna@x.com")
	fmt.Println(err, errors.Is(err, user.ErrMailTaken))
	// register Anna2: email already taken true

	_, err = svc.Find("nobody@x.com")
	fmt.Println(err, errors.Is(err, user.ErrNotFound))
	// user not found true

	svc.Register("Boris", "boris@x.com") // To: boris@x.com , msg: Welcome, Boris!
	list, _ := svc.Users()
	fmt.Println(list) // [{1 Anna anna@x.com} {2 Boris boris@x.com}]

	var wg sync.WaitGroup
	newRepo := storage.NewMemoryRepo()
	svc = user.NewService(newRepo, &notify.Fake{})
	for i := 0; i < 100; i++ {
		wg.Go(func() {
			svc.Register("user", fmt.Sprintf("u%d@x.com", i))
		})
	}

	wg.Wait()
	list, _ = svc.Users()
	for i, u := range list {
		if u.ID != i+1 {
			fmt.Println("bad id at", i, u)
		}
	}
	fmt.Println(len(list))
}
