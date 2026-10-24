package user_test

import (
	"errors"
	"myfirstgo/week06/day04-tests/internal/notify"
	"myfirstgo/week06/day04-tests/internal/storage"
	"myfirstgo/week06/day04-tests/internal/user"
	"testing"
)

func TestRegister(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		email    string
		wantErr  bool
	}{
		{name: "normal", userName: "Anna", email: "anna@x.com", wantErr: false},
		{name: "empty name", userName: "", email: "boris@x.com", wantErr: true},
		{name: "spaces name", userName: " ", email: "clara@x.com", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := storage.NewMemoryRepo()
			fake := &notify.Fake{}
			svc := user.NewService(repo, fake)

			_, err := svc.Register(tt.userName, tt.email)
			if tt.wantErr && err == nil {
				t.Fatalf("Register(%q, %q): got nil error, want error", tt.userName, tt.email)
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("Register(%q, %q): unexpected error: %v", tt.userName, tt.email, err)
			}

			if tt.name == "empty name" {
				if len(fake.Sent) != 0 {
					t.Errorf("fake.Sent has %d entries, want 0 (no welcome email should be sent)", len(fake.Sent))
				}
			}
		})
	}
}

func TestRegister_SendsWelcome(t *testing.T) {
	repo := storage.NewMemoryRepo()
	fake := &notify.Fake{}
	svc := user.NewService(repo, fake)

	_, err := svc.Register("Anna", "anna@x.com")
	if err != nil {
		t.Fatalf("Register: unexpected error: %v", err)
	}

	if len(fake.Sent) != 1 {
		t.Fatalf("fake.Sent has %d entries, want 1", len(fake.Sent))
	}
	const want = "anna@x.com: Welcome, Anna!"
	if fake.Sent[0] != want {
		t.Errorf("fake.Sent[0] = %q, want %q", fake.Sent[0], want)
	}
}

func TestRegister_DuplicateEmail(t *testing.T) {
	repo := storage.NewMemoryRepo()
	fake := &notify.Fake{}
	svc := user.NewService(repo, fake)

	_, err := svc.Register("Anna", "anna@x.com")
	if err != nil {
		t.Fatalf("first Register: unexpected error: %v", err)
	}

	_, err = svc.Register("Anna", "anna@x.com")
	if !errors.Is(err, user.ErrMailTaken) {
		t.Fatalf("second Register: got error %v, want ErrMailTaken", err)
	}

	if len(fake.Sent) != 1 {
		t.Errorf("fake.Sent has %d entries, want 1 (second welcome should not be sent)", len(fake.Sent))
	}
}

func TestRegister_NotifyFails(t *testing.T) {
	repo := storage.NewMemoryRepo()
	fake := &notify.Fake{Fail: true}
	svc := user.NewService(repo, fake)

	got, err := svc.Register("Anna", "anna@x.com")
	if err == nil {
		t.Fatalf("Register: got nil error, want notify failure")
	}

	if got.ID == 0 {
		t.Errorf("Register: got user ID 0, want non-zero (user should be saved)")
	}

	users, err := svc.Users()
	if err != nil {
		t.Fatalf("Users: unexpected error: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("Users() returned %d users, want 1", len(users))
	}
}

//         myfirstgo/week06/day04-tests            coverage: 0.0% of statements
// myfirstgo/week06/day04-tests/internal/notify            coverage: 0.0% of statements
// ok      myfirstgo/week06/day04-tests/internal/storage   0.894s  coverage: 100.0% of statements
// ok      myfirstgo/week06/day04-tests/internal/user      0.966s  coverage: 91.7% of statements
