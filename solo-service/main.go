package main

import (
	"log"
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"
)

var Count = 0

func main() {
	//  Socket to talk to clients
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Bind("tcp://*:5554")

	for {
		log.Println("waiting for request")
		//  Wait for next request from client
		msg, _ := responder.RecvMessage(0)
		log.Printf("Received request: [%s]", msg)

		//  Do some 'work'
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		Count++
		//  Send reply back to client
		log.Printf("number of jobs done: %d", Count)
		responder.SendMessage("done!")
	}
}
