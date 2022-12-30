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
		/*
			CORS 허용을 위해서 모든 도메인을 허용한다면 보안에 이슈가 생깁니다.
			보통 운영되는 시스템의 경우는 특정한 도메인만을 허용하고 그 이외의 요청은 거부하도록 설정합니다.
		*/
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

func (p *Router) Idx(swaggerHost string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.Use(logger.GinLogger())
	e.Use(logger.GinRecovery(true))
	e.Use(CORS())
	e.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = swaggerHost

	/*
		엔드포인트 구성에 대해서 전반적인 코멘트 드립니다.
		1. REST API의 성숙도 모델에 대해서 공부해보시면 좋을 것 같습니다.

		2. 일반적으로 HTTP URI에 new, modify 와 같은 행위는 들어가지 않습니다.
			복수형의 단어로 구성을 하고, 동일한 URI 내에서 http method만 변경하여 행위를 표현하는 것이 일반적인 REST API의 구성 방식입니다.
			e.g.
			GET v1/menus -> 메뉴 목록을 조회.
			GET v1/menus/1 -> 1번 메뉴를 조회.
			POST v1/menus -> 메뉴를 생성.
			PATCH v1/menus/1 -> 1번 메뉴에 대해서 업데이트
			DELETE v1/menus/1 -> 1번 메뉴에 대해서 삭제
	*/

	/*
		Group을 사용해 연관된 것끼리 묶어주신 점 좋습니다. 이렇게 구성한다면 API가 확장되어도 어느정도 가독성을 지킬 수 있어 보입니다.
	*/
	recipantGroup := e.Group("/recipant")
	{
		menuGroup := recipantGroup.Group("/menu")
		{
			menuGroup.POST("", p.ct.CreateMenu)
			menuGroup.PUT("/:name", p.ct.UpdateMenu)
			menuGroup.DELETE("/:name", p.ct.DeleteMenu)
		}
		orderGroup := recipantGroup.Group("/order")
		{
			orderGroup.GET("/list", p.ct.GetOrderList)
			orderGroup.PUT("/:seq/:status", p.ct.ChangeOrderStatus)
		}
	}
	ordererGroup := e.Group("/orderer")
	{
		menuGroup := ordererGroup.Group("/menu")
		{
			menuGroup.GET("/list", p.ct.GetMenuList)
		}
		orderGroup := ordererGroup.Group("/order")
		{
			orderGroup.GET("/list", p.ct.GetOrderList)
			orderGroup.POST("", p.ct.CreateOrder)
			orderGroup.PUT("/:seq/:type", p.ct.ChangeOrderMenu)
		}
		reviewGroup := ordererGroup.Group("/review")
		{
			reviewGroup.GET("/list/:menu", p.ct.GetReviewList)
			reviewGroup.POST("", p.ct.CreateReview)
		}
	}

	return e
}
