package service

import (
	"encoding/json"
	"log"

	"github.com/alanwade2001/go-sepa-portal/internal/model"
	q "github.com/alanwade2001/go-sepa-q"
)

var DEST_PORTAL_INITIATION_ACCEPTED string = "topic:portal.initiation.accepted"
var DEST_PORTAL_INITIATION_REJECTED string = "topic:portal.initiation.rejected"
var DEST_PORTAL_INITIATION_APPROVED string = "topic:portal.initiation.approved"
var DEST_PORTAL_INITIATION_CANCELLED string = "topic:portal.initiation.cancelled"

type Message struct {
	sender q.MessageSender
}

func NewMessage(sender q.MessageSender) *Message {
	message := &Message{
		sender: sender,
	}

	return message
}

func (s *Message) SendAccepted(initn *model.Initiation) error {
	return s.sendInitiation(DEST_PORTAL_INITIATION_ACCEPTED, initn)
}

func (s *Message) SendRejected(initn *model.Initiation) error {
	return s.sendInitiation(DEST_PORTAL_INITIATION_REJECTED, initn)
}

func (s *Message) SendApproved(initn *model.Initiation) error {
	return s.sendInitiation(DEST_PORTAL_INITIATION_APPROVED, initn)
}

func (s *Message) SendCancelled(initn *model.Initiation) error {
	return s.sendInitiation(DEST_PORTAL_INITIATION_CANCELLED, initn)
}

func (s *Message) sendInitiation(dest string, initn *model.Initiation) error {
	bytes, err := json.MarshalIndent(initn, "", "  ")

	if err != nil {
		return err
	}

	log.Printf("sending initiation dest:[%s] id:[%d]", dest, initn.ID)

	return s.sender.SendMessage(dest, bytes)
}
