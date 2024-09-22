package router

import (
	"net/http"

	"github.com/mrdhira/project-taurus/port/httpExt/controller"
	"github.com/mrdhira/project-taurus/port/httpExt/middleware"
)

func V1Group(
	middleware middleware.IMiddleware,
	accountCtrl controller.IV1AccountController,
	accountTransactionCtrl controller.IV1AccountTransactionController,
	userCtrl controller.IV1UserController,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(
		"/account",
		middleware.WithValidateAccessToken(
			V1AccountGroup(accountCtrl, accountTransactionCtrl),
		),
	)
	mux.Handle("/user", V1UserGroup(userCtrl))

	return mux
}

func V1AccountGroup(
	accountCtrl controller.IV1AccountController,
	accountTransactionCtrl controller.IV1AccountTransactionController,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", accountCtrl.CreateAccount)
	mux.HandleFunc("GET /", accountCtrl.GetAccountsByUser)
	mux.HandleFunc("GET /{account_id}", accountCtrl.GetAccountByID)
	mux.HandleFunc("PATCH /{account_id}", accountCtrl.UpdateAccount)
	mux.HandleFunc("DELETE /{account_id}", accountCtrl.DeleteAccount)

	mux.Handle("/{account_id}/transaction", V1AccountTransactionGroup(accountTransactionCtrl))

	return mux
}

func V1AccountTransactionGroup(
	accountTransactionCtrl controller.IV1AccountTransactionController,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", accountTransactionCtrl.CreateTransaction)
	mux.HandleFunc("GET /", accountTransactionCtrl.GetTransactionsByAccount)
	mux.HandleFunc("PATCH /{transaction_id}", accountTransactionCtrl.UpdateTransaction)
	mux.HandleFunc("DELETE /{transaction_id}", accountTransactionCtrl.DeleteTransaction)

	return mux
}

func V1UserGroup(
	userCtrl controller.IV1UserController,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", userCtrl.Register)
	mux.HandleFunc("POST /login", userCtrl.Login)

	return mux
}
