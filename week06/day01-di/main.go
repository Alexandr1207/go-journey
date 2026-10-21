package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Job struct {
	Index int
	Value int
}

type Result struct {
	Index int
	Out   int
}

func worker(id int, job <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range job {
		results <- Result{Index: i.Index, Out: i.Value * i.Value}
	}
}

func SquareAll(nums []int, workers int) []int {
	jobs := make(chan Job)
	result := make(chan Result)

	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, jobs, result, &wg)
	}

	go func() {
		for i, v := range nums {
			jobs <- Job{Index: i, Value: v}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	out := make([]int, len(nums))
	for r := range result {
		out[r.Index] = r.Out
	}
	return out
}

type Notifier interface {
	Send(to, msg string) error
}

type ConsoleNotifier struct{}

func (ConsoleNotifier) Send(to, msg string) error {
	fmt.Println("To:", to, ", msg:", msg)
	return nil
}

type EmailNotifier struct {
	From string
}

func (e EmailNotifier) Send(to, msg string) error {
	if !strings.Contains(to, "@") {
		return errors.New("Bad adress")
	}
	fmt.Println("email from:", e.From, ". To:", to, ", msg:", msg)
	return nil
}

type UserService struct {
	notifier Notifier
	users    []string
}

func NewUserService(n Notifier) *UserService {
	return &UserService{notifier: n}
}

func (s *UserService) Register(name, email string) error {
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

type FakeNotifier struct {
	Sent []string
	Fail bool
}

func (f *FakeNotifier) Send(to, msg string) error {
	if f.Fail {
		return errors.New("fake failure")
	}
	f.Sent = append(f.Sent, to+": "+msg)
	return nil
}

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(SquareAll(nums, 3))

	var n Notifier

	n = ConsoleNotifier{}
	n.Send("Alex", "Hello world")

	n = EmailNotifier{From: "Katya"}
	n.Send("Alice@mail", "CHAO")

	var n1 Notifier
	var n2 Notifier

	n1 = ConsoleNotifier{}
	n2 = EmailNotifier{From: "noreply@site.com"}

	r1 := NewUserService(n1)
	r2 := NewUserService(n2)

	r1.Register("Anna", "anna@mail")
	r2.Register("Sasha", "sasha@mail")

	fmt.Println(r2.Register("Ura", "borisx.com"))

	fake := &FakeNotifier{}
	r1 = NewUserService(fake)

	r1.Register("Anna", "anna@mail")
	fmt.Println(fake.Sent)

	fake.Fail = true
	r1.Register("Boris", "boris@")
}
