package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	address := "scanme.nmap.org"
	fmt.Println("Running a TCP port scanner on", address)

	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			socket := fmt.Sprintf("scanme.nmap.org:%d", i)
			conn, err := net.Dial("tcp", socket)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println("Port", i, "is open")
		}(i)
	}
	wg.Wait()
}
