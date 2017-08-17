package main

import "testing"

func TestHello(t *testing.T) {
	msg, err := hello("foo")
	if err != nil {
		t.Fatal(err)
	}

	if g, e := msg, "Hello, foo !"; g != e {
		t.Fatalf("got mesage %v, want %v", g, e)
	}
}
