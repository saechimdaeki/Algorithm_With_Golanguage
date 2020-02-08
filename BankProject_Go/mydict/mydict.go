package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("찾지 못하였다 ")
var errWordExists = errors.New("그단어는 이미있어 ")

//search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}
	return "", errNotFound
}

// add a word to the dicitonary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
