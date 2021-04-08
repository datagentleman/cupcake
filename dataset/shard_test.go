package dataset

import (
	"testing"
)

func TestShard(t *testing.T) {
	createShard(1, 100, "./", "shard_1")

	if !FileExists("./shard_1.data") {
		t.Errorf("Shard file was not created")
	}
}
