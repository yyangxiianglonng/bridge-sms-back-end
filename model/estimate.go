package model

import (
	"time"
)

/**
 * 见积信息结构体,用于生成见积信息表.
 */
type Estimate struct {
	Id                int64     `xorm:"pk unique autoincr" json:"id"`
	EstimateCode      *string   `xorm:"comment('見積コード')" json:"estimate_code"`
	EstimateName      *string   `xorm:"comment('見積名')" json:"estimate_name"`
	ProjectCode       *string   `xorm:"comment('案件コード')" json:"project_code"`
	ProjectName       *string   `xorm:"comment('案件名')" json:"project_name"`
	CustomerName      *string   `xorm:"comment('得意先名')" json:"customer_name"`
	EstimateStartDate time.Time `xorm:"comment('見積開始日')" json:"estimate_start_date"`
	EstimateEndDate   time.Time `xorm:"comment('見積終了日')" json:"estimate_end_date"`
	Work1             *string   `xorm:"comment('作業内容①')" json:"work1"`
	Work2             *string   `xorm:"comment('作業内容②')" json:"work2"`
	Work3             *string   `xorm:"comment('作業内容③')" json:"work3"`
	Deliverables1     *string   `xorm:"comment('成果物①')" json:"deliverables1"`
	Deliverables2     *string   `xorm:"comment('成果物②')" json:"deliverables2"`
	Deliverables3     *string   `xorm:"comment('成果物③')" json:"deliverables3"`
	Media1            *string   `xorm:"comment('媒体①')" json:"media1"`
	Media2            *string   `xorm:"comment('媒体②')" json:"media2"`
	Media3            *string   `xorm:"comment('媒体③')" json:"media3"`
	Quantity1         *string   `xorm:"comment('部数/数量①')" json:"quantity1"`
	Quantity2         *string   `xorm:"comment('部数/数量②')" json:"quantity2"`
	Quantity3         *string   `xorm:"comment('部数/数量③')" json:"quantity3"`
	DeliveryDate1     *string   `xorm:"comment('納品予定日①')" json:"delivery_date1"`
	DeliveryDate2     *string   `xorm:"comment('納品予定日②')" json:"delivery_date2"`
	DeliveryDate3     *string   `xorm:"comment('納品予定日③')" json:"delivery_date3"`
	WorkSpace         *string   `xorm:"comment('作業場所')" json:"work_space"`
	SubTotal          *string   `xorm:"comment('見積合計金額')" json:"sub_total"`
	Tax               *string   `xorm:"comment('消費税')" json:"tax"`
	Total             *string   `xorm:"comment('税込合計金額')" json:"total"`
	InitialTotal      *string   `xorm:"comment('イニシャル合計')" json:"initial_total"`
	RunningTotal      *string   `xorm:"comment('ランニング合計')" json:"running_total"`
	Supplement        *string   `xorm:"comment('補足事項')" json:"supplement"`
	Remarks           *string   `xorm:"comment('得意先名')" json:"remarks"`
	PaymentConditions *string   `xorm:"comment('支払条件')" json:"payment_conditions"`
	Other             *string   `xorm:"comment('その他費用')" json:"other"`
	EstimatePdfNum    *string   `xorm:"comment('見積書No.')" json:"estimate_pdf_num"`
	CreatedAt         time.Time `xorm:"created comment('作成時間')" json:"created_at"`
	CreatedBy         *string   `xorm:"comment('作成者')" json:"created_by"`
	ModifiedAt        time.Time `xorm:"updated comment('更新時間')" json:"modified_at"`
	ModifiedBy        *string   `xorm:"comment('更新者')" json:"modified_by"`
	DeletedAt         time.Time `xorm:"deleted comment('削除時間')" json:"deleted_at"`
	DeletedBy         *string   `xorm:"comment('削除者')" json:"deleted_by"`
}

/**
 * 见积详细信息结构体,用于生成见积详细信息表.
 */
