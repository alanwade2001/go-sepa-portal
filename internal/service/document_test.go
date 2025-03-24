package service_test

import (
	"testing"

	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	"github.com/alanwade2001/go-sepa-portal/internal/data"
	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/alanwade2001/go-sepa-portal/internal/service"
	mrepos "github.com/alanwade2001/go-sepa-portal/mocks/internal_/repository"
	mservice "github.com/alanwade2001/go-sepa-portal/mocks/internal_/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Pass(t *testing.T) {
	repos := mrepos.NewIInitiation(t)
	msg := mservice.NewIMessage(t)
	ctrl := mservice.NewIControl(t)
	store := mservice.NewIStore(t)
	decoder := mservice.NewIPain001Decoder(t)

	store.On("StoreDocument", mock.AnythingOfType("string")).Return(&model.Document{
		ID:      1,
		Content: "content",
	}, nil)

	decoder.On("Decode", mock.AnythingOfType("string")).Return(&pain_001_001_03.Document{
		CstmrCdtTrfInitn: &pain_001_001_03.CustomerCreditTransferInitiationV03{
			GrpHdr: &pain_001_001_03.GroupHeader32{
				MsgId:   "msgid",
				CreDtTm: "2025-03-23T18:42:20Z",
				NbOfTxs: "1",
				CtrlSum: 10.01,
			},
		},
	}, nil)

	repos.On("Persist", mock.Anything).Return(&data.Initiation{
		ID: 1,
	}, nil)

	msg.On("Send", mock.Anything).Return(nil)

	ctrl.On("Check", mock.Anything).Return(&model.CheckResult{
		Pass: true,
		Msg:  "",
	}, nil)

	docSvc := service.NewDocument(repos, msg, ctrl, store, decoder)

	newInitn, err := docSvc.InitiateDocument("dummy xml")

	assert.NoError(t, err)
	assert.Equal(t, "msgid", newInitn.MsgId)
	assert.Equal(t, "", newInitn.RjctdRsn)

}

func Test_Fail(t *testing.T) {
	repos := mrepos.NewIInitiation(t)
	msg := mservice.NewIMessage(t)
	ctrl := mservice.NewIControl(t)
	store := mservice.NewIStore(t)
	decoder := mservice.NewIPain001Decoder(t)

	store.On("StoreDocument", mock.AnythingOfType("string")).Return(&model.Document{
		ID:      1,
		Content: "content",
	}, nil)

	decoder.On("Decode", mock.AnythingOfType("string")).Return(&pain_001_001_03.Document{
		CstmrCdtTrfInitn: &pain_001_001_03.CustomerCreditTransferInitiationV03{
			GrpHdr: &pain_001_001_03.GroupHeader32{
				MsgId:   "msgid",
				CreDtTm: "2025-03-23T18:42:20Z",
				NbOfTxs: "1",
				CtrlSum: 10.01,
			},
		},
	}, nil)

	repos.On("Persist", mock.Anything).Return(&data.Initiation{
		ID: 1,
	}, nil)

	//msg.On("SendAccepted", mock.Anything).Times(0)
	msg.On("Send", mock.Anything).Return(nil)

	ctrl.On("Check", mock.Anything).Return(&model.CheckResult{
		Pass: false,
		Msg:  "error",
	}, nil)

	docSvc := service.NewDocument(repos, msg, ctrl, store, decoder)

	newInitn, err := docSvc.InitiateDocument("dummy xml")

	assert.NoError(t, err)
	assert.Equal(t, "msgid", newInitn.MsgId)
	assert.Equal(t, "error", newInitn.RjctdRsn)

}
