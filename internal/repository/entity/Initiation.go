package entity

import "gorm.io/gorm"

type Initiation struct {
	Model   gorm.Model `gorm:"embedded"`
	MsgID   string
	NbOfTxs string
	CtrlSum float64
	State   string
	DocID   uint
}
