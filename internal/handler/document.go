package handler

import (
	"log"
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
	newInitiation, err := d.service.InitiateDocument(string(data))

	code := http.StatusCreated

	if err != nil {
		code = http.StatusInternalServerError
	}

	// Add the new initiations to the slice.
	c.IndentedJSON(code, newInitiation)

}
