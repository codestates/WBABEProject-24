package controller

import (
	"fmt"
	"net/http"

	"codestates.wba-01/archoi/backend/oos/model"
	"github.com/gin-gonic/gin"
)

const (
	GET_ORDER_LIST_QUERY_STATUS         = "status"
	CHANGE_ORDER_MENU_PARAM_SEQ         = "seq"
	CHANGE_ORDER_MENU_PARAM_TYPE        = "type"
	CHANGE_ORDER_MENU_PARAM_TYPE_ADD    = "add"
	CHANGE_ORDER_MENU_PARAM_TYPE_CHANGE = "change"
	CHANGE_ORDER_MENU_PARAM_STATUS      = "status"
)

func (ctl *Controller) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 메뉴 체크
	if err := ctl.md.MenuModel.IsOrderable(order.MenuList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 주문 저장
	seq, err := ctl.md.OrderModel.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "orderSeq": seq})
	return
}

func (ctl *Controller) GetOrderList(c *gin.Context) {
	status := c.Query(GET_ORDER_LIST_QUERY_STATUS)
	// 주문 상태(진행/완료)에 따른 주문 리스트 얻기
	orderList, err := ctl.md.OrderModel.FindOrderListInStatus(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderList)
}

func (ctl *Controller) ChangeOrderMenu(c *gin.Context) {
	seq := c.Param(CHANGE_ORDER_MENU_PARAM_SEQ)
	changeType := c.Param(CHANGE_ORDER_MENU_PARAM_TYPE)
	// seq에 해당하는 주문 찾기
	order, err := ctl.md.OrderModel.FindOrderBySeq(seq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// JSON으로 전달된 menuList 저장
	menuListUpdate := model.OrderMenuList{}
	if err := c.ShouldBindJSON(&menuListUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 변경 타입 체크
	switch changeType {
	case CHANGE_ORDER_MENU_PARAM_TYPE_ADD:
		order.MenuList = append(order.MenuList, menuListUpdate.MenuList...)
	case CHANGE_ORDER_MENU_PARAM_TYPE_CHANGE:
		order.MenuList = menuListUpdate.MenuList
	default:
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid type of change"})
		return
	}
	// 주문 변경 가능 체크
	if order.IsChangeable() == false {
		if changeType == CHANGE_ORDER_MENU_PARAM_TYPE_CHANGE {
			// 주문 변경 불가
			c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("Order Not changeale current status: [%s]", order.Status)})
			return
		} else if changeType == CHANGE_ORDER_MENU_PARAM_TYPE_ADD {
			// 신규 주문으로 처리
			if err := ctl.md.MenuModel.IsOrderable(menuListUpdate.MenuList); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
				return
			}
			newSeq, err := ctl.md.OrderModel.CreateOrder(order)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"msg": "success", "orderSeq": newSeq})
			return
		}
	}
	// 주문 업데이트
	if err := ctl.md.OrderModel.UpdateOrderBySeq(order.Seq, order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}

func (ctl *Controller) ChangeOrderStatus(c *gin.Context) {
	seq := c.Param(CHANGE_ORDER_MENU_PARAM_SEQ)
	status := c.Param(CHANGE_ORDER_MENU_PARAM_STATUS)
	// seq에 해당하는 주문 찾기
	order, err := ctl.md.OrderModel.FindOrderBySeq(seq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	order.Status = status
	// 주문 업데이트
	if err := ctl.md.OrderModel.UpdateOrderBySeq(order.Seq, order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}
