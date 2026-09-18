package input

import (
	"bufio"
	"os"
)

type StdinSource struct {
	reader *os.File
}

func NewStdinSource() *StdinSource {
	return &StdinSource{
		reader: os.Stdin,
	}
}

func (s *StdinSource) Read(fn func(string) error) error {
	scanner := bufio.NewScanner(s.reader)

	for scanner.Scan() {
		url := scanner.Text()

		if url == "" {
			continue
		}

		if err := fn(url); err != nil {
			return err
		}
	}

	return scanner.Err()
}
