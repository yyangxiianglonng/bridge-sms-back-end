package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"main/model"
	"main/service"
	"main/utils"
	"time"
)

func (de *DeliveryController) BeforeActivation(ba mvc.BeforeActivation) {
	//通过project_code获取对应的注文列表
	ba.Handle("GET", "/all/{project_code}", "GerAllByProjectCode")
	//通过order_code获取对应的注文列表
	ba.Handle("GET", "/one/{delivery_code}", "GetOneByDeliveryCode")
}

type DeliveryController struct {
	Context         iris.Context
	DeliveryService service.DeliveryService
	Session         *sessions.Session
}

/**
 * url: /v1/delivery/all/{project_code}
 * type：GET
 * descs：通过案件CD获取所有纳品功能
 */
func (de *DeliveryController) GerAllByProjectCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/delivery/all/{project_code} Controller:DeliveryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := de.Context.GetHeader("Authorization")
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

	projectCode := de.Context.Params().Get("project_code")
	delivery := de.DeliveryService.GetDeliveries(projectCode)

	if delivery == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_DELIVERYGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_DELIVERYGET),
			},
		}
	}

	//将查询到的纳品数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range delivery {
		respList = append(respList, item.DeliveryToRespDesc())
	}

	//返回纳品列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_DELIVERYGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_DELIVERYGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/delivery/one/{delivery_code}
 * type：GET
 * descs：通过纳品CD获取某一条纳品信息
 */
func (de *DeliveryController) GetOneByDeliveryCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/delivery/one/{delivery_code} Controller:DeliveryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := de.Context.GetHeader("Authorization")
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

	deliveryCode := de.Context.Params().Get("delivery_code")
	delivery := de.DeliveryService.GetDelivery(deliveryCode)

	if delivery == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_DELIVERYGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_DELIVERYGET),
			},
		}
	}

	//将查询到的纳品数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range delivery {
		respList = append(respList, item.DeliveryToRespDesc())
	}

	//返回纳品列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_DELIVERYGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_DELIVERYGET),
			"data":    &respList,
		},
	}
}

/**
 * 即将添加的注文记录实体
 */
type AddDeliveryEntity struct {
	Id            int64  `json:"id"`
	DeliveryCode  string `json:"delivery_code"`
	EstimateCode  string `json:"estimate_code"`
	ProjectCode   string `json:"project_code"`
	ProjectName   string `json:"project_name"`
	CustomerName  string `json:"customer_name"`
	Deliverables1 string `json:"deliverables1"`
	Deliverables2 string `json:"deliverables2"`
	Deliverables3 string `json:"deliverables3"`
	Quantity1     string `json:"quantity1"`
	Quantity2     string `json:"quantity2"`
	Quantity3     string `json:"quantity3"`
	Memo1         string `json:"memo1"`
	Memo2         string `json:"memo2"`
	Memo3         string `json:"memo3"`
	DeliveryDate  string `json:"delivery_date"`
	Remarks       string `json:"remarks"`
	CreatedBy     string `json:"created_by"`
	IsDelete      int64  `json:"is_delete"`
}

/**
 * url: /v1/delivery
 * type：POST
 * descs：添加纳品功能
 */
func (de *DeliveryController) Post() mvc.Result {
	const COMMENT = "method:Post url:/v1/delivery Controller:DeliveryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := de.Context.GetHeader("Authorization")
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

	var deliveryEntity AddDeliveryEntity
	err = de.Context.ReadJSON(&deliveryEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEDETAILADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEDETAILADD),
			},
		}
	}

	var deliveryInfo model.Delivery

	deliveryInfo.DeliveryCode = deliveryEntity.DeliveryCode
	deliveryInfo.EstimateCode = deliveryEntity.EstimateCode
	deliveryInfo.ProjectCode = deliveryEntity.ProjectCode
	deliveryInfo.ProjectName = deliveryEntity.ProjectName
	deliveryInfo.CustomerName = deliveryEntity.CustomerName
	deliveryInfo.Deliverables1 = deliveryEntity.Deliverables1
	deliveryInfo.Deliverables2 = deliveryEntity.Deliverables2
	deliveryInfo.Deliverables3 = deliveryEntity.Deliverables3
	deliveryInfo.Quantity1 = deliveryEntity.Quantity1
	deliveryInfo.Quantity2 = deliveryEntity.Quantity2
	deliveryInfo.Quantity3 = deliveryEntity.Quantity3
	deliveryInfo.Memo1 = deliveryEntity.Memo1
	deliveryInfo.Memo2 = deliveryEntity.Memo2
	deliveryInfo.Memo3 = deliveryEntity.Memo3
	deliveryInfo.DeliveryDate = deliveryEntity.DeliveryDate
	deliveryInfo.Remarks = deliveryEntity.Remarks
	deliveryInfo.CreatedBy = deliveryEntity.CreatedBy
	deliveryInfo.IsDelete = deliveryEntity.IsDelete

	isSuccess := de.DeliveryService.SaveDelivery(deliveryInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEDETAILADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEDETAILADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ESTIMATEDETAILADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEDETAILADD),
		},
	}
}

/**
 * url: /v1/delivery
 * type：PUT
 * descs：更新纳品功能
 */
func (de *DeliveryController) Put() mvc.Result {
	const COMMENT = "method:PUT url:/v1/delivery Controller:DeliveryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := de.Context.GetHeader("Authorization")
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

	var deliveryEntity AddDeliveryEntity
	err = de.Context.ReadJSON(&deliveryEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_DELIVERYUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_DELIVERYUPDATE),
			},
		}
	}

	var deliveryInfo model.Delivery

	deliveryInfo.DeliveryCode = deliveryEntity.DeliveryCode
	deliveryInfo.EstimateCode = deliveryEntity.EstimateCode
	deliveryInfo.ProjectCode = deliveryEntity.ProjectCode
	deliveryInfo.ProjectName = deliveryEntity.ProjectName
	deliveryInfo.CustomerName = deliveryEntity.CustomerName
	deliveryInfo.Deliverables1 = deliveryEntity.Deliverables1
	deliveryInfo.Deliverables2 = deliveryEntity.Deliverables2
	deliveryInfo.Deliverables3 = deliveryEntity.Deliverables3
	deliveryInfo.Quantity1 = deliveryEntity.Quantity1
	deliveryInfo.Quantity2 = deliveryEntity.Quantity2
	deliveryInfo.Quantity3 = deliveryEntity.Quantity3
	deliveryInfo.Memo1 = deliveryEntity.Memo1
	deliveryInfo.Memo2 = deliveryEntity.Memo2
	deliveryInfo.Memo3 = deliveryEntity.Memo3
	deliveryInfo.DeliveryDate = deliveryEntity.DeliveryDate
	deliveryInfo.Remarks = deliveryEntity.Remarks
	deliveryInfo.CreatedBy = deliveryEntity.CreatedBy
	deliveryInfo.IsDelete = deliveryEntity.IsDelete

	isSuccess := de.DeliveryService.UpdateDelivery(deliveryEntity.DeliveryCode, deliveryInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_DELIVERYUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_DELIVERYUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_DELIVERYUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_DELIVERYUPDATE),
		},
	}
}
