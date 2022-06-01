package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
)

var (
	confFileName = "conf.json"
	conf         = &Conf{}
)

func init() {
	loadConfiguration()
}

func main() {
	listener, err := net.Listen("tcp", conf.Listener)
	checkErr(err, "failed to listen: ")
	defer func() {
		err = listener.Close()
		checkErr(err, "failed to close listener: ")
	}()

	for {
		conn, err := listener.Accept()
		checkErr(err, "failed to accept connection: ")
		
		server := conf.Servers[getRand(len(conf.Servers))]
		log.Printf("server=%s", server)

		go func() {
			err = proxy(server, conn)
			checkErr(err, "proxy error")
		}()
	}
}

func proxy(server string, conn net.Conn) error {
	conn2, err := net.Dial("tcp", server)
	if err != nil {
		return fmt.Errorf("failed to connect server %s: %v", server, err)
	}

	// Load Balancer -> Server
	go io.Copy(conn, conn2)

	// Server -> Load balancer
	go io.Copy(conn2, conn)

	return nil
}

func loadConfiguration() {
	f, err := os.Open(confFileName)
	checkErr(err, "")

	data, err := io.ReadAll(f)
	checkErr(err, "")

	err = json.Unmarshal(data, conf)
	checkErr(err, "")
}

func checkErr(e error, message string) {
	if e != nil {
		log.Println(message, e)
	}
}

func getRand(max int) int {
	return rand.Intn(max)
}

type Conf struct {
	Name     string   `json:"name"`
	Listener string   `json:"listener"`
	Servers  []string `json:"servers"`
}
