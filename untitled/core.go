package main

import "errors"

var storage = make(map[string]string)

var ErrorNoSuchKey = errors.New("no such key")

func Put(key string, val string) error {
	storage[key] = val
	return nil
}

func Get(key string) (string, error) {
	val, ok := storage[key]
	if !ok {
		return "", ErrorNoSuchKey
	}

	return val, nil
}

func Delete(key string) error {
	delete(storage, key)
	return nil
}
