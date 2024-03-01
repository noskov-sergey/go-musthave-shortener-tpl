package storage

import (
	"bufio"
	"encoding/json"
	"go-musthave-shortener-tpl/internal/app/models"
	"log"
	"os"
)

type Writer struct {
	file   *os.File
	writer *bufio.Writer
}

func NewWriter(fileName string) (*Writer, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &Writer{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func (w *Writer) WriteData(key, url string) error {
	line := models.BackupModel{
		URI:         key,
		OriginalUri: url,
	}
	data, err := json.Marshal(&line)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if _, err = w.writer.Write(data); err != nil {
		log.Fatal(err)
		return err
	}
	if err = w.writer.WriteByte('\n'); err != nil {
		log.Fatal(err)
		return err
	}
	return w.writer.Flush()
	return nil
}

func (w *Writer) Close() error {
	return w.file.Close()
}

type Reader struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewReader(fileName string) (*Reader, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &Reader{
		file:    file,
		scanner: bufio.NewScanner(file),
	}, nil
}

func (c *Reader) ReadFile() error {
	for c.scanner.Scan() {
		backupData := models.BackupModel{}
		line := c.scanner.Bytes()
		err := json.Unmarshal(line, &backupData)
		if err != nil {
			return err
		}
		err = RealStorage.ReadBackup(backupData.URI, backupData.OriginalUri)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Reader) Close() error {
	return c.file.Close()
}
