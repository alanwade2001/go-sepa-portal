package service

import (
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/alanwade2001/go-sepa-portal/internal/data"
	q "github.com/alanwade2001/go-sepa-q"
)

var DEST_PORTAL_INITIATION_ACCEPTED string = "topic:portal.initiation.accepted"
var DEST_PORTAL_INITIATION_REJECTED string = "topic:portal.initiation.rejected"
var DEST_PORTAL_INITIATION_APPROVED string = "topic:portal.initiation.approved"
var DEST_PORTAL_INITIATION_CANCELLED string = "topic:portal.initiation.cancelled"

type MessageSender interface {
	Send(v interface{}) error
}

type Message struct {
	sender q.MessageSender
}

type IMessage interface {
	Send(initn *data.Initiation) error
}

func NewMessage(sender q.MessageSender) IMessage {
	message := &Message{
		sender: sender,
	}

	return message
}

func (s *Message) Send(initn *data.Initiation) error {
	switch initn.State {
	case data.AcceptedState:
		return s.SendAccepted(initn)
	case data.ApprovedState:
		return s.SendApproved(initn)
	case data.RejectedState:
		return s.SendRejected(initn)
	case data.CancelledState:
		return s.SendCancelled(initn)
	default:
		return errors.New("unknown initiation state")
	}
}

func (s *Message) SendAccepted(initn *data.Initiation) error {
	return s.sendInitiation(DEST_PORTAL_INITIATION_ACCEPTED, initn)
}

func (s *Message) SendRejected(initn *data.Initiation) error {
	return s.sendInitiation(DEST_PORTAL_INITIATION_REJECTED, initn)
}

func (s *Message) SendApproved(initn *data.Initiation) error {
	return s.sendInitiation(DEST_PORTAL_INITIATION_APPROVED, initn)
}

func (s *Message) SendCancelled(initn *data.Initiation) error {
	return s.sendInitiation(DEST_PORTAL_INITIATION_CANCELLED, initn)
}

func (s *Message) sendInitiation(dest string, initn *data.Initiation) error {
	bytes, err := json.MarshalIndent(initn, "", "  ")

	if err != nil {
		return err
	}

	slog.Info("sending initiation", "dest", dest, "initn.ID", initn.ID)

	return s.sender.SendMessage(dest, bytes)
}
