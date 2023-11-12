package maps

const (
	ErrNotFound   = DictionaryErr("Could not find the word you are looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("Could not update word as it does not exist")
)
/*
	We made the errors constant; this required us to create our own DictionaryErr type which 
	implements the error interface.
	Simply put, it makes the errors more reusable and immutable.

*/
type DictionaryErr string
// https://dave.cheney.net/2016/04/07/constant-errors
func (e DictionaryErr) Error() string {
	return string(e)
}

// a Dictionary type which acts as a thin wrapper around map.
// With the custom type defined, we can create the Search method.
/*
	An interesting property of maps is that you can modify them without
	passing as an address to it (e.g &myMap)
	NOTE : A map value is a pointer to a runtime.hmap structure.
	So when you pass a map to a function/method, you are indeed copying it, but just the pointer part,
	not the underlying data structure that contains the data.

	To create a map :

	var dictionary = map[string]string{}
	OR
	var dictionary = make(map[string]string)


*/
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return " ", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update( word , newDefinition string)error{
	_ ,err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

func(d Dictionary) Delete(Word string){
	// Go has a built-in function delete that works on maps
	delete(d , Word)
}
