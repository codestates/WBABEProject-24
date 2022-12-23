package controller

import (
	"net/http"

	"codestates.wba-01/archoi/backend/oos/model"
	"github.com/gin-gonic/gin"
)

func (ctl *Controller) CreateMenu(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := ctl.md.MenuModel.CreateMenu(menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}

func (ctl *Controller) UpdateMenuByName(c *gin.Context) {
	menuName := model.MenuForUpdate{}
	if err := c.ShouldBindUri(&menuName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	menu, err := ctl.md.MenuModel.FindMenuByName(menuName.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	updateForMenu := menu
	if err := c.ShouldBindJSON(&updateForMenu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := ctl.md.MenuModel.UpdateMenuByName(menuName.Name, model.Menu(updateForMenu)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}

func (ctl *Controller) DeleteMenuByName(c *gin.Context) {
	menuName := model.MenuForUpdate{}
	if err := c.ShouldBindUri(&menuName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := ctl.md.MenuModel.DeleteMenuByName(menuName.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}

func (ctl *Controller) GetMenuIsDeletedFalseOrderBy(c *gin.Context) {
	orderBy := c.Query("orderBy")
	menuList, err := ctl.md.MenuModel.FindMenuIsDeletedFalseOrderBy(orderBy, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menuList)
}
