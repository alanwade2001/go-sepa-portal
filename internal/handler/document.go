package handler

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/alanwade2001/go-sepa-infra/routing"
	"github.com/alanwade2001/go-sepa-portal/internal/service"
	"github.com/gin-gonic/gin"
)

type Document struct {
	service *service.Document
}

func NewDocument(service *service.Document, r *routing.Router) *Document {
	document := &Document{
		service: service,
	}

	r.Router.POST("/documents", document.PostDocument)

	return document
}

// postInitiation adds an initiations from JSON received in the request body.
func (d *Document) PostDocument(c *gin.Context) {

	data, _ := c.GetRawData()

	log.Println("initiating doc")
	if newInitiation, err := d.service.InitiateDocument(string(data)); err != nil {
		slog.Error("failed to post document", "Error", err)
		c.IndentedJSON(http.StatusInternalServerError, newInitiation)
	} else {
		c.IndentedJSON(http.StatusCreated, newInitiation)
	}
}
