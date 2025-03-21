package service

import (
	"encoding/xml"
	"log/slog"

	"strings"

	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"

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

	slog.Debug("store doc")
	if storedDoc, err = d.store.StoreDocument(content); err != nil {
		return nil, err
	} else {
		xmlDocument := &pain_001_001_03.Document{}
		if err := xml.NewDecoder(strings.NewReader(content)).Decode(xmlDocument); err != nil {
			slog.Error("xml decode content", "error", err)
			return nil, err
		} else if result, err = d.control.Check(xmlDocument); err != nil {
			slog.Error("control check", "error", err)
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
			slog.Error("model mapping", "error", err)
			return nil, err
		} else if newInitn, err := newInitiation.ToEntity(); err != nil {
			slog.Error("entity mapping", "error", err)
			return nil, err
		} else if persisted, err := d.initiationRepos.Perist(newInitn); err != nil {
			slog.Error("initiation persisting", "error", err)
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
