package consumer

import (
	"context"
	"encoding/json"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/save_account"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type balanceConsumer struct {
	Consumer                     *kafka.Consumer
	updateAccountsBalanceUseCase save_account.UpdateAccountsBalanceUseCase
}

type BalanceConsumerMsg struct {
	Payload save_account.UpdateAccountsBalanceInputDTO `json:"Payload"`
}

func NewBalanceConsumer(consumer *kafka.Consumer, updateAccountsBalanceUseCase save_account.UpdateAccountsBalanceUseCase) *balanceConsumer {
	return &balanceConsumer{
		Consumer:                     consumer,
		updateAccountsBalanceUseCase: updateAccountsBalanceUseCase,
	}
}

func (c *balanceConsumer) InitConsumer() {
	msgChan := make(chan *ckafka.Message)
	go c.Consumer.Consume(msgChan)

	for msg := range msgChan {
		var balanceConsumerMsg BalanceConsumerMsg

		json.Unmarshal(msg.Value, &balanceConsumerMsg)

		c.updateAccountsBalanceUseCase.Execute(
			context.Background(),
			balanceConsumerMsg.Payload,
		)
	}
}
