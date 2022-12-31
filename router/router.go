package router

import (
	ctl "codestates.wba-01/archoi/backend/oos/controller"
	"codestates.wba-01/archoi/backend/oos/logger"

	"codestates.wba-01/archoi/backend/oos/docs"
	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl}
	return r, nil
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort()
			return
		}
		c.Next()
	}
}

func (p *Router) Idx(swaggerHost string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.Use(logger.GinLogger())
	e.Use(logger.GinRecovery(true))
	e.Use(CORS())
	e.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = swaggerHost

	v1 := e.Group("/v1")
	{
		menuGroup := v1.Group("/menus")
		{
			menuGroup.GET("", p.ct.GetMenuList)
			menuGroup.POST("", p.ct.CreateMenu)
			menuGroup.PUT("/:name", p.ct.UpdateMenu)
			menuGroup.DELETE("/:name", p.ct.DeleteMenu)
		}
		orderGroup := v1.Group("/orders")
		{
			orderGroup.GET("", p.ct.GetOrderList)
			orderGroup.POST("", p.ct.CreateOrder)
			orderGroup.PUT("/menu/:seq/:type", p.ct.ChangeOrderMenu)
			orderGroup.PUT("/status/:seq/:status", p.ct.ChangeOrderStatus)
		}
		reviewGroup := v1.Group("/reviews")
		{
			reviewGroup.GET("/:menu", p.ct.GetReviewList)
			reviewGroup.POST("", p.ct.CreateReview)
		}
	}

	return e
}
