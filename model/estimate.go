package model

import (
	"time"
)

/**
 * 见积信息结构体,用于生成见积信息表.
 */
type Estimate struct {
	Id                int64     `xorm:"pk unique autoincr" json:"id"`
	EstimateCode      string    `json:"estimate_code"`
	ProjectCode       string    `json:"project_code"`
	ProjectName       string    `json:"project_name"`
	CustomerName      string    `json:"customer_name"`
	EstimateStartDate time.Time `json:"estimate_start_date"`
	EstimateEndDate   time.Time `json:"estimate_end_date"`
	Work1             string    `json:"work_1"`
	Work2             string    `json:"work_2"`
	Work3             string    `json:"work_3"`
	Deliverables1     string    `json:"deliverables_1"`
	Deliverables2     string    `json:"deliverables_2"`
	Deliverables3     string    `json:"deliverables_3"`
	Media1            string    `json:"media_1"`
	Media2            string    `json:"media_2"`
	Media3            string    `json:"media_3"`
	Quantity1         string    `json:"quantity_1"`
	Quantity2         string    `json:"quantity_2"`
	Quantity3         string    `json:"quantity_3"`
	DeliveryDate1     string    `json:"delivery_date_1"`
	DeliveryDate2     string    `json:"delivery_date_2"`
	DeliveryDate3     string    `json:"delivery_date_3"`
	WorkSpace         string    `json:"work_space"`
	SubTotal          string    `json:"sub_total"`
	Tax               string    `json:"tax"`
	Total             string    `json:"total"`
	Supplement        string    `json:"supplement"`
	Remarks           string    `json:"remarks"`
	PaymentConditions string    `json:"payment_conditions"`
	Other             string    `json:"other"`
	CreatedAt         time.Time `xorm:"created" json:"created_at"`
	CreatedBy         string    `json:"created_by"`
	ModifiedAt        time.Time `xorm:"updated" json:"modified_at"`
	ModifiedBy        string    `json:"modified_by"`
	DeletedAt         time.Time `xorm:"deleted" json:"deleted_at"`
	DeletedBy         string    `json:"deleted_by"`
	IsDelete          int64     `json:"is_delete"`
}

/**
 * 见积详细信息结构体,用于生成见积详细信息表.
 */
type EstimateDetail struct {
	Id                  int64     `xorm:"pk unique autoincr" json:"id"`
	EstimateDetailsCode string    `json:"estimate_details_code"`
	EstimateCode        string    `json:"estimate_code"`
	ProductCode         string    `json:"product_code"`
	ProductName         string    `json:"product_name"`
	Quantity            string    `json:"quantity"`
	Price               string    `json:"price"`
	SubTotal            string    `json:"sub_total"`
	Tax                 string    `json:"tax"`
	Total               string    `json:"total"`
	MainFlag            bool      `json:"main_flag"`
	CreatedAt           time.Time `xorm:"created" json:"created_at"`
	CreatedBy           string    `json:"created_by"`
	ModifiedAt          time.Time `xorm:"updated" json:"modified_at"`
	ModifiedBy          string    `json:"modified_by"`
	DeletedAt           time.Time `xorm:"deleted" json:"deleted_at"`
	DeletedBy           string    `json:"deleted_by"`
	IsDelete            int64     `json:"is_delete"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (estimate *Estimate) EstimateToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":                  estimate.Id,
		"estimate_code":       estimate.EstimateCode,
		"project_code":        estimate.ProjectCode,
		"project_name":        estimate.ProjectName,
		"customer_name":       estimate.CustomerName,
		"estimate_start_date": estimate.EstimateStartDate,
		"estimate_end_date":   estimate.EstimateEndDate,
		"work_1":              estimate.Work1,
		"work_2":              estimate.Work2,
		"work_3":              estimate.Work3,
		"deliverables_1":      estimate.Deliverables1,
		"deliverables_2":      estimate.Deliverables2,
		"deliverables_3":      estimate.Deliverables3,
		"media_1":             estimate.Media1,
		"media_2":             estimate.Media2,
		"media_3":             estimate.Media3,
		"quantity_1":          estimate.Quantity1,
		"quantity_2":          estimate.Quantity2,
		"quantity_3":          estimate.Quantity3,
		"delivery_date_1":     estimate.DeliveryDate1,
		"delivery_date_2":     estimate.DeliveryDate2,
		"delivery_date_3":     estimate.DeliveryDate3,
		"work_space":          estimate.WorkSpace,
		"sub_total":           estimate.SubTotal,
		"tax":                 estimate.Tax,
		"total":               estimate.Total,
		"supplement":          estimate.Supplement,
		"remarks":             estimate.Remarks,
		"payment_conditions":  estimate.PaymentConditions,
		"other":               estimate.Other,
		"created_at":          estimate.CreatedAt,
		"created_by":          estimate.CreatedBy,
		"modified_at":         estimate.ModifiedAt,
		"modified_by":         estimate.ModifiedBy,
		"deleted_at":          estimate.DeletedAt,
		"deleted_by":          estimate.DeletedBy,
		"is_delete":           estimate.IsDelete,
	}
	return
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (estimateDetail *EstimateDetail) EstimateDetailToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":                    estimateDetail.Id,
		"estimate_details_code": estimateDetail.EstimateDetailsCode,
		"estimate_code":         estimateDetail.EstimateCode,
		"product_code":          estimateDetail.ProductCode,
		"product_name":          estimateDetail.ProductName,
		"quantity":              estimateDetail.Quantity,
		"price":                 estimateDetail.Price,
		"sub_total":             estimateDetail.SubTotal,
		"tax":                   estimateDetail.Tax,
		"total":                 estimateDetail.Total,
		"main_flag":             estimateDetail.MainFlag,
		"created_at":            estimateDetail.CreatedAt,
		"created_by":            estimateDetail.CreatedBy,
		"modified_at":           estimateDetail.ModifiedAt,
		"modified_by":           estimateDetail.ModifiedBy,
		"deleted_at":            estimateDetail.DeletedAt,
		"deleted_by":            estimateDetail.DeletedBy,
		"is_delete":             estimateDetail.IsDelete,
	}
	return
}
