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
	Name        string  `json:"name" bson:"name" uri:"name" binding:"required"`
	Price       int     `json:"price" bson:"price" binding:"required"`
	HotGrade    int     `json:"hotGrade" bson:"hotGrade" binding:"required"`
	IsAvailable *bool   `json:"isAvailable" bson:"isAvailable" binding:"required"`
	IsRecommend *bool   `json:"isRecommend" bson:"isRecommend"`
	IsDeleted   *bool   `json:"isDeleted" bson:"isDeleted"`
	AvgScore    float32 `json:"avgScore" bson:"avgScore"`
	OrderCount  int     `json:"orderCount" bson:"orderCount"`
	/*
		일반적으로 created_at, updated_at 두개의 필드를 동시에 저장합니다.
		그래야 오브젝트가 언제 수정되었는지 파악할 수 있고, 무슨 일이 발생했는지에 대한 히스토리 추적도 가능합니다.
	*/
	CreateDate primitive.DateTime `json:"createDate" bson:"createDate"`
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

/*
1. CanOrder라는 네이밍을 이용하는 건 어떨까요? 조금 더 직관적이라고 생각해요.
*/
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
	/*
		Menu struct에서 Name 필드의 바인딩이 required로 되어 있습니다.
		빈값이 들어온다면 자동으로 에러를 반환할 것 같은데 따로 처리하신 이유가 있을까요?

		다음의 링크에서 아래의 내용을 확인해보실 수 있습니다.
		You can also specify that specific fields are required. If a field is decorated with binding:"required" and has an empty value when binding, an error will be returned.

		https://github.com/gin-gonic/gin#model-binding-and-validation
	*/
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
	} else {
		return results, fmt.Errorf("Invalid sort type")
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
	/*
		메뉴의 이름을 변경하고 싶지 않다면, 애초에 값을 받을 수 없도록 구성하는 것은 어떨까요?
		예를들면, update용 메뉴 struct를 새로 생성하고 name 필드를 제외하는 방법이 있을 것 같습니다.
	*/
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
	/*
		이렇게 체크하기 보다는, FindMenuByName 함수 안에서 IsDeleted가 True인 것은 제외하도록 필터를 구성하는 것은 어떨까요?
		그렇게 한다면 메뉴를 삭제하는 로직에서는, 삭제된 것인지에 대해서는 신경을 쓰지 않아도 됩니다.
	*/
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

/*
IncreaseOrderCount와 같이 좀더 직관적이고 간단하게 네이밍 할 수 있겠습니다.
*/
func (p *menuModel) UpdateMenuByNameIncOrderCount(name string) error {
	menu, err := p.FindMenuByName(name)
	if err != nil {
		return err
	}
	menu.OrderCount = menu.OrderCount + 1
	if err := p.UpdateMenuByName(name, menu); err != nil {
		return err
	}
	return nil
}
