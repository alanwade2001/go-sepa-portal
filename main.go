package main

import (
	inf "github.com/alanwade2001/go-sepa-infra"
	"github.com/alanwade2001/go-sepa-portal/internal/handler"
	"github.com/alanwade2001/go-sepa-portal/internal/repository"
	"github.com/alanwade2001/go-sepa-portal/internal/service"
)

type App struct {
	Infra             *inf.Infra
	Repository        repository.IInitiation
	ControlService    service.IControl
	StoreService      service.IStore
	DocumentService   service.IDocument
	DocumentHandler   handler.IDocument
	InitiationService service.IInitiation
	InitiationHandler handler.IInitiation
}

func NewApp() *App {
	infra := inf.NewInfra()
	repos := repository.NewInitiation(infra.Persist)
	message := service.NewMessage(infra.Stomp)
	initnSvc := service.NewInitiation(repos, message)

	store := service.NewStore()
	control := service.NewControl()

	pain001Decoder := &service.Pain001Decoder{}

	docSvc := service.NewDocument(repos, message, control, store, pain001Decoder)
	docHdlr := handler.NewDocument(docSvc)
	docHdlr.Register(infra.Router)
	initnHdlr := handler.NewInitiation(initnSvc)
	initnHdlr.Register(infra.Router)

	app := &App{
		Infra:             infra,
		Repository:        repos,
		DocumentService:   docSvc,
		StoreService:      store,
		DocumentHandler:   docHdlr,
		InitiationService: initnSvc,
		InitiationHandler: initnHdlr,
	}

	return app
}

func (a *App) Run() {
	a.Infra.RunWithTLS()
}

func (a *App) Cleanup() {
	//a.ControlService.Cleanup()
}

func main() {
	app := NewApp()
	defer app.Cleanup()

	app.Run()

}
