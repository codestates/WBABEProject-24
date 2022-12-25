package controller

import (
	"net/http"

	"codestates.wba-01/archoi/backend/oos/model"
	"github.com/gin-gonic/gin"
)

const (
	GET_ORDER_LIST_QUERY_STATUS    = "status"
	CHANGE_ORDER_MENU_PARAM_SEQ    = "seq"
	CHANGE_ORDER_MENU_PARAM_TYPE   = "type"
	CHANGE_ORDER_MENU_PARAM_STATUS = "status"
)

func (ctl *Controller) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	seq, err := ctl.srv.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "orderSeq": seq})
	return
}

func (ctl *Controller) GetOrderList(c *gin.Context) {
	status := c.Query(GET_ORDER_LIST_QUERY_STATUS)
	orderList, err := ctl.srv.GetOrderList(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderList)
}

func (ctl *Controller) ChangeOrderMenu(c *gin.Context) {
	seq := c.Param(CHANGE_ORDER_MENU_PARAM_SEQ)
	changeType := c.Param(CHANGE_ORDER_MENU_PARAM_TYPE)
	menuListUpdate := model.OrderMenuList{}
	if err := c.ShouldBindJSON(&menuListUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 주문 메뉴 변경
	retSeq, err := ctl.srv.ChangeOrderMenu(seq, changeType, menuListUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 새로운 주문으로 처리된 경우
	if retSeq != seq {
		c.JSON(http.StatusOK, gin.H{"msg": "success", "newOrderSeq:": retSeq})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}

func (ctl *Controller) ChangeOrderStatus(c *gin.Context) {
	seq := c.Param(CHANGE_ORDER_MENU_PARAM_SEQ)
	status := c.Param(CHANGE_ORDER_MENU_PARAM_STATUS)
	// 주문 업데이트
	if err := ctl.srv.ChangeOrderStatus(seq, status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}
