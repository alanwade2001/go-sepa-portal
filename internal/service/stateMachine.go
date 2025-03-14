package service

import (
	"context"
	"fmt"

	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/looplab/fsm"
)

type InitiationSM struct {
	FSM *fsm.FSM
}

func NewInitiationSM(initialState model.InitiationState) *InitiationSM {
	s := &InitiationSM{}

	s.FSM = fsm.NewFSM(
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
	fmt.Printf("The submission is %s\n", e.Dst)
}
