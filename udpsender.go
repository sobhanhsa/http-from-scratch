package main

import "fmt"
import "log"
import "net"
import "bufio"
import "os"

func main() {

	udpAddr , err := net.ResolveUDPAddr("udp","localhost:42069")

	if err != nil {
		log.Fatal("error","error",err)
	}


	conn , err := net.DialUDP("udp",nil,udpAddr)

	
	if err != nil {
		log.Fatal("error","error",err)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print(">")
		l,err := reader.ReadString('\n')
	
		if err != nil {
			log.Fatal("error","error",err)
		}

		conn.Write([]byte(l))

	}
}
