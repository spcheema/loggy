package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jz222/loggy/models"
	"github.com/jz222/loggy/services/service"
	"github.com/jz222/loggy/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type serviceController struct{}

// Service contains all controllers related to services.
var Service serviceController

func (s *serviceController) Create(c *gin.Context) {
	var newService models.Service

	err := json.NewDecoder(c.Request.Body).Decode(&newService)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	userData, ok := c.Get("user")
	if !ok {
		utils.RespondWithError(c, http.StatusInternalServerError, "could not parse user data")
		return
	}

	newService.OrganizationID = userData.(models.User).OrganizationID

	createdService, err := service.Create(newService)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, createdService)
}

func (s *serviceController) Delete(c *gin.Context) {
	id := c.Param("id")

	serviceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userData, ok := c.Get("user")
	if !ok {
		utils.RespondWithError(c, http.StatusInternalServerError, "could not parse user data")
		return
	}

	filter := bson.M{"_id": serviceID, "organizationId": userData.(models.User).OrganizationID}

	count, err := service.Delete(filter)
	if err != nil {
		utils.RespondWithError(c, http.StatusForbidden, err.Error())
		return
	}

	if count == 0 {
		utils.RespondWithError(c, http.StatusBadRequest, "the service with the id "+id+" does not exist")
		return
	}

	utils.RespondWithSuccess(c)
}
