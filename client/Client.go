package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bwesterb/go-pow"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		log.Fatalln(err.Error())
	}
	proof, err := pow.Fulfil(string(buff[0:n]), []byte("WordOfWisdom"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	conn.Write([]byte(proof))
	quote := make([]byte, 1024)
	n, err = conn.Read([]byte(quote))
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(string(quote[0:n]))
}
