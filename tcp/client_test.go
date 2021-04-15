package tcp

import (
	"fmt"
	"testing"
)

func Test_Connect(t *testing.T) {
	_, err := Connect("127.0.0.1", 2001)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
