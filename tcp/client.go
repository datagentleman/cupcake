package tcp

import (
	syscall "golang.org/x/sys/unix"
)

type Client struct {
	connection int
}

func (c *Client) Hello() (string, error) {
	_, err := syscall.Write(c.connection, Message("HELLO"))
	if err != nil {
		return "", err
	}

	r, err := ReadMessage(c)
	if err != nil {
		return "", err
	}

	return string(r), nil
}

func Connect(ip string, port int) (*Client, error) {
	conn, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, err
	}

	err = syscall.Connect(conn, IPv4Address(ip, port))
	if err != nil {
		return nil, err
	}

	client := &Client{connection: conn}
	return client, nil
}
