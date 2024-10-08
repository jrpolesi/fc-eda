package find_account

import "github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"

type FindAccountInputDTO struct {
	ID string `json:"id"`
}

type FindAccountOutputDTO struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type FindAccountUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewFindAccountUseCase(a gateway.AccountGateway) *FindAccountUseCase {
	return &FindAccountUseCase{
		AccountGateway: a,
	}
}

func (uc *FindAccountUseCase) Execute(input FindAccountInputDTO) (*FindAccountOutputDTO, error) {
	account, err := uc.AccountGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	output := &FindAccountOutputDTO{
		ID:      account.ID,
		Balance: account.Balance,
	}
	return output, nil
}
