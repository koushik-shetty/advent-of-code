package utils

import (
	"bufio"
	"log"
	"os"
)

type Input struct {
	FileName  string
	BatchSize int
	rdr       *bufio.Scanner
	file      *os.File
}

func NewInput(file string, batchSize int) Input {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Err: %#v\n\n", err.Error())
		os.Exit(1)
	}

	rdr := bufio.NewScanner(f)
	return Input{
		FileName:  file,
		BatchSize: batchSize,
		rdr:       rdr,
		file:      f,
	}
}

func (i *Input) ReadBatch() ([]string, error) {
	lines := make([]string, 0)
	l := 0
	for l < i.BatchSize && i.rdr.Scan() {
		lines = append(lines, i.rdr.Text())
		l++
	}
	if err := i.rdr.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func (i *Input) close() {
	i.file.Close()
}
