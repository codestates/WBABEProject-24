package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client    *mongo.Client
	MenuModel *menuModel
	orderCol  *mongo.Collection
	reviewCol *mongo.Collection
}

type Order struct {
	Num      int    `json:"num" bson:"num"`
	Date     string `json:"date" bson:"date"`
	MenuName string `json:"menuName" bson:"menuName"`
	Address  string `json:"address" bson:"address"`
}

type Review struct {
	MenuName string `json:"menuName" bson:"menuName"`
	Score    int    `json:"score" bson:"score"`
	Comment  string `json:"comment" bson:"comment"`
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
		r.orderCol = db.Collection("order")
		r.MenuModel = NewMenuModel(db.Collection("menu"))
		r.reviewCol = db.Collection("review")
	}

	return r, nil
}
