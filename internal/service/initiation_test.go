package service_test

import (
	"testing"

	"github.com/alanwade2001/go-sepa-portal/internal/data"
	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/alanwade2001/go-sepa-portal/internal/service"
	mrepos "github.com/alanwade2001/go-sepa-portal/mocks/internal_/repository"
	mocks "github.com/alanwade2001/go-sepa-portal/mocks/internal_/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindByID(t *testing.T) {

	repos := mrepos.NewIInitiation(t)
	repos.On("FindByID", mock.AnythingOfType("string")).Return(&data.Initiation{
		ID:       1,
		MsgId:    "msgid",
		NbOfTxs:  "1",
		CreDtTm:  "2025-03-24T18:10:55Z",
		CtrlSum:  10.01,
		State:    "Approved",
		DocID:    1,
		RjctdRsn: "",
	}, nil)

	msg := mocks.NewIMessage(t)

	svc := service.NewInitiation(repos, msg)

	mdl, err := svc.FindByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, mdl)

	assert.Equal(t, "msgid", mdl.MsgId)

}

func TestFindAll(t *testing.T) {

	repos := mrepos.NewIInitiation(t)
	repos.On("FindAll").Return([]*data.Initiation{
		{
			ID:       1,
			MsgId:    "msgid",
			NbOfTxs:  "1",
			CreDtTm:  "2025-03-24T18:10:55Z",
			CtrlSum:  10.01,
			State:    "Approved",
			DocID:    1,
			RjctdRsn: "",
		},
		{
			ID:       2,
			MsgId:    "msgid2",
			NbOfTxs:  "2",
			CreDtTm:  "2025-03-24T18:10:55Z",
			CtrlSum:  10.02,
			State:    "Approved",
			DocID:    2,
			RjctdRsn: "",
		},
		{
			ID:       3,
			MsgId:    "msgid3",
			NbOfTxs:  "3",
			CreDtTm:  "2025-03-24T18:10:55Z",
			CtrlSum:  10.03,
			State:    "Rejected",
			DocID:    3,
			RjctdRsn: "error",
		},
	}, nil)

	msg := mocks.NewIMessage(t)

	svc := service.NewInitiation(repos, msg)

	mdls, err := svc.FindAll()

	assert.NoError(t, err)
	assert.NotNil(t, mdls)

	assert.Len(t, mdls, 3)
	assert.Equal(t, mdls[2].State, model.RejectedState)

}

func TestAccept(t *testing.T) {

	repos := mrepos.NewIInitiation(t)
	repos.On("FindByID", mock.AnythingOfType("string")).Return(&data.Initiation{
		ID:       1,
		MsgId:    "msgid",
		NbOfTxs:  "1",
		CreDtTm:  "2025-03-24T18:10:55Z",
		CtrlSum:  10.01,
		State:    "Approved",
		DocID:    1,
		RjctdRsn: "",
	}, nil)

	repos.On("Persist", mock.Anything).Return(&data.Initiation{
		ID:       1,
		MsgId:    "msgid",
		NbOfTxs:  "1",
		CreDtTm:  "2025-03-24T18:10:55Z",
		CtrlSum:  10.01,
		State:    "Approved",
		DocID:    1,
		RjctdRsn: "",
	}, nil)

	msg := mocks.NewIMessage(t)
	msg.On("Send", mock.Anything).Return(nil)

	svc := service.NewInitiation(repos, msg)

	mdl, err := svc.SendInitiationAccept("1")

	assert.NoError(t, err)
	assert.NotNil(t, mdl)

	assert.Equal(t, model.AcceptedState, mdl.State)

}
