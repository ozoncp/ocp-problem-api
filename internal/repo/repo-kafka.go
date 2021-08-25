package repo

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
)

type kafkaConnector func()

type repoKafka struct {
	RepoRemover
	brokers []string
	config *sarama.Config
	producer sarama.SyncProducer
}

func (rk *repoKafka) AddEntities(_ context.Context, problems []utils.Problem) error {
	var errMethod error

	producer, err := rk.getProducer()
	if err != nil {
		return err
	}

	producerMessages := make([]*sarama.ProducerMessage, 0, len(problems))
	for _, problem := range problems{
		message, err := prepareMessage("add", problem)
		if err != nil {
			errMethod = utils.NewWrappedError(err.Error(), errMethod)
		} else {
			producerMessages = append(producerMessages, message)
		}
	}

	if err := producer.SendMessages(producerMessages); err != nil {
		errMethod = utils.NewWrappedError(err.Error(), errMethod)
	}

	return errMethod
}

func (rk *repoKafka) UpdateEntity(_ context.Context, problem utils.Problem) error {
	var errMethod error

	producer, err := rk.getProducer()
	if err != nil {
		return err
	}

	message, err := prepareMessage("update", problem)
	if err != nil {
		errMethod = utils.NewWrappedError(err.Error(), errMethod)
	}

	if _, _, err := producer.SendMessage(message); err != nil {
		errMethod = utils.NewWrappedError(err.Error(), errMethod)
	}

	return errMethod
}

func (rk *repoKafka) ListEntities(_ context.Context, _, _ uint64) ([]utils.Problem, error) {
	return nil, nil
}

func (rk *repoKafka) DescribeEntity(_ context.Context, _ uint64) (*utils.Problem, error) {
	return nil, nil
}

func (rk *repoKafka) RemoveEntity(_ context.Context, entityId uint64) error {
	var errMethod error

	producer, err := rk.getProducer()
	if err != nil {
		return err
	}

	message, err := prepareMessage("remove", entityId)
	if err != nil {
		errMethod = utils.NewWrappedError(err.Error(), errMethod)
	}

	if _, _, err := producer.SendMessage(message); err != nil {
		errMethod = utils.NewWrappedError(err.Error(), errMethod)
	}

	return errMethod
}

func (rk *repoKafka) getProducer() (sarama.SyncProducer, error) {
	if rk.producer != nil {
		return rk.producer, nil
	}

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(rk.brokers, config)
	if err != nil {
		return nil, err
	}

	rk.producer = producer
	return rk.producer, nil
}

func NewRepoKafka(brokers []string) RepoRemover {
	return &repoKafka{brokers: brokers}
}


func prepareMessage(topic string, message interface{}) (*sarama.ProducerMessage, error) {
	b, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(string(b)),
	}

	return msg, nil
}
