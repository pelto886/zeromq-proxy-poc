package main

import (
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

const (
	recvTimeout = 90 * time.Millisecond
)

var dropCount = 0

func main() {
	//  Socket to talk to server
	log.Println("starting requests...")
	requester, _ := zmq.NewSocket(zmq.REQ)
	requester.SetRcvtimeo(recvTimeout)
	defer requester.Close()
	requester.Connect("tcp://localhost:5554")

	for i := 1; i <= 100; i++ {
		log.Println("Sending request ", i, "…")
		requester.SendMessage("do work!")

		reply, err := requester.RecvMessage(0)
		if err != nil {
			log.Println("Error receiving reply ", i, err)
			dropCount++
		}
		log.Println("dropCount: ", dropCount)
		log.Println("Received reply ", i, reply)
	}
}
