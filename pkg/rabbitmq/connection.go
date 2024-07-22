package rabbitmq

import (
	"sync"
	"time"
	"github.com/DanilMankiev/SofiaApplication/pkg/logger"
	"github.com/rabbitmq/amqp091-go"
)

const ( 
	reconnectDelay = 5 * time.Second
	reInitDelay = 2 * time.Second
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
	go client.reconnect(addr)
	return &client
}

func (client *Client) connect(addr string) (*amqp091.Connection,error){
	conn, err:=amqp091.Dial(addr)
	if err!=nil{
		client.logger.Errorf("Failed connect rabbitmq")
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

func( client *Client) reconnect(addr string){
	for{ 
		client.m.Lock()
		client.isReady = false
		client.m.Unlock()

		client.logger.Infof("Attempting to connect")
		
		conn,err:=client.connect(addr)

		if err!=nil{
			client.logger.Errorf("Failed to connect. Reconnecting...")
			
			select {
			case <- client.done:
				return
			case <- time.After(reconnectDelay):
			}
			continue
		}
		if done:= client.reInit(conn); done{
			break
		}
	}
}

func (client *Client) reInit(connection *amqp091.Connection) bool{ 
	for {
		client.m.Lock()
		client.isReady=false
		client.m.Unlock()

		err:=client.init(connection)
		
		if err!=nil{
			client.logger.Infof("Failed to init: %s. Reinitialization...", err.Error())

			select{
			case <- client.done:
				return true
			case <- client.notifyConnClose:
				client.logger.Errorf("Connection closed. Reconnection...")
				return false
			case <- time.After(reInitDelay):
			}
			continue
		}

		select{ 
		case <- client.done: 
			return true
		case <- client.notifyConnClose: 
			client.logger.Infof("Connection closed. Reconnecting...")
			return false 
		case <- client.notifyChanClose:
			client.logger.Infof("Channel closed. Reconnecting...")
		}
	}

}

func (client *Client) init(connection *amqp091.Connection) error{
	channel,err:= connection.Channel()
	if err!=nil{
		return err
	}
	err=channel.Confirm(false)

	if err!=nil{
		return err
	}

	for _,v:= range client.cfg.Exchanges{
		err =channel.ExchangeDeclare(v.Name,v.Kind,v.Durable,v.AutoDelete,v.Internal,v.NoWait,v.Args) 
		if err!=nil{
			return err
		}
	}
	for _,v:= range client.cfg.Queues{
		_,err=channel.QueueDeclare(v.Name,v.Durable,v.AutoDelete,v.Exclusive,v.NoWait,v.Args)
		if err!=nil{
			return err
		}
	}
	for _,v := range client.cfg.Bindings{
		err:= channel.QueueBind(v.QueueName,v.RoutingKey,v.ExchangeName,v.NoWait,v.Args)
		if err!=nil{
			return err
		}
	}
	client.changeChan(channel)
	client.m.Lock()
	client.isReady=true
	client.m.Unlock()
	client.logger.Infof("Setup complete")
	return nil
	}
