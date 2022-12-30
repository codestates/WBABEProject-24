package service

import (
	"fmt"

	"codestates.wba-01/archoi/backend/oos/model"
)

const (
	CHANGE_ORDER_TYPE_ADD    = "add"
	CHANGE_ORDER_TYPE_CHANGE = "change"
)

func (srv *Service) CreateOrder(order model.Order) (string, error) {
	// 메뉴 체크
	if err := srv.md.MenuModel.CanOrder(order.MenuList); err != nil {
		return "", err
	}
	// 주문 저장
	seq, err := srv.md.OrderModel.CreateOrder(order)
	if err != nil {
		return "", err
	}
	return seq, nil
}

func (srv *Service) GetOrderList(status string) ([]model.Order, error) {
	// 주문 상태(진행/완료)에 따른 주문 리스트 얻기
	orderList, err := srv.md.OrderModel.FindOrderListInStatus(status)
	if err != nil {
		return nil, err
	}
	return orderList, nil
}

func (srv *Service) ChangeOrderMenu(orderSeq, changeType string, menuListUpdate model.OrderMenuList) (string, error) {
	// seq에 해당하는 주문 찾기
	order, err := srv.md.OrderModel.FindOrderBySeq(orderSeq)
	if err != nil {
		return "", err
	}
	// 변경 타입 체크
	switch changeType {
	case CHANGE_ORDER_TYPE_ADD:
		order.MenuList = append(order.MenuList, menuListUpdate.MenuList...)
	case CHANGE_ORDER_TYPE_CHANGE:
		order.MenuList = menuListUpdate.MenuList
	default:
		return "", fmt.Errorf("Invalid type of change")
	}
	// 주문 변경 가능 체크
	if order.IsChangeable() == false {
		if changeType == CHANGE_ORDER_TYPE_CHANGE {
			// 주문 변경 불가
			return "", fmt.Errorf("Order Not changeable current status: [%s]", order.Status)
		} else if changeType == CHANGE_ORDER_TYPE_ADD {
			// 신규 주문으로 처리
			if err := srv.md.MenuModel.CanOrder(menuListUpdate.MenuList); err != nil {
				return "", err
			}
			newSeq, err := srv.md.OrderModel.CreateOrder(order)
			if err != nil {
				return "", err
			}
			return newSeq, nil
		}
	}
	// 주문 업데이트
	if err := srv.md.OrderModel.UpdateOrderBySeq(order.Seq, order); err != nil {
		return "", err
	}
	return order.Seq, nil
}

func (srv *Service) ChangeOrderStatus(seq, status string) error {
	// seq에 해당하는 주문 찾기
	order, err := srv.md.OrderModel.FindOrderBySeq(seq)
	if err != nil {
		return err
	}
	// 주문 상태 변경
	switch status {
	case model.ORDER_STATUS_COOKING:
	case model.ORDER_STATUS_DELIVERY:
	case model.ORDER_STATUS_RECEIPT:
	case model.ORDER_STATUS_WAITING:
	case model.ORDER_STATUS_COMPLETED:
		// 메뉴의 OrderCount 증가
		for _, name := range order.MenuList {
			srv.md.MenuModel.IncreaseOrderCount(name)
		}
	default:
		return fmt.Errorf("Invalid Order status")
	}
	order.Status = status
	// 주문 업데이트
	if err := srv.md.OrderModel.UpdateOrderBySeq(order.Seq, order); err != nil {
		return err
	}
	return nil
}
