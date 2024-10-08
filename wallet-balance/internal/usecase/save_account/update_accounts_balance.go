package save_account

import (
	"context"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/uow"
)

type UpdateAccountsBalanceInputDTO struct {
	AccountIDFrom      string  `json:"account_id_from"`
	AccountIDTo        string  `json:"account_id_to"`
	BalanceAccountFrom float64 `json:"balance_account_id_from"`
	BalanceAccountTo   float64 `json:"balance_account_id_to"`
}

type UpdateAccountsBalanceUseCase struct {
	Uow uow.UowInterface
}

func NewUpdateAccountsBalanceUseCase(Uow uow.UowInterface) *UpdateAccountsBalanceUseCase {
	return &UpdateAccountsBalanceUseCase{
		Uow: Uow,
	}
}

func (uc *UpdateAccountsBalanceUseCase) Execute(ctx context.Context, input UpdateAccountsBalanceInputDTO) error {
	accountFrom := entity.Account{
		ID:      input.AccountIDFrom,
		Balance: input.BalanceAccountFrom,
	}

	accountTo := entity.Account{
		ID:      input.AccountIDTo,
		Balance: input.BalanceAccountTo,
	}

	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {
		err := uc.updateAccountBalance(ctx, accountFrom)
		if err != nil {
			return err
		}

		err = uc.updateAccountBalance(ctx, accountTo)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (uc *UpdateAccountsBalanceUseCase) updateAccountBalance(ctx context.Context, input entity.Account) error {
	accountExists := true
	accountRepository := uc.getAccountRepository(ctx)

	account, err := accountRepository.FindByID(input.ID)
	if err != nil {
		accountExists = false
		account = entity.NewAccount()
		account.ID = input.ID
	}

	account.UpdateBalance(input.Balance)

	if !accountExists {
		return accountRepository.Save(account)
	}

	return accountRepository.Update(account)
}

func (uc *UpdateAccountsBalanceUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}