type EstimateDetail struct {
	Id                  int64     `xorm:"pk unique autoincr" json:"id"`
	Index               *int      `json:"index"`
	EstimateDetailsCode *string   `json:"estimate_details_code"`
	EstimateCode        *string   `json:"estimate_code"`
	ProductCode         *string   `json:"product_code"`
	ProductName         *string   `json:"product_name"`
	Quantity            *string   `json:"quantity"`
	Price               *string   `json:"price"`
	SubTotal            *string   `json:"sub_total"`
	Tax                 *string   `json:"tax"`
	Total               *string   `json:"total"`
	Remarks             *string   `json:"remarks"`
	MainFlag            *bool     `json:"main_flag"`
	CreatedAt           time.Time `xorm:"created" json:"created_at"`
	CreatedBy           *string   `json:"created_by"`
	ModifiedAt          time.Time `xorm:"updated" json:"modified_at"`
	ModifiedBy          *string   `json:"modified_by"`
	DeletedAt           time.Time `xorm:"deleted" json:"deleted_at"`
	DeletedBy           *string   `json:"deleted_by"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (estimate *Estimate) EstimateToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":                  estimate.Id,
		"estimate_code":       estimate.EstimateCode,
		"estimate_name":       estimate.EstimateName,
		"project_code":        estimate.ProjectCode,
		"project_name":        estimate.ProjectName,
		"customer_name":       estimate.CustomerName,
		"estimate_start_date": estimate.EstimateStartDate,
		"estimate_end_date":   estimate.EstimateEndDate,
		"work1":               estimate.Work1,
		"work2":               estimate.Work2,
		"work3":               estimate.Work3,
		"deliverables1":       estimate.Deliverables1,
		"deliverables2":       estimate.Deliverables2,
		"deliverables3":       estimate.Deliverables3,
		"media1":              estimate.Media1,
		"media2":              estimate.Media2,
		"media3":              estimate.Media3,
		"quantity1":           estimate.Quantity1,
		"quantity2":           estimate.Quantity2,
		"quantity3":           estimate.Quantity3,
		"delivery_date1":      estimate.DeliveryDate1,
		"delivery_date2":      estimate.DeliveryDate2,
		"delivery_date3":      estimate.DeliveryDate3,
		"work_space":          estimate.WorkSpace,
		"sub_total":           estimate.SubTotal,
		"tax":                 estimate.Tax,
		"total":               estimate.Total,
		"initial_total":       estimate.InitialTotal,
		"running_total":       estimate.RunningTotal,
		"supplement":          estimate.Supplement,
		"remarks":             estimate.Remarks,
		"payment_conditions":  estimate.PaymentConditions,
		"other":               estimate.Other,
		"estimate_pdf_num":    estimate.EstimatePdfNum,
		"created_at":          estimate.CreatedAt,
		"created_by":          estimate.CreatedBy,
		"modified_at":         estimate.ModifiedAt,
		"modified_by":         estimate.ModifiedBy,
		"deleted_at":          estimate.DeletedAt,
		"deleted_by":          estimate.DeletedBy,
	}
	return
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (estimateDetail *EstimateDetail) EstimateDetailToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":                    estimateDetail.Id,
		"index":                 estimateDetail.Index,
		"estimate_details_code": estimateDetail.EstimateDetailsCode,
		"estimate_code":         estimateDetail.EstimateCode,
		"product_code":          estimateDetail.ProductCode,
		"product_name":          estimateDetail.ProductName,
		"quantity":              estimateDetail.Quantity,
		"price":                 estimateDetail.Price,
		"sub_total":             estimateDetail.SubTotal,
		"tax":                   estimateDetail.Tax,
		"total":                 estimateDetail.Total,
		"remarks":               estimateDetail.Remarks,
		"main_flag":             estimateDetail.MainFlag,
		"created_at":            estimateDetail.CreatedAt,
		"created_by":            estimateDetail.CreatedBy,
		"modified_at":           estimateDetail.ModifiedAt,
		"modified_by":           estimateDetail.ModifiedBy,
		"deleted_at":            estimateDetail.DeletedAt,
		"deleted_by":            estimateDetail.DeletedBy,
	}
	return
}
