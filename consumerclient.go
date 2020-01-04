package rabbitmq

// ConsumerClient exported
// ConsumerClient ...
type ConsumerClient struct {
	Client *Client
}

// NewConsumerClient exported
// NewConsumerClient ...
func NewConsumerClient(host string, port int, username string, password string) *ConsumerClient {
	return &ConsumerClient{
		Client: NewClient(host, port, username, password)}
}

// ConsumeMessages exported
// ConsumeMessages ...
func (cc *ConsumerClient) ConsumeMessages(consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args map[string]interface{}) {
	cc.Client.Channel.Consume(cc.Client.Queue.Name, consumer, autoAck, exclusive, noLocal, noWait, args)
}
