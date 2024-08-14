package main

import (
	"fmt"
	"net"
)

func main() {
	address := "scanme.nmap.org"
	fmt.Println("Running a TCP port scanner on", address)

	for i := 1; i <= 1024; i++ {
		socket := fmt.Sprintf("scanme.nmap.org:%d", i)

		conn, err := net.Dial("tcp", socket)

		if err != nil {
			continue

		}
		conn.Close()
		fmt.Println("Port", i, "is open")
	}
}
