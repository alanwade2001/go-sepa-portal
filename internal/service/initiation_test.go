package service_test

import (
	"testing"
	"time"

	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/alanwade2001/go-sepa-portal/internal/repository/entity"
	"github.com/alanwade2001/go-sepa-portal/internal/service"
	mrepos "github.com/alanwade2001/go-sepa-portal/mocks/internal_/repository"
	mocks "github.com/alanwade2001/go-sepa-portal/mocks/internal_/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestFindByID(t *testing.T) {

	repos := mrepos.NewIInitiation(t)
	repos.On("FindByID", mock.AnythingOfType("string")).Return(&entity.Initiation{
		Model:           gorm.Model{ID: 1},
		MsgID:           "msgid",
		NbOfTxs:         1,
		CreDtTm:         time.Now(),
		CtrlSum:         10.01,
		State:           "Approved",
		DocID:           1,
		RejectionReason: "",
	}, nil)

	msg := mocks.NewIMessage(t)

	svc := service.NewInitiation(repos, msg)

	mdl, err := svc.FindByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, mdl)

	assert.Equal(t, "msgid", mdl.MsgID)

}

func TestFindAll(t *testing.T) {

	repos := mrepos.NewIInitiation(t)
	repos.On("FindAll").Return([]*entity.Initiation{
		{
			Model:           gorm.Model{ID: 1},
			MsgID:           "msgid",
			NbOfTxs:         1,
			CreDtTm:         time.Now(),
			CtrlSum:         10.01,
			State:           "Approved",
			DocID:           1,
			RejectionReason: "",
		},
		{
			Model:           gorm.Model{ID: 2},
			MsgID:           "msgid2",
			NbOfTxs:         2,
			CreDtTm:         time.Now(),
			CtrlSum:         10.02,
			State:           "Approved",
			DocID:           2,
			RejectionReason: "",
		},
		{
			Model:           gorm.Model{ID: 3},
			MsgID:           "msgid3",
			NbOfTxs:         3,
			CreDtTm:         time.Now(),
			CtrlSum:         10.03,
			State:           "Rejected",
			DocID:           3,
			RejectionReason: "error",
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
	repos.On("FindByID", mock.AnythingOfType("string")).Return(&entity.Initiation{
		Model:           gorm.Model{ID: 1},
		MsgID:           "msgid",
		NbOfTxs:         1,
		CreDtTm:         time.Now(),
		CtrlSum:         10.01,
		State:           string(model.InitiatedState),
		DocID:           1,
		RejectionReason: "",
	}, nil)

	repos.On("Perist", mock.Anything).Return(&entity.Initiation{
		Model:           gorm.Model{ID: 1},
		MsgID:           "msgid",
		NbOfTxs:         1,
		CreDtTm:         time.Now(),
		CtrlSum:         10.01,
		State:           string(model.AcceptedState),
		DocID:           1,
		RejectionReason: "",
	}, nil)

	msg := mocks.NewIMessage(t)
	msg.On("Send", mock.Anything).Return(nil)

	svc := service.NewInitiation(repos, msg)

	mdl, err := svc.SendInitiationAccept("1")

	assert.NoError(t, err)
	assert.NotNil(t, mdl)

	assert.Equal(t, model.AcceptedState, mdl.State)

}
