package v1

import (
	"github.com/DanilMankiev/SofiaApplication/internal/service"
	"github.com/gin-gonic/gin"

)
type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) NewRouter(router *gin.Engine){
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	auth:=router.Group("/auth")
	{
		auth.POST("/sign-up",h.signUp)
		auth.POST("/sign-in",h.signIn)

	}
	api:=router.Group("/api/v1")
	{
		category:=api.Group("/category")
		{
			category.POST("/", h.createCategory)
			category.GET("/", h.getAllCategory)
			
			categoryID:=category.Group("/:id")
			{
				categoryID.GET("", h.getCategoryByID)
				categoryID.PUT("",h.updateCategory)
				categoryID.DELETE("", h.deleteCategory)
			}
		}	
	}
	
}