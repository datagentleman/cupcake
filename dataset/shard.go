package dataset

import (
	"path"
	"sync"
)

type Shard struct {
	ID           int32
	RecordsLimit int32
	offsets      *Offsets
	storage      *File

	mux          sync.Mutex
	recordsCount int32
}

func (s *Shard) Add(data []byte) (int32, error) {
	// generate id for new record
	s.mux.Lock()
	s.recordsCount += 1
	recordID := s.recordsCount + ((s.ID - 1) * s.RecordsLimit)
	offsetID := s.recordsCount
	s.mux.Unlock()

	offset, err := s.storage.Write(data)
	if err != nil {
		return -1, err
	}

	s.offsets.Add(offsetID, offset)
	return recordID, nil
}

func (s *Shard) Get(id int32) ([]byte, error) {
	id -= ((s.ID - 1) * s.RecordsLimit)
	offset := s.offsets.Get(id)

	data, err := s.storage.Read(offset)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func createShard(id, recordsLimit int32, dir, name string) (*Shard, error) {
	recordsLimit += 1

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
