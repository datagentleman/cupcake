package tcp

import (
	syscall "golang.org/x/sys/unix"
)

type Client struct {
	connection int
}

func Connect(ip string, port int) (*Client, error) {
	conn, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, err
	}

	addr := IPv4Address(ip, port)
	err = syscall.Connect(conn, addr)
	if err != nil {
		return nil, err
	}

	client := &Client{connection: conn}
	return client, nil
}
