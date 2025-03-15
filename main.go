package main

import (
	inf "github.com/alanwade2001/go-sepa-infra"
	"github.com/alanwade2001/go-sepa-portal/internal/handler"
	"github.com/alanwade2001/go-sepa-portal/internal/repository"
	"github.com/alanwade2001/go-sepa-portal/internal/service"
)

type App struct {
	Infra             *inf.Infra
	Repository        *repository.Initiation
	ControlService    *service.Control
	DocumentService   *service.Document
	DocumentHandler   *handler.Document
	InitiationService *service.Initiation
	InitiationHandler *handler.Initiation
}

func NewApp() *App {
	infra := inf.NewInfra()
	repos := repository.NewInitiation(infra.Persist)
	message := service.NewMessage(infra.Stomp)
	initnSvc := service.NewInitiation(repos, message)

	store := service.NewStore()
	control := service.NewControl()

	docSvc := service.NewDocument(repos, message, control, store)
	docHdlr := handler.NewDocument(docSvc, infra.Router)
	initnHdlr := handler.NewInitiation(initnSvc, infra.Router)

	app := &App{
		Infra:             infra,
		Repository:        repos,
		DocumentService:   docSvc,
		DocumentHandler:   docHdlr,
		InitiationService: initnSvc,
		InitiationHandler: initnHdlr,
	}

	return app
}

func (a *App) Run() {
	a.Infra.RunWithTLS()
}

func main() {
	app := NewApp()
	app.Run()
}
