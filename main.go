package main 

import (
	"fmt"
	"os"
	"bytes"
	"io"
)

func main () {
	file, err := os.Open("messages.txt")

	if err != nil {
		fmt.Println("error opening the file")
	}
	
	lines := getLinesChannel(file)

	for line := range(lines) {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {
	out := make(chan string, 1)
	var str string = ""

	go func(){
		defer close(out)
		defer file.Close()
		for  {
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
