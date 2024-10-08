package save_account

import (
	"context"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveAccountUseCase_Execute(t *testing.T) {
	mockUow := &mocks.UowMock{}

	useCase := NewUpdateAccountsBalanceUseCase(mockUow)

	accountDbFrom := &entity.Account{
		ID: "123",
	}

	accountDbTo := &entity.Account{
		ID: "456",
	}

	input := UpdateAccountsBalanceInputDTO{
		AccountIDFrom:      accountDbFrom.ID,
		AccountIDTo:        accountDbTo.ID,
		BalanceAccountFrom: 100,
		BalanceAccountTo:   200,
	}

	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	err := useCase.Execute(context.TODO(), input)
	assert.Nil(t, err)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
