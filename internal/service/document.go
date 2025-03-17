package service

import (
	"encoding/xml"
	"log"

	"strings"

	s "github.com/alanwade2001/go-sepa-iso/schema"

	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/alanwade2001/go-sepa-portal/internal/repository"
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

	var storedDoc *model.Document
	var result *model.CheckResult

	log.Println("store doc")
	if storedDoc, err = d.store.StoreDocument(content); err != nil {
		return nil, err
	} else {
		xmlDocument := &s.Pain001Document{}
		xml.NewDecoder(strings.NewReader(content)).Decode(xmlDocument)

		if result, err = d.control.Check(xmlDocument); err != nil {
			return nil, err
		}

		var state model.InitiationState
		if result.Pass {
			state = model.AcceptedState
		} else {
			state = model.RejectedState
		}

		gh := xmlDocument.CstmrCdtTrfInitn.GrpHdr

		if newInitiation, err := model.NewInitiation(gh, state, storedDoc.ID, result.Msg); err != nil {
			return nil, err
		} else if newInitn, err := newInitiation.ToEntity(); err != nil {
			return nil, err
		} else if persisted, err := d.initiationRepos.Perist(newInitn); err != nil {
			return nil, err
		} else {
			newInitiation.ID = persisted.Model.ID

			var sender SenderFunc

			if state == model.AcceptedState {
				sender = d.message.SendAccepted
			} else {
				sender = d.message.SendRejected
			}

			if err := sender(newInitiation); err != nil {
				return nil, err
			} else {
				return newInitiation, err
			}
		}
	}
}
