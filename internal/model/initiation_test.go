package model

import (
	"testing"
	"time"

	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"

	"github.com/stretchr/testify/assert"
)

var dt = time.Date(2025, time.March, 17, 21, 0, 30, 0, time.Local)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		gh       *pain_001_001_03.GroupHeader32
		expected *Initiation
	}{
		{
			desc: "",
			gh: &pain_001_001_03.GroupHeader32{
				MsgId:   "msg-1",
				CreDtTm: dt.Format("2006-03-15T12:12:12"),
				NbOfTxs: "1",
				CtrlSum: 10.01,
				InitgPty: &pain_001_001_03.PartyIdentification32{
					Id: &pain_001_001_03.Party6Choice{
						OrgId: &pain_001_001_03.OrganisationIdentification4{
							Othr: []*pain_001_001_03.GenericOrganisationIdentification1{
								{
									Id: "IP12345",
								},
							},
						},
					},
				},
			},
			expected: &Initiation{
				ID:              0,
				MsgID:           "msg-1",
				CtrlSum:         10.01,
				NbOfTxs:         1,
				CreDtTm:         dt.Format("2006-03-15T12:12:12"),
				State:           AcceptedState,
				DocID:           6,
				RejectionReason: "",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := NewInitiation(tC.gh, AcceptedState, 6, "")

			assert.NoError(t, err)

			assert.Equal(t, actual, tC.expected)
		})
	}
}
