package client

import (
	req "banking/dto"
	"banking/model"
	"context"
	crand "crypto/rand"
	"fmt"
	"math/big"
	"strconv"

	"github.com/opentracing/opentracing-go"
)

type BankingClient struct {
	// userData model.DataBankingInterface
}

func NewBanking() model.BusinessBankingInterface {
	return &BankingClient{}
}

func (c *BankingClient) RegisterUser(ctx context.Context, reqUser req.UserRequest) (resUser model.Saldo, err error) {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "RegisterUser-Client")
	defer trace.Finish()
	conn, err := conn()
	if err != nil {
		trace.LogKV("Error Conn", err)
		return model.Saldo{}, fmt.Errorf(err.Error())
	}

	defer conn.Close()
	fmt.Println("req-client", reqUser.Nama)

	Norek := GenerateNorek()
	NorekStr := strconv.Itoa(Norek)

	trace.LogKV("Request-RegisterUser-Client", reqUser)
	stmt, errQuery := conn.ExecContext(context.Background(), "INSERT INTO public.user (nama, nik, no_hp, no_rekening) VALUES ($1, $2, $3, $4)", reqUser.Nama, reqUser.NIK, reqUser.NoHp, NorekStr)
	if errQuery != nil {
		trace.LogKV("Error Db Insert", errQuery)
		return model.Saldo{}, fmt.Errorf(errQuery.Error())
	}
	rowsA, errRowsA := stmt.RowsAffected()
	if errRowsA != nil {
		trace.LogKV("Error Rows A", err)
		return model.Saldo{}, fmt.Errorf(errRowsA.Error())
	}
	var noRek string
	if rowsA != 0 {
		stmtNorek, errStmt := conn.ExecContext(context.Background(), "INSERT INTO public.saldo (no_rekening, saldo) VALUES ($1, $2)", NorekStr, 0)
		if errStmt != nil {
			trace.LogKV("Error Db", errStmt)
			return model.Saldo{}, fmt.Errorf(errQuery.Error())
		}
		rowsB, errRows := stmtNorek.RowsAffected()
		if errRows != nil {
			trace.LogKV("Error Rows", errRows)
			return model.Saldo{}, fmt.Errorf(errRows.Error())
		}
		if rowsB != 0 {
			noRek = NorekStr
		}
	}
	trace.LogKV("Response-RegisterUser-Client", noRek)
	return model.Saldo{NoRekening: noRek}, nil
}

func (c *BankingClient) SetorSaldo(ctx context.Context, reqSetor req.SetorSaldo) (resSetor model.Saldo, err error) {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "SetorSaldo-Client")
	defer trace.Finish()
	conn, err := conn()
	if err != nil {
		trace.LogKV("Error Conn", err)
		return model.Saldo{}, fmt.Errorf(err.Error())
	}

	defer conn.Close()
	// fmt.Println("req-client", reqUser.Nama)

	trace.LogKV("Request-SetorSaldo-Client", reqSetor)
	var saldoBaru int64

	errQuery := conn.QueryRowContext(context.Background(), "UPDATE public.saldo SET saldo = saldo + $1 WHERE no_rekening = $2 RETURNING saldo", reqSetor.Nominal, reqSetor.NoRekening).Scan(&saldoBaru)
	if errQuery != nil {
		trace.LogKV("Error Db", errQuery)
		return model.Saldo{}, fmt.Errorf(errQuery.Error())
	}
	Result := model.Saldo{Saldo: int(saldoBaru), NoRekening: reqSetor.NoRekening}
	trace.LogKV("Response-SetorSaldo-Client", Result)
	return Result, nil
}

func (c *BankingClient) TarikSaldo(ctx context.Context, reqTarik req.TarikSaldo) (resTarik model.Saldo, err error) {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "TarikSaldo-Client")
	defer trace.Finish()
	conn, err := conn()
	if err != nil {
		trace.LogKV("Error Conn", err)
		return model.Saldo{}, fmt.Errorf(err.Error())
	}

	defer conn.Close()
	// fmt.Println("req-client", reqUser.Nama)

	trace.LogKV("Request-TarikSaldo-Client", reqTarik)
	var saldoBaru int64

	errQuery := conn.QueryRowContext(context.Background(), "UPDATE public.saldo SET saldo = saldo - $1 WHERE no_rekening = $2 RETURNING saldo", reqTarik.Nominal, reqTarik.NoRekening).Scan(&saldoBaru)
	if errQuery != nil {
		trace.LogKV("Error Db", errQuery)
		return model.Saldo{}, fmt.Errorf(errQuery.Error())
	}
	var Result = model.Saldo{Saldo: int(saldoBaru), NoRekening: reqTarik.NoRekening}
	trace.LogKV("Request-TarikSaldo-Client", Result)
	return Result, nil
}

func (c *BankingClient) CekSaldo(ctx context.Context, reqNorek string) (resTarik model.Saldo, err error) {
	trace, ctx := opentracing.StartSpanFromContext(ctx, "CekSaldo-Client")
	defer trace.Finish()
	conn, err := conn()
	if err != nil {
		trace.LogKV("Error Conn", err)
		return model.Saldo{}, fmt.Errorf(err.Error())
	}

	defer conn.Close()
	fmt.Println("Masuk req-client", reqNorek)

	trace.LogKV("Request-CekSaldo-Client", reqNorek)
	var saldoBaru int64

	errQuery := conn.QueryRowContext(context.Background(), "SELECT saldo FROM public.saldo WHERE no_rekening= $1", reqNorek).Scan(&saldoBaru)
	if errQuery != nil {
		fmt.Println("err---Cek", errQuery)
		trace.LogKV("Error Db", errQuery)
		return model.Saldo{}, fmt.Errorf(errQuery.Error())
	}
	var Result = model.Saldo{Saldo: int(saldoBaru), NoRekening: reqNorek}
	trace.LogKV("Response-CekSaldo-Client", Result)
	return Result, nil
}

func GenerateNorek() int {
	n, _ := crand.Int(crand.Reader, big.NewInt(9000000)) // Angka antara 0 - 8999999
	return int(n.Int64()) + 1000000                      // Ubah ke rentang 1000000 - 9999999
}
