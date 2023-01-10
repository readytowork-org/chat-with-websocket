package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FirebaseController struct {
	logger   infrastructure.Logger
	services services.FirebaseService
}

func NewFirebaseController(
	logger infrastructure.Logger,
	services services.FirebaseService,

) FirebaseController {
	return FirebaseController{
		logger:   logger,
		services: services,
	}

}

func (fc FirebaseController) CreateUser(c *gin.Context) {
	var newUser models.FirebaseAuthUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		msg := "Error validating user input"
		fc.logger.Zap.Info(msg, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": msg,
		})
	}

	regsiteredUser, err := fc.services.CreateUser(newUser)
	if err != nil {
		msg := "Error validating user input"
		fc.logger.Zap.Info(msg, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": msg})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
		"data":    regsiteredUser,
	})
}
