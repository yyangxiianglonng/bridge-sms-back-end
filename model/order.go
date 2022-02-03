package model

import "time"

/**
 * 注文信息结构体,用于生成注文信息表.
 */
type Order struct {
	Id                   int64     `xorm:"pk unique autoincr" json:"id"`
	OrderCode            *string   `json:"order_code"`
	EstimateCode         *string   `json:"estimate_code"`
	EstimateName         *string   `json:"estimate_name"`
	ProjectCode          *string   `json:"project_code"`
	ProjectName          *string   `json:"project_name"`
	EstimateOfOrder      *string   `json:"estimate_of_order"`
	CustomerName         *string   `json:"customer_name"`
	CustomerAddress      *string   `json:"customer_address"`
	Work                 *string   `json:"work"`
	Deliverables         *string   `json:"deliverables"`
	WorkTime             *string   `json:"work_time"`
	Personnel1           *string   `json:"personnel1"`
	Personnel2           *string   `json:"personnel2"`
	DeliverableSpace     *string   `json:"deliverable_space"`
	Commission           *string   `json:"commission"`
	PaymentDate          *string   `json:"payment_date"`
	AcceptanceConditions *string   `json:"acceptance_conditions"`
	Other                *string   `json:"other"`
	Note                 *string   `json:"note"`
	OrderPdfNum          *string   `json:"order_pdf_num"`
	InvoiceOrderPdfNum   *string   `json:"invoice_order_pdf_num"`
	InvoiceOrderDate     time.Time `xorm:"comment('注文請書日付')" json:"invoice_order_date"`
	CreatedAt            time.Time `xorm:"created" json:"created_at"`
	CreatedBy            *string   `json:"created_by"`
	ModifiedAt           time.Time `xorm:"updated" json:"modified_at"`
	ModifiedBy           *string   `json:"modified_by"`
	DeletedAt            time.Time `xorm:"deleted" json:"deleted_at"`
	DeletedBy            *string   `json:"deleted_by"`
	IsDelete             *int64    `json:"is_delete"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (order *Order) OrderToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":                    order.Id,
		"order_code":            order.OrderCode,
		"estimate_code":         order.EstimateCode,
		"estimate_name":         order.EstimateName,
		"project_code":          order.ProjectCode,
		"project_name":          order.ProjectName,
		"estimate_of_order":     order.EstimateOfOrder,
		"customer_name":         order.CustomerName,
		"customer_address":      order.CustomerAddress,
		"work":                  order.Work,
		"deliverables":          order.Deliverables,
		"work_time":             order.WorkTime,
		"personnel1":            order.Personnel1,
		"personnel2":            order.Personnel2,
		"deliverable_space":     order.DeliverableSpace,
		"commission":            order.Commission,
		"payment_date":          order.PaymentDate,
		"acceptance_conditions": order.AcceptanceConditions,
		"other":                 order.Other,
		"note":                  order.Note,
		"order_pdf_num":         order.OrderPdfNum,
		"invoice_order_pdf_num": order.InvoiceOrderPdfNum,
		"invoice_order_date":    order.InvoiceOrderDate,
		"created_at":            order.CreatedAt,
		"created_by":            order.CreatedBy,
		"modified_at":           order.ModifiedAt,
		"modified_by":           order.ModifiedBy,
		"deleted_at":            order.DeletedAt,
		"deleted_by":            order.DeletedBy,
		"is_delete":             order.IsDelete,
	}
	return
}
