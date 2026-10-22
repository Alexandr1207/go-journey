package user

import (
	"errors"
	"fmt"
	"strings"
)

type Notifier interface {
	Send(to, msg string) error
}

type Service struct {
	notifier Notifier
	users    []string
}

func NewService(n Notifier) *Service {
	return &Service{notifier: n}
}

func (s *Service) Register(name, email string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("empty name")
	}
	s.users = append(s.users, name)
	err := s.notifier.Send(email, "Welcome, "+name+"!")
	if err != nil {
		return fmt.Errorf("register %s: %w", name, err)
	}
	return nil
}

func (s *Service) Users() []string {
	usersCopy := make([]string, len(s.users))
	copy(usersCopy, s.users) // copy, because we don't need to change source
	return usersCopy
}

// var _ = notify.Console{}

// imports myfirstgo/week06/day02-layout/internal/notify from service.go: import cycle not allowed
