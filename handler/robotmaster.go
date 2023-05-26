package handler

import (
	"net/http"
	"sofosrobotics/helper"
	"sofosrobotics/robotmasters"

	"github.com/gin-gonic/gin"
)

type robotmasterHandler struct {
	robotmasterService robotmasters.Service
}

func RobotMasterHandler(robotmasterService robotmasters.Service) *robotmasterHandler {
	return &robotmasterHandler{robotmasterService}
}

func (h *robotmasterHandler) FindAllRobotic(c *gin.Context) {

	FindAllRobotic, err := h.robotmasterService.FindAllRobotic()

	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := robotmasters.FormatAll(FindAllRobotic)
	response := helper.APIResponse("Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *robotmasterHandler) Recomendation(c *gin.Context) {

	Recomendation, err := h.robotmasterService.Recomendation()

	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := robotmasters.FormatAll(Recomendation)
	response := helper.APIResponse("Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *robotmasterHandler) FindAllAutomation(c *gin.Context) {

	FindAllAutomation, err := h.robotmasterService.FindAllAutomation()

	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := robotmasters.FormatAll(FindAllAutomation)
	response := helper.APIResponse("Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *robotmasterHandler) Byid(c *gin.Context) {
	var input robotmasters.IDINPUT

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.robotmasterService.Byid(input.ID)

	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := robotmasters.FormatRobotMaster(newUser, "")

	response := helper.APIResponse("Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
