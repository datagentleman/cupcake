package dataset

type Offset struct {
	Start int64
	End   int64
}

type Offsets struct {
	storage *File
	offsets map[int32]*Offset
	maxSize int32
}

func (o *Offsets) Add(key int32, offset *Offset) {
	o.offsets[key] = offset
}

func (o *Offsets) Get(key int32) *Offset {
	return o.offsets[key]
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
