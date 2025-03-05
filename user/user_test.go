// user_test.go
package user

import (
	"testing"
)

// Test setting a valid email
func TestSetEmail_Valid(t *testing.T) {
	u := User{Name: "Alice"}

	err := u.SetEmail("alice@example.com")
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}

	if u.GetEmail() != "alice@example.com" {
		t.Fatalf("expected email to be 'alice@example.com', got: %s", u.GetEmail())
	}
}

// Test setting an invalid email (empty string)
func TestSetEmail_Invalid(t *testing.T) {
	u := User{Name: "Alice"}

	err := u.SetEmail("")
	if err == nil {
		t.Fatal("expected error, but got none")
	}

	expected := ""
	if u.GetEmail() != expected {
		t.Fatalf("expected email to be empty, got: %s", u.GetEmail())
	}
}

// Test the getter directly
func TestGetEmail(t *testing.T) {
	u := User{Name: "Alice", email: "bob@example.com"}

	if u.GetEmail() != "bob@example.com" {
		t.Fatalf("expected email to be 'bob@example.com', got: %s", u.GetEmail())
	}
}
