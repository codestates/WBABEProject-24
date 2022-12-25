package model

import (
	"context"
	"fmt"
	"time"

	"codestates.wba-01/archoi/backend/oos/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 매움 정도 상수 정의
type HotGrade int

const (
	MildHot HotGrade = 1 + iota
	SlightHot
	MediumHot
	VeryHot
	ExtreamHot
)

type MenuOrderBy struct {
	target string // 정렬 기준 필드 이름
	sort   int    // 1: accending, -1: descending
}

type Menu struct {
	Name        string             `json:"name" bson:"name" uri:"name" binding:"required"`
	Price       int                `json:"price" bson:"price" binding:"required"`
	HotGrade    int                `json:"hotGrade" bson:"hotGrade" binding:"required"`
	IsAvailable *bool              `json:"isAvailable" bson:"isAvailable" binding:"required"`
	IsRecommend *bool              `json:"isRecommend" bson:"isRecommend"`
	IsDeleted   *bool              `json:"isDeleted" bson:"isDeleted"`
	AvgScore    float32            `json:"avgScore" bson:"avgScore"`
	OrderCount  int                `json:"orderCount" bson:"orderCount"`
	CreateDate  primitive.DateTime `json:"createDate" bson:"createDate"`
}

type menuModel struct {
	col        *mongo.Collection
	orderByMap map[string]MenuOrderBy
}

func NewMenuModel(col *mongo.Collection) *menuModel {
	m := new(menuModel)
	m.col = col
	m.orderByMap = make(map[string]MenuOrderBy)
	m.orderByMap["recommend"] = MenuOrderBy{target: "isRecommend", sort: -1}
	m.orderByMap["score"] = MenuOrderBy{target: "avgScore", sort: -1}
	m.orderByMap["mostOrder"] = MenuOrderBy{target: "orderCount", sort: -1}
	m.orderByMap["new"] = MenuOrderBy{target: "createDate", sort: -1}
	return m
}

func (p *menuModel) IsOrderable(menuameList []string) error {
	for _, name := range menuameList {
		menu, err := p.FindMenuByName(name)
		// 존재하는 메뉴인지 체크
		if err != nil {
			return err
		}
		// 주문 가능한 메뉴인지 체크
		if *menu.IsAvailable == false || *menu.IsDeleted == true {
			return fmt.Errorf("Menu %s Not Available", menu.Name)
		}
	}
	return nil
}

func (p *menuModel) CreateMenu(menu Menu) error {
	if menu.Name == "" {
		return fmt.Errorf("Require Menu name")
	}
	// 이미 존재하는 메뉴인지 체크
	if _, err := p.FindMenuByName(menu.Name); err == nil {
		return fmt.Errorf("Menu Already exist")
	}
	// 메뉴 기본값 초기화
	menu.AvgScore = 0
	menu.OrderCount = 0
	menu.IsRecommend = util.NewFalse()
	menu.IsDeleted = util.NewFalse()
	menu.CreateDate = primitive.NewDateTimeFromTime(time.Now())
	// 메뉴 저장
	_, err := p.col.InsertOne(context.TODO(), menu)
	if err != nil {
		return err
	}
	return nil
}

func (p *menuModel) FindMenuByName(name string) (Menu, error) {
	var result Menu
	filter := bson.D{{"name", name}}
	err := p.col.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return result, fmt.Errorf("No Menu found given name: %s\n", name)
	} else if err != nil {
		return result, err
	}
	return result, nil
}

func (p *menuModel) FindMenuIsDeletedSortBy(exceptDeleted bool, sortBy string) ([]Menu, error) {
	var results []Menu
	filter := bson.D{}
	// 필터 설정 : 삭제된 메뉴 제외
	if exceptDeleted == true {
		filter = bson.D{{"isDeleted", false}}
	}
	// 정렬 옵션 설정
	var opts *options.FindOptions
	if v, ok := p.orderByMap[sortBy]; ok {
		opts = options.Find().SetSort(bson.D{{v.target, v.sort}})
	}
	cur, err := p.col.Find(context.TODO(), filter, opts)
	if err != nil {
		return results, err
	}
	if err = cur.All(context.TODO(), &results); err != nil {
		return results, err
	}
	return results, nil
}

func (p *menuModel) UpdateMenuByName(name string, menu Menu) error {
	if name != menu.Name {
		return fmt.Errorf("Menu name Can not be changed")
	}
	filter := bson.D{{"name", name}}
	updateResult, err := p.col.UpdateOne(context.TODO(), filter, bson.D{{"$set", menu}})
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("No Menu Updated")
	}
	return nil
}

func (p *menuModel) DeleteMenuByName(name string) error {
	menuForDelete, err := p.FindMenuByName(name)
	if err != nil {
		return err
	}
	if *menuForDelete.IsDeleted == true {
		return fmt.Errorf("Already deleted Menu")
	}
	// 삭제 flag를 true로 설정 후, 업데이트
	menuForDelete.IsDeleted = util.NewTrue()
	if err := p.UpdateMenuByName(name, menuForDelete); err != nil {
		return err
	}
	return nil
}
