package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)

			conn, err := net.Dial("tcp", address)
			defer conn.Close()

			if err != nil {
				return
			}

			fmt.Printf("%d open\n", j)
		}(i)
	}
}
