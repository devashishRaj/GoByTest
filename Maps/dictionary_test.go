package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "okay"}
	t.Run("Known Word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "okay"
		assertSearch(t, got, want)
	})
	t.Run("Unknown word", func(t *testing.T) {
		// map can two values , the key and a boolean if the key was present in the map
		_, got := dictionary.Search("unknown")

		// Errors can be converted to a string with the .Error()
		// We are also protecting assertStrings with if to ensure we don't call .Error() on nil.
		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "okay"
		err:=dictionary.Update(word, definition)

		assertError(t, err ,  ErrWordDoesNotExists)
	})
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "okay"
		dictionary := Dictionary{word: definition}
		newDefinition := "new okay"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
}

func TestUpdate(t *testing.T){
	word := "test"
	definition := "okay"
	dictionary := Dictionary{word : definition}
	newDefinition := "Updated okay"
	err:= dictionary.Update(word ,newDefinition)
	assertError(t , err , nil)
	assertDefinition(t , dictionary , word , newDefinition)
}

func TestDelete(t *testing.T){
	word := "test"
	dictionary := Dictionary{word:"okay"}
	dictionary.Delete(word)
	_,err := dictionary.Search(word)
	assertError(t , err , ErrNotFound)
}


func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}
	assertString(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertSearch(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q wanted %q given %q", got, want, "test")
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q wanted %q", got, want)
	}
}
