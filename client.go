package rabbitmq

import (
	"errors"
	"strconv"

	"github.com/streadway/amqp"
)

// Client exported
// Client ...
type Client struct {
	ServerHost string
	ServerPort int

	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      *amqp.Queue

	username string
	password string
}

// NewClient exported
// NewClient ...
func NewClient(host string, port int, username string, password string) *Client {
	return &Client{
		ServerHost: host,
		ServerPort: port,

		username: username,
		password: password}
}

// ConnectToServer exported
// ConnectToServer ...
func (client *Client) ConnectToServer() error {
	conn, err := amqp.Dial(getConnectionURL(client))

	if err != nil {
		return err
	}

	client.Connection = conn

	return nil
}

// OpenChannel exported
// OpenChannel ...
func (client *Client) OpenChannel() error {
	if client.Connection != nil {
		ch, err := client.Connection.Channel()

		if err != nil {
			return err
		}

		client.Channel = ch

		return nil
	}

	return errors.New("rabbbitmq connection not opened")
}

// QueueDeclare exported
// QueueDeclare ...
func (client *Client) QueueDeclare(name string, durable bool, deleteWhenUnused bool, exclusive bool, noWait bool, args map[string]interface{}) error {
	if client.Channel != nil {
		q, err := client.Channel.QueueDeclare(name, durable, deleteWhenUnused, exclusive, noWait, args)

		if err != nil {
			return err
		}

		client.Queue = &q

		return nil
	}

	return errors.New("rabbbitmq channel not initialized")
}

func getConnectionURL(client *Client) string {
	return "amqp://" + client.username + ":" + client.password + "@" + client.ServerHost + ":" + strconv.Itoa(client.ServerPort) + "/"
}
