package controller

import (
	"main/model"
	"main/service"
	"main/utils"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func (in *InvoiceController) BeforeActivation(ba mvc.BeforeActivation) {
	//通过project_code获取对应的注文列表
	ba.Handle("GET", "/all/{project_code}", "GetAllByProjectCode")
	//通过order_code获取对应的注文列表
	ba.Handle("GET", "/one/{invoice_code}", "GetOneByInvoiceCode")
}

type InvoiceController struct {
	Context        iris.Context
	InvoiceService service.InvoiceService
	Session        *sessions.Session
}

/**
 * url: /v1/invoice
 * type：GET
 * descs：获取所有请求功能
 */
func (in *InvoiceController) Get() mvc.Result {
	const COMMENT = "method:Get url:/v1/invoice Controller:InvoiceController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := in.Context.GetHeader("Authorization")
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

	invlice := in.InvoiceService.GetInvoiceAll()
	if invlice == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_INVOICEGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_INVOICEGET),
			},
		}
	}

	//将查询到的见积数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range invlice {
		respList = append(respList, item.InvoiceToRespDesc())
	}

	//返回见积列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_INVOICEGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_INVOICEGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/invoice/all/{project_code}
 * type：GET
 * descs：通过案件CD获取所有请求书功能
 */
func (in *InvoiceController) GetAllByProjectCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/invoice/all/{project_code} Controller:InvoiceController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := in.Context.GetHeader("Authorization")
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

	projectCode := in.Context.Params().Get("project_code")
	invoice := in.InvoiceService.GetInvoices(projectCode)

	if invoice == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_INVOICEGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_INVOICEGET),
			},
		}
	}

	//将查询到的请求书数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range invoice {
		respList = append(respList, item.InvoiceToRespDesc())
	}

	//返回请求书列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_INVOICEGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_INVOICEGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/invoice/one/{invoice_code}
 * type：GET
 * descs：通过检收CD获取某一条请求书信息
 */
func (in *InvoiceController) GetOneByInvoiceCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/invoice/one/{invoice_code} Controller:InvoiceController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := in.Context.GetHeader("Authorization")
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

	acceptanceCode := in.Context.Params().Get("invoice_code")
	invoice := in.InvoiceService.GetInvoice(acceptanceCode)

	if invoice == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_INVOICEGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_INVOICEGET),
			},
		}
	}

	//将查询到的检收数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range invoice {
		respList = append(respList, item.InvoiceToRespDesc())
	}

	//返回检收列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_INVOICEGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_INVOICEGET),
			"data":    &respList,
		},
	}
}

/**
 * 即将添加的请求书记录实体
 */
type AddInvoiceEntity struct {
	Id           int64  `json:"id"`
	InvoiceCode  string `json:"invoice_code"`
	DeliveryCode string `json:"delivery_code"`
	EstimateCode string `json:"estimate_code"`
	ProjectCode  string `json:"project_code"`
	ProjectName  string `json:"project_name"`
	CustomerName string `json:"customer_name"`
	InvoiceDate  string `json:"invoice_date"`
	SubTotal     string `json:"sub_total"`
	Tax          string `json:"tax"`
	Total        string `json:"total"`
	BankName     string `json:"bank_name"`
	BankNumber   string `json:"bank_number"`
	BankUser     string `json:"bank_user"`
	Remarks      string `json:"remarks"`
	Note         string `json:"note"`
	CreatedBy    string `json:"created_by"`
	ModifiedBy   string `json:"modified_by"`
}

/**
 * url: /v1/invoice
 * type：POST
 * descs：添加请求功能
 */
func (in *InvoiceController) Post() mvc.Result {
	const COMMENT = "method:Post url:/v1/invoice Controller:InvoiceController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := in.Context.GetHeader("Authorization")
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

	var invoiceEntity AddInvoiceEntity
	err = in.Context.ReadJSON(&invoiceEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_INVOICEADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_INVOICEADD),
			},
		}
	}

	var invoiceInfo model.Invoice

	invoiceInfo.InvoiceCode = invoiceEntity.InvoiceCode
	invoiceInfo.DeliveryCode = invoiceEntity.DeliveryCode
	invoiceInfo.EstimateCode = invoiceEntity.EstimateCode
	invoiceInfo.ProjectCode = invoiceEntity.ProjectCode
	invoiceInfo.ProjectName = invoiceEntity.ProjectName
	invoiceInfo.CustomerName = invoiceEntity.CustomerName
	invoiceInfo.InvoiceDate = invoiceEntity.InvoiceDate
	invoiceInfo.SubTotal = invoiceEntity.SubTotal
	invoiceInfo.Tax = invoiceEntity.Tax
	invoiceInfo.Total = invoiceEntity.Total
	invoiceInfo.BankName = invoiceEntity.BankName
	invoiceInfo.BankNumber = invoiceEntity.BankNumber
	invoiceInfo.BankUser = invoiceEntity.BankUser
	invoiceInfo.Remarks = invoiceEntity.Remarks
	invoiceInfo.Note = invoiceEntity.Note
	invoiceInfo.CreatedBy = invoiceEntity.CreatedBy

	isSuccess := in.InvoiceService.SaveInvoice(invoiceInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_INVOICEADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_INVOICEADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_INVOICEADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_INVOICEADD),
		},
	}
}

/**
 * url: /v1/invoice
 * type：PUT
 * descs：更新请求功能
 */
func (in *InvoiceController) Put() mvc.Result {
	const COMMENT = "method:Put url:/v1/invoice Controller:InvoiceController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := in.Context.GetHeader("Authorization")
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

	var invoiceEntity AddInvoiceEntity
	err = in.Context.ReadJSON(&invoiceEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_INVOICEUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_INVOICEUPDATE),
			},
		}
	}

	var invoiceInfo model.Invoice

	invoiceInfo.InvoiceCode = invoiceEntity.InvoiceCode
	invoiceInfo.DeliveryCode = invoiceEntity.DeliveryCode
	invoiceInfo.EstimateCode = invoiceEntity.EstimateCode
	invoiceInfo.ProjectCode = invoiceEntity.ProjectCode
	invoiceInfo.ProjectName = invoiceEntity.ProjectName
	invoiceInfo.CustomerName = invoiceEntity.CustomerName
	invoiceInfo.InvoiceDate = invoiceEntity.InvoiceDate
	invoiceInfo.SubTotal = invoiceEntity.SubTotal
	invoiceInfo.Tax = invoiceEntity.Tax
	invoiceInfo.Total = invoiceEntity.Total
	invoiceInfo.BankName = invoiceEntity.BankName
	invoiceInfo.BankNumber = invoiceEntity.BankNumber
	invoiceInfo.BankUser = invoiceEntity.BankUser
	invoiceInfo.Remarks = invoiceEntity.Remarks
	invoiceInfo.Note = invoiceEntity.Note
	invoiceInfo.ModifiedBy = invoiceEntity.ModifiedBy

	isSuccess := in.InvoiceService.UpdateInvoice(invoiceEntity.InvoiceCode, invoiceInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_INVOICEUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_INVOICEUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_INVOICEUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_INVOICEUPDATE),
		},
	}
}
