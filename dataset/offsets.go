package dataset

type Offset struct {
	Start int64
	End   int64
}

type Offsets struct {
	storage *File
	offsets []*Offset
	maxSize int32
}

func (o *Offsets) Add(index int32, offset *Offset) {
	o.offsets[index] = offset
}

func createOffsets(maxSize int32, path string) (*Offsets, error) {
	file, err := CreateFile(path)
	if err != nil {
		return nil, err
	}

	o := &Offsets{
		storage: file,
		maxSize: maxSize,
		offsets: make([]*Offset, maxSize, maxSize),
	}

	return o, nil
}
