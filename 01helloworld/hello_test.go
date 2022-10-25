package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bully","")
		want := "Hello, Bully"

		assertCorrectMessage(t,got,want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, world"

		assertCorrectMessage(t,got,want)
	})
	t.Run("test in spanish", func(t *testing.T) {
		got := Hello("Bully","spanish")
		want := "Hola, Bully"

		assertCorrectMessage(t,got,want)
	})
	t.Run("test in french", func(t *testing.T) {
		got := Hello("Bully","french")
		want := "Bonjour, Bully"

		assertCorrectMessage(t,got,want)
	})
	t.Run("test in hindi", func(t *testing.T) {
		got := Hello("Bully","hindi")
		want := "Namaste, Bully"

		assertCorrectMessage(t,got,want)
	})

}

func assertCorrectMessage(t testing.TB,got, want string){
	t.Helper()
	if(got!=want){
		t.Errorf("got %q want %q", got, want)
	}

}
