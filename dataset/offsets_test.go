package dataset

import (
	"reflect"
	"testing"
)

func Test_Offsets_Add(t *testing.T) {
	o, _ := createOffsets(100, "./shard_1.offset")

	offset1 := &Offset{Key: 1, Start: 10, End: 20}
	offset2 := &Offset{Key: 2, Start: 21, End: 30}

	o.Add(1, offset1)
	o.Add(2, offset2)

	o.offsets = make(map[int32]*Offset)
	o.readAll()

	if !reflect.DeepEqual(o.offsets[1], offset1) {
		t.Errorf("Offsets are not properly saved to file")
	}
}
