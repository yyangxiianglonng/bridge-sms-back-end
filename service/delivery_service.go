package service

import (
	"main/model"

	"xorm.io/xorm"
)

/*
* 纳品服务接口
 */
type DeliveryService interface {
	GetDeliveries(projectCode string) []*model.Delivery
	GetDelivery(deliveryCode string) []*model.Delivery
	SaveDelivery(delivery model.Delivery) bool
	UpdateDelivery(deliveryCode string, delivery model.Delivery) bool
}

/*
* 实例化纳品服务:服务器
 */
func NewDeliveryService(engine *xorm.Engine) DeliveryService {
	return &deliveryService{
		Engine: engine,
	}
}

/*
* 纳品服务实现结构体
 */
type deliveryService struct {
	Engine *xorm.Engine
}

/*
* 请求某个案件下的所有纳品列表数据
 */
func (de *deliveryService) GetDeliveries(projectCode string) (deliveryList []*model.Delivery) {
	err := de.Engine.Where("is_delete = ?", 0).And("project_code = ?", projectCode).Find(&deliveryList)

	if err != nil {
		panic(err.Error())
	}
	return
}

/*
* 痛殴纳品CD获取某个纳品信息
 */
func (de *deliveryService) GetDelivery(deliveryCode string) (delivery []*model.Delivery) {
	err := de.Engine.Where("is_delete = ?", 0).And("delivery_code = ?", deliveryCode).Find(&delivery)

	if err != nil {
		panic(err.Error())
	}
	return
}

/*
*　保存纳品信息
 */
func (de *deliveryService) SaveDelivery(delivery model.Delivery) bool {
	_, err := de.Engine.Insert(&delivery)
	return err == nil
}

/*
* 更新纳品信息
 */
func (de *deliveryService) UpdateDelivery(deliveryCode string, delivery model.Delivery) bool {
	_, err := de.Engine.Where("delivery_code = ?", deliveryCode).Update(&delivery)
	return err == nil
}
