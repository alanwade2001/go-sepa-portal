package service

import (
	"context"
	"log"

	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/alanwade2001/go-sepa-portal/internal/repository"
)

type Initiation struct {
	repository *repository.Initiation
	message    *Message
}

func NewInitiation(repository *repository.Initiation, message *Message) *Initiation {
	initiation := &Initiation{
		repository: repository,
		message:    message,
	}

	return initiation
}

func (i *Initiation) FindAll() ([]*model.Initiation, error) {
	initns, err := i.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return model.FromEntities(initns)
}

func (i *Initiation) FindByID(id string) (*model.Initiation, error) {
	initn, err := i.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	initiation := &model.Initiation{}
	if err := initiation.FromEntity(initn); err != nil {
		return nil, err
	}

	return initiation, err
}

func (i *Initiation) SendInitiationAccept(id string) (*model.Initiation, error) {
	initiation, err := i.SendInitiationEvent(id, model.AcceptEvent)

	if err == nil {
		i.message.SendAccepted(initiation)
	}

	return initiation, err
}
func (i *Initiation) SendInitiationReject(id string) (*model.Initiation, error) {
	initiation, err := i.SendInitiationEvent(id, model.RejectEvent)

	if err == nil {
		i.message.SendRejected(initiation)
	}

	return initiation, err
}
func (i *Initiation) SendInitiationCancel(id string) (*model.Initiation, error) {
	initiation, err := i.SendInitiationEvent(id, model.CancelEvent)

	if err == nil {
		i.message.SendCancelled(initiation)
	}

	return initiation, err
}
func (i *Initiation) SendInitiationApprove(id string) (*model.Initiation, error) {
	if initiation, err := i.SendInitiationEvent(id, model.ApproveEvent); err != nil {
		log.Printf("Not sending approved to queue: [%v]", err)
		return nil, err
	} else {
		log.Printf("sending approved to queue")
		i.message.SendApproved(initiation)
		return initiation, nil
	}

}

func (i *Initiation) SendInitiationEvent(id string, evt model.InitiationEvent) (*model.Initiation, error) {
	initn, err := i.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	currentState := model.InitiationState(initn.State)

	sm := NewInitiationSM(currentState)
	err = sm.FSM.Event(context.Background(), string(evt))
	if err != nil {
		return nil, err
	}

	initn.State = sm.FSM.Current()

	updated, err := i.repository.Perist(initn)

	if err != nil {
		return nil, err
	}

	initiation := &model.Initiation{}
	if err := initiation.FromEntity(updated); err != nil {
		return nil, err
	}

	return initiation, nil

}
