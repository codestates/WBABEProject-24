package service

import (
	"codestates.wba-01/archoi/backend/oos/model"
)

/*
특별히 서비스 레이어를 분리하신 이유가 있을까요?
현재 코드만 살펴본다면 모델의 함수를 호출하는 것 이외에는 별다른 로직을 수행하고 있지 않습니다.
서비스 레이어를 도입하신 이유에 대해서 한번 더 공부해보시면 좋을 것 같습니다.
*/
func (srv *Service) CreateMenu(menu model.Menu) error {
	if err := srv.md.MenuModel.CreateMenu(menu); err != nil {
		return err
	}
	return nil
}

func (srv *Service) FindMenuByName(name string) (model.Menu, error) {
	model, err := srv.md.MenuModel.FindMenuByName(name)
	if err != nil {
		return model, err
	}
	return model, nil
}

func (srv *Service) UpdateMenuByName(name string, updateForMenu model.Menu) error {
	if err := srv.md.MenuModel.UpdateMenuByName(name, updateForMenu); err != nil {
		return err
	}
	return nil
}

func (srv *Service) DeleteMenuByName(name string) error {
	if err := srv.md.MenuModel.DeleteMenuByName(name); err != nil {
		return err
	}
	return nil
}

func (srv *Service) GetMenuIsDeletedFalseSortBy(sort string) ([]model.Menu, error) {
	menuList, err := srv.md.MenuModel.FindMenuIsDeletedSortBy(true, sort)
	if err != nil {
		return nil, err
	}
	return menuList, nil
}
