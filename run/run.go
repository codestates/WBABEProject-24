package run

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"codestates.wba-01/archoi/backend/oos/conf"
	"codestates.wba-01/archoi/backend/oos/controller"
	"codestates.wba-01/archoi/backend/oos/logger"
	"codestates.wba-01/archoi/backend/oos/model"
	"codestates.wba-01/archoi/backend/oos/router"
	"codestates.wba-01/archoi/backend/oos/service"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func serverShutdown(mapi *http.Server, graceful bool) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Warn("Server Shutdown...")
	shutdown := make(chan bool, 1)
	var ctx context.Context
	var cancel context.CancelFunc
	if graceful {
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
		shutdown <- true
	}
	defer cancel()

	if err := mapi.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown Error:", err)
	}
	select {
	case <-ctx.Done():
		logger.Info("Server Shutdown timeout of 5 seconds")
	case <-shutdown:
		logger.Info("Server Shutdown immediately")
	}
	logger.Info("Server Shutdown Complete!")
}

func Run() {
	// config 파일 경로 옵션 설정
	var configFlag = flag.String("config", "./conf/config.toml", "toml file to use for configuration")
	// 실행 옵션 파싱
	flag.Parse()
	// config 초기화
	cf := conf.NewConfig(*configFlag)

	// 로그 초기화
	if err := logger.InitLogger(cf); err != nil {
		fmt.Printf("logger.InitLogger error: %v\n", err)
		return
	}
	logger.Debug("Server Ready...")

	// MVC 초기화
	if md, err := model.NewModel(cf.DB.Host); err != nil {
		// 모델 초기화
		logger.Error("model.NewModel", err)
		panic(fmt.Errorf("model.NewModel error: %v", err))
	} else if srv, err := service.NewSRV(md); err != nil {
		logger.Error("service.NewSRV", err)
		panic(fmt.Errorf("service.NewSRV error: %v", err))
	} else if ctl, err := controller.NewCTL(srv); err != nil {
		// 컨트롤러 초기화
		logger.Error("controller.NewCTL", err)
		panic(fmt.Errorf("controller.NewCTL error: %v", err))
	} else if rt, err := router.NewRouter(ctl); err != nil {
		// 라우터 초기화
		logger.Error("router.NewRouter", err)
		panic(fmt.Errorf("router.NewRouter error: %v", err))
	} else {
		// 웹서버 설정
		mapi := &http.Server{
			Addr:           cf.Web.Port,
			Handler:        rt.Idx(cf.Swagger.Host),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		// 웹서버 실행
		g.Go(func() error {
			return mapi.ListenAndServe()
		})
		logger.Info("Server Start...!")

		// 서버 종료 (우아한 종료: graceful를 true로 설정)
		serverShutdown(mapi, false)
	}
	// 종료 대기
	if err := g.Wait(); err != nil {
		logger.Error(err)
	}
}
