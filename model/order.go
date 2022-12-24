package model

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 주문 상태 상수 정의
const (
	ORDER_STATUS_GROUP_ACTIVE    = "active"
	ORDER_STATUS_GROUP_DEACTIVE  = "deactive"
	ORDER_STATUS_GROUP_COMPELETE = "complete"
	ORDER_STATUS_GROUP_ALL       = "all"

	ORDER_STATUS_WAITING   = "대기"
	ORDER_STATUS_RECEIPT   = "주문"
	ORDER_STATUS_COOKING   = "조리"
	ORDER_STATUS_DELIVERY  = "배달"
	ORDER_STATUS_COMPLETED = "완료"
)

type Order struct {
	Seq      string             `json:"seq" bson:"seq"`
	Count    uint32             `json:"count" bson:"count"`
	MenuList []string           `json:"menuList" bson:"menuList" binding:"required"`
	Address  string             `json:"address" bson:"address" binding:"required"`
	Phone    string             `json:"phone" bson:"phone" binding:"required"`
	Status   string             `json:"status" bson:"status"`
	Date     primitive.DateTime `json:"date" bson:"date"`
}

func (o *Order) IsChangeable() bool {
	if o.Status == ORDER_STATUS_WAITING || o.Status == ORDER_STATUS_RECEIPT {
		return true
	}
	return false
}

type OrderMenuList struct {
	MenuList []string `json:"menuList" bson:"menuList" binding:"required"`
}

type orderModel struct {
	col              *mongo.Collection
	orderCounter     uint32
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
	m.orderStatusGroup[ORDER_STATUS_GROUP_DEACTIVE] = []string{ORDER_STATUS_COMPLETED}
	m.orderStatusGroup[ORDER_STATUS_GROUP_COMPELETE] = []string{ORDER_STATUS_COMPLETED}
	return m
}

// 주문 일련번호 생성 함수
func createSeqStr(orderCount uint32) string {
	return fmt.Sprintf("%d-%010d", time.Now().Unix(), orderCount)
}

func (p *orderModel) CreateOrder(order Order) (string, error) {
	// 주문 기본값 초기화
	order.Count = atomic.AddUint32(&p.orderCounter, 1)
	order.Seq = createSeqStr(order.Count)
	order.Status = ORDER_STATUS_WAITING
	order.Date = primitive.NewDateTimeFromTime(time.Now())
	// 주문 저장
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
	updateResult, err := p.col.UpdateOne(context.TODO(), filter, bson.D{{"$set", order}})
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("No Order Updated")
	}
	return nil
}
