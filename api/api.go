package api

import (
	_ "crud/api/docs"
	"crud/api/handler"
	"crud/config"
	"crud/pkg/helper"
	"crud/storage"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(cfg *config.Config, r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandlerV1(cfg, storage)

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	v1 := r.Group("v1")

	r.Use(customCORSMiddleware())
	v1.Use(SecurityMiddleware())

	r.POST("/user", handlerV1.CreateUser)
	r.GET("/user/:id", handlerV1.GetUserById)
	r.GET("/user", handlerV1.GetUserList)
	r.PUT("/user/:id", handlerV1.UpdateUser)
	r.DELETE("/user/:id", handlerV1.DeleteUser)

	r.POST("/subscription", handlerV1.CreateSubscription)
	r.GET("/subscription/:id", handlerV1.GetSubscriptionById)
	r.GET("/subscription", handlerV1.GetSubscriptionList)
	r.PUT("/subscription/:id", handlerV1.UpdateSubscription)
	r.DELETE("/subscription/:id", handlerV1.DeleteSubscription)

	r.POST("/user-subscription", handlerV1.CreateUserSubscription)
	r.GET("/user-subscription/:id", handlerV1.GetUserSubscriptionById)
	r.GET("/user-subscription", handlerV1.GetUserSubscriptionList)
	r.GET("/user-id-subscription/:id", handlerV1.GetUserSubscriptionByUserId)
	r.GET("/has-acces/:id", handlerV1.HasAccess)
	r.GET("/make-access/:id", handlerV1.MakeActive)
	// r.PUT("/user-subscription/:id", handlerV1.UpdateSubscription)
	// r.DELETE("/usersubscription/:id", handlerV1.DeleteSubscription)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var cfg config.Config

		if len(c.Request.Header["Authorization"]) > 0 {
			token := c.Request.Header["Authorization"][0]
			_, err := helper.ParseClaims(token, cfg.AuthSecretKey)

			if err != nil {
				c.JSON(http.StatusUnauthorized, struct {
					Code int
					Err  string
				}{
					Code: http.StatusUnauthorized,
					Err:  errors.New("error access denied").Error(),
				})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, struct {
				Code int
				Err  string
			}{
				Code: http.StatusUnauthorized,
				Err:  errors.New("error access denied").Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
