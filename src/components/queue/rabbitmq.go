package queue

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/streadway/amqp"
)

const (
	CHANBUFFER   = 1000
	ExchangeType = "topic"
	Retry        = 3
)

type RMQ struct {
	conn           *amqp.Connection
	channel        *amqp.Channel
	exchange       string
	routingKey     string
	queues         map[string]bool
	addr           string
	queueName      string
	UnSendMsgQueue *Queue
	Delivery       <-chan amqp.Delivery
}

// 初始化 rmqAddr参数格式：amqp://用户名:密码@地址:端口号/host
func InitRMQ(rmqAddr, exchange, routingkey string) (*RMQ, error) {

	conn, err := amqp.Dial(rmqAddr)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	rmq := &RMQ{
		conn:    conn,
		channel: channel,
		addr:    rmqAddr,
	}
	rmq.UnSendMsgQueue = NewQueue()

	if exchange == "" {
		return nil, errors.New("exchange is empty")
	}
	err = rmq.channel.ExchangeDeclare(exchange, ExchangeType, true, false, false, true, nil)
	if err != nil {
		return nil, err
	}
	rmq.exchange = exchange

	if routingkey == "" {
		return nil, errors.New("RoutingKey is empty")
	}

	rmq.routingKey = routingkey

	err = rmq.Ping()
	if err != nil {
		return nil, err
	}
	rmq.queues = make(map[string]bool)

	return rmq, nil
}

func InitQuenue(rmqAddr, quenueName string) (*RMQ, error) {

	conn, err := amqp.Dial(rmqAddr)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	rmq := &RMQ{
		conn:    conn,
		channel: channel,
		addr:    rmqAddr,
	}
	rmq.UnSendMsgQueue = NewQueue()

	_, err = rmq.channel.QueueDeclare(quenueName, true, false, false, true, nil)
	if err != nil {
		return nil, err
	}
	rmq.queueName = quenueName

	err = rmq.Ping()
	if err != nil {
		return nil, err
	}
	rmq.queues = make(map[string]bool)

	return rmq, nil
}

func ReconnectRMQ(rmq *RMQ) (err error) {
	conn, err := amqp.Dial(rmq.addr)
	if err != nil {
		return err
	}
	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	rmq.Close()
	rmq.channel = channel
	rmq.conn = conn

	err = rmq.channel.ExchangeDeclare(rmq.exchange, ExchangeType, true, false, false, true, nil)
	if err != nil {
		return err
	}

	err = rmq.Ping()
	if err != nil {
		return err
	}

	rmq.queues = make(map[string]bool)

	return nil
}

// 测试连接是否正常
func (rmq *RMQ) Ping() (err error) {

	if rmq.channel == nil {
		return errors.New("RabbitMQ is not initialize")
	}

	err = rmq.channel.ExchangeDeclare("ping.ping", "topic", false, true, false, true, nil)
	if err != nil {
		return err
	}

	msgContent := "ping.ping"

	err = rmq.channel.Publish("ping.ping", "ping.ping", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})

	if err != nil {
		return err
	}

	err = rmq.channel.ExchangeDelete("ping.ping", false, false)

	return err
}

// 发布消息
func (rmq *RMQ) Publish(msg []byte) error {
	var err error

	for i := 0; i < Retry; i++ {
		err = rmq.channel.Publish(rmq.exchange, rmq.routingKey, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}

	return err
}

func (rmq *RMQ) DealUnSendMsgQueue() error {
	for !rmq.UnSendMsgQueue.Empty() {

		sendmsg := rmq.UnSendMsgQueue.Peek()
		if reflect.TypeOf(sendmsg) == nil || reflect.TypeOf(sendmsg).String() != "[]uint8" {
			rmq.UnSendMsgQueue.Remove()
			return fmt.Errorf("sendmsg[%s] type[%s] isnot []uint8", sendmsg, reflect.TypeOf(sendmsg))
		}

		err := rmq.Publish(sendmsg.([]byte))
		if err != nil {
			return err
		}
		rmq.UnSendMsgQueue.Remove()
	}
	return nil
}

// 监听接收到的消息
func (rmq *RMQ) Receive(queue string) error {
	var err error

	if queue == "" {
		return errors.New("queue is empty")
	}
	if _, ok := rmq.queues[queue]; !ok {
		_, err = rmq.channel.QueueDeclare(queue, true, false, false, true, nil)
		if err != nil {
			return err
		}
		err = rmq.channel.QueueBind(queue, rmq.routingKey, rmq.exchange, true, nil)
		if err != nil {
			return err
		}
		rmq.queues[queue] = true
	}

	//No-AutoAck
	msgs, err := rmq.channel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	rmq.Delivery = msgs

	return nil
}

// 关闭连接
func (rmq *RMQ) Close() {
	rmq.channel.Close()
	rmq.conn.Close()
	rmq.queues = make(map[string]bool)
}
