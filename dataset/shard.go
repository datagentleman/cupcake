package dataset

import (
	"bytes"
	"encoding/binary"
	"path"
	"sync"
	"unsafe"
)

type Shard struct {
	ID           int32
	RecordsLimit int32
	offsets      *Offsets
	storage      *File

	mux          sync.Mutex
	recordsCount int32
}

type record struct {
	id   int32
	data []byte
}

func (r record) Bytes() []byte {
	b := new(bytes.Buffer)
	recordSize := int32(len(r.data)) + int32(unsafe.Sizeof(r.id))

	binary.Write(b, binary.BigEndian, recordSize)
	binary.Write(b, binary.BigEndian, r.id)
	binary.Write(b, binary.BigEndian, r.data)

	return b.Bytes()
}

func recordFromBytes(data []byte) *record {
	var recordSize int32

	r := &record{}
	b := bytes.NewBuffer(data)

	binary.Read(b, binary.BigEndian, &recordSize)
	binary.Read(b, binary.BigEndian, &r.id)

	r.data = b.Bytes()
	return r
}

func (s *Shard) Add(data []byte) (int32, error) {
	s.mux.Lock()
	s.recordsCount += 1
	id := s.recordsCount + ((s.ID - 1) * s.RecordsLimit)
	s.mux.Unlock()

	r := &record{id: id, data: data}
	offset, err := s.storage.Write(r.Bytes())
	if err != nil {
		return -1, err
	}

	s.offsets.Add(id, offset)
	return id, nil
}

func (s *Shard) Get(id int32) (*record, error) {
	offset := s.offsets.Get(id)

	b, err := s.storage.Read(offset)
	if err != nil {
		return nil, err
	}

	r := recordFromBytes(b)
	return r, nil
}

func createShard(id, recordsLimit int32, dir, name string) (*Shard, error) {
	file, err := CreateFile(path.Join(dir, name+".data"))
	if err != nil {
		return nil, err
	}

	o, err := createOffsets(recordsLimit, path.Join(dir, name+".offset"))
	if err != nil {
		return nil, err
	}

	s := &Shard{
		ID:           id,
		RecordsLimit: recordsLimit,
		offsets:      o,
		storage:      file,
		recordsCount: 0,
	}

	return s, nil
}
