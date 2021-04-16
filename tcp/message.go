package tcp

import (
	"bytes"
	"encoding/binary"
	syscall "golang.org/x/sys/unix"
)

func Message(data ...interface{}) []byte {
	message := new(bytes.Buffer)

	for _, d := range data {
		switch d.(type) {
		case string:
			d = []byte(d.(string))
		case int:
			d = int32(d.(int))
		}

		binary.Write(message, binary.BigEndian, dataSize(d))
		binary.Write(message, binary.BigEndian, d)
	}

	return message.Bytes()
}

func ReadMessage(c *Client) ([]byte, error) {
	dataSize := make([]byte, 4)
	_, err := syscall.Read(c.connection, dataSize)
	if err != nil {
		return nil, err
	}

	data := make([]byte, binary.BigEndian.Uint32(dataSize))
	_, _, err = syscall.Recvfrom(c.connection, data, syscall.MSG_WAITALL)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func dataSize(data interface{}) int32 {
	switch data := data.(type) {
	case int:
		return 4
	case bool, int8, uint8, *bool, *int8, *uint8:
		return 1
	case []bool:
		return int32(len(data))
	case []int8:
		return int32(len(data))
	case []uint8:
		return int32(len(data))
	case int16, uint16, *int16, *uint16:
		return 2
	case []int16:
		return int32(2 * len(data))
	case []uint16:
		return int32(2 * len(data))
	case int32, uint32, *int32, *uint32:
		return 4
	case []int32:
		return int32(4 * len(data))
	case []uint32:
		return int32(4 * len(data))
	case int64, uint64, *int64, *uint64:
		return 8
	case []int64:
		return int32(8 * len(data))
	case []uint64:
		return int32(8 * len(data))
	case string:
		return int32(len(data))
	default:
		panic("UNKNOWN DATA TYPE")
	}

}
