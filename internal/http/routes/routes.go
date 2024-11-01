package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"book-store/internal/models"
)

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type apiRoutes struct {
	router   *gin.RouterGroup
	dbClient *gorm.DB
}

func NewRoute(router *gin.RouterGroup, dbClient *gorm.DB) *apiRoutes {
	return &apiRoutes{router: router, dbClient: dbClient}
}

func (r *apiRoutes) Routes() {
	V1Route := r.router.Group("v1")
	{
		V1BookRoute := V1Route.Group("book")
		V1BookRoute.GET("/books", func(ctx *gin.Context) {

			var books []models.Book
			r.dbClient.Find(&books)

			ctx.JSON(http.StatusOK, BaseResponse{
				Status:  true,
				Message: "Ok",
				Data:    books,
			})
		})
	}
}
