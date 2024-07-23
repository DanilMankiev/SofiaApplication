package v1

import (
	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) register(c *gin.Context){
	var input entity.RegiterInput

	if err:=c.BindJSON(&input);err!=nil{
		newResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	
	err:=h.services.Authorization.Register(input)
	if err!=nil{
		newResponse(c, http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,StatusResponse{
		Status: "OK",
	})
}

func (h *Handler) signIn ( c *gin.Context){
	
}

func(h *Handler) sendCodeEmail(c *gin.Context){
	var email string

	err:=h.services.Authorization.SendCodeEmail(email)
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{
		Status: "Sent",
	})	
	

}

func(h *Handler) sendCodeSMS(c *gin.Context){
	var phone string

	err:=h.services.Authorization.SendCodeSms(phone)
	if err!=nil{
		newResponse(c, http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{
		Status: "Sent",
	})

}