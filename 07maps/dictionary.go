package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("word already exists")
var ErrWordDoesNotExists = errors.New("cannot update word because it does not exist")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
		// return "", errors.New("abeeyyy")
	}

	return definition, nil
}

func (d Dictionary) Update(word string,newMeaning string) error{
	_,err:=d.Search(word)

	switch err{
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word]=newMeaning
	default:
		return err
	}
	
	return nil

}

func (d Dictionary) Delete(word string){
	delete(d,word)
}

func (d Dictionary) Add(word,meaning string) error{
	
	_,err:=d.Search(word)

	switch err{
	case ErrNotFound:
		d[word]=meaning
	case nil:
		return ErrWordExists
	default:
		return err
	}	

	return nil
}

var myDict = Dictionary{"testKey": "test value"}
