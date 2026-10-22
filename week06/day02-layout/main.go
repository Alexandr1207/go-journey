package main

import (
	"fmt"
	"myfirstgo/week06/day02-layout/internal/notify"
	"myfirstgo/week06/day02-layout/internal/user"
)

func main() {
	svc := user.NewService(notify.Email{From: "noreply@site.com"})
	fmt.Println(svc.Register("Anna", "anna@x.com"))  // email from: noreply@site.com . To: anna@x.com , msg: Welcome, Anna!
	fmt.Println(svc.Register("", "x@x.com"))         // <nil>
	fmt.Println(svc.Register("Boris", "borisx.com")) // empty name // register Boris: Bad adress
	fmt.Println(svc.Users())                         // [Anna Boris]

	svt := user.NewService(&notify.Telegram{})
	fmt.Println(svt.Register("Alex", "alex@mail")) // [telegram] chat=alex@mail : Welcome, Alex! // <nil>
}
