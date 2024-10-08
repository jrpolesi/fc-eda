package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	account := NewAccount()

	assert.NotNil(t, account)
	assert.Equal(t, account.Balance, float64(0))

	_, err := uuid.Parse(account.ID)
	assert.Nil(t, err)
}

func TestUpdateBalance(t *testing.T) {
	account := NewAccount()
	account.UpdateBalance(100)
	assert.Equal(t, float64(100), account.Balance)
}
