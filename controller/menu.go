package controller

import (
	"net/http"

	"codestates.wba-01/archoi/backend/oos/model"
	"github.com/gin-gonic/gin"
)

const (
	UPDATE_MENU_PARAM_NAME   = "name"
	DELETE_MENU_PARAM_NAME   = "name"
	GET_MENU_LIST_QUERY_SORT = "sort"
)

// CreateMenu godoc
// @Summary call CreateMenu, Create the Menu object and return result message by string.
// @Description Menu 객체를 생성하기 위한 기능
// @name CreateMenu
// @Accept json
// @Produce json
// @Param menu body model.Menu true "Menu data"
// @Router /recipant/menu [post]
// @Success 200 {object} controller.SuccessResultJSON
// @failure 400 {object} controller.ErrorResultJSON
func (ctl *Controller) CreateMenu(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	if err := ctl.srv.CreateMenu(menu); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, SuccessResultJSON{Message: "success"})
	return
}

// UpdateMenu godoc
// @Summary call UpdateMenu, Update the Menu object and return result message by string.
// @Description Menu 객체를 업데이트하기 위한 기능
// @name UpdateMenu
// @Accept json
// @Produce json
// @Param name path string true "Menu name for update"
// @Param menu body model.Menu true "Menu data for update"
// @Router /recipant/menu/{name} [put]
// @Success 200 {object} controller.SuccessResultJSON
// @failure 400 {object} controller.ErrorResultJSON
func (ctl *Controller) UpdateMenu(c *gin.Context) {
	menuName := c.Param(UPDATE_MENU_PARAM_NAME)
	menu, err := ctl.srv.FindMenuByName(menuName)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	updateForMenu := model.MenuForUpdate(menu)
	if err := c.ShouldBindJSON(&updateForMenu); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	if err := ctl.srv.UpdateMenuByName(menuName, updateForMenu); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResultJSON{Message: "success"})
	return
}

// DeleteMenu godoc
// @Summary call DeleteMenu, Delete the Menu object and return result message by string.
// @Description Menu 객체를 삭제하기 위한 기능
// @name DeleteMenu
// @Produce json
// @Param name path string true "Menu name for delete"
// @Router /recipant/menu/{name} [delete]
// @Success 200 {object} controller.SuccessResultJSON
// @failure 400 {object} controller.ErrorResultJSON
func (ctl *Controller) DeleteMenu(c *gin.Context) {
	menuName := c.Param(DELETE_MENU_PARAM_NAME)
	if err := ctl.srv.DeleteMenuByName(menuName); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResultJSON{Message: "success"})
	return
}

// GetMenuList godoc
// @Summary call GetMenuList, return the Menu object list.
// @Description Menu 객체 리스트를 반환하기 위한 기능
// @name GetMenuList
// @Produce json
// @Param sort query string false "sort type [recommend|score|most|new]"
// @Router /orderer/menu/list [get]
// @Success 200 {object} controller.SuccessResultJSON{data=[]model.Menu} "data: 메뉴 리스트"
// @failure 400 {object} controller.ErrorResultJSON
func (ctl *Controller) GetMenuList(c *gin.Context) {
	sort := c.Query(GET_MENU_LIST_QUERY_SORT)
	menuList, err := ctl.srv.GetMenuListSortBy(sort)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResultJSON{Message: "success", Data: menuList})
	return
}
