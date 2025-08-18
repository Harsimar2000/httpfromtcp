package main 

import (
	"fmt"
	"net"
	"bytes"
	"io"
)

func main () {
	ln, err := net.Listen("tcp", ":42069")

	if err != nil {
		fmt.Println("error listening to the tcp connection")
	}
	
	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("error Accecpting the connection")
		}

		for line := range getLinesChannel(conn) {
			fmt.Printf("read: %s\n", line)
		}
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {
	out := make(chan string, 2)
	var str string = ""

	go func(){
		defer close(out)
		defer file.Close()
		for {
			
			data := make([]byte, 8)
			count, err := file.Read(data)
			data = data[:count]

			if err != nil {
				return
			}

			if i:=bytes.IndexByte(data,'\n'); i!=-1 {
				str += string(data[:i])
				data = data[i+1:]
				out <- str
				str = ""
			}
			str += string(data)
		}
		if len(str) != 0 {
			out<-str
		}
	} ()
	return out
}
