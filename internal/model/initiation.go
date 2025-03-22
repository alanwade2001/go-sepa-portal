package model

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	"github.com/alanwade2001/go-sepa-portal/internal/repository/entity"
)

// initiations represents data about a record initiations.
type Initiation struct {
	ID              uint            `json:"id"`
	MsgID           string          `json:"msgId"`
	CtrlSum         float64         `json:"ctrlSum"`
	NbOfTxs         uint            `json:"nbOfTxs"`
	CreDtTm         string          `json:"creDtTm"`
	State           InitiationState `json:"state"`
	DocID           uint            `json:"docId"`
	RejectionReason string          `json:"rejectionReason"`
}

func NewInitiation(gh *pain_001_001_03.GroupHeader32, state InitiationState, docID uint, reason string) (*Initiation, error) {
	if nbOfTxs, err := strconv.ParseUint(gh.NbOfTxs, 10, 64); err != nil {
		return nil, err
	} else {

		initn := &Initiation{
			MsgID:           gh.MsgId,
			CtrlSum:         gh.CtrlSum,
			NbOfTxs:         uint(nbOfTxs),
			CreDtTm:         gh.CreDtTm,
			State:           state,
			DocID:           docID,
			RejectionReason: reason,
		}
		return initn, nil

	}
}

func (i *Initiation) ToEntity() (*entity.Initiation, error) {

	if creDtTm, err := time.Parse("2006-01-02T15:04:05", i.CreDtTm); err != nil {
		slog.Error("Failed to parse creDtTm", "error", err)
		return nil, err
	} else {

		newInitn := &entity.Initiation{
			CtrlSum:         i.CtrlSum,
			CreDtTm:         creDtTm,
			MsgID:           i.MsgID,
			NbOfTxs:         i.NbOfTxs,
			State:           string(i.State),
			DocID:           i.DocID,
			RejectionReason: i.RejectionReason,
		}
		return newInitn, nil
	}
}

func (i *Initiation) FromEntity(ent *entity.Initiation) error {

	i.ID = ent.Model.ID
	i.MsgID = ent.MsgID
	i.CtrlSum = ent.CtrlSum
	i.NbOfTxs = ent.NbOfTxs
	i.State = InitiationState(ent.State)
	i.DocID = ent.DocID
	i.CreDtTm = ent.CreDtTm.Format("2006-03-15T12:12:12")
	i.RejectionReason = ent.RejectionReason

	return nil
}

func FromEntities(ents []*entity.Initiation) ([]*Initiation, error) {
	models := make([]*Initiation, len(ents))

	for _, ent := range ents {
		mdl := &Initiation{}

		if err := mdl.FromEntity(ent); err != nil {
			return nil, err
		} else {
			models = append(models, mdl)
		}
	}

	return models, nil
}

type InitiationEvent string
type InitiationState string

var AcceptEvent = InitiationEvent("Accept")
var RejectEvent = InitiationEvent("Reject")
var ApproveEvent = InitiationEvent("Approve")
var CancelEvent = InitiationEvent("Cancel")

var InitiatedState = InitiationState("Initiated")
var AcceptedState = InitiationState("Accepted")
var RejectedState = InitiationState("Rejected")
var CancelledState = InitiationState("Cancelled")
var ApprovedState = InitiationState("Approved")
