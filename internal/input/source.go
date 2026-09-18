package input

import (
	"bufio"
	"io"
)

type Source interface {
	Read(func(string) error) error
}

func readLines(reader io.Reader, fn func(string) error) error {
	scanner := bufio.NewScanner(reader)

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
