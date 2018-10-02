package main

import (
	"fmt"
	"net"
	"encoding/gob"
)

type Quote struct {
	Symbol string
	Ltp float64
}

func server() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("connected...")
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	msg := new(Quote)
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println("error.. server recv.")
	} else {
		fmt.Println("Rcvd : " , msg.Symbol)
	}

	c.Close()
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	quote := Quote{"INFY", 745.6}
	fmt.Println(quote)
	//msg := "hello testing tcp client send"
	err = gob.NewEncoder(c).Encode(quote)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
}
