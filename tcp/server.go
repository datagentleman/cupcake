package tcp

import (
	"fmt"
	syscall "golang.org/x/sys/unix"
	"net"
)

func handleClientConnection(connection int) {
	syscall.Write(connection, []byte("Hello from server :)"))
}

func parseIPv4Address(ip string, port int) *syscall.SockaddrInet4 {
	ip4 := net.ParseIP(ip).To4()

	return &syscall.SockaddrInet4{
		Port: port,
		Addr: [4]byte{ip4[0], ip4[1], ip4[2], ip4[3]},
	}
}

func StartServer(ip string, port int) error {
	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)

	err := syscall.Bind(fd, parseIPv4Address(ip, port))
	if err != nil {
		return err
	}

	err = syscall.Listen(fd, 10)
	if err != nil {
		return err
	}

	for {
		c, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Print(err)
		}
		go handleClientConnection(c)
	}
}
