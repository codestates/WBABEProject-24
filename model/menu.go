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
	/*
		메뉴의 이름은 충분히 중복될 수 있습니다. 동일한 이름을 가지는 메뉴를 생성하고 싶다면, 어떻게 처리할 수 있을까요?
	*/
	Name        string             `json:"name" bson:"name" uri:"name" binding:"required"`
	Price       int                `json:"price" bson:"price" binding:"required"`
	HotGrade    int                `json:"hotGrade" bson:"hotGrade" binding:"required"`
	IsAvailable *bool              `json:"isAvailable" bson:"isAvailable" binding:"required"`
	IsRecommend *bool              `json:"isRecommend" bson:"isRecommend"`
	IsDeleted   *bool              `json:"isDeleted" bson:"isDeleted"`
	AvgScore    float32            `json:"avgScore" bson:"avgScore"`
	OrderCount  int                `json:"orderCount" bson:"orderCount"`
	CreatedAt   primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt   primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
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
	m.orderByMap["most"] = MenuOrderBy{target: "orderCount", sort: -1}
	m.orderByMap["new"] = MenuOrderBy{target: "createDate", sort: -1}
	return m
}

func (p *menuModel) CanOrder(menuameList []string) error {
	for _, name := range menuameList {
		// 주문 가능한 메뉴인지 체크
		if _, err := p.FindMenuByName(name, true, true); err != nil {
			return err
		}
	}
	return nil
}

func (p *menuModel) CreateMenu(menu Menu) error {
	// 이미 존재하는 메뉴인지 체크
	if _, err := p.FindMenuByName(menu.Name, false, false); err == nil {
		return fmt.Errorf("Menu Already exist")
	}
	// 메뉴 기본값 초기화
	menu.AvgScore = 0
	menu.OrderCount = 0
	menu.IsRecommend = util.NewFalse()
	menu.IsDeleted = util.NewFalse()
	menu.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	menu.UpdatedAt = menu.CreatedAt
	// 메뉴 저장
	_, err := p.col.InsertOne(context.TODO(), menu)
	if err != nil {
		return err
	}
	return nil
}

func (p *menuModel) FindMenuByName(name string, availableOnly, notDeletedOnly bool) (Menu, error) {
	var result Menu
	filter := bson.D{{"name", name}}
	err := p.col.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return result, fmt.Errorf("No Menu found given name: %s\n", name)
	} else if err != nil {
		return result, err
	}
	if notDeletedOnly && *result.IsDeleted == true {
		return result, fmt.Errorf("Deleted Menu")
	}
	if availableOnly && *result.IsAvailable == false {
		return result, fmt.Errorf("Menu Not available")
	}

	return result, nil
}

func (p *menuModel) FindMenuListSortBy(sortBy string) ([]Menu, error) {
	var results []Menu
	// 정렬 옵션 설정
	var opts *options.FindOptions
	if v, ok := p.orderByMap[sortBy]; ok {
		opts = options.Find().SetSort(bson.D{{v.target, v.sort}})
	} else {
		return results, fmt.Errorf("Invalid sort type")
	}
	cur, err := p.col.Find(context.TODO(), bson.D{{"isDeleted", false}}, opts)
	if err != nil {
		return results, err
	}
	if err = cur.All(context.TODO(), &results); err != nil {
		return results, err
	}
	return results, nil
}

func (p *menuModel) UpdateMenuByName(name string, menu Menu) error {
	/*
		메뉴의 이름을 변경하고 싶지 않다면, 애초에 값을 받을 수 없도록 구성하는 것은 어떨까요?
		예를들면, update용 메뉴 struct를 새로 생성하고 name 필드를 제외하는 방법이 있을 것 같습니다.
	*/
	if name != menu.Name {
		return fmt.Errorf("Menu name Can not be changed")
	}
	filter := bson.D{{"name", name}}
	menu.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
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
	// 존재하는 메뉴인지, 이미 삭제된 메뉴인지 체크
	menuForDelete, err := p.FindMenuByName(name, false, true)
	if err != nil {
		return err
	}
	// 삭제 flag를 true로 설정 후, 업데이트
	menuForDelete.IsDeleted = util.NewTrue()
	menuForDelete.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	if err := p.UpdateMenuByName(name, menuForDelete); err != nil {
		return err
	}
	return nil
}

func (p *menuModel) IncreaseOrderCount(name string) error {
	menu, err := p.FindMenuByName(name, false, false)
	if err != nil {
		return err
	}
	menu.OrderCount = menu.OrderCount + 1
	if err := p.UpdateMenuByName(name, menu); err != nil {
		return err
	}
	return nil
}
