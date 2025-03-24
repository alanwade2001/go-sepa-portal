package service

import (
	"context"
	"log/slog"

	"github.com/alanwade2001/go-sepa-portal/internal/data"
	"github.com/looplab/fsm"
)

type InitiationSM struct {
	fsm *fsm.FSM
}

type IInitiationSM interface {
	FireEvent(evt data.InitiationEvent) (data.InitiationState, error)
}

func NewInitiationSM(initialState data.InitiationState) IInitiationSM {
	s := &InitiationSM{}

	s.fsm = fsm.NewFSM(
		string(initialState),
		fsm.Events{
			{Name: string(data.AcceptEvent), Src: []string{string(data.InitiatedState)}, Dst: string(data.AcceptedState)},
			{Name: string(data.RejectEvent), Src: []string{string(data.InitiatedState)}, Dst: string(data.RejectedState)},
			{Name: string(data.ApproveEvent), Src: []string{string(data.AcceptedState)}, Dst: string(data.ApprovedState)},
			{Name: string(data.CancelEvent), Src: []string{string(data.InitiatedState), string(data.AcceptedState)}, Dst: string(data.CancelledState)},
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

func (s *InitiationSM) FireEvent(evt data.InitiationEvent) (data.InitiationState, error) {
	if err := s.fsm.Event(context.Background(), string(evt)); err != nil {
		return "", err
	} else {
		state := s.fsm.Current()
		iState := data.InitiationState(state)
		return iState, nil
	}

}
