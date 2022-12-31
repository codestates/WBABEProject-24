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

// CreateOrder godoc
// @Summary call CreateOrder, Create the Order object and return result message by string.
// @Description 주문 객체를 생성하기 위한 기능
// @name CreateOrder
// @Accept json
// @Produce json
// @Param order body model.Order true "Order data"
// @Router /v1/orders [post]
// @Success 200 {object} controller.SuccessResultJSON{data=string} "data: 생성된 주문 일련번호"
// @failure 400 {object} controller.ErrorResultJSON
func (ctl *Controller) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	seq, err := ctl.srv.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, SuccessResultJSON{Message: "success", Data: seq})
	return
}

// GetOrderList godoc
// @Summary call GetOrderList, return the Order object list.
// @Description 주문 상태에 해당하는 리스트를 반환하는 기능
// @name GetOrderList
// @Produce json
// @Param status query string true "Order status [active|complete|all]"
// @Router /v1/orders [get]
// @Success 200 {object} controller.SuccessResultJSON{data=[]model.Order} "data: status에 해당하는 주문 리스트"
// @failure 400 {object} controller.ErrorResultJSON
func (ctl *Controller) GetOrderList(c *gin.Context) {
	status := c.Query(GET_ORDER_LIST_QUERY_STATUS)
	orderList, err := ctl.srv.GetOrderList(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResultJSON{Message: "success", Data: orderList})
}

// ChangeOrderMenu godoc
// @Summary call ChangeOrderMenu, Change the Order menu and return result message by string.
// @Description 주문의 메뉴를 변경하기 위한 기능
// @name ChangeOrderMenu
// @Accept json
// @Produce json
// @Param menuList body model.OrderMenuList true "Menu name list"
// @Param seq path string true "Order sequence number"
// @Param type path string true "Order change type [add|change]"
// @Router /v1/orders/menu/{seq}/{type} [put]
// @Success 200 {object} controller.SuccessResultJSON
// @Success 200 {object} controller.SuccessResultJSON{data=string} "새로운 주문 일련번호"
// @failure 400 {object} controller.ErrorResultJSON
func (ctl *Controller) ChangeOrderMenu(c *gin.Context) {
	seq := c.Param(CHANGE_ORDER_MENU_PARAM_SEQ)
	changeType := c.Param(CHANGE_ORDER_MENU_PARAM_TYPE)
	menuListUpdate := model.OrderMenuList{}
	if err := c.ShouldBindJSON(&menuListUpdate); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	// 주문 메뉴 변경
	retSeq, err := ctl.srv.ChangeOrderMenu(seq, changeType, menuListUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	// 새로운 주문으로 처리된 경우
	if retSeq != seq {
		c.JSON(http.StatusOK, SuccessResultJSON{Message: "success", Data: retSeq})
		return
	}
	c.JSON(http.StatusOK, SuccessResultJSON{Message: "success"})
	return
}

// ChangeOrderStatus godoc
// @Summary call ChangeOrderStatus, Change the Order status and return result message by string.
// @Description 주문의 상태를 변경하기 위한 기능
// @name ChangeOrderStatus
// @Produce json
// @Param seq path string true "Order sequence number"
// @Param status path string true "status value to change [대기|주문|조리|배달|완료]"
// @Router /v1/orders/status/{seq}/{status} [put]
// @Success 200 {object} controller.SuccessResultJSON
// @failure 400 {object} controller.ErrorResultJSON
func (ctl *Controller) ChangeOrderStatus(c *gin.Context) {
	seq := c.Param(CHANGE_ORDER_MENU_PARAM_SEQ)
	status := c.Param(CHANGE_ORDER_MENU_PARAM_STATUS)
	// 주문 업데이트
	if err := ctl.srv.ChangeOrderStatus(seq, status); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResultJSON{Message: "success"})
	return
}
