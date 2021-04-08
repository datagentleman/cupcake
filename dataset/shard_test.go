package dataset

import (
	"testing"
)

func Test_Create(t *testing.T) {
	createShard(1, 100, "./", "shard_1")

	if !FileExists("./shard_1.data") {
		t.Errorf("Shard file was not created")
	}

	if !FileExists("./shard_1.offset") {
		t.Errorf("Offset file for given shard was not created")
	}
}

func Test_Add(t *testing.T) {
	r1 := []byte("record 1")
	r2 := []byte("record 2")

	s1, _ := createShard(1, 100, "./", "shard_1")
	s1.Add(r1)
	s1.Add(r2)

	s2, _ := createShard(2, 100, "./", "shard_2")
	s2.Add(r1)
	s2.Add(r2)
}
