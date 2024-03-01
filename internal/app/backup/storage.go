package storage

import (
	"errors"
	"fmt"
	"go-musthave-shortener-tpl/internal/app/config"
	"log"
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

var RealStorage = newStorage()

func (c *Storage) Add(url string) (string, error) {
	short := make([]rune, 8)
	for i := range short {
		short[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	key := string(short)
	c.links[key] = url
	if c.links[key] != url {
		fmt.Println("Error for Add key to storage")
	}
	if config.Fileparams.W == 1 {
		Writer, err := NewWriter(config.Fileparams.String())
		if err != nil {
			log.Fatal(err)
		}
		defer Writer.Close()
		err = Writer.WriteData(key, url)
		if err != nil {
			log.Fatal(err)
		}
	}
	return key, nil
}

func (c *Storage) Get(key string) (string, error) {
	url, ok := c.links[key]
	if !ok {
		return "", errors.New("key not exist")
	}
	return url, nil
}

func (c *Storage) ReadBackup(uri string, originalUri string) error {
	c.links[uri] = originalUri
	return nil
}
