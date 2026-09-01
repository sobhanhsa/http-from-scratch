package main

import "fmt"
import "os"
import "log"
import "bytes"
import "io"

func ReadLinesByReader(f io.ReadCloser) <- chan string {

	out := make(chan string,1)

	go func() {
		defer f.Close()
		defer close(out)



		str := ""

		for {
			data := make([]byte,8)

			n,err := f.Read(data)

			if err != nil {
				break
			}
	
			data = data[:n]
	
			if i := bytes.IndexByte(data,'\n'); i != -1 {
				str += string(data[:i])
				data = data[i+1:]	
				out <- str
				str=""
			}

			str += string(data)

		
		}


		if len(str) != 0 {
			out <- str
		}


	}()

	return out

}


func main() {

	file,error := os.Open("message.txt")

	if error != nil {
		log.Fatal("error","error",error)
	}


	linesCh :=  ReadLinesByReader(file)

	for l := range linesCh {
		fmt.Printf("read : %s\n",l)
		
		
	}

}
