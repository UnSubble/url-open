package input

import "os"

type FileSource struct {
	file *os.File
}

func NewFileSource(path string) (*FileSource, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &FileSource{
		file: file,
	}, nil
}

func (s *FileSource) Read(fn func(string) error) error {
	defer s.file.Close()

	return readLines(s.file, fn)
}
