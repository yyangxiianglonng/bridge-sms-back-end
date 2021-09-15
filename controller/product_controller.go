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

type ProductController struct {
	Context        iris.Context
	ProductService service.ProductService
	Session        sessions.Session
}

func (pd *ProductController) BeforeActivation(ba mvc.BeforeActivation) {

	//通过商品ID检索商品信息功能
	ba.Handle("GET", "/{id}", "GetOneByProductId")

	//通过商品ID删除商品信息功能
	ba.Handle("DELETE", "/{id}", "DeleteProduct")
}

/**
 * url: /v1/product
 * type：GET
 * descs：获取所有商品信息功能
 */
func (pd *ProductController) Get() mvc.Result {
	const COMMENT = "method:Get url:/v1/product Controller:CustomerController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := pd.Context.GetHeader("Authorization")
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

	productList := pd.ProductService.GetProducts()
	if len(productList) == 0 {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PRODUCTGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PRODUCTGET),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, product := range productList {
		respList = append(respList, product.ProductToRespDesc())
	}

	//返回案件列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_PRODUCTGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_PRODUCTGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/product/{id}
 * type：GET
 * descs：通过商品ID检索商品信息功能
 */
func (pd *ProductController) GetOneByProductId() mvc.Result {
	const COMMENT = "method:Get url:/v1/customer/{customer_code} Controller:ProductController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := pd.Context.GetHeader("Authorization")
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

	productId, _ := strconv.ParseInt(pd.Context.Params().Get("id"), 10, 64)
	product := pd.ProductService.GetProduct(productId)

	if product == nil {
		iris.New().Logger().Error(COMMENT + " ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PRODUCTGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PRODUCTGET),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range product {
		respList = append(respList, item.ProductToRespDesc())
	}

	//返回商品列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_PRODUCTGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_PRODUCTGET),
			"data":    &respList,
		},
	}
}

/*
* 即将添加的顾客记录实体
 */
type AddProductEntity struct {
	Id          int64  `json:"id"`
	ProductName string `json:"product_name"`
	Category1   string `json:"category1"`
	Category2   string `json:"category2"`
	Category3   string `json:"category3"`
	Price       string `json:"price"`
	CreatedBy   string `json:"created_by"`
	IsDelete    int64  `json:"is_delete"`
}

/*
* url: /v1/product
* type: POST
* descs: 添加商品功能
 */
func (pd *ProductController) Post() mvc.Result {
	const COMMENT = "method:Post url:/v1/product Controller:ProductController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := pd.Context.GetHeader("Authorization")
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

	var productEntity AddProductEntity
	err = pd.Context.ReadJSON(&productEntity)

	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RESPMSG_FAIL,
				"type":    utils.RESPMSG_ERROR_PRODUCTADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PRODUCTADD),
			},
		}

	}

	var productInfo model.Product
	productInfo.ProductName = productEntity.ProductName
	productInfo.Category1 = productEntity.Category1
	productInfo.Category2 = productEntity.Category2
	productInfo.Category3 = productEntity.Category3
	productInfo.Price = productEntity.Price
	productInfo.CreatedBy = productEntity.CreatedBy
	productInfo.IsDelete = productEntity.IsDelete

	isSuccess := pd.ProductService.SaveProduct(productInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PRODUCTADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PRODUCTADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_PRODUCTADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_PRODUCTADD),
		},
	}
}

/**
 * url: /v1/product
 * type：Put
 * descs：更新顾客信息功能
 */
func (pd *ProductController) Put() mvc.Result {
	const COMMENT = "method:Put url:/v1/product Controller:ProductController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := pd.Context.GetHeader("Authorization")
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

	var productEntity AddProductEntity
	err = pd.Context.ReadJSON(&productEntity)

	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RESPMSG_FAIL,
				"type":    utils.RESPMSG_ERROR_PRODUCTUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PRODUCTUPDATE),
			},
		}

	}

	var productInfo model.Product
	productInfo.Id = productEntity.Id
	productInfo.ProductName = productEntity.ProductName
	productInfo.Category1 = productEntity.Category1
	productInfo.Category2 = productEntity.Category2
	productInfo.Category3 = productEntity.Category3
	productInfo.Price = productEntity.Price
	productInfo.CreatedBy = productEntity.CreatedBy
	productInfo.IsDelete = productEntity.IsDelete

	isSuccess := pd.ProductService.UpdateProduct(productEntity.Id, productInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PRODUCTUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PRODUCTUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_PRODUCTUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_PRODUCTUPDATE),
		},
	}
}

/**
 * url: /v1/product/{id}
 * type：DELETE
 * descs：保存见积详细功能
 */
func (pd *ProductController) DeleteProduct() mvc.Result {
	const COMMENT = "method:Delete url:/v1/category/{id} Controller:EstimateController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := pd.Context.GetHeader("Authorization")
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

	productId, _ := strconv.ParseInt(pd.Context.Params().Get("id"), 10, 64)
	isSuccess := pd.ProductService.DeleteProduct(productId)

	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PRODUCTDELETE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PRODUCTDELETE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_PRODUCTDELETE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_PRODUCTDELETE),
		},
	}
}
