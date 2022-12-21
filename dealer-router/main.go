package main

import (
	"fmt"
	"log"

	"github.com/pebbe/zmq4"
)

func main() {
	//  Socket to talk to clients
	frontend, _ := zmq4.NewSocket(zmq4.ROUTER)
	defer frontend.Close()
	frontend.Bind("tcp://*:5554")

	//  Socket to talk to workers
	backend, _ := zmq4.NewSocket(zmq4.DEALER)
	defer backend.Close()
	backend.Bind("tcp://*:5555")

	//  Initialize poll set
	poller := zmq4.NewPoller()
	poller.Add(frontend, zmq4.POLLIN)
	poller.Add(backend, zmq4.POLLIN)

	//  Switch messages between sockets
	for {
		log.Println("polling")
		sockets, err := poller.Poll(-1)
		if err != nil {
			break //  Context has been shut down
		}
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case frontend:
				fmt.Println("frontend triggered.")
				msg, err := frontend.RecvMessage(0)
				if err != nil {
					break
				}
				backend.SendMessage(msg, 0)
			case backend:
				fmt.Println("backend triggered.")
				msg, err := backend.RecvMessage(0)
				if err != nil {
					break
				}
				frontend.SendMessage(msg)
			}
		}
	}
}
