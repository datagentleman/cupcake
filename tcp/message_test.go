package tcp

import (
	"fmt"
	"testing"
)

func Test_CreateMessage(t *testing.T) {
	a := "😍"
	b := []int8{1, 30, 3}

	m := Message(a, b)
	fmt.Printf("%v\n", m)
}
