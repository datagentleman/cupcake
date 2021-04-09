package dataset

import (
	"os"
	"sync"
)

type File struct {
	*os.File

	Path   string
	mux    sync.Mutex
	offset int64
}

// threadsafe
func (f *File) Write(data []byte) (*Offset, error) {
	len := int64(len(data))

	// reserve space for given data
	f.mux.Lock()
	start := f.offset
	f.offset += len
	f.mux.Unlock()

	_, err := f.WriteAt(data, int64(start))
	if err != nil {
		return nil, err
	}

	return &Offset{Start: start, End: start + len}, nil
}

// threadsafe
func (f *File) ReadOffset(o *Offset) ([]byte, error) {
	data := make([]byte, o.End-o.Start)

	_, err := f.ReadAt(data, o.Start)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CreateFile(path string) (*File, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}

	return &File{File: file, Path: path, offset: 0}, nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}
