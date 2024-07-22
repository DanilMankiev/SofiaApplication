package rabbitmq

import (
	"log"
	"sync"

	"github.com/DanilMankiev/SofiaApplication/pkg/logger"
	"github.com/rabbitmq/amqp091-go"
)

type Client struct {
	m       *sync.Mutex
	cfg Config
	logger  logger.Logger

	connection *amqp091.Connection
	channel    *amqp091.Channel

	done chan bool

	notifyConnClose chan *amqp091.Error
	notifyChanClose chan *amqp091.Error
	notifyConfirm   chan amqp091.Confirmation

	isReady bool
}

func New(addr string,logger logger.Logger, cfg *Config) *Client{
	client:=Client{
		m: &sync.Mutex{},
		logger: logger,
		done: make(chan bool),
		cfg: *cfg,
	}
	return &client
}

func (client *Client) connect(addr string) (*amqp091.Connection,error){
	conn, err:=amqp091.Dial(addr)
	if err!=nil{
		client.logger.Errorf("failed connect rabbitmq")
		return nil, err
	}

	client.changeConnect(conn)
	client.logger.Infof("Connected")
	return conn,nil
}

func (client *Client) changeConnect(connection *amqp091.Connection){
	client.connection = connection
	client.notifyConnClose=make(chan *amqp091.Error,1)
	client.connection.NotifyClose(client.notifyChanClose)
}
func (client *Client) changeChan(channel *amqp091.Channel){
	client.channel = channel
	client.notifyChanClose=make(chan *amqp091.Error, 1)
	client.notifyConfirm=make( chan amqp091.Confirmation,1)
	client.channel.NotifyClose(client.notifyChanClose)
	client.channel.NotifyPublish(client.notifyConfirm)
	
	
}