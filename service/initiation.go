package service

import (
	"github.com/alanwade2001/go-sepa-portal/data"
	"github.com/alanwade2001/go-sepa-portal/repository"
)

type Initiation struct {
	repository repository.IInitiation
	message    IMessage
}

type IInitiation interface {
	FindAll() ([]*data.Initiation, error)
	FindByID(id string) (*data.Initiation, error)
	SendInitiationAccept(id string) (*data.Initiation, error)
	SendInitiationReject(id string) (*data.Initiation, error)
	SendInitiationCancel(id string) (*data.Initiation, error)
	SendInitiationApprove(id string) (*data.Initiation, error)
}

func NewInitiation(repository repository.IInitiation, message IMessage) IInitiation {
	initiation := &Initiation{
		repository: repository,
		message:    message,
	}

	return initiation
}

func (i *Initiation) FindAll() ([]*data.Initiation, error) {
	return i.repository.FindAll()
}

func (i *Initiation) FindByID(id string) (*data.Initiation, error) {
	return i.repository.FindByID(id)
}

func (i *Initiation) SendInitiationAccept(id string) (*data.Initiation, error) {
	return i.SendInitiationEvent(id, data.AcceptEvent)
}

func (i *Initiation) SendInitiationReject(id string) (*data.Initiation, error) {
	return i.SendInitiationEvent(id, data.RejectEvent)
}

func (i *Initiation) SendInitiationCancel(id string) (*data.Initiation, error) {
	return i.SendInitiationEvent(id, data.CancelEvent)
}

func (i *Initiation) SendInitiationApprove(id string) (*data.Initiation, error) {
	return i.SendInitiationEvent(id, data.ApproveEvent)
}

func (i *Initiation) SendInitiationEvent(id string, evt data.InitiationEvent) (*data.Initiation, error) {
	if initn, err := i.repository.FindByID(id); err != nil {
		return nil, err
	} else {
		currentState := data.InitiationState(initn.State)
		sm := NewInitiationSM(currentState)

		if initn.State, err = sm.FireEvent(evt); err != nil {
			return nil, err
		} else if updated, err := i.repository.Persist(initn); err != nil {
			return nil, err
		} else if err := i.message.Send(updated); err != nil {
			return nil, err
		} else {
			return updated, nil
		}
	}
}
