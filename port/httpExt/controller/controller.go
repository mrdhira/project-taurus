package controller

import "net/http"

type IHealthCheckController interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type IV1AccountController interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccountByID(w http.ResponseWriter, r *http.Request)
	GetAccountsByUser(w http.ResponseWriter, r *http.Request)
	UpdateAccount(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)
}

type IV1AccountTransactionController interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	GetTransactionsByAccount(w http.ResponseWriter, r *http.Request)
	UpdateTransaction(w http.ResponseWriter, r *http.Request)
	DeleteTransaction(w http.ResponseWriter, r *http.Request)
}

type IV1UserController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}
