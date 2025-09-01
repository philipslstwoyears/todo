package server

import (
	"fmt"
	"net"
)

func main() {
	serverConn, err := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 10001})
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}
	defer serverConn.Close()
	for {
		buf := make([]byte, 1024)
		n, _, err := serverConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("Received:", string(buf[:n]))
	}
}
