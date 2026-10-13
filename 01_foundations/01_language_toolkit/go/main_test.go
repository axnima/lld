package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// 1. Validation runs before dependencies; %w preserves the typed error.
func TestNotifyRejectsInvalidInput(t *testing.T) {
	service := NotificationService{
		sender: FailingSender{},
	}

	err := service.Notify(Notification{
		Recipient: "John",
		Message:   " ",
	})

	var validationErr ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("expected ValidationError, got %v", err)
	}

	if validationErr.Field != "message" {
		t.Fatalf("expected message field, got %q", validationErr.Field)
	}
}

// 2. Temporary files isolate tests; checking content verifies delivery and ordering.
func TestFileSenderWritesValidNotifications(t *testing.T) {
	path := filepath.Join(t.TempDir(), "notifications.log")

	service := NotificationService{
		sender: FileSender{Path: path},
		now: func() time.Time {
			return time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
		},
		nextID: generateID(),
	}

	first := Notification{Recipient: "John", Message: "Hello"}
	last := Notification{Recipient: "Noah", Message: "Welcome"}

	results := service.NotifyAll([]Notification{
		first,
		{Recipient: "John", Message: " "},
		last,
	})

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	if results[0] != nil || results[1] == nil || results[2] != nil {
		t.Fatalf("expected [nil, error, nil], got %v", results)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want := first.Summary() + "\n" + last.Summary() + "\n"
	if string(content) != want {
		t.Fatalf("expected file content %q, got %q", want, content)
	}
}

// 3. O_CREATE creates files, not parent directories; errors.Is checks the cause.
func TestNotifyReturnsFileOpenError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing", "notifications.log")

	service := NotificationService{
		sender: FileSender{Path: path},
		now:    time.Now,
		nextID: generateID(),
	}

	err := service.Notify(Notification{
		Recipient: "John",
		Message:   "Hello",
	})

	if err == nil {
		t.Fatal("expected an error because the parent directory does not exist")
	}

	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected a missing-path error, got %v", err)
	}
}

// 4. A recording sender observes behavior without writing to a file or console.
type recordingSender struct {
	received []Notification
	err      error
}

func (s *recordingSender) Send(n Notification) error {
	s.received = append(s.received, n)
	return s.err
}

func TestBatchContinuesAndSkipsInvalidInput(t *testing.T) {
	failure := errors.New("delivery failed")
	sender := &recordingSender{err: failure}
	var ids []string
	clockCalls := 0
	generator := generateID()
	service := NotificationService{
		sender: sender,
		now:    func() time.Time { clockCalls++; return time.Time{} },
		nextID: func() string { id := generator(); ids = append(ids, id); return id },
	}
	first := Notification{Recipient: "John", Message: "Hello"}
	last := Notification{Recipient: "Noah", Message: "Welcome"}
	results := service.NotifyAll([]Notification{first, {Message: "Hello"}, last})
	if len(results) != 3 {
		t.Fatalf("results: %v", results)
	}
	if !errors.Is(results[0], failure) || !errors.Is(results[2], failure) {
		t.Fatalf("delivery errors lost: %v", results)
	}
	var validationErr ValidationError
	if !errors.As(results[1], &validationErr) || validationErr.Field != "recipient" {
		t.Fatalf("expected recipient validation error, got %v", results[1])
	}
	if len(sender.received) != 2 || sender.received[0] != first || sender.received[1] != last {
		t.Fatalf("wrong delivery order or invalid input delivered: %v", sender.received)
	}
	if len(ids) != 2 || ids[0] != "attempt_1" || ids[1] != "attempt_2" || clockCalls != 2 {
		t.Fatalf("unexpected dependency calls: ids=%v clock=%d", ids, clockCalls)
	}
	if got := service.NotifyAll(nil); len(got) != 0 || len(ids) != 2 {
		t.Fatalf("empty input produced results or an ID: %v", got)
	}
}

// 5. Missing and explicitly false preferences have different meanings.
func TestNotificationEnabled(t *testing.T) {
	for _, tc := range []struct {
		name        string
		preferences map[string]bool
		want        bool
	}{
		{"enabled", map[string]bool{"John": true}, true},
		{"disabled", map[string]bool{"John": false}, false},
		{"missing", map[string]bool{}, true},
		{"nil", nil, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := NotificationEnabled(tc.preferences, "John"); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}

}
