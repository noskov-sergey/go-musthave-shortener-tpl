package server

import (
	"errors"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Storage struct {
	links map[string]string
}

func newStorage() *Storage {
	return &Storage{
		links: map[string]string{},
	}
}

var storage = newStorage()

func (c *Storage) Add(url string) string {
	short := make([]rune, 8)
	for i := range short {
		short[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	key := string(short)
	c.links[key] = url
	return key
}

func (c *Storage) Get(key string) (string, error) {
	url, ok := c.links[key]
	if ok == false {
		return "", errors.New("Key not exist")
	}
	return url, nil
}
