package v1

import (
	"net/http"
	"strconv"

	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCategory(c *gin.Context){	
	var input entity.Category

	if err:=c.BindJSON(&input);err!=nil{
		newResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	id,err:=h.services.Category.CreateCategory(input)
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK, IDResponse{ID: id})
}

func (h *Handler) getAllCategory(c *gin.Context){
	
	categories,err:= h.services.Category.GetAllCategorys()
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}
func (h *Handler) getCategoryByID(c *gin.Context){

	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	category,err:=h.services.Category.GetCategoryById(id)
	if err!=nil{
		newResponse(c, http.StatusInternalServerError,err.Error())
		return
	}
	
	c.JSON(http.StatusOK, category)
	
}
func (h *Handler) updateCategory(c *gin.Context){
	
}
func (h *Handler) deleteCategory(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newResponse(c,http.StatusBadGateway,err.Error())
		return
	}
	err=h.services.Category.DeleteCategory(id)
	if err!=nil{
		newResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,StatusResponse{Status: "ok"})
}
