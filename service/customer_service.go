package service

import (
	"github.com/kataras/iris/v12"
	"main/model"
	"xorm.io/xorm"
)

/**
*	顾客服务接口
 */
type CustomerService interface {
	GetCustomers() []*model.Customer
	GetCustomer(customerCode string) []*model.Customer
	SaveCustomer(customer model.Customer) bool
	UpdateCustomer(customerCode string, customer model.Customer) bool
}

/**
 * 实例化顾客服务:服务器
 */
func NewCustomerService(engine *xorm.Engine) CustomerService {
	return &customerService{
		Engine: engine,
	}
}

/**
 * 顾客服务实现结构体
 */
type customerService struct {
	Engine *xorm.Engine
}

/**
 * 请求顾客列表数据
 */
func (pr *customerService) GetCustomers() (customerList []*model.Customer) {
	err := pr.Engine.Where("is_delete = ?", 0).Find(&customerList)

	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
	return
}

/**
 * 通过客户CD获取单个客户信息服务
 */
func (pr *customerService) GetCustomer(customerCode string) (customer []*model.Customer) {
	err := pr.Engine.Where("customer_code = ?", customerCode).Find(&customer)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
	return
}

/**
 * 添加客户信息服务
 */
func (pr *customerService) SaveCustomer(customer model.Customer) bool {
	_, err := pr.Engine.Insert(&customer)
	return err == nil
}

/**
 * 更新客户信息服务
 */
func (pr *customerService) UpdateCustomer(customerCode string, customer model.Customer) bool {
	_, err := pr.Engine.Where("customer_code = ?", customerCode).Update(customer)
	return err == nil
}
