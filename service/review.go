package service

import (
	"fmt"

	"codestates.wba-01/archoi/backend/oos/model"
)

func (srv *Service) CreateReview(review model.Review) error {
	// 이미 리뷰가 존재하는지 체크
	if _, err := srv.md.ReviewModel.FindReviewByOrderSeqAndMenuName(review.OrderSeq, review.MenuName); err == nil {
		return fmt.Errorf("Review already exists")
	}
	// 주문번호에 해당하는 주문이 존재하는지 체크
	order, err := srv.md.OrderModel.FindOrderBySeq(review.OrderSeq)
	if err != nil {
		return err
	}
	// 주문이 완료된 상태인지 체크
	if order.Status != model.ORDER_STATUS_COMPLETED {
		return fmt.Errorf("Reviews can only be written on completed orders")
	}
	// 주문에 해당 메뉴가 존재하는지 체크
	if order.IsContainMenu(review.MenuName) == false {
		return fmt.Errorf("Not found Menu in Order[%s]", order.Seq)
	}
	// 리뷰 저장
	if err := srv.md.ReviewModel.CreateReview(review); err != nil {
		return err
	}
	// 해당 메뉴의 평점 업데이트
	avgScore, err := srv.md.ReviewModel.GetAvgScore(review.MenuName)
	if err != nil {
		return err
	}
	menu, err := srv.md.MenuModel.FindMenuByName(review.MenuName, false, false)
	if err != nil {
		return err
	}
	menu.AvgScore = avgScore
	if err := srv.md.MenuModel.UpdateMenuByName(menu.Name, model.MenuForUpdate(menu)); err != nil {
		return err
	}
	return nil
}

func (srv *Service) GetReviewList(menuName string) ([]model.Review, error) {
	reviewList, err := srv.md.ReviewModel.FindReviewListByMenu(menuName)
	if err != nil {
		return reviewList, err
	}

	return reviewList, nil
}
