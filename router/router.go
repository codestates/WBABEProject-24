package router

import (
	ctl "codestates.wba-01/archoi/backend/oos/controller"
	"codestates.wba-01/archoi/backend/oos/logger"

	"github.com/gin-gonic/gin"
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
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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

func (p *Router) Idx() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.Use(logger.GinLogger())
	e.Use(logger.GinRecovery(true))
	e.Use(CORS())

	menuGroup := e.Group("/menu")
	{
		menuGroup.GET("/list", p.ct.GetMenuList)
		menuGroup.POST("", p.ct.CreateMenu)
		menuGroup.PUT("/:name", p.ct.UpdateMenuByName)
		menuGroup.DELETE("/:name", p.ct.DeleteMenuByName)
	}

	return e
}
