package service

import (
	"main/model"
	"xorm.io/xorm"
)

/*
* 注文服务接口
 */
type OrderService interface {
	GetOrders(projectCode string) []*model.Order
	GetOrder(orderCode string) []*model.Order
	SaveOrder(order model.Order) bool
	UpdateOrder(orderCode string, order model.Order) bool
}

/**
 * 实例化注文服务:服务器
 */
func NewOrderService(engine *xorm.Engine) OrderService {
	return &orderService{
		Engine: engine,
	}
}

/**
 * 注文服务实现结构体
 */
type orderService struct {
	Engine *xorm.Engine
}

/*
* 请求某个案件下的所有注文列表数据
 */
func (or *orderService) GetOrders(projectCode string) (orderList []*model.Order) {

	err := or.Engine.Where("is_delete = ?", 0).And("project_code = ?", projectCode).Find(&orderList)

	if err != nil {
		panic(err.Error())
	}
	return
}

/**
 * 通过注文CD获取注文信息
 */
func (or *orderService) GetOrder(orderCode string) (order []*model.Order) {
	err := or.Engine.Where("is_delete = ?", 0).And("order_code = ?", orderCode).Find(&order)

	if err != nil {
		panic(err.Error())
	}
	return
}

/*
* 保存注文信息
 */
func (or *orderService) SaveOrder(order model.Order) bool {
	_, err := or.Engine.Insert(&order)
	return err == nil
}

/*
* 更新注文信息
 */
func (or *orderService) UpdateOrder(orderCode string, order model.Order) bool {
	_, err := or.Engine.Where("order_code = ?", orderCode).Update(&order)
	return err == nil
}
