package notify

import (
	"errors"
	"fmt"
	"strings"
)

type Console struct{}

func (Console) Send(to, msg string) error {
	fmt.Println("To:", to, ", msg:", msg)
	return nil
}

type Email struct {
	From string
}

func (e Email) Send(to, msg string) error {
	if !strings.Contains(to, "@") {
		return errors.New("Bad adress")
	}
	fmt.Println("email from:", e.From, ". To:", to, ", msg:", msg)
	return nil
}

type Fake struct {
	Sent []string
	Fail bool
}

func (f *Fake) Send(to, msg string) error {
	if f.Fail {
		return errors.New("fake failure")
	}
	f.Sent = append(f.Sent, to+": "+msg)
	return nil
}

type Telegram struct{}

func (t Telegram) Send(to, msg string) error {
	fmt.Println("[telegram] chat="+to, ": "+msg)
	return nil
}

// var _ user.Service

// imports myfirstgo/week06/day02-layout/internal/notify from service.go: import cycle not allowed
