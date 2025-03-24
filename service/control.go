package service

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	"github.com/alanwade2001/go-sepa-iso/schema"
	"github.com/alanwade2001/go-sepa-portal/data"
	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

type Control struct {
	p1Handler *xsdvalidate.XsdHandler
}

type IControl interface {
	Check(doc *pain_001_001_03.Document) (*data.CheckResult, error)
}

func NewControl() IControl {
	xsdvalidate.Init()
	if handler, err := schema.NewPain001XsdHandler(); err != nil {
		slog.Error("failed to create pain001 xsd handler", "err", err)
		panic(err)
	} else {
		control := &Control{
			p1Handler: handler,
		}

		return control
	}
}

func (c *Control) Cleanup() {
	c.p1Handler.Free()
	xsdvalidate.Cleanup()
}

func (c *Control) Check(doc *pain_001_001_03.Document) (*data.CheckResult, error) {

	if result, err := c.ControlGrpHdrCtrlSum(doc); !result.Pass || err != nil {
		return result, err
	}

	if result, err := c.ControlGrpHdrNbOfTxs(doc); !result.Pass || err != nil {
		return result, err
	}

	for _, pmtInf := range doc.CstmrCdtTrfInitn.PmtInf {
		if result, err := c.ControlPmtInfCtrlSum(pmtInf); !result.Pass || err != nil {
			return result, err
		}

		if result, err := c.ControlPmtInfNbOfTxs(pmtInf); !result.Pass || err != nil {
			return result, err
		}
	}

	return data.NewPassResult(), nil
}

func (c *Control) ControlGrpHdrCtrlSum(doc *pain_001_001_03.Document) (*data.CheckResult, error) {
	ghCtrlSum := doc.CstmrCdtTrfInitn.GrpHdr.CtrlSum

	pmtInves := doc.CstmrCdtTrfInitn.PmtInf

	var ctrlSum float64 = 0.0
	for _, pmtInf := range pmtInves {
		ctrlSum = ctrlSum + pmtInf.CtrlSum
	}

	if ghCtrlSum == ctrlSum {
		return data.NewPassResult(), nil
	} else {
		return data.NewFailResult(fmt.Sprintf("ctrlSum does not match: expected=[%f], actual=[%f]", ghCtrlSum, ctrlSum), nil), nil
	}
}

func (c *Control) ControlGrpHdrNbOfTxs(doc *pain_001_001_03.Document) (*data.CheckResult, error) {
	ghNbOfTxs, _ := strconv.Atoi(doc.CstmrCdtTrfInitn.GrpHdr.NbOfTxs)

	pmtInves := doc.CstmrCdtTrfInitn.PmtInf

	var nbOfTxs int = 0
	for _, pmtInf := range pmtInves {
		no, _ := strconv.Atoi(pmtInf.NbOfTxs)
		nbOfTxs = nbOfTxs + no
	}

	if ghNbOfTxs == nbOfTxs {
		return data.NewPassResult(), nil
	} else {
		return data.NewFailResult(fmt.Sprintf("nbOfTxs does not match: expected=[%d], actual=[%d]", ghNbOfTxs, nbOfTxs), nil), nil
	}
}

func (c *Control) ControlPmtInfCtrlSum(pmtInf *pain_001_001_03.PaymentInstructionInformation3) (*data.CheckResult, error) {

	piCtrlSum := pmtInf.CtrlSum
	cdtTrfTxInves := pmtInf.CdtTrfTxInf

	var ctrlSum float64 = 0.0
	for _, cdtTrfTxInf := range cdtTrfTxInves {
		ctrlSum = ctrlSum + cdtTrfTxInf.Amt.InstdAmt.Value
	}

	if piCtrlSum == ctrlSum {
		return data.NewPassResult(), nil
	} else {
		return data.NewFailResult(fmt.Sprintf("ctrlSum does not match: expected=[%f], actual=[%f]", piCtrlSum, pmtInf.CtrlSum), nil), nil
	}

}

func (c *Control) ControlPmtInfNbOfTxs(pmtInf *pain_001_001_03.PaymentInstructionInformation3) (*data.CheckResult, error) {
	var result *data.CheckResult
	if piNbOfTxs, err := strconv.Atoi(pmtInf.NbOfTxs); err != nil {
		result = data.NewFailResult("Failed to parse Payment NbOfTxs", err)
	} else {

		nbOfTxs := len(pmtInf.CdtTrfTxInf)
		pass := piNbOfTxs == nbOfTxs

		if pass {
			result = data.NewPassResult()
		} else {
			result = data.NewFailResult(fmt.Sprintf("nbOfTxs does not match: expected=[%d], actual=[%d]", piNbOfTxs, nbOfTxs), nil)
		}
	}

	return result, nil
}
