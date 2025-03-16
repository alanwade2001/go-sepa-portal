package service

import (
	"encoding/xml"
	"log"

	"strings"

	s "github.com/alanwade2001/go-sepa-iso/schema"

	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/alanwade2001/go-sepa-portal/internal/repository"
	"github.com/alanwade2001/go-sepa-portal/internal/repository/entity"
)

type Document struct {
	initiationRepos *repository.Initiation
	message         *Message
	control         *Control
	store           *Store
}

func NewDocument(initiationRepos *repository.Initiation, message *Message, control *Control, store *Store) *Document {
	document := &Document{
		initiationRepos: initiationRepos,
		message:         message,
		control:         control,
		store:           store,
	}

	return document
}

func (d *Document) InitiateDocument(content string) (newInitiation *model.Initiation, err error) {
	newInitiation = &model.Initiation{}
	var storedDoc *model.Document

	log.Println("store doc")
	if storedDoc, err = d.store.StoreDocument(content); err != nil {
		return nil, err
	}

	xmlDocument := &s.Pain001Document{}
	xml.NewDecoder(strings.NewReader(content)).Decode(xmlDocument)

	accepted := d.control.Check(xmlDocument)
	state := model.AcceptedState

	if !accepted {
		state = model.RejectedState
	}

	gh := xmlDocument.CstmrCdtTrfInitn.GrpHdr
	newInitiation.CtrlSum = gh.CtrlSum
	newInitiation.MsgID = gh.MsgId
	newInitiation.NbOfTxs = gh.NbOfTxs
	newInitiation.State = state
	newInitiation.DocID = storedDoc.ID

	newInitn := &entity.Initiation{
		CtrlSum: newInitiation.CtrlSum,
		MsgID:   newInitiation.MsgID,
		NbOfTxs: newInitiation.NbOfTxs,
		State:   string(newInitiation.State),
		DocID:   newInitiation.DocID,
	}

	persisted, err := d.initiationRepos.Perist(newInitn)

	if err != nil {
		return nil, err
	}

	newInitiation.ID = persisted.Model.ID

	if accepted {
		err = d.message.SendAccepted(newInitiation)
	} else {
		err = d.message.SendRejected(newInitiation)
	}

	// if err != nil {
	// 	log.Printf("error sending to queue: %t, %v", accepted, err)
	// }

	return newInitiation, err
}
