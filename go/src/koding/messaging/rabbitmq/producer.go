package rabbitmq

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

type Producer struct {
	conn       *amqp.Connection
	channel    *amqp.Channel
	deliveries <-chan amqp.Delivery
	tag        string
	handler    func(deliveries <-chan amqp.Delivery)
	done       chan error
	session    Session
}

type PublishingOptions struct {
	RoutingKey, Tag      string
	Mandatory, Immediate bool
}

func NewProducer(e Exchange, q Queue, po PublishingOptions) (*Producer, error) {
	if po.Tag == "" {
		return nil, errors.New("Tag is not defined in consumer options")
	}

	p := &Producer{
		conn:    nil,
		channel: nil,
		tag:     po.Tag,
		session: Session{
			Exchange:          e,
			Queue:             q,
			PublishingOptions: po,
		},
	}

	err := p.connect()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Producer) connect() error {

	var err error
	// get connection
	p.conn, err = amqp.Dial(getConnectionString())
	if err != nil {
		return err
	}
	handleErrors(p.conn)
	// getting channel
	p.channel, err = p.conn.Channel()
	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) Publish(publishing amqp.Publishing) error {
	e := p.session.Exchange
	q := p.session.Queue
	po := p.session.PublishingOptions

	routingKey := po.RoutingKey
	// if exchange name is empty, this means we are gonna publish
	// this mesage to a queue, every queue has a binding to default exchange
	if e.Name == "" {
		routingKey = q.Name
	}

	err := p.channel.Publish(
		e.Name,       // publish to an exchange(it can be default exchange)
		routingKey,   // routing to 0 or more queues
		po.Mandatory, // mandatory, if no queue than err
		po.Immediate, // immediate, if no consumer than err
		publishing,
		// amqp.Publishing {
		//        // Application or exchange specific fields,
		//        // the headers exchange will inspect this field.
		//        Headers Table

		//        // Properties
		//        ContentType     string    // MIME content type
		//        ContentEncoding string    // MIME content encoding
		//        DeliveryMode    uint8     // Transient (0 or 1) or Persistent (2)
		//        Priority        uint8     // 0 to 9
		//        CorrelationId   string    // correlation identifier
		//        ReplyTo         string    // address to to reply to (ex: RPC)
		//        Expiration      string    // message expiration spec
		//        MessageId       string    // message identifier
		//        Timestamp       time.Time // message timestamp
		//        Type            string    // message type name
		//        UserId          string    // creating user id - ex: "guest"
		//        AppId           string    // creating application id

		//        // The application specific payload of the message
		//        Body []byte
		// }
	)

	return err
}

func (p *Producer) Shutdown() error {
	err := shutdown(p.conn, p.channel, p.tag)
	// change fmt => log
	defer fmt.Println("Producer shutdown OK")
	return err
}

func (p *Producer) RegisterSignalHandler() {
	registerSignalHandler(p)
}
