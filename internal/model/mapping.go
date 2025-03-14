package model

import "github.com/alanwade2001/go-sepa-portal/internal/repository/entity"

func ToInitiation(initn *entity.Initiation) *Initiation {
	initiation := Initiation{
		ID:      initn.Model.ID,
		MsgID:   initn.MsgID,
		CtrlSum: initn.CtrlSum,
		NbOfTxs: initn.NbOfTxs,
		State:   InitiationState(initn.State),
		DocID:   initn.DocID,
	}

	return &initiation
}

func ToInitiations(initns []*entity.Initiation) []*Initiation {
	initiations := []*Initiation{}

	for _, v := range initns {
		initiation := ToInitiation(v)
		initiations = append(initiations, initiation)
	}

	return initiations
}
