package handler

import (
	"net/http"
	"sofosrobotics/helper"
	"sofosrobotics/robotorders"

	"github.com/gin-gonic/gin"
)

type robotorderHandler struct {
	robotorderService robotorders.Service
}

func NewRobotorderHandler(robotorderService robotorders.Service) *robotorderHandler {
	return &robotorderHandler{robotorderService}
}

func (h *robotorderHandler) Save(c *gin.Context) {
	var input robotorders.Input

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var dataInput robotorders.Robotorder
	dataInput.Idrobotproductpartdevice = input.Idrobotproductpartdevice
	dataInput.Userid = input.Userid
	dataInput.Kodeinvoice = input.Kodeinvoice
	dataInput.Namecustomer = input.Namecustomer
	dataInput.Namarobotmaster = input.Namarobotmaster
	dataInput.Addresscustomer = input.Addresscustomer
	dataInput.Delivery = input.Delivery
	dataInput.Pricedelivery = input.Pricedelivery
	dataInput.Pricemethod = input.Pricemethod
	dataInput.Pricemethodsn = input.Pricemethodsn
	dataInput.Subtotal = input.Subtotal
	dataInput.Ppn = input.Ppn
	dataInput.Totals = input.Totals
	dataInput.Pesan = input.Pesan
	dataInput.Deliverydesc = input.Deliverydesc
	dataInput.Pricemethodadmin = input.Pricemethodadmin
	dataInput.Recommended = input.Recommended

	newUser, err := h.robotorderService.Save(dataInput)

	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := robotorders.FormatRobotorder(newUser)

	response := helper.APIResponse("Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *robotorderHandler) ByUserid(c *gin.Context) {
	var input robotorders.Userid

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.robotorderService.FindAllUserid(input.Userid)

	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := robotorders.FormatAll(newUser)

	response := helper.APIResponse("Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *robotorderHandler) Kodeinvoice(c *gin.Context) {
	var input robotorders.Kodeinvoice

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.robotorderService.FindAllKodeinvoice(input.Kodeinvoice)

	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := robotorders.FormatAll(newUser)

	response := helper.APIResponse("Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
