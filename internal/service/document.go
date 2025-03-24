package service

import (
	"encoding/xml"
	"log/slog"

	"strings"

	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"

	"github.com/alanwade2001/go-sepa-portal/internal/data"
	"github.com/alanwade2001/go-sepa-portal/internal/repository"
)

type Document struct {
	initiationRepos repository.IInitiation
	message         IMessage
	control         IControl
	store           IStore
	decoder         IPain001Decoder
}

func NewDocument(initiationRepos repository.IInitiation, message IMessage, control IControl, store IStore, decoder IPain001Decoder) IDocument {
	document := &Document{
		initiationRepos: initiationRepos,
		message:         message,
		control:         control,
		store:           store,
		decoder:         decoder,
	}

	return document
}

type IDocument interface {
	InitiateDocument(content string) (newInitiation *data.Initiation, err error)
}

type IPain001Decoder interface {
	Decode(content string) (*pain_001_001_03.Document, error)
}

type Pain001Decoder struct {
}

func (d Pain001Decoder) Decode(content string) (*pain_001_001_03.Document, error) {
	xmlDocument := &pain_001_001_03.Document{}
	if err := xml.NewDecoder(strings.NewReader(content)).Decode(xmlDocument); err != nil {
		return nil, err
	} else {
		return xmlDocument, nil
	}
}

func (d *Document) InitiateDocument(content string) (newInitiation *data.Initiation, err error) {

	var storedDoc *data.Document
	var result *data.CheckResult

	slog.Debug("store doc")
	if storedDoc, err = d.store.StoreDocument(content); err != nil {
		return nil, err
	} else {
		if xmlDocument, err := d.decoder.Decode(content); err != nil {
			slog.Error("xml decode content", "error", err)
			return nil, err
		} else if result, err = d.control.Check(xmlDocument); err != nil {
			slog.Error("control check", "error", err)
			return nil, err
		} else {

			var state data.InitiationState
			if result.Pass {
				state = data.AcceptedState
			} else {
				state = data.RejectedState
			}

			gh := xmlDocument.CstmrCdtTrfInitn.GrpHdr
			dataInitn := &data.Initiation{
				MsgId:    gh.MsgId,
				CreDtTm:  gh.CreDtTm,
				NbOfTxs:  gh.NbOfTxs,
				CtrlSum:  gh.CtrlSum,
				State:    state,
				RjctdRsn: result.Msg,
				DocID:    storedDoc.ID,
			}

			if persisted, err := d.initiationRepos.Persist(dataInitn); err != nil {
				slog.Error("initiation persisting", "error", err)
				return nil, err
			} else {
				dataInitn.ID = persisted.ID

				if err := d.message.Send(dataInitn); err != nil {
					return nil, err
				} else {

					return dataInitn, err
				}
			}
		}
	}
}
