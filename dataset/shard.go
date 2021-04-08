package dataset

import (
	"path"
)

type Shard struct {
	Id           uint32
	RecordsLimit uint32
}

func createShard(id, recordsLimit uint32, dir, name string) (*Shard, error) {
	_, err := CreateFile(path.Join(dir, name+".data"))
	if err != nil {
		return nil, err
	}

	return &Shard{Id: id, RecordsLimit: recordsLimit}, nil
}
