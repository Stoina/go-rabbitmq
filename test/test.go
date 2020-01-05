package main

import (
	"log"
	"os"
	"time"

	"github.com/Stoina/go-rabbitmq"
)

func main() {

	if len(os.Args) > 1 {
		testModul := os.Args[1]

		if testModul == "sender" {
			testSenderClient()
		} else if testModul == "consumer" {
			testConsumerClient()
		}
	}

	testConsumerClient()
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

			for {
				defer senderClient.Client.Connection.Close()

				err = senderClient.SendTxtMessage("", false, false, "Hello World")

				if err != nil {
					log.Println("Send Txt Message failed: " + err.Error())
				}

				time.Sleep(15000)
			}
		}

		log.Println("Declare queue failed: " + err.Error())

	}

	senderClient.Client.Connection.Close()
}

func testConsumerClient() {
	consumerClient := rabbitmq.NewConsumerClient("localhost", 5672, "guest", "guest")

	err := consumerClient.Client.ConnectToServer()

	if err != nil {
		log.Println("Connection to RabbitMQ-Server failed: " + err.Error())
	}

	err = consumerClient.Client.OpenChannel()

	if err == nil {

		err = consumerClient.Client.QueueDeclare("hello", false, false, false, false, nil)

		if err == nil {

			defer consumerClient.Client.Connection.Close()

			for {
				msgs, err := consumerClient.ConsumeMessages("", true, false, false, false, nil)

				if err != nil {
					log.Printf("Failed Conusme Messages: %s", err.Error())
				}

				for d := range msgs {
					log.Printf("Received a message: %s", d.Body)
				}

				log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
			}
		}

		log.Println("Declare queue failed: " + err.Error())

	}

	consumerClient.Client.Connection.Close()
}
