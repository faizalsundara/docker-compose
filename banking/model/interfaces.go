package model

import (
	req "banking/dto"
	"context"
)

type BusinessBankingInterface interface {
	RegisterUser(ctx context.Context, reqUser req.UserRequest) (resUser Saldo, err error)
	SetorSaldo(ctx context.Context, reqSetor req.SetorSaldo) (resSetor Saldo, err error)
	TarikSaldo(ctx context.Context, reqTarik req.TarikSaldo) (resTarik Saldo, err error)
	CekSaldo(ctx context.Context, reqNorek string) (resCekSaldo Saldo, err error)
}
