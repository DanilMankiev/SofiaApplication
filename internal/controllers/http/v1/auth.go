package v1

import (
	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context){
	var input entity.RegiterInput

	if err:=c.BindJSON(&input);err!=nil{
		newResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	err:=h.services.Authorization.SignUp(input)
	if err!=nil{
		newResponse(c, http.StatusInternalServerError,err.Error())
		return
	}

}

func (h *Handler) signIn ( c * gin.Context){
	
}