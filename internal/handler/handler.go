package handler

import (
	"net/http"

	"Diplom/internal/model"
	"Diplom/internal/processingData"

	"github.com/gin-gonic/gin"
)

func statusPage(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "./data/status_page.html")

	return
}

func api(c *gin.Context) {

	var result model.ResultT

	resultSet := processingData.GetResultData()
	billing := resultSet.Billing

	if len(resultSet.SMS) > 0 && len(resultSet.MMS) > 1 && len(resultSet.VoiceCall) > 0 && len(resultSet.Email) > 0 &&
		len(resultSet.Support) > 1 && len(resultSet.Incidents) > 1 && billing.CheckoutPage == true ||
		billing.CheckoutPage == false {

		result.Data = resultSet
		result.Status = true
	} else {
		result.Status = false
		result.Error = "Error on collect data"
		result.Data = resultSet
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(http.StatusOK, result)

	return
}
