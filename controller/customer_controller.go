package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"main/model"
	"main/service"
	"main/utils"
)

type CustomerController struct {
	Context         iris.Context
	CustomerService service.CustomerService
	Session         sessions.Session
}

func (cu *CustomerController) BeforeActivation(ba mvc.BeforeActivation) {

	//通过customer_code检索顾客信息功能
	ba.Handle("GET", "/{customer_code}", "GetOneByCustomerCode")
}

/**
 * url: /v1/customer
 * type：GET
 * descs：获取所有顾客信息功能
 */
func (cu *CustomerController) Get() mvc.Result {
	const COMMENT = "method:Get url:/v1/customer Controller:CustomerController" + " "
	iris.New().Logger().Info(COMMENT + "Start")
	customerList := cu.CustomerService.GetCustomers()
	if len(customerList) == 0 {
		iris.New().Logger().Error(COMMENT + "ERR")
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
	for _, customer := range customerList {
		respList = append(respList, customer.CustomerToRespDesc())
	}

	//返回案件列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CUSTOMERGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CUSTOMERGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/customer/{customer_code}
 * type：GET
 * descs：通过顾客CD检索顾客信息功能
 */
func (cu *CustomerController) GetOneByCustomerCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/customer/{customer_code} Controller:CustomerController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	customerCode := cu.Context.Params().Get("customer_code")
	customer := cu.CustomerService.GetCustomer(customerCode)

	if customer == nil {
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
	for _, item := range customer {
		respList = append(respList, item.CustomerToRespDesc())
	}

	//返回案件列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CUSTOMERGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CUSTOMERGET),
			"data":    &respList,
		},
	}
}

/*
* 即将添加的顾客记录实体
 */
type AddCustomerEntity struct {
	CustomerCode  string `json:"customer_code"`
	CustomerName  string `json:"customer_name"`
	Department1   string `json:"department1"`
	Department2   string `json:"department2"`
	Department3   string `json:"department3"`
	PersonnelName string `json:"personnel_name"`
	PostalNumber  string `json:"postal_number"`
	Address1      string `json:"address1"`
	Address2      string `json:"address2"`
	Address3      string `json:"address3"`
	Telephone     string `json:"telephone"`
	Fix           string `json:"fix"`
	CreatedBy     string `json:"created_by"`
	IsDelete      int64  `json:"is_delete"`
}

/*
* url: /v1/customer
* type: POST
* descs: 添加顾客功能
 */
func (cu *CustomerController) Post() mvc.Result {
	const COMMENT = "method:Post /v1/customer Controller:CustomerController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	var customerEntity AddCustomerEntity
	err := cu.Context.ReadJSON(&customerEntity)

	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RESPMSG_FAIL,
				"type":    utils.RESPMSG_ERROR_CUSTOMERADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CUSTOMERADD),
			},
		}

	}

	var customerInfo model.Customer
	customerInfo.CustomerCode = customerEntity.CustomerCode
	customerInfo.CustomerName = customerEntity.CustomerName
	customerInfo.Department1 = customerEntity.Department1
	customerInfo.Department2 = customerEntity.Department2
	customerInfo.Department3 = customerEntity.Department3
	customerInfo.PersonnelName = customerEntity.PersonnelName
	customerInfo.PostalNumber = customerEntity.PostalNumber
	customerInfo.Address1 = customerEntity.Address1
	customerInfo.Address2 = customerEntity.Address2
	customerInfo.Address3 = customerEntity.Address3
	customerInfo.Telephone = customerEntity.Telephone
	customerInfo.Fix = customerEntity.Fix
	customerInfo.CreatedBy = customerEntity.CreatedBy
	customerInfo.IsDelete = customerEntity.IsDelete

	isSuccess := cu.CustomerService.SaveCustomer(customerInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CUSTOMERADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CUSTOMERADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CUSTOMERADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CUSTOMERADD),
		},
	}
}

/**
 * url: /v1/customer
 * type：PUT
 * descs：更新顾客信息功能
 */
func (cu *CustomerController) Put() mvc.Result {
	const COMMENT = "method:Put /v1/customer Controller:CustomerController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	var customerEntity AddCustomerEntity
	err := cu.Context.ReadJSON(&customerEntity)

	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RESPMSG_FAIL,
				"type":    utils.RESPMSG_ERROR_CUSTOMERUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CUSTOMERUPDATE),
			},
		}

	}

	var customerInfo model.Customer
	customerInfo.CustomerCode = customerEntity.CustomerCode
	customerInfo.CustomerName = customerEntity.CustomerName
	customerInfo.Department1 = customerEntity.Department1
	customerInfo.Department2 = customerEntity.Department2
	customerInfo.Department3 = customerEntity.Department3
	customerInfo.PersonnelName = customerEntity.PersonnelName
	customerInfo.PostalNumber = customerEntity.PostalNumber
	customerInfo.Address1 = customerEntity.Address1
	customerInfo.Address2 = customerEntity.Address2
	customerInfo.Address3 = customerEntity.Address3
	customerInfo.Telephone = customerEntity.Telephone
	customerInfo.Fix = customerEntity.Fix
	customerInfo.CreatedBy = customerEntity.CreatedBy
	customerInfo.IsDelete = customerEntity.IsDelete

	isSuccess := cu.CustomerService.UpdateCustomer(customerInfo.CustomerCode, customerInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_CUSTOMERUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CUSTOMERUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_CUSTOMERUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CUSTOMERUPDATE),
		},
	}
}
