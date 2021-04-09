package dataset

import (
	"bytes"
	"testing"
)

func Test_Shard_Create(t *testing.T) {
	createShard(1, 100, "./", "shard_1")

	if !FileExists("./shard_1.data") {
		t.Errorf("Shard file was not created")
	}

	if !FileExists("./shard_1.offset") {
		t.Errorf("Offset file for given shard was not created")
	}
}

func Test_Shard_Add(t *testing.T) {
	r1 := []byte("record 1")
	r2 := []byte("record 2")

	s1, _ := createShard(1, 2, "./", "shard_1")
	id1, _ := s1.Add(r1)
	id2, _ := s1.Add(r2)

	d1, _ := s1.Get(id1)
	d2, _ := s1.Get(id2)

	if bytes.Compare(d1, r1) != 0 || bytes.Compare(d2, r2) != 0 {
		t.Errorf("Corrupted data while writing to shard")
	}
}
