package model

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"codestates.wba-01/archoi/backend/oos/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
상수로 정의해주신 점 좋습니다.
*/
// 주문 상태 상수 정의
const (
	ORDER_STATUS_GROUP_ACTIVE    = "active"
	ORDER_STATUS_GROUP_COMPELETE = "complete"
	ORDER_STATUS_GROUP_ALL       = "all"

	ORDER_STATUS_WAITING   = "대기"
	ORDER_STATUS_RECEIPT   = "주문"
	ORDER_STATUS_COOKING   = "조리"
	ORDER_STATUS_DELIVERY  = "배달"
	ORDER_STATUS_COMPLETED = "완료"
)

type Order struct {
	Seq       string             `json:"seq" bson:"seq"`
	MenuList  []string           `json:"menuList" bson:"menuList" binding:"required"`
	Address   string             `json:"address" bson:"address" binding:"required"`
	Phone     string             `json:"phone" bson:"phone" binding:"required"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}

func (o *Order) CanChangeStatus() bool {
	if o.Status == ORDER_STATUS_WAITING || o.Status == ORDER_STATUS_RECEIPT {
		return true
	}
	return false
}

func (o *Order) IsContainMenu(menuName string) bool {
	for _, v := range o.MenuList {
		if v == menuName {
			return true
		}
	}
	return false
}

type OrderMenuList struct {
	MenuList []string `json:"menuList" bson:"menuList" binding:"required"`
}

type orderModel struct {
	col              *mongo.Collection
	orderCounter     int32
	orderStatusGroup map[string][]string
}

// orderModel 객체 생성자
func NewOrderModel(col *mongo.Collection) *orderModel {
	m := new(orderModel)
	m.col = col
	m.orderCounter = 0
	m.orderStatusGroup = make(map[string][]string)
	m.orderStatusGroup[ORDER_STATUS_GROUP_ALL] =
		[]string{
			ORDER_STATUS_COOKING,
			ORDER_STATUS_DELIVERY,
			ORDER_STATUS_RECEIPT,
			ORDER_STATUS_WAITING,
			ORDER_STATUS_COMPLETED}
	m.orderStatusGroup[ORDER_STATUS_GROUP_ACTIVE] =
		[]string{
			ORDER_STATUS_COOKING,
			ORDER_STATUS_DELIVERY,
			ORDER_STATUS_RECEIPT,
			ORDER_STATUS_WAITING}
	m.orderStatusGroup[ORDER_STATUS_GROUP_COMPELETE] = []string{ORDER_STATUS_COMPLETED}
	return m
}

func (p *orderModel) CreateOrder(order Order) (string, error) {
	// 주문 기본값 초기화
	/*
		이렇게 seq를 구성한다면 주문에 대해서 중복이 발생할 가능성은 없나요?
		동일한 시간에 주문이 들어온다면 중복된 주문 번호가 생기지는 않을까요?
		----------------
		시퀀스 문자열은 "현재시간-주문카운트"의 형식이며,
		주문카운트는 원자적 증가 연산을 사용하기에 중복되지 않을 것이라고 생각합니다.
		하지만 InsertOne 함수 호출이 실패하여 주문카운트를 감소시키면서 주문카운트가 중복될 가능성이 있어 보입니다.
		일단 주문카운트 감소 로직을 제거하는 것으로 주문 번호 중복 문제는 해결될 것으로 생각합니다.
	*/
	order.Seq = util.CreateSeqStr(uint32(atomic.AddInt32(&p.orderCounter, 1)))
	order.Status = ORDER_STATUS_WAITING
	order.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	order.UpdatedAt = order.CreatedAt
	_, err := p.col.InsertOne(context.TODO(), order)
	if err != nil {
		return "", err
	}
	// 주문 일련번호 반환
	return order.Seq, nil
}

func (p *orderModel) FindOrderBySeq(seq string) (Order, error) {
	var result Order
	filter := bson.D{{"seq", seq}}
	err := p.col.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return result, fmt.Errorf("No Order found given seq: %s\n", seq)
	} else if err != nil {
		return result, err
	}
	return result, nil
}

func (p *orderModel) FindOrderListInStatus(statusGroup string) ([]Order, error) {
	var results []Order
	// 필터 설정
	statusList := p.orderStatusGroup[ORDER_STATUS_GROUP_ALL]
	if v, ok := p.orderStatusGroup[statusGroup]; ok {
		statusList = v
	}
	filter := bson.D{{"status", bson.D{{"$in", statusList}}}}
	// 정렬 옵션 설정
	opts := options.Find().SetSort(bson.D{{"date", -1}})
	cur, err := p.col.Find(context.TODO(), filter, opts)
	if err != nil {
		return results, err
	}
	if err = cur.All(context.TODO(), &results); err != nil {
		return results, err
	}
	return results, nil
}

func (p *orderModel) UpdateOrderBySeq(seq string, order Order) error {
	filter := bson.D{{"seq", seq}}
	order.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	updateResult, err := p.col.UpdateOne(context.TODO(), filter, bson.D{{"$set", order}})
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("No Order Updated")
	}
	return nil
}
