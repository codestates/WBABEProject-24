package controller

import (
	"codestates.wba-01/archoi/backend/oos/service"
)

type Controller struct {
	srv *service.Service
}

func NewCTL(rep *service.Service) (*Controller, error) {
	r := &Controller{srv: rep}
	return r, nil
}
