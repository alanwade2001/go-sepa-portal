package model

// initiations represents data about a record initiations.
type Initiation struct {
	ID      uint            `json:"id"`
	MsgID   string          `json:"msgId"`
	CtrlSum float64         `json:"ctrlSum"`
	NbOfTxs string          `json:"nbOfTxs"`
	State   InitiationState `json:"state"`
	DocID   uint            `json:"docId"`
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
