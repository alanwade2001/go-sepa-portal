package entity

import (
	"time"

	"gorm.io/gorm"
)

type Initiation struct {
	Model           gorm.Model `gorm:"embedded"`
	MsgID           string
	NbOfTxs         uint
	CreDtTm         *time.Time
	CtrlSum         float64
	State           string
	DocID           uint
	RejectionReason string
}
