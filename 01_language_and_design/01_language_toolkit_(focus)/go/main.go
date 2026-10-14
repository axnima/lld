package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Revision order: definitions below, followed by matching examples in main.
// Run: go run .

// 1. Structs, methods, and validation.
// A struct groups data; a method attaches behavior to that type.
type Notification struct {
	Recipient string
	Message   string
}

func (n Notification) Summary() string {
	return "Notified : {" + n.Recipient + " : " + n.Message + "}"
}

func (n Notification) Validate() error {
	if strings.TrimSpace(n.Recipient) == "" {
		return ValidationError{Field: "recipient"}
	} else if strings.TrimSpace(n.Message) == "" {
		return ValidationError{Field: "message"}
	}
	return nil
}

// 2. Structured errors: implementing Error() string satisfies Go's error interface.
// errors.As finds an error by type; errors.Is checks a particular error value.
type ValidationError struct {
	Field string
}

func (e ValidationError) Error() string {
	return e.Field + " is required"
}

// 3. Interfaces: matching method signatures satisfy an interface implicitly.
type Sender interface {
	Send(n Notification) error
}

type ConsoleSender struct {
}

type FailingSender struct {
}

type FileSender struct {
	Path string
}

// Compile-time checks: these types must implement Sender.
var _ Sender = ConsoleSender{}
var _ Sender = FailingSender{}
var _ Sender = FileSender{}

func (c ConsoleSender) Send(n Notification) error {
	if err := n.Validate(); err != nil {
		return err
	}
	fmt.Println(n.Summary())
	return nil
}

func (f FailingSender) Send(n Notification) error {
	return errors.New("Delivery unavailable")
}

// 4. Resource cleanup: defer runs on return, including after a write failure.
// The named result lets cleanup report a close error without hiding a write error.
func (f FileSender) Send(n Notification) (err error) {
	file, err := os.OpenFile(
		f.Path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0600,
	)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()

	_, err = fmt.Fprintln(file, n.Summary())
	return err
}

// 5. Composition and injected dependencies.
// Supply all three dependencies before sending valid input; nil functions panic.
type NotificationService struct {
	sender Sender
	now    func() time.Time
	nextID func() string
}

func (s NotificationService) Notify(n Notification) error {
	if err := n.Validate(); err != nil {
		// %w preserves the error for errors.As; %v would keep only its text.
		return fmt.Errorf("notify %q: %w", n.Recipient, err)
	}
	id := s.nextID()
	fmt.Println("Delivery attempt:", id, "at:", s.now())
	return s.sender.Send(n)
}

// 6. Collections: preserve one result per input and continue after errors.
// Calling the service method preserves validation for every Sender implementation.
func (s NotificationService) NotifyAll(items []Notification) []error {
	var errorlist []error
	fmt.Println("Starting Notification Broadcast")
	for _, value := range items {
		errorlist = append(errorlist, s.Notify(value))
	}
	return errorlist
}

// 7. Values versus pointers: a value receiver changes a copy of the struct.
func (n Notification) RenameRecipient(s string) {
	n.Recipient = s
}

func (n *Notification) RenameRecipientInPlace(s string) {
	n.Recipient = s
}

// A struct copy containing a slice still shares the slice's backing array.
type NotificationBatch struct {
	Items []Notification
}

// 8. Missing versus false: the second map result reports whether the key exists.
// Reading a nil map is safe; writing to a nil map panics.
func NotificationEnabled(preferences map[string]bool, s string) bool {
	if enabled, exists := preferences[s]; exists && !enabled {
		return false
	}
	return true
}

// 9. Closures: each generator captures its own counter, retained between calls.
// shortcut: IDs are local to one sequential generator, upgrade for concurrent/global IDs.
func generateID() func() string {
	sequence := 0

	return func() string {
		sequence += 1
		return fmt.Sprintf("attempt_%d", sequence)
	}
}

func main() {
	fmt.Println("1. Structs and validation")
	valid := Notification{Recipient: "John", Message: "Hello"}
	invalid := Notification{Recipient: "John", Message: " "}
	for _, n := range []Notification{valid, invalid, {Message: "Hello"}} {
		fmt.Println("Validation:", n.Validate()) // nil means success.
	}

	fmt.Println("2. Structured errors and wrapping")
	originalErr := ValidationError{Field: "message"}
	for _, err := range []error{
		fmt.Errorf("notify: %v", originalErr),
		fmt.Errorf("notify: %w", originalErr),
	} {
		var validationErr ValidationError
		fmt.Println(err, "matches ValidationError:", errors.As(err, &validationErr)) // false, true.
	}

	fmt.Println("3. Interfaces and interchangeable senders")
	// Sender describes behavior; this loop accepts two different concrete types.
	for _, sender := range []Sender{ConsoleSender{}, FailingSender{}} {
		if err := sender.Send(valid); err != nil {
			fmt.Println("Delivery failed:", err)
		}
	}

	fmt.Println("4–6. Composition, clocks, IDs, and batches")
	service := NotificationService{
		sender: ConsoleSender{},
		now: func() time.Time {
			return time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
		},
		nextID: generateID(), // Call the factory once; retain the returned function.
	}
	// Invalid input consumes no ID: two attempts, same time, attempt_1 and attempt_2.
	fmt.Println("Batch results:", service.NotifyAll([]Notification{valid, invalid, valid}))

	fmt.Println("7. Values, pointers, and shared slice data")
	original := valid
	copied := original
	copied.Recipient = "Alice"
	fmt.Println(original.Recipient, copied.Recipient) // John Alice.
	original.RenameRecipient("Bob")
	fmt.Println(original.Recipient) // John: only the receiver copy changed.
	original.RenameRecipientInPlace("Carol")
	fmt.Println(original.Recipient) // Carol: the pointer reaches the original.

	first := NotificationBatch{Items: []Notification{valid}}
	second := first
	second.Items[0].Recipient = "Alice"
	fmt.Println("Shared:", first.Items[0].Recipient) // Alice: same backing array.
	// Copy N elements into a new array: O(N) time/space; nested references stay shared.
	second.Items = append([]Notification(nil), first.Items...)
	second.Items[0].Recipient = "Noah"
	fmt.Println("Independent:", first.Items[0].Recipient, second.Items[0].Recipient) // Alice Noah.

	fmt.Println("8. Missing versus false")
	preferences := map[string]bool{"John": true, "Holly": false}
	enabled, exists := preferences["Bob"]
	fmt.Println("Raw lookup:", enabled, exists) // false false.
	for _, name := range []string{"John", "Holly", "Bob"} {
		fmt.Println(name, NotificationEnabled(preferences, name)) // true, false, true.
	}
	fmt.Println("Nil map default:", NotificationEnabled(nil, "Bob")) // true.

	fmt.Println("9. Equality and map keys")
	a, b := valid, valid
	fmt.Println("Equal structs:", a == b) // All fields are comparable strings.
	counts := map[Notification]int{a: 1}
	fmt.Println("Equal key:", counts[b]) // 1.
	// map[NotificationBatch]int is invalid: its slice field is not comparable.
	// *NotificationBatch is a valid key, but compares identity, not contents.

	fmt.Println("10. File delivery and cleanup")
	fileService := NotificationService{
		sender: FileSender{Path: "notifications.log"},
		now:    time.Now, // Pass the function without calling it.
		nextID: generateID(),
	}
	// Each run appends two lines in the working directory.
	for i, err := range fileService.NotifyAll([]Notification{valid, invalid, valid}) {
		if err != nil {
			fmt.Printf("Notification %d failed: %v\n", i+1, err)
		}
	}
}
