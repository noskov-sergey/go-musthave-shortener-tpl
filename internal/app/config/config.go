package config

import (
	"errors"
	"strconv"
	"strings"
)

var BaseUrl string

type NetAddress struct {
	Host string
	Port int
}

func NewNetAddress() *NetAddress {
	return &NetAddress{
		Host: "localhost",
		Port: 8080,
	}
}

func (n *NetAddress) String() string {
	return n.Host + ":" + strconv.Itoa(n.Port)
}

func (n *NetAddress) Set(src string) error {
	hp := strings.Split(src, ":")
	if len(hp) != 2 {
		return errors.New("need address in a form host:port")
	}
	port, err := strconv.Atoi(hp[1])
	if err != nil {
		return err
	}
	n.Host = hp[0]
	n.Port = port
	return nil
}
