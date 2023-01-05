package main

import (
	//"fmt"
	"log"
	"net"
	"os"
	"time"
)

func dummy() {
	listner, err := net.Listen("tcp", "localhost:4000")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for {
		con, err := listner.Accept()

		if err != nil {
			log.Fatal(err)
			continue
		}

		go func() {
			for {
				str := time.Now().Format("15:05:05")
				con.Write([]byte(str + "\n"))
				time.Sleep(time.Second)

			}
		}()

	}
}
