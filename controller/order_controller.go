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

func (or *OrderController) BeforeActivation(ba mvc.BeforeActivation) {
	//通过project_code获取对应的注文列表
	ba.Handle("GET", "/all/{project_code}", "GetAllByProjectCode")
	//通过order_code获取对应的注文列表
	ba.Handle("GET", "/one/{order_code}", "GetOneByOrderCode")
}

type OrderController struct {
	Context      iris.Context
	OrderService service.OrderService
	Session      *sessions.Session
}

/**
 * url: /v1/order/all/{project_code}
 * type：GET
 * descs：通过案件CD获取所有注文功能
 */
func (or *OrderController) GetAllByProjectCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/order/all/{project_code} Controller:OrderController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := or.Context.GetHeader("Authorization")
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

	projectCode := or.Context.Params().Get("project_code")
	order := or.OrderService.GetOrders(projectCode)

	if order == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERGET),
			},
		}
	}

	//将查询到的注文数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range order {
		respList = append(respList, item.OrderToRespDesc())
	}

	//返回注文列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ORDERGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ORDERGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/order/one/{order_code}
 * type：GET
 * descs：通过注文CD获取某一条注文信息
 */
func (or *OrderController) GetOneByOrderCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/order/one/{order_code} Controller:OrderController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := or.Context.GetHeader("Authorization")
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

	orderCode := or.Context.Params().Get("order_code")
	order := or.OrderService.GetOrder(orderCode)

	if order == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERGET),
			},
		}
	}

	//将查询到的注文数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range order {
		respList = append(respList, item.OrderToRespDesc())
	}

	//返回注文列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ORDERGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ORDERGET),
			"data":    &respList,
		},
	}
}

/**
 * 即将添加的注文记录实体
 */
type AddOrderEntity struct {
	Id                   int64  `json:"id"`
	OrderCode            string `json:"order_code"`
	EstimateCode         string `json:"estimate_code"`
	ProjectCode          string `json:"project_code"`
	ProjectName          string `json:"project_name"`
	CustomerName         string `json:"customer_name"`
	CustomerAddress      string `json:"customer_address"`
	Work                 string `json:"work"`
	Deliverables         string `json:"deliverables"`
	WorkTime             string `json:"work_time"`
	Personnel1           string `json:"personnel1"`
	Personnel2           string `json:"personnel2"`
	DeliverableSpace     string `json:"deliverable_space"`
	Commission           string `json:"commission"`
	PaymentDate          string `json:"payment_date"`
	AcceptanceConditions string `json:"acceptance_conditions"`
	Other                string `json:"other"`
	Note                 string `json:"note"`
	CreatedBy            string `json:"created_by"`
	IsDelete             int64  `json:"is_delete"`
}

/**
 * url: /v1/order
 * type：POST
 * descs：添加注文功能
 */
func (or *OrderController) Post() mvc.Result {
	const COMMENT = "method:Post url:/v1/order Controller:OrderController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := or.Context.GetHeader("Authorization")
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

	var orderEntity AddOrderEntity
	err = or.Context.ReadJSON(&orderEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERADD),
			},
		}
	}

	var orderInfo model.Order

	orderInfo.OrderCode = orderEntity.OrderCode
	orderInfo.EstimateCode = orderEntity.EstimateCode
	orderInfo.ProjectCode = orderEntity.ProjectCode
	orderInfo.ProjectName = orderEntity.ProjectName
	orderInfo.CustomerName = orderEntity.CustomerName
	orderInfo.CustomerAddress = orderEntity.CustomerAddress
	orderInfo.Work = orderEntity.Work
	orderInfo.Deliverables = orderEntity.Deliverables
	orderInfo.WorkTime = orderEntity.WorkTime
	orderInfo.Personnel1 = orderEntity.Personnel1
	orderInfo.Personnel2 = orderEntity.Personnel2
	orderInfo.DeliverableSpace = orderEntity.DeliverableSpace
	orderInfo.Commission = orderEntity.Commission
	orderInfo.PaymentDate = orderEntity.PaymentDate
	orderInfo.AcceptanceConditions = orderEntity.AcceptanceConditions
	orderInfo.Other = orderEntity.Other
	orderInfo.Note = orderEntity.Note
	orderInfo.CreatedBy = orderEntity.CreatedBy
	orderInfo.IsDelete = orderEntity.IsDelete

	isSuccess := or.OrderService.SaveOrder(orderInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ORDERADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ORDERADD),
		},
	}
}

/**
 * url: /v1/order
 * type：Put
 * descs：更新注文功能
 */
func (or *OrderController) Put() mvc.Result {
	const COMMENT = "method:Put url:/v1/order Controller:OrderController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := or.Context.GetHeader("Authorization")
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

	var orderEntity AddOrderEntity
	err = or.Context.ReadJSON(&orderEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERUPDATE),
			},
		}
	}

	var orderInfo model.Order

	orderInfo.OrderCode = orderEntity.OrderCode
	orderInfo.EstimateCode = orderEntity.EstimateCode
	orderInfo.ProjectCode = orderEntity.ProjectCode
	orderInfo.ProjectName = orderEntity.ProjectName
	orderInfo.CustomerName = orderEntity.CustomerName
	orderInfo.CustomerAddress = orderEntity.CustomerAddress
	orderInfo.Work = orderEntity.Work
	orderInfo.Deliverables = orderEntity.Deliverables
	orderInfo.WorkTime = orderEntity.WorkTime
	orderInfo.Personnel1 = orderEntity.Personnel1
	orderInfo.Personnel2 = orderEntity.Personnel2
	orderInfo.DeliverableSpace = orderEntity.DeliverableSpace
	orderInfo.Commission = orderEntity.Commission
	orderInfo.PaymentDate = orderEntity.PaymentDate
	orderInfo.AcceptanceConditions = orderEntity.AcceptanceConditions
	orderInfo.Other = orderEntity.Other
	orderInfo.Note = orderEntity.Note
	orderInfo.CreatedBy = orderEntity.CreatedBy
	orderInfo.IsDelete = orderEntity.IsDelete

	isSuccess := or.OrderService.UpdateOrder(orderEntity.OrderCode, orderInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ORDERUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ORDERUPDATE),
		},
	}
}
