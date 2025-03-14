package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	//"sepa/portal/main"
	"testing"
)

var a *App

func TestMain(m *testing.M) {

	a = NewApp()

	code := m.Run()
	os.Exit(code)
}

func clearTable() {
	a.Infra.Persist.DB.Exec("DELETE FROM INITIATIONS")
	a.Infra.Persist.DB.Exec("ALTER SEQUENCE initiations_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Infra.Router.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/initiations", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentInitiation(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/initiations/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Initiation not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Initiation not found'. Got '%s'", m["error"])
	}
}

var GOOD_XML string = `
<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.001.001.03">
    <CstmrCdtTrfInitn>
        <GrpHdr>
            <MsgId>MSG_ID1</MsgId>
            <CreDtTm>2021-09-28T13:16:37.219430288</CreDtTm>
            <NbOfTxs>1</NbOfTxs>
            <CtrlSum>11</CtrlSum>
            <InitgPty>
                <Id>
                    <PrvtId>
                        <Othr>
                            <Id>IP-123456</Id>
                        </Othr>
                    </PrvtId>
                </Id>
            </InitgPty>
        </GrpHdr>
        <PmtInf>
            <PmtInfId>PIID-1</PmtInfId>
            <PmtMtd>TRF</PmtMtd>
            <NbOfTxs>1</NbOfTxs>
            <CtrlSum>11</CtrlSum>
            <PmtTpInf/>
            <ReqdExctnDt>2025-01-10</ReqdExctnDt>
            <Dbtr>
                <Nm>Mr. Debtor</Nm>
            </Dbtr>
            <DbtrAcct>
                <Id>
                    <IBAN>IE30BOFI90909012345678</IBAN>
                </Id>
            </DbtrAcct>
            <DbtrAgt>
                <FinInstnId>
                    <BIC>BOFIIE2D</BIC>
                </FinInstnId>
            </DbtrAgt>
            <ChrgBr>CRED</ChrgBr>
            <CdtTrfTxInf>
                <PmtId>
                    <EndToEndId>E2EID-2</EndToEndId>
                </PmtId>
                <Amt>
                    <InstdAmt Ccy="EUR">11</InstdAmt>
                </Amt>
                <CdtrAgt>
                    <FinInstnId>
                        <BIC>BOFIIE2D</BIC>
                    </FinInstnId>
                </CdtrAgt>
                <Cdtr>
                    <Nm>Mr. Cdtr</Nm>
                </Cdtr>
                <CdtrAcct>
                    <Id>
                        <IBAN>IE16BOFI90909187654321</IBAN>
                    </Id>
                </CdtrAcct>
            </CdtTrfTxInf>
        </PmtInf>
    </CstmrCdtTrfInitn>
</Document>
`

// func TestCreateInitiation(t *testing.T) {

// 	docId := 2

// 	mockReturnBody := make(map[string]interface{})
// 	mockReturnBody["id"] = docId
// 	mockReturnBody["content"] = "hello world"

// 	body, _ := json.Marshal(mockReturnBody)

// 	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(201)
// 		_, err := w.Write(body)

// 		assert.NoError(t, err)
// 	}))

// 	defer testServer.Close()

// 	clearTable()

// 	a.IStoreService.Address = testServer.URL
// 	log.Printf("StoreService.Address: [%s]", a.StoreService.Address)

// 	var xmlStr = []byte(GOOD_XML)
// 	req, _ := http.NewRequest("POST", "/documents", bytes.NewBuffer(xmlStr))
// 	req.Header.Set("Content-Type", "application/xml")

// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusCreated, response.Code)

// 	var bodyMap map[string]interface{}
// 	json.Unmarshal(response.Body.Bytes(), &bodyMap)

// 	assert.Equal(t, 11.0, bodyMap["ctrlSum"])
// 	assert.Equal(t, "1", bodyMap["nbOfTxs"])
// 	assert.Equal(t, "Accepted", bodyMap["state"])
// 	assert.Equal(t, "MSG_ID1", bodyMap["msgId"])
// 	assert.Equal(t, 1.0, bodyMap["id"])
// 	assert.Equal(t, float64(docId), bodyMap["docId"])
// }
