package service

import (
	"strconv"

	"github.com/alanwade2001/go-sepa-iso/gen"
	"github.com/alanwade2001/go-sepa-iso/schema"
)

type Control struct {
}

func NewControl() *Control {
	control := &Control{}

	return control
}

func (c *Control) Check(doc *schema.P1Document) bool {

	if !c.ControlGrpHdrCtrlSum(doc) {
		return false
	}

	if !c.ControlGrpHdrNbOfTxs(doc) {
		return false
	}

	for _, pmtInf := range doc.CstmrCdtTrfInitn.PmtInf {
		if !c.ControlPmtInfCtrlSum(pmtInf) {
			return false
		}

		if !c.ControlPmtInfNbOfTxs(pmtInf) {
			return false
		}
	}

	return true
}

func (c *Control) ControlGrpHdrCtrlSum(doc *schema.P1Document) bool {
	ghCtrlSum := doc.CstmrCdtTrfInitn.GrpHdr.CtrlSum

	pmtInves := doc.CstmrCdtTrfInitn.PmtInf

	var ctrlSum float64 = 0.0
	for _, pmtInf := range pmtInves {
		ctrlSum = ctrlSum + pmtInf.CtrlSum
	}

	return ghCtrlSum == ctrlSum
}

func (c *Control) ControlGrpHdrNbOfTxs(doc *schema.P1Document) bool {
	ghNbOfTxs, _ := strconv.Atoi(doc.CstmrCdtTrfInitn.GrpHdr.NbOfTxs)

	pmtInves := doc.CstmrCdtTrfInitn.PmtInf

	var nbOfTxs int = 0
	for _, pmtInf := range pmtInves {
		no, _ := strconv.Atoi(pmtInf.NbOfTxs)
		nbOfTxs = nbOfTxs + no
	}

	return ghNbOfTxs == nbOfTxs
}

func (c *Control) ControlPmtInfCtrlSum(pmtInf *gen.PaymentInstructionInformation3) bool {
	piCtrlSum := pmtInf.CtrlSum

	cdtTrfTxInves := pmtInf.CdtTrfTxInf

	var ctrlSum float64 = 0.0
	for _, cdtTrfTxInf := range cdtTrfTxInves {
		ctrlSum = ctrlSum + cdtTrfTxInf.Amt.InstdAmt.Value
	}

	return piCtrlSum == ctrlSum
}

func (c *Control) ControlPmtInfNbOfTxs(pmtInf *gen.PaymentInstructionInformation3) bool {
	piNbOfTxs, _ := strconv.Atoi(pmtInf.NbOfTxs)

	nbOfTxs := len(pmtInf.CdtTrfTxInf)

	return piNbOfTxs == nbOfTxs
}
