package playground

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting()
	want := "I'm super excited to learn GO!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
