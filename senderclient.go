package rabbitmq

import (
	"errors"

	"github.com/streadway/amqp"
)

// SenderClient exported
// SenderClient ...
type SenderClient struct {
	Client *Client
}

// NewSenderClient exported
// NewSenderClient ...
func NewSenderClient(host string, port int, username string, password string) *SenderClient {
	return &SenderClient{
		Client: NewClient(host, port, username, password)}
}

// SendTxtMessage exported
// SendTxtMessage ...
func (sc *SenderClient) SendTxtMessage(exchange string, mandatory bool, immediate bool, msg string) error {

	if sc.Client.Queue != nil {
		return sendMessage(sc, exchange, sc.Client.Queue.Name, mandatory, immediate, "text/plain", msg)
	}

	return errors.New("rabbitmq queue not initialized")
}

func sendMessage(sc *SenderClient, exchange string, routingKey string, mandatory bool, immediate bool, contentType string, body string) error {

	if sc.Client.Channel != nil {
		return sc.Client.Channel.Publish(exchange, routingKey, mandatory, immediate, amqp.Publishing{
			ContentType: contentType,
			Body:        []byte(body)})
	}

	return errors.New("rabbbitmq channel not initialized")
}
