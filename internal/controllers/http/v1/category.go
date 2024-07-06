package v1

import (
	"net/http"

	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCategory(c *gin.Context){	
	var input entity.Category

	if err:=c.BindJSON(input);err!=nil{
		newResponse(c,http.StatusBadRequest,err.Error())
		return

	id,err:=h.services.Category.CreateCategory(input)
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK, IDResponse{ID: id})
	
}

}
func (h *Handler) getAllCategory(c *gin.Context){
	
}
func (h *Handler) getCategoryByID(c *gin.Context){
	
}
func (h *Handler) updateCategory(c *gin.Context){
	
}
func (h *Handler) deleteCategory(c *gin.Context){
	
}
