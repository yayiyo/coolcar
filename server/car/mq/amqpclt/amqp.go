package amqpclt

import (
	"context"
	"encoding/json"

	carpb "coolcar/car/api/gen/v1"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Publisher struct {
	ch       *amqp.Channel
	exchange string
}

func NewPublisher(conn *amqp.Connection, exchange string) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "create channel")
	}

	err = declareExchange(ch, exchange)
	if err != nil {
		return nil, errors.Wrap(err, "declare exchange")
	}

	return &Publisher{
		ch:       ch,
		exchange: exchange,
	}, nil
}

func (p *Publisher) Publish(ctx context.Context, car *carpb.CarEntity) error {
	b, err := json.Marshal(car)
	if err != nil {
		return errors.Wrap(err, "failed to marshal")
	}

	return p.ch.PublishWithContext(ctx, p.exchange,
		"", false, false, amqp.Publishing{
			Body: b,
		},
	)
}

type Subscriber struct {
	conn     *amqp.Connection
	exchange string
	logger   *zap.Logger
}

func NewSubscriber(conn *amqp.Connection, exchange string, logger *zap.Logger) (*Subscriber, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "create channel")
	}
	defer ch.Close()

	err = declareExchange(ch, exchange)
	if err != nil {
		return nil, errors.Wrap(err, "declare exchange")
	}
	return &Subscriber{
		conn:     conn,
		exchange: exchange,
		logger:   logger,
	}, nil
}

func (s *Subscriber) SubscribeRaw(ctx context.Context) (<-chan amqp.Delivery, func(), error) {
	ch, err := s.conn.Channel()
	if err != nil {
		return nil, func() {}, errors.Wrap(err, "failed to create channel")
	}
	//defer ch.Close()
	closeCh := func() {
		err := ch.Close()
		if err != nil {
			s.logger.Error("failed to close channel", zap.Error(err))
		}
	}

	q, err := ch.QueueDeclare("", false, true, false, false, nil)
	if err != nil {
		return nil, closeCh, errors.Wrap(err, "failed to declare queue")
	}

	cleanUp := func() {
		_, err := ch.QueueDelete(q.Name, false, false, false)
		if err != nil {
			s.logger.Error("failed to delete queue", zap.String("name", q.Name), zap.Error(err))
		}
		closeCh()
	}

	err = ch.QueueBind(q.Name, "", s.exchange, false, nil)
	if err != nil {
		return nil, cleanUp, errors.Wrap(err, "failed to bind queue")
	}

	msg, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, cleanUp, errors.Wrap(err, "failed to start consume queue")
	}
	return msg, cleanUp, nil
}

func (s *Subscriber) Subscribe(ctx context.Context) (chan *carpb.CarEntity, func(), error) {
	msgChan, f, err := s.SubscribeRaw(ctx)
	if err != nil {
		return nil, f, err
	}

	carCh := make(chan *carpb.CarEntity)
	go func() {
		for msg := range msgChan {
			var car = &carpb.CarEntity{}
			err = json.Unmarshal(msg.Body, car)
			if err != nil {
				s.logger.Error("failed to unmarshal", zap.String("data", string(msg.Body)), zap.Error(err))
				continue
			}
			carCh <- car
		}
		close(carCh)
	}()

	return carCh, f, nil
}

func declareExchange(ch *amqp.Channel, exchange string) error {
	return ch.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
}
