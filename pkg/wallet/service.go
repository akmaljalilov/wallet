package wallet

import (
	"github.com/akmaljalilov/wallet/pkg/types"
)

var dbAccounts = []*types.Account{
	{
		ID: int64(1),
	},
}

func FindAccountById(id int64) (*types.Account, error) {
	for _, account := range dbAccounts {
		if account.ID ==id {
			return account, nil
		}
	}
	return nil, types.ErrAccountNotFound{}
}
