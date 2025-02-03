package main

import (
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {
	name := "James"
	expected := User{name: name}
	actual := NewUser(name)

	if actual != expected {
		t.Error("Expected User is not same as actual user")
	}
}

func TestUser_Get(t *testing.T) {
	name := "James"
	user := NewUser(name)
	message := "You look great"

	expected := fmt.Sprintf("Hello %s - %s", user.name, message)
	actual := user.Get(message)

	if actual != expected {
		t.Error("Expected User is not same as actual user")
	}
}
