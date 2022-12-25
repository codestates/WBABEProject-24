package controller

import (
	"fmt"
	"net/http"

	"codestates.wba-01/archoi/backend/oos/model"
	"github.com/gin-gonic/gin"
)

const (
	GET_REVIEW_LIST_PARAM_MENU_NAME = "menu"
)

func (ctl *Controller) CreateReview(c *gin.Context) {
	var review model.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 이미 리뷰가 존재하는지 체크
	if _, err := ctl.md.ReviewModel.FindReviewByOrderSeqAndMenuName(review.OrderSeq, review.MenuName); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Review already exists"})
		return
	}
	// 주문번호에 해당하는 주문이 존재하는지 체크
	order, err := ctl.md.OrderModel.FindOrderBySeq(review.OrderSeq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 주문이 완료된 상태인지 체크
	if order.Status != model.ORDER_STATUS_COMPLETED {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Reviews can only be written on completed orders"})
		return
	}
	// 주문에 해당 메뉴가 존재하는지 체크
	if order.IsContainMenu(review.MenuName) == false {
		c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("Not found Menu in Order[%s]", order.Seq)})
		return
	}
	// 리뷰 저장
	if err := ctl.md.ReviewModel.CreateReview(review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 해당 메뉴의 평점 업데이트
	avgScore, err := ctl.md.ReviewModel.GetAvgScore(review.MenuName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	menu, err := ctl.md.MenuModel.FindMenuByName(review.MenuName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	menu.AvgScore = avgScore
	if err := ctl.md.MenuModel.UpdateMenuByName(menu.Name, menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}

func (ctl *Controller) GetReviewList(c *gin.Context) {
	menuName := c.Param(GET_REVIEW_LIST_PARAM_MENU_NAME)
	reviewList, err := ctl.md.ReviewModel.FindReviewListByMenu(menuName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviewList)
}
