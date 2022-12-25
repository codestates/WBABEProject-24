package service

import (
	"codestates.wba-01/archoi/backend/oos/model"
)

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
