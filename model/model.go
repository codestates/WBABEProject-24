package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client      *mongo.Client
	MenuModel   *menuModel
	OrderModel  *orderModel
	ReviewModel *reviewModel
}

func NewModel(mgUrl string) (*Model, error) {
	r := &Model{}

	var err error
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("oos")
		r.OrderModel = NewOrderModel(db.Collection("order"))
		r.MenuModel = NewMenuModel(db.Collection("menu"))
		r.ReviewModel = NewReviewModel(db.Collection("review"))
	}

	return r, nil
}
