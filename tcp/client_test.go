package tcp

import (
	"fmt"
	"testing"
)

func Test_Connect(t *testing.T) {
	client, err := Connect("127.0.0.1", 2001)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	r, err := client.Hello()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	fmt.Printf("HELLO:%s\n", r)
}
