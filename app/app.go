package app

import (
	"github.com/ednailson/httping-go"
	"github.com/ednailson/serasa-challenge/controller"
	"github.com/ednailson/serasa-challenge/controller/crypto"
	"github.com/ednailson/serasa-challenge/database"
	"net/http"
)

type App struct {
	server    httping.IServer
	closeFunc httping.ServerCloseFunc
}

func LoadApp(cfg Config) (*App, error) {
	db, err := database.NewDatabase(cfg.Database)
	if err != nil {
		return nil, err
	}
	cryptoModule, err := crypto.NewCrypto(cfg.Key, cfg.Nonce)
	if err != nil {
		return nil, err
	}
	ctrl := controller.NewController(cfg.MainframeUrl, db, cryptoModule)
	server := httping.NewHttpServer("", cfg.Port)
	loadRoutes(ctrl, server)
	return &App{
		server: server,
	}, nil
}

func (a *App) Run() <-chan error {
	closeFunc, chErr := a.server.RunServer()
	a.closeFunc = closeFunc
	return chErr
}

func (a *App) Close() {
	if a.closeFunc != nil {
		a.closeFunc()
	}
}

func loadRoutes(ctrl *controller.Controller, server httping.IServer) {
	v1Route := server.NewRoute(nil, "/v1")
	server.NewRoute(v1Route, "/update").POST(func(request httping.HttpRequest) httping.IResponse {
		err := ctrl.UpdateData()
		if err != nil {
			return NewResponse(http.StatusInternalServerError, map[string]string{"server": "server internal error"})
		}
		return NewResponse(http.StatusNoContent, nil)
	})
	server.NewRoute(v1Route, "/negativation").GET(func(request httping.HttpRequest) httping.IResponse {
		document, ok := request.Query["cpf"]
		if !ok {
			return NewResponse(http.StatusBadRequest, map[string]string{"cpf": "cpf is a required parameter"})
		}
		negativations, err := ctrl.NegativationByDocument(document[0])
		if err != nil {
			return NewResponse(http.StatusInternalServerError, map[string]string{"server": "server internal error"})
		}
		return NewResponse(http.StatusOK, negativations)
	})
}
