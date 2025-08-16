package main 

import (
	"fmt"
	"os"
	"bytes"
)

func main () {
	file, err := os.Open("messages.txt")
	var str string = ""

	if err != nil {
		fmt.Println("error opening the file")
	}

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
			fmt.Printf("read: %s\n", str)
			str = ""
		}
		str += string(data)
	}
	if len(str) != 0 {
		fmt.Printf("read: %s\n", str)
	}
}
