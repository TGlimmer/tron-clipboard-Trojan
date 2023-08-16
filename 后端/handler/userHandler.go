package handler

import (
	"GinHello/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getadress(context *gin.Context) {
	var address model.AddressModel
	if err := context.ShouldBindQuery(&address); err != nil {
		context.String(http.StatusBadRequest, "error")
		log.Panicln("err ->", err.Error())
	}
	newAddress := address.FindAddress()
	if newAddress != "" {
		context.JSON(200, gin.H{"address": newAddress})
	} else {
		context.JSON(404, gin.H{"error": "address not found"})
	}
}
