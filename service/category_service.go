package service

import (
	"main/model"

	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

/**
* 商品分类接口
 */
type CategoryService interface {
	GetCategories() []*model.Category
	GetParentCategories() []*model.Category
	GetCategory(id int64) []*model.Category
	SaveCategory(category model.Category) bool
	UpdateCategory(id int64, category model.Category) bool
	DeleteCategory(id int64) bool
}

/**
 * 实例化商品分类服务:服务器
 */
func NewCategoryService(engine *xorm.Engine) CategoryService {
	return &categoryService{
		Engine: engine,
	}
}

/**
* 商品分类服务结构体
 */
type categoryService struct {
	Engine *xorm.Engine
}

/**
 * 请求商品分类列表数据
 */
func (ca *categoryService) GetCategories() (categoryList []*model.Category) {
	err := ca.Engine.Where("is_delete = ?", 0).OrderBy("parent_id").Find(&categoryList)

	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
	return
}

/**
 * 请求父商品分类列表数据
 */
func (ca *categoryService) GetParentCategories() (parentCategoryList []*model.Category) {
	err := ca.Engine.Where("is_delete = ?", 0).And("parent_flg = ?", true).Find(&parentCategoryList)

	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
	return
}

/**
 * 通过商品ID获取单个商品分类信息服务
 */
func (ca *categoryService) GetCategory(id int64) (category []*model.Category) {
	err := ca.Engine.Where("is_delete = ?", 0).And("id = ?", id).Find(&category)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
	return
}

/**
 * 保存商品分类服务
 */
func (ca *categoryService) SaveCategory(category model.Category) bool {
	_, err := ca.Engine.Insert(&category)
	return err == nil
}

/**
 * 更新商品分类服务
 */
func (ca *categoryService) UpdateCategory(id int64, category model.Category) bool {
	_, err := ca.Engine.Where("id = ?", id).Update(category)
	return err == nil
}

/**
 * 删除商品分类服务
 */
func (ca *categoryService) DeleteCategory(id int64) bool {
	var category model.Category
	_, err := ca.Engine.Where("id = ?", id).Delete(&category)
	return err == nil
}
