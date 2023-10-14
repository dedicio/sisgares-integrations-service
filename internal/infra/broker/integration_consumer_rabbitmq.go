package broker

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/dedicio/sisgares-integrations-service/internal/dto"
	"github.com/dedicio/sisgares-integrations-service/internal/entity"
	"github.com/dedicio/sisgares-integrations-service/internal/infra/repository"
	"github.com/dedicio/sisgares-integrations-service/internal/usecase"
	rbmq_utils "github.com/dedicio/sisgares-integrations-service/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

type Consumer struct {
	conn       *amqp.Connection
	queueName  string
	repository entity.IntegrationRepository
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	return rbmq_utils.DeclareExchange(channel, entity.ExchangeName)
}

func NewConsumer(
	db *sql.DB,
	conn *amqp.Connection,
) (Consumer, error) {
	repository := repository.NewIntegrationRepositoryPostgresql(db)
	consumer := Consumer{
		conn:       conn,
		repository: repository,
	}
	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) Consume(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := rbmq_utils.DeclareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, topic := range topics {
		err = ch.QueueBind(
			queue.Name,
			topic,
			entity.ExchangeName,
			false,
			nil,
		)

		if err != nil {
			return err
		}
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			order := &dto.OrderMessagingDTO{}
			err := json.Unmarshal(d.Body, order)
			if err != nil {
				log.Printf("Error on unmarshal order: %s", err)
				continue
			}

			usecase.NewIntegrateOnOrderCreatedUseCase(consumer.repository).Execute(order)
		}
	}()
	<-forever
	return nil
}
