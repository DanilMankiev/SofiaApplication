package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

const (
	publishTimeout = time.Second*30
	republishDelay = time.Second*5 
)

func (client *Client) Publish(exchange, routingKey string, mandatory, immediate bool, body []byte) error{
	client.m.Lock()
	if !client.isReady{
		client.m.Unlock()
		client.logger.Errorf("Not ready to publish: %s", errNotConnected.Error())
		return errNotConnected
	}
	client.m.Unlock()

	data,err:=json.Marshal(body)
	if err!=nil{
		client.logger.Errorf("Cannot marshal body msg: %s", err.Error())
		return err
	}
	
	ctx,cancel := context.WithTimeout(context.Background(), publishTimeout)
	defer cancel()

	for {
		err:=client.channel.PublishWithContext(
			ctx,
			exchange,
			routingKey,
			mandatory,
			immediate,
			amqp091.Publishing{
				ContentType: "application/json",
				Body: data,
			},
		)
		if err!=nil{
			client.logger.Errorf("Failed to publish: %s. Republishing...",err.Error())
			
			select{
			case <-time.After(republishDelay):
			case <-ctx.Done():
				client.logger.Errorf("Timeout publishing")
				return errTimeoutPublishing
			}
			continue
		}
		return nil	
	}
}