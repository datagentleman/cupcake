package dataset

import (
	"os"
)

type File struct {
	*os.File
	Path string
}

func CreateFile(path string) (*File, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}

	return &File{File: file, Path: path}, nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}
