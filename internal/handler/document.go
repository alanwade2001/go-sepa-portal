package handler

import (
	"log/slog"
	"net/http"

	"github.com/alanwade2001/go-sepa-infra/routing"
	"github.com/alanwade2001/go-sepa-portal/internal/service"
	"github.com/gin-gonic/gin"
)

type Document struct {
	Service service.IDocument
}

type IDocument interface {
	PostDocument(c *gin.Context)
	Register(r *routing.Router)
}

func NewDocument(service service.IDocument) IDocument {
	document := &Document{
		Service: service}

	return document
}

func (d Document) Register(r *routing.Router) {
	r.Router.POST("/documents", d.PostDocument)
}

// postInitiation adds an initiations from JSON received in the request body.
func (d *Document) PostDocument(c *gin.Context) {

	if data, err := c.GetRawData(); err != nil {
		c.XML(http.StatusInternalServerError, err)
	} else if newInitiation, err := d.Service.InitiateDocument(string(data)); err != nil {
		slog.Error("failed to post document", "Error", err)
		c.XML(http.StatusInternalServerError, newInitiation)
	} else {
		c.XML(http.StatusCreated, newInitiation)
	}
}
