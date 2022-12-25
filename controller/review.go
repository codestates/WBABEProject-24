package controller

import (
	"net/http"

	"codestates.wba-01/archoi/backend/oos/model"
	"github.com/gin-gonic/gin"
)

const (
	GET_REVIEW_LIST_PARAM_MENU_NAME = "menu"
)

// CreateReview godoc
// @Summary call CreateReview, Create the Review object and return result message by string.
// @Description 리뷰 객체를 생성하기 위한 기능
// @name CreateReview
// @Accept json
// @Produce json
// @Param review body model.Review true "Review data"
// @Router /orderer/review [post]
// @Success 200 {object} controller.SuccessResultJSON
// @failure 400 {object} controller.ErrorResultJSON
func (ctl *Controller) CreateReview(c *gin.Context) {
	var review model.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	if err := ctl.srv.CreateReview(review); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResultJSON{Message: "success"})
	return
}

// GetReviewList godoc
// @Summary call GetReviewList, return the List of reviews corresponding to the menu.
// @Description 메뉴에 해당하는 리뷰 리스트를 반환하기 위한 기능
// @name GetReviewList
// @Produce json
// @Param menu path string true "Menu name"
// @Router /orderer/review/list/{menu} [get]
// @Success 200 {object} controller.SuccessResultJSON{data=[]model.Review} "data: 메뉴에 해당하는 리뷰 리스트"
// @failure 400 {object} controller.ErrorResultJSON
func (ctl *Controller) GetReviewList(c *gin.Context) {
	menuName := c.Param(GET_REVIEW_LIST_PARAM_MENU_NAME)
	reviewList, err := ctl.srv.GetReviewList(menuName)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResultJSON{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResultJSON{Message: "success", Data: reviewList})
}
