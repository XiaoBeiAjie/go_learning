package main
import "testing"

func TestHello(t *testing.T) {
	got := Hello("a_j1e")
	want := "Hello, a_j1e"
	
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
