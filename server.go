package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"bytes"
	"strings"
	"strconv"
	"math/rand"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "44415"
	CONN_TYPE = "tcp"
)

var (
	stockPrice = make(map[string]int)
)

func quoteHandler(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		fmt.Println(err.Error())
	}

	commandLength := bytes.Index(buffer, []byte{0})
	commandText := string(buffer[:commandLength - 1])
	stock := strings.Split(commandText, ",")[0]

	if price, exists := stockPrice[stock]; exists {
		conn.Write([]byte(strconv.Itoa(price)))
	} else {
		stockPrice[stock] = rand.Intn(1000 - 20) + 20
		conn.Write([]byte(strconv.Itoa(stockPrice[stock])))
	}

	conn.Close()
}

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Cannot listen on port: %s", CONN_PORT)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rand.Seed(time.Now().Unix())

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		go quoteHandler(conn)
	}

}