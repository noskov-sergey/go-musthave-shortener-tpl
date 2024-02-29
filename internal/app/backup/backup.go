package backup

import (
	"encoding/json"
	server "go-musthave-shortener-tpl/internal/app"
	"os"
)

type Producer struct {
	file    *os.File
	encoder *json.Encoder
}

func NewProducer(fileName string) (*Producer, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &Producer{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (p *Producer) WriteData(storage *server.Storage) error {
	return p.encoder.Encode(&storage)
}

func (p *Producer) Close() error {
	return p.file.Close()
}

type Consumer struct {
	file    *os.File
	decoder *json.Decoder
}

func NewConsumer(fileName string) (*Consumer, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

func (c *Consumer) ReadEvent() (*server.Storage, error) {
	storage := &server.Storage{}
	if err := c.decoder.Decode(&storage); err != nil {
		return nil, err
	}

	return storage, nil
}

func (c *Consumer) Close() error {
	return c.file.Close()
}
