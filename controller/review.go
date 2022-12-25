package controller

import (
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
	if err := ctl.srv.CreateReview(review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}

func (ctl *Controller) GetReviewList(c *gin.Context) {
	menuName := c.Param(GET_REVIEW_LIST_PARAM_MENU_NAME)
	reviewList, err := ctl.srv.GetReviewList(menuName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviewList)
}
