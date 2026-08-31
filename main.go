package main

import "fmt"
import "os"
import "log"

func main() {

	file,error := os.Open("message.txt")

	if error != nil {
		log.Fatal("error","error",error)
	}

	for {
		data := make([]byte,8)

		n,err := file.Read(data)

		if err != nil {
			break
		}

		fmt.Printf("read %s\n",string(data[:n]))
	}


}
