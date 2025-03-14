package handler

import (
	"net/http"

	"github.com/alanwade2001/go-sepa-infra/routing"
	"github.com/alanwade2001/go-sepa-portal/internal/model"
	"github.com/alanwade2001/go-sepa-portal/internal/service"
	"github.com/gin-gonic/gin"
)

type Initiation struct {
	service *service.Initiation
}

func NewInitiation(service *service.Initiation, r *routing.Router) *Initiation {
	initiation := &Initiation{
		service: service,
	}

	r.Router.GET("/initiations", initiation.GetInitiation)
	r.Router.GET("/initiations/:id", initiation.GetInitiationByID)

	r.Router.PUT("/initiations/:id/approve", initiation.PutInitiationApprove)
	r.Router.PUT("/initiations/:id/cancel", initiation.PutInitiationCancel)
	r.Router.PUT("/initiations/:id/accept", initiation.PutInitiationAccept)
	r.Router.PUT("/initiations/:id/reject", initiation.PutInitiationReject)

	return initiation
}

// getInitiation responds with the list of all initiationss as JSON.
func (i *Initiation) GetInitiation(c *gin.Context) {
	initiations, err := i.service.FindAll()

	code := http.StatusOK

	if err != nil {
		code = http.StatusInternalServerError
	}

	c.IndentedJSON(code, initiations)
}

// getInitiationByID locates the initiations whose ID value matches the id
// parameter sent by the client, then returns that initiations as a response.
func (i *Initiation) GetInitiationByID(c *gin.Context) {
	id := c.Param("id")
	initn, err := i.service.FindByID(id)

	if err != nil {
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
		return
	}

	c.IndentedJSON(http.StatusOK, initiation)
}
