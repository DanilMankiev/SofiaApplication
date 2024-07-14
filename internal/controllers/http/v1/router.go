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
		auth.POST("/register",h.Register)
		auth.POST("/login",h.signIn)

	}
	api:=router.Group("/api/v1")
	{
		admin:=api.Group("/admin")
		{
			admin.POST("/login")
			
			category:=admin.Group("/category")
			{
				category.POST("/", h.createCategory)

				categoryID:=category.Group("/:id")
				{
					categoryID.PUT("",h.updateCategory)
					categoryID.DELETE("", h.deleteCategory)
				}
			}
		}
		category:=api.Group("/category")
		{
			category.GET("/", h.getAllCategory)
			
			categoryID:=category.Group("/:id")
			{
				categoryID.GET("", h.getCategoryByID)
			}
		}
		// company:=api.Group("/company")
		// {
		// 	company.POST("/",h.createCompany)
		// 	company.GET("/",h.getAllCompany)

		// 	companyID:=company.Group("/id")
		// 	{
		// 		companyID.GET("/", h.getCompanyByID)
		// 		companyID.PUT("/",h.updateCompany)
		// 		companyID.DELETE("/",h.deleteCompany)
		// 	}
		// }	
	}
	
}