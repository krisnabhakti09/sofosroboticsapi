package handler

import (
	"net/http"
	"sofosrobotics/helper"
	"sofosrobotics/robotproductpartdevices"

	"github.com/gin-gonic/gin"
)

type robotproductpartdeviceHandler struct {
	robotproductpartdeviceService robotproductpartdevices.Service
}

func RobotproductpartdeviceHandler(robotproductpartdeviceService robotproductpartdevices.Service) *robotproductpartdeviceHandler {
	return &robotproductpartdeviceHandler{robotproductpartdeviceService}
}

func (h *robotproductpartdeviceHandler) FindAllIDRobotMasterAndRobotPartDevice(c *gin.Context) {
	var input robotproductpartdevices.GetInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	FindAllRobotic, err := h.robotproductpartdeviceService.FindAllIDRobotMasterAndRobotPartDevice(input.Idrobotmaster, input.Idrobotpartdevice)

	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := robotproductpartdevices.FormatAll(FindAllRobotic)
	response := helper.APIResponse("Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *robotproductpartdeviceHandler) FindAllIDRobotProductPartDevice(c *gin.Context) {
	var input robotproductpartdevices.GetID

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	FindAllRobotic, err := h.robotproductpartdeviceService.FindAllIDRobotProductPartDevice(input.ID)

	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := robotproductpartdevices.FormatRobotProductDevice(FindAllRobotic)
	response := helper.APIResponse("Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
