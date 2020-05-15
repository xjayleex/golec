package dict

import (
	"errors"
)

type Dictionary map[string]string

var (
	errNotFound = errors.New("Not Found")
 	errCantUpdate = errors.New("Cant Update non_existring word")
 	errWordExists = errors.New( "Word already exist")
)
func (d Dictionary) Search(word string) (string, error){
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d  Dictionary) Add(word, def string) error {
	_, err:= d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, newDef string) error{
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = newDef
	case errNotFound:
		return errCantUpdate
	}
	return nil
}
// Delete a word(key)
func (d Dictionary) Delete(word string){
	delete(d, word)
}