package service_test

import (
	"testing"

	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	"github.com/alanwade2001/go-sepa-portal/data"
	"github.com/alanwade2001/go-sepa-portal/mocks/github.com/alanwade2001/go-sepa-portal/repository"
	mock_svc "github.com/alanwade2001/go-sepa-portal/mocks/github.com/alanwade2001/go-sepa-portal/service"
	"github.com/alanwade2001/go-sepa-portal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Pass(t *testing.T) {
	repos := repository.NewMockIInitiation(t)
	msg := mock_svc.NewMockIMessage(t)
	ctrl := mock_svc.NewMockIControl(t)
	store := mock_svc.NewMockIStore(t)
	decoder := mock_svc.NewMockIPain001Decoder(t)

	store.On("StoreDocument", mock.AnythingOfType("string")).Return(&data.Document{
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

	ctrl.On("Check", mock.Anything).Return(&data.CheckResult{
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
	repos := repository.NewMockIInitiation(t)
	msg := mock_svc.NewMockIMessage(t)
	ctrl := mock_svc.NewMockIControl(t)
	store := mock_svc.NewMockIStore(t)
	decoder := mock_svc.NewMockIPain001Decoder(t)

	store.On("StoreDocument", mock.AnythingOfType("string")).Return(&data.Document{
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

	ctrl.On("Check", mock.Anything).Return(&data.CheckResult{
		Pass: false,
		Msg:  "error",
	}, nil)

	docSvc := service.NewDocument(repos, msg, ctrl, store, decoder)

	newInitn, err := docSvc.InitiateDocument("dummy xml")

	assert.NoError(t, err)
	assert.Equal(t, "msgid", newInitn.MsgId)
	assert.Equal(t, "error", newInitn.RjctdRsn)

}
