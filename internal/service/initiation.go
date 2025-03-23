package service

import (
	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/alanwade2001/go-sepa-portal/internal/repository"
)

type Initiation struct {
	repository repository.IInitiation
	message    IMessage
}

type IInitiation interface {
	FindAll() ([]*model.Initiation, error)
	FindByID(id string) (*model.Initiation, error)
	SendInitiationAccept(id string) (*model.Initiation, error)
	SendInitiationReject(id string) (*model.Initiation, error)
	SendInitiationCancel(id string) (*model.Initiation, error)
	SendInitiationApprove(id string) (*model.Initiation, error)
}

func NewInitiation(repository repository.IInitiation, message IMessage) IInitiation {
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
	if initn, err := i.repository.FindByID(id); err != nil {
		return nil, err
	} else {
		initiation := &model.Initiation{}
		if err := initiation.FromEntity(initn); err != nil {
			return nil, err
		} else {
			return initiation, err
		}
	}
}

func (i *Initiation) SendInitiationAccept(id string) (*model.Initiation, error) {
	return i.SendInitiationEvent(id, model.AcceptEvent)
}

func (i *Initiation) SendInitiationReject(id string) (*model.Initiation, error) {
	return i.SendInitiationEvent(id, model.RejectEvent)
}

func (i *Initiation) SendInitiationCancel(id string) (*model.Initiation, error) {
	return i.SendInitiationEvent(id, model.CancelEvent)
}

func (i *Initiation) SendInitiationApprove(id string) (*model.Initiation, error) {
	return i.SendInitiationEvent(id, model.ApproveEvent)
}

func (i *Initiation) SendInitiationEvent(id string, evt model.InitiationEvent) (*model.Initiation, error) {
	if initn, err := i.repository.FindByID(id); err != nil {
		return nil, err
	} else {
		currentState := model.InitiationState(initn.State)
		sm := NewInitiationSM(currentState)

		if initn.State, err = sm.FireEvent(evt); err != nil {
			return nil, err
		} else if updated, err := i.repository.Perist(initn); err != nil {
			return nil, err
		} else {
			initiation := &model.Initiation{}
			if err := initiation.FromEntity(updated); err != nil {
				return nil, err
			} else if err := i.message.Send(initiation); err != nil {
				return nil, err
			} else {
				return initiation, nil
			}
		}

	}

}
