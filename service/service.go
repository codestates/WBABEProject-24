package service

import "codestates.wba-01/archoi/backend/oos/model"

type Service struct {
	md *model.Model
}

func NewSRV(rep *model.Model) (*Service, error) {
	r := &Service{md: rep}
	return r, nil
}
