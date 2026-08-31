package main

import "fmt"
import "os"
import "log"
import "bytes"

func main() {

	file,error := os.Open("message.txt")

	if error != nil {
		log.Fatal("error","error",error)
	}

	str := ""

	for {
		data := make([]byte,8)

		n,err := file.Read(data)

		if err != nil {
			break
		}

		data = data[:n]

		if i := bytes.IndexByte(data,'\n'); i != -1 {
			str += string(data[:i])
			data = data[i+1:]	

			fmt.Printf("read (i : %d ) : %s\n",i,str)
			str=""
		}

		str += string(data)

	}


	if len(str) != 0 {
		fmt.Printf("read :%s",str)
	}


}
