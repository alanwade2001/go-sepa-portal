package handler

import (
	"net/http"

	"github.com/alanwade2001/go-sepa-infra/routing"
	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/alanwade2001/go-sepa-portal/internal/service"
	"github.com/gin-gonic/gin"
)

type Initiation struct {
	service service.IInitiation
}

type IInitiation interface {
	Register(r *routing.Router)
	GetInitiation(c *gin.Context)
	GetInitiationByID(c *gin.Context)
	PutInitiationApprove(c *gin.Context)
	PutInitiationCancel(c *gin.Context)
	PutInitiationAccept(c *gin.Context)
	PutInitiationReject(c *gin.Context)
}

func NewInitiation(service service.IInitiation) IInitiation {
	initiation := &Initiation{
		service: service,
	}

	return initiation
}

func (i Initiation) Register(r *routing.Router) {

	r.Router.GET("/initiations", i.GetInitiation)
	r.Router.GET("/initiations/:id", i.GetInitiationByID)

	r.Router.PUT("/initiations/:id/approve", i.PutInitiationApprove)
	r.Router.PUT("/initiations/:id/cancel", i.PutInitiationCancel)
	r.Router.PUT("/initiations/:id/accept", i.PutInitiationAccept)
	r.Router.PUT("/initiations/:id/reject", i.PutInitiationReject)

}

// getInitiation responds with the list of all initiationss as JSON.
func (i *Initiation) GetInitiation(c *gin.Context) {

	if initiations, err := i.service.FindAll(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, initiations)
	} else {
		c.IndentedJSON(http.StatusOK, initiations)
	}

}

// getInitiationByID locates the initiations whose ID value matches the id
// parameter sent by the client, then returns that initiations as a response.
func (i *Initiation) GetInitiationByID(c *gin.Context) {
	id := c.Param("id")
	if initn, err := i.service.FindByID(id); err != nil {
		message := "Initiation not found"
		payload := map[string]string{"error": message}
		c.IndentedJSON(http.StatusNotFound, payload)
	} else {
		c.IndentedJSON(http.StatusOK, initn)
	}

}

func (i *Initiation) PutInitiationApprove(c *gin.Context) {
	id := c.Param("id")
	initiation, err := i.service.SendInitiationApprove(id)
	i.PutResponse(c, initiation, err)
}

func (i *Initiation) PutInitiationCancel(c *gin.Context) {
	id := c.Param("id")
	initiation, err := i.service.SendInitiationCancel(id)
	i.PutResponse(c, initiation, err)
}

func (i *Initiation) PutInitiationAccept(c *gin.Context) {
	id := c.Param("id")
	initiation, err := i.service.SendInitiationAccept(id)
	i.PutResponse(c, initiation, err)
}

func (i *Initiation) PutInitiationReject(c *gin.Context) {
	id := c.Param("id")
	initiation, err := i.service.SendInitiationReject(id)
	i.PutResponse(c, initiation, err)
}

func (i *Initiation) PutResponse(c *gin.Context, initiation *model.Initiation, err error) {

	if err != nil {
		m := make(map[string]interface{})
		m["error"] = err
		c.IndentedJSON(http.StatusBadRequest, m)
	} else {
		c.IndentedJSON(http.StatusOK, initiation)
	}
}
