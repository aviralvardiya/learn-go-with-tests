package dictionary

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("adding new word", func(t *testing.T) {
		testDict := Dictionary{}
		word := "test"
		meaning := "this is just a test"
		err := testDict.Add(word, meaning)

		assertErrors(t, err, nil)

		assertDefinition(t, testDict, word, meaning)
	})
	t.Run("adding existing word", func(t *testing.T) {

		word := "test"
		meaning := "this is just a test"
		testDict := Dictionary{word: meaning}
		err := testDict.Add(word, "new meaning")

		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, testDict, word, meaning)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("updating a word that exist", func(t *testing.T) {
		word := "test"
		meaning := "this is the old meaning"
		myDict := Dictionary{word: meaning}
		newMeaning := "this is the new meaning"
		myDict.Update(word, newMeaning)

		assertDefinition(t, myDict, word, newMeaning)

	})
	t.Run("updating a word that does not exist",func(t *testing.T) {
		word:="test"
		meaning:="this is a test"
		myDict:=Dictionary{}
		err:= myDict.Update(word,meaning)

		assertErrors(t,err,ErrWordDoesNotExists)
	})

}

func TestDelete(t *testing.T){
	word:="test"
	myDict:=Dictionary{word:"test definition"}
	myDict.Delete(word)

	_,err:=myDict.Search(word)
	if(err!=ErrNotFound){
		t.Errorf("expected %q to be deleted",word)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, meaning string) {
	t.Helper()
	want := meaning
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find the added word, ", err)
	}
	assertStrings(t, got, want)
}

func TestSearch(t *testing.T) {
	t.Run("searching known word", func(t *testing.T) {

		got, _ := myDict.Search("testKey")
		want := "test value"

		assertStrings(t, got, want)
	})
	t.Run("searching unknown word", func(t *testing.T) {

		_, err := myDict.Search("unknown word")
		want := ErrNotFound
		if err == nil {
			t.Fatalf("error needed but got none")
		}
		assertErrors(t, err, want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
