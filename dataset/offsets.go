package dataset

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io/ioutil"
)

var offsetSizeInBytes = 20

type Offset struct {
	Key   int32
	Start int64
	End   int64
}

type Offsets struct {
	storage *File
	offsets map[int32]*Offset
	maxSize int32
}

func (o *Offsets) Add(key int32, offset *Offset) error {
	o.offsets[key] = offset

	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, key)
	binary.Write(b, binary.BigEndian, offset.Start)
	binary.Write(b, binary.BigEndian, offset.End)

	_, err := o.storage.Write(b.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (o *Offsets) Get(key int32) *Offset {
	return o.offsets[key]
}

func (o *Offsets) readAll() error {
	data, err := ioutil.ReadAll(o.storage)
	if err != nil {
		return err
	}

	corrupted := len(data) % offsetSizeInBytes
	if corrupted != 0 {
		return errors.New("Corrupted offset file")
	}

	buf := bytes.NewBuffer(data)
	for {
		b := buf.Next(offsetSizeInBytes)
		if len(b) != offsetSizeInBytes {
			break
		}

		offset := offsetFromBytes(b)
		o.offsets[offset.Key] = offset
	}

	return nil
}

func offsetFromBytes(data []byte) *Offset {
	o := &Offset{}
	b := bytes.NewBuffer(data)

	binary.Read(b, binary.BigEndian, &o.Key)
	binary.Read(b, binary.BigEndian, &o.Start)
	binary.Read(b, binary.BigEndian, &o.End)

	return o
}

func createOffsets(maxSize int32, path string) (*Offsets, error) {
	file, err := CreateFile(path)
	if err != nil {
		return nil, err
	}

	o := &Offsets{
		storage: file,
		maxSize: maxSize,
		offsets: make(map[int32]*Offset, maxSize),
	}
	return o, nil
}
