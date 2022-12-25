package model

import (
	"context"
	"sync/atomic"

	"codestates.wba-01/archoi/backend/oos/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Review struct {
	ReviewSeq string `json:"reviewNum" bson:"reviewNum"`
	OrderSeq  string `json:"orderSeq" bson:"orderSeq" binding:"required"`
	MenuName  string `json:"menuName" bson:"menuName" binding:"required"`
	Score     int    `json:"score" bson:"score" binding:"required"`
	Comment   string `json:"comment" bson:"comment" binding:"required"`
}

type reviewModel struct {
	col           *mongo.Collection
	reviewCounter uint32
}

func NewReviewModel(col *mongo.Collection) *reviewModel {
	m := new(reviewModel)
	m.col = col
	return m
}

func (p *reviewModel) CreateReview(review Review) error {
	count := atomic.AddUint32(&p.reviewCounter, 1)
	review.ReviewSeq = util.CreateSeqStr(count)
	// 리뷰 저장
	_, err := p.col.InsertOne(context.TODO(), review)
	if err != nil {
		return err
	}
	return nil
}

func (p *reviewModel) FindReviewByOrderSeqAndMenuName(orderSeq, menuName string) (Review, error) {
	var result Review
	// 필터 설정
	filter := bson.D{{"orderSeq", orderSeq}, {"menuName", menuName}}
	err := p.col.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (p *reviewModel) FindReviewListByMenu(menuName string) ([]Review, error) {
	var results []Review
	// 필터 설정
	filter := bson.D{{"menuName", menuName}}
	// 정렬 옵션 설정
	opts := options.Find().SetSort(bson.D{{"ReviewSeq", -1}})
	cur, err := p.col.Find(context.TODO(), filter, opts)
	if err != nil {
		return results, err
	}
	if err = cur.All(context.TODO(), &results); err != nil {
		return results, err
	}
	return results, nil
}

func (p *reviewModel) GetAvgScore(menuName string) (float32, error) {
	match := bson.D{
		{"$match", bson.D{{"menuName", menuName}}},
	}
	group := bson.D{{"$group", bson.D{
		{"_id", "$menuName"},
		{"average_score", bson.D{{"$avg", "$score"}}},
		{"total_score", bson.D{{"$sum", "$score"}}},
	}}}
	cursor, err := p.col.Aggregate(context.TODO(), mongo.Pipeline{match, group})
	if err != nil {
		return 0, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return 0, err
	}
	if len(results) == 0 {
		return 0, nil
	}
	avg := float32(results[0]["average_score"].(float64))
	return avg, nil
}
