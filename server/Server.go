package main

import (
	"log"
	"net"
	"time"

	"github.com/bwesterb/go-pow"
	"github.com/google/uuid"
)

const quote = "If you don't know where you are going, any road will get you there."

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	log.Printf("listen on: %s \n", conn.RemoteAddr().String())
	req := pow.NewRequest(10, []byte(uuid.NewString()))
	conn.Write([]byte(req))
	var n int
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	proof := make([]byte, 1024)
	for n == 0 {
		var err error
		n, err = conn.Read(proof)
		if err != nil {
			conn.Write([]byte(err.Error()))
			return
		}
	}
	ok, err := pow.Check(req, string(proof[0:n]), []byte("WordOfWisdom"))
	if !ok {
		conn.Write([]byte(err.Error()))
		return
	}
	conn.Write([]byte(quote))
}
