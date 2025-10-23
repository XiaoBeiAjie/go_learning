package main
import "testing"


func TestHello(t *testing.T) {
	assertCorrentMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("a_j1e", "English")
		want := "Hello, a_j1e"
		assertCorrentMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrentMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("a_j1e", "Spanish")
		want := "Hola, a_j1e"
		assertCorrentMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("a_j1e", "French")
		want := "Bonjour, a_j1e"
		assertCorrentMessage(t, got, want)
	})
}
