package service

import (
	"context"
	"log/slog"

	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/looplab/fsm"
)

type InitiationSM struct {
	fsm *fsm.FSM
}

type IInitiationSM interface {
	FireEvent(evt model.InitiationEvent) (string, error)
}

func NewInitiationSM(initialState model.InitiationState) IInitiationSM {
	s := &InitiationSM{}

	s.fsm = fsm.NewFSM(
		string(initialState),
		fsm.Events{
			{Name: string(model.AcceptEvent), Src: []string{string(model.InitiatedState)}, Dst: string(model.AcceptedState)},
			{Name: string(model.RejectEvent), Src: []string{string(model.InitiatedState)}, Dst: string(model.RejectedState)},
			{Name: string(model.ApproveEvent), Src: []string{string(model.AcceptedState)}, Dst: string(model.ApprovedState)},
			{Name: string(model.CancelEvent), Src: []string{string(model.InitiatedState), string(model.AcceptedState)}, Dst: string(model.CancelledState)},
		},
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { s.enterState(e) },
		},
	)

	return s
}

func (s *InitiationSM) enterState(e *fsm.Event) {
	slog.Info("State Machine", "state", e.Dst)
}

func (s *InitiationSM) FireEvent(evt model.InitiationEvent) (string, error) {
	if err := s.fsm.Event(context.Background(), string(evt)); err != nil {
		return "", err
	} else {
		state := s.fsm.Current()
		return state, nil
	}

}
