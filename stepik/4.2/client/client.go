package client

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	for i := 0; i < 3; i++ {
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(strings.ToUpper(string(buf[:read])))
	}
}
