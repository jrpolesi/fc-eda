package database

import (
	"database/sql"
	"testing"
	"time"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table accounts (id varchar(255), balance float, created_at date, updated_at date)")

	s.accountDB = NewAccountDB(db)
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount()
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	accountID := "1"

	s.db.Exec(
		"Insert into accounts (id, balance, created_at, updated_at) values (?, ?, ?, ?)",
		accountID, 100.0, time.Now(), time.Now(),
	)

	account, err := s.accountDB.FindByID(accountID)

	s.Nil(err)
	s.Equal(accountID, account.ID)
	s.Equal(100.0, account.Balance)
	s.NotNil(account.CreatedAt)
	s.NotNil(account.UpdatedAt)
}

func (s *AccountDBTestSuite) TestUpdate() {
	account := &entity.Account{
		ID:        "1",
		Balance:   100.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.db.Exec(
		"Insert into accounts (id, balance, created_at, updated_at) values (?, ?, ?, ?)",
		account.ID, account.Balance, account.CreatedAt, account.UpdatedAt,
	)

	account.Balance = 200.0
	err := s.accountDB.Update(account)
	s.Nil(err)

	var accountDb entity.Account
	row := s.db.QueryRow("Select balance from accounts where id = ?", account.ID)

	err = row.Scan(&accountDb.Balance)
	s.Nil(err)
	s.Equal(200.0, accountDb.Balance)
}
