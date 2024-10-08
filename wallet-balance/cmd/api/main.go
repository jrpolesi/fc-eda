package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/consumer"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/database"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/find_account"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/save_account"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web/webserver"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/kafka"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/uow"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "wallet-balance-mysql", "3306", "balance"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to MySQL")

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	findAccountUseCase := find_account.NewFindAccountUseCase(accountDb)

	webserver := webserver.NewWebServer(":3003")

	accountHandler := web.NewWebBalanceHandler(*findAccountUseCase)

	webserver.AddHandler("/balances/{account_id}", "GET", accountHandler.FindBalance)

	kafkaConsumer := kafka.NewConsumer(&configMap, []string{"balances"})

	updateAccountsBalanceUseCase := save_account.NewUpdateAccountsBalanceUseCase(uow)

	balanceConsumer := consumer.NewBalanceConsumer(kafkaConsumer, *updateAccountsBalanceUseCase)

	go balanceConsumer.InitConsumer()

	fmt.Println("Server is running")
	webserver.Start()
}
