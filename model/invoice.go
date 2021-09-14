package model

import "time"

type Invoice struct {
	Id           int64     `xorm:"pk unique autoincr"json:"id"`
	InvoiceCode  string    `json:"invoice_code"`
	DeliveryCode string    `json:"delivery_code"`
	ProjectCode  string    `json:"project_code"`
	ProjectName  string    `json:"project_name"`
	CustomerName string    `json:"customer_name"`
	InvoiceDate  string    `json:"invoice_date"`
	SubTotal     string    `json:"sub_total"`
	Tax          string    `json:"tax"`
	Total        string    `json:"total"`
	BankName     string    `json:"bank_name"`
	BankNumber   string    `json:"bank_number"`
	BankUser     string    `json:"bank_user"`
	Remarks      string    `json:"remarks"`
	Note         string    `json:"note"`
	CreatedAt    time.Time `xorm:"created" json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	ModifiedAt   time.Time `xorm:"updated" json:"modified_at"`
	ModifiedBy   string    `json:"modified_by"`
	DeletedAt    time.Time `xorm:"deleted" json:"deleted_at"`
	DeletedBy    string    `json:"deleted_by"`
	IsDelete     int64     `json:"is_delete"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (invoice *Invoice) InvoiceToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":            invoice.Id,
		"invoice_code":  invoice.InvoiceCode,
		"delivery_code": invoice.DeliveryCode,
		"project_code":  invoice.ProjectCode,
		"project_name":  invoice.ProjectName,
		"customer_name": invoice.CustomerName,
		"invoice_date":  invoice.InvoiceDate,
		"sub_total":     invoice.SubTotal,
		"tax":           invoice.Tax,
		"total":         invoice.Total,
		"bank_name":     invoice.BankName,
		"bank_number":   invoice.BankNumber,
		"bank_user":     invoice.BankUser,
		"remarks":       invoice.Remarks,
		"note":          invoice.Note,
		"created_by":    invoice.CreatedBy,
		"modified_at":   invoice.ModifiedAt,
		"modified_by":   invoice.ModifiedBy,
		"deleted_at":    invoice.DeletedAt,
		"deleted_by":    invoice.DeletedBy,
		"is_delete":     invoice.IsDelete,
	}
	return
}
