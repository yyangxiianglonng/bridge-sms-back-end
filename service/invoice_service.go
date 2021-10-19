package service

import (
	"main/model"

	"xorm.io/xorm"
)

/*
* 请求书服务接口
 */
type InvoiceService interface {
	GetInvoiceAll() []*model.Invoice
	GetInvoices(projectCode string) []*model.Invoice
	GetInvoice(invoiceCode string) []*model.Invoice
	SaveInvoice(invoice model.Invoice) bool
	UpdateInvoice(invoiceCode string, invoice model.Invoice) bool
}

/**
 * 实例化请求书服务:服务器
 */
func NewInvoiceService(engine *xorm.Engine) InvoiceService {
	return &invoiceService{
		Engine: engine,
	}
}

/**
 *请求书服务实现结构体
 */
type invoiceService struct {
	Engine *xorm.Engine
}

/**
 * 获取所有请求列表数据
 */
func (in *invoiceService) GetInvoiceAll() (invoiceList []*model.Invoice) {
	err := in.Engine.Where("is_delete = ?", 0).Find(&invoiceList)

	if err != nil {
		panic(err.Error())
	}
	return
}

/*
* 请求某个案件下的所有请求书列表数据
 */
func (in *invoiceService) GetInvoices(projectCode string) (invoiceList []*model.Invoice) {
	err := in.Engine.Where("is_delete = ?", 0).And("project_code = ?", projectCode).Find(&invoiceList)

	if err != nil {
		panic(err.Error())
	}
	return
}

/**
 * 通过请求书CD获取请求书信息
 */
func (in *invoiceService) GetInvoice(invoiceCode string) (invoice []*model.Invoice) {
	err := in.Engine.Where("is_delete = ?", 0).And("invoice_code = ?", invoiceCode).Find(&invoice)

	if err != nil {
		panic(err.Error())
	}
	return
}

/*
 * 保存请求书信息
 */
func (in *invoiceService) SaveInvoice(invoice model.Invoice) bool {
	_, err := in.Engine.Insert(&invoice)
	return err == nil
}

/*
 * 更新请求书信息
 */
func (in *invoiceService) UpdateInvoice(invoiceCode string, invoice model.Invoice) bool {
	_, err := in.Engine.Where("invoice_code = ?", invoiceCode).Update(&invoice)
	return err == nil
}
