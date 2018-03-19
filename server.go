package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8081"
	CONN_TYPE = "tcp"
)

var (
	stockPrice = make(map[string]float64)
)

func quoteHandler(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		fmt.Println(err.Error())
	}

	commandLength := bytes.Index(buffer, []byte{0})
	commandText := string(buffer[:commandLength-1])
	commandComponents := strings.Split(commandText, ",")
	stock := commandComponents[0]
	userId := commandComponents[1]

	if _, exists := stockPrice[stock]; !exists {
		stockPrice[stock] = (rand.Float64() * 1000) + 1
	}

	responseString := strconv.FormatFloat(stockPrice[stock], 'f', 2, 64) + ","
	responseString += stock + "," + userId + ","
	responseString += strconv.FormatInt(int64(time.Nanosecond)*int64(time.Now().UnixNano())/int64(time.Millisecond), 10) + ","
	responseString += strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)
	conn.Write([]byte(responseString))

	conn.Close()
}

func main() {
	l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if err != nil {
		fmt.Println("Cannot listen on port: ", CONN_PORT)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("listening on " + CONN_HOST+":"+CONN_PORT)

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
