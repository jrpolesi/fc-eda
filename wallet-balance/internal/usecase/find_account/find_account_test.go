package find_account

import (
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFindAccountUseCase_Execute(t *testing.T) {
	accountGatewayMock := &mocks.AccountGatewayMock{}
	useCase := NewFindAccountUseCase(accountGatewayMock)

	input := FindAccountInputDTO{
		ID: "123",
	}

	accountDb := &entity.Account{
		ID:      "123",
		Balance: 100,
	}

	accountGatewayMock.
		On("FindByID", input.ID).
		Return(accountDb, nil).
		Times(1)

	output, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, accountDb.ID, output.ID)
	assert.Equal(t, accountDb.Balance, output.Balance)

	accountGatewayMock.AssertExpectations(t)
}
