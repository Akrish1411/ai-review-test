package main

import "testing"

func TestSayHello(t *testing.T) {
	got := sayHello()
	want := "Hello World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSayGoodbye(t *testing.T) {
	got := sayGoodbye()
	want := "Goodbye World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
