package service

import (
	"github.com/kataras/iris/v12"
	"main/model"
	"xorm.io/xorm"
)

type ProductService interface {
	GetProducts() []*model.Product
	GetProduct(id int64) []*model.Product
	SaveProduct(product model.Product) bool
	UpdateProduct(id int64, product model.Product) bool
	DeleteProduct(id int64) bool
}

func NewProductService(engine *xorm.Engine) ProductService {
	return &productService{
		Engine: engine,
	}
}

type productService struct {
	Engine *xorm.Engine
}

/*
* 请求商品列表
 */
func (pd *productService) GetProducts() (productList []*model.Product) {
	err := pd.Engine.Where("is_delete = ?", 0).Find(&productList)

	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
	return
}

/*
* 通过商品ID获取商品信息
 */
func (pd *productService) GetProduct(id int64) (product []*model.Product) {
	err := pd.Engine.Where("id = ?", id).Find(&product)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
	return
}

/*
* 添加商品信息
 */
func (pd *productService) SaveProduct(product model.Product) bool {
	_, err := pd.Engine.Insert(&product)
	return err == nil
}

/*
* 更新商品信息
 */
func (pd *productService) UpdateProduct(id int64, product model.Product) bool {
	_, err := pd.Engine.Where("id = ?", id).Update(product)
	return err == nil
}

/**
 * 删除商品服务
 */
func (pd *productService) DeleteProduct(id int64) bool {
	var product model.Product
	_, err := pd.Engine.Where("id = ?", id).Delete(&product)
	return err == nil
}
