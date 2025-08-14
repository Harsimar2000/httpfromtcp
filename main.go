package main 

import (
	"fmt"
	"os"
)

func main () {
	file, err := os.Open("messages.txt")
	data := make([]byte, 8)
	
	if err != nil {
		fmt.Println("error opening the file")
	}
	
	for  {
		count, err := file.Read(data)

		if err != nil {
			return
		}
	
		fmt.Printf("read: %s\n", data[:count])
	}
}
