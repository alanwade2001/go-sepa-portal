package data

type Initiation struct {
	ID uint `xml:"id" gorm:"primaryKey"`

	MsgId    string          `xml:"GroupHeader>MsgId"`
	CreDtTm  string          `xml:"GroupHeader>CreDtTm"`
	NbOfTxs  string          `xml:"GroupHeader>NbOfTxs"`
	CtrlSum  float64         `xml:"GroupHeader>CtrlSum"`
	State    InitiationState `xml:"Submission>State"`
	RjctdRsn string          `xml:"Submission>RejectReason,omitempty"`
	DocID    uint            `xml:"Document>id"`
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

type Document struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

type CheckResult struct {
	Pass bool
	Msg  string
	Err  error
}
