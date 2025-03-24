package handler_test

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/alanwade2001/go-sepa-portal/internal/data"
	"github.com/alanwade2001/go-sepa-portal/internal/handler"
	mocks "github.com/alanwade2001/go-sepa-portal/mocks/internal_/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	slog.Info("Line 1")
	service := mocks.NewIDocument(t)

	initn := &data.Initiation{
		ID: 1,
	}
	slog.Info("Line 2")
	service.On("InitiateDocument", "<hello>alan</hello>").Return(initn, nil)

	slog.Info("Line 3")
	w := httptest.NewRecorder()

	slog.Info("Line 4")
	c := GetTestGinContext(w)

	slog.Info("Line 5")
	handler := handler.NewDocument(service)

	slog.Info("Line 6")
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/xml")

	slog.Info("Line 7")
	c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte("<hello>alan</hello>")))

	slog.Info("Line 8")
	handler.PostDocument(c)

	slog.Info("Line 9")
	assert.Equal(t, http.StatusCreated, w.Code)
}

// mock gin context
func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}
