package main

import (
	"log"

	"github.com/Stoina/go-rabbitmq"
)

func main() {
	testSenderClient()
}

func testSenderClient() {
	senderClient := rabbitmq.NewSenderClient("localhost", 5672, "guest", "guest")

	err := senderClient.Client.ConnectToServer()

	if err != nil {
		log.Println("Connection to RabbitMQ-Server failed: " + err.Error())
	}

	log.Println("Successfully connected to RabbitMQ-Server")

	err = senderClient.Client.OpenChannel()

	if err == nil {

		err = senderClient.Client.QueueDeclare("hello", false, false, false, false, nil)

		if err == nil {

			err = senderClient.SendTxtMessage("", false, false, "Hello World")

			if err == nil {
				defer senderClient.Client.Connection.Close()

				for {

				}
			}

			log.Println("Send Txt Message failed: " + err.Error())
		}

		log.Println("Declare queue failed: " + err.Error())

	}

	senderClient.Client.Connection.Close()
}

func testConsumerClient() {

}
