package main

import (
	"bufio"
	"errors"
	"net"
	"strings"
	"time"
)

const prompt = "server> "

func handleConnection(conn net.Conn) {
	methods := map[string]func(*ClientMessage){
		"GET": getHandler,
	}

	for {
		conn.Write([]byte(prompt))
		text, _ := bufio.NewReader(conn).ReadString(byte('\n'))

		if len(text) > len("\n") {
			println("server received:", text)
			method, err := parseInput(text, methods)
			if err != nil {
				conn.Write([]byte("Invalid Query!\n"))
			} else {
				conn.Write([]byte(method + " request received!\n"))
				time.Sleep(1000 * time.Millisecond)
				conn.Write([]byte("Done!\n"))
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}
}

// ClientMessage action for server to take
type ClientMessage struct {
	method     string
	statusCode string
}

func parseInput(s string, methods map[string]func(*ClientMessage)) (string, error) {
	content := strings.SplitAfterN(s, " ", 2)
	m := strings.Trim(content[0], "\r\n ")
	if _, ok := methods[m]; ok {
		// return handler(m, content), nil
		return m, nil
	}
	return "", errors.New("Invalid query")
}

func getHandler(m *ClientMessage) {}

func main() {
	ln, _ := net.Listen("tcp", "localhost:1234")
	conn, _ := ln.Accept()
	handleConnection(conn)
}
