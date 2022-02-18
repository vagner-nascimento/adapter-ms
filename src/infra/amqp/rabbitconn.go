package amqp

import (
	"log"
	"os"

	stAmq "github.com/streadway/amqp"
)

type rabbitConnection struct {
	conn *stAmq.Connection
}

func (rb *rabbitConnection) isConnected() bool {
	return rb.conn != nil && !rb.conn.IsClosed()
}

func (rb *rabbitConnection) connect() {
	connStr := os.Getenv("AMQP_CONN")
	if len(connStr) == 0 {
		log.Fatalf("FATAL - ENV AMQP_CONN not informed")
	}

	var err error
	if rb.conn, err = stAmq.Dial(connStr); err != nil {
		log.Fatalf("FATAL - error on connect on rabbit mq server: " + err.Error())
	}

}

var singConn rabbitConnection

func getChannel() (*stAmq.Channel, error) {
	if !singConn.isConnected() {
		singConn.connect()
	}

	return singConn.conn.Channel()
}

func ListenAmqpCloseConn() <-chan *stAmq.Error {
	ch := make(chan *stAmq.Error)

	go func() {
		if singConn.isConnected() {
			singConn.conn.NotifyClose(ch)
		} else {
			ch <- &stAmq.Error{
				Reason: "connection is alredy closed",
			}
		}
	}()

	return ch
}
