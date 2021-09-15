package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"main/model"
	"main/service"
	"main/utils"
	"strconv"
	"time"
)

type CategoryController struct {
	Context         iris.Context
	CategoryService service.CategoryService
	Session         *sessions.Session
}

func (ca *CategoryController) BeforeActivation(ba mvc.BeforeActivation) {

	//通过商品分类ID检索商品分类信息功能
	ba.Handle("GET", "/{id}", "GetOneById")

	//通过商品分类ID删除指定商品分类信息功能
	ba.Handle("DELETE", "/{id}", "DeleteCategory")
}

/**
 * url: /v1/category
 * type：GET
 * descs：获取所有商品分类功能
 */
func (ca *CategoryController) Get() mvc.Result {
	const COMMENT = "method:Get url:/v1/category Controller:CategoryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := ca.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

	categoryList := ca.CategoryService.GetCategories()
	iris.New().Logger().Info(categoryList)
	if len(categoryList) == 0 {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORIEGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORIEGET),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, category := range categoryList {
		respList = append(respList, category.CategoryToRespDesc())
	}

	//返回案件列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CATEGORIEGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CATEGORIEGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/category/Parent
 * type：GET
 * descs：获取所有父商品分类功能
 */
func (ca *CategoryController) GetParent() mvc.Result {
	const COMMENT = "method:Get url:/v1/category/Parent Controller:CategoryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := ca.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

	parentCategoryList := ca.CategoryService.GetParentCategories()
	if len(parentCategoryList) == 0 {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORIEGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORIEGET),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, parentCategory := range parentCategoryList {
		respList = append(respList, parentCategory.CategoryToRespDesc())
	}
	//返回案件列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CATEGORIEGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CATEGORIEGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/category/{id}
 * type：GET
 * descs：通过商品分类ID检索商品分类信息功能
 */
func (ca *CategoryController) GetOneById() mvc.Result {
	const COMMENT = "method:Get url:/v1/category/{id} Controller:CategoryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := ca.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

	//把商品分类ID从string类型变换为int64类型
	categoryId, _ := strconv.ParseInt(ca.Context.Params().Get("id"), 10, 64)
	category := ca.CategoryService.GetCategory(categoryId)

	if category == nil {
		iris.New().Logger().Error(COMMENT + " ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CUSTOMERGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CUSTOMERGET),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range category {
		respList = append(respList, item.CategoryToRespDesc())
	}

	//返回案件列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CATEGORIEGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CATEGORIEGET),
			"data":    &respList,
		},
	}
}

/*
* 即将添加的商品分类实体
 */
type AddCategoryEntity struct {
	Id           int64  `json:"id"`
	CategoryName string `json:"category_name"`
	ParentId     string `json:"parent_id"`
	ParentFlg    bool   `json:"parent_flg"`
	Sorting      string `json:"sorting"`
	CreatedBy    string `json:"created_by"`
	IsDelete     int64  `json:"is_delete"`
}

/*
* url: /v1/category
* type: POST
* descs: 添加商品分类功能
 */
func (ca *CategoryController) Post() mvc.Result {
	const COMMENT = "method:Post /v1/category Controller:CategoryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := ca.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

	var categoryEntity AddCategoryEntity
	err = ca.Context.ReadJSON(&categoryEntity)

	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RESPMSG_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORYADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORYADD),
			},
		}

	}

	var categoryInfo model.Category
	categoryInfo.Id = categoryEntity.Id
	categoryInfo.CategoryName = categoryEntity.CategoryName
	categoryInfo.ParentId = categoryEntity.ParentId
	categoryInfo.ParentFlg = categoryEntity.ParentFlg
	categoryInfo.Sorting = categoryEntity.Sorting
	categoryInfo.CreatedBy = categoryEntity.CreatedBy
	categoryInfo.IsDelete = categoryEntity.IsDelete

	isSuccess := ca.CategoryService.SaveCategory(categoryInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORYADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORYADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CATEGORYADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CATEGORYADD),
		},
	}
}

/**
 * url: /v1/category
 * type：PUT
 * descs：更新商品分类信息功能
 */
func (ca *CategoryController) Put() mvc.Result {
	const COMMENT = "method:Put url:/v1/category Controller:CategoryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := ca.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

	var categoryEntity AddCategoryEntity
	err = ca.Context.ReadJSON(&categoryEntity)

	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RESPMSG_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORYUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORYUPDATE),
			},
		}

	}

	var categoryInfo model.Category
	categoryInfo.Id = categoryEntity.Id
	categoryInfo.CategoryName = categoryEntity.CategoryName
	categoryInfo.ParentId = categoryEntity.ParentId
	categoryInfo.ParentFlg = categoryEntity.ParentFlg
	categoryInfo.Sorting = categoryEntity.Sorting
	categoryInfo.CreatedBy = categoryEntity.CreatedBy
	categoryInfo.IsDelete = categoryEntity.IsDelete

	isSuccess := ca.CategoryService.UpdateCategory(categoryEntity.Id, categoryInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORYUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORYUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CATEGORYUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CATEGORYUPDATE),
		},
	}
}

/**
 * url: /v1/category/{id}
 * type：DELETE
 * descs：保存见积详细功能
 */
func (ca *CategoryController) DeleteCategory() mvc.Result {
	const COMMENT = "method:Delete url:/v1/category/{id} Controller:EstimateController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := ca.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

	categoryId, _ := strconv.ParseInt(ca.Context.Params().Get("id"), 10, 64)
	isSuccess := ca.CategoryService.DeleteCategory(categoryId)

	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CATEGORIEDELETE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORIEDELETE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CATEGORIEDELETE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CATEGORIEDELETE),
		},
	}
}
