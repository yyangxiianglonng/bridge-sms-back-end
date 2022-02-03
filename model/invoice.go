package model

import "time"

type Invoice struct {
	Id            int64     `xorm:"pk unique autoincr" json:"id"`
	InvoiceCode   string    `xorm:"comment('請求ード')" json:"invoice_code"`
	DeliveryCode  string    `xorm:"comment('納品コード')" json:"delivery_code"`
	EstimateCode  string    `xorm:"comment('見積コード')" json:"estimate_code"`
	ProjectCode   string    `xorm:"comment('案件コード')" json:"project_code"`
	ProjectName   string    `xorm:"comment('案件名')" json:"project_name"`
	CustomerName  string    `xorm:"comment('得意先名')" json:"customer_name"`
	InvoiceDate   string    `xorm:"comment('請求日付')" json:"invoice_date"`
	SubTotal      string    `xorm:"comment('税抜合計')" json:"sub_total"`
	Tax           string    `xorm:"comment('消費税')" json:"tax"`
	Total         string    `xorm:"comment('合計')" json:"total"`
	BankName      string    `xorm:"comment('銀行名')" json:"bank_name"`
	BankNumber    string    `xorm:"comment('口座')" json:"bank_number"`
	BankUser      string    `xorm:"comment('口座名')" json:"bank_user"`
	Remarks       string    `xorm:"comment('備考')" json:"remarks"`
	Note          string    `xorm:"comment('注')" json:"note"`
	InvoicePdfNum string    `xorm:"comment('請求書No.')" json:"invoice_pdf_num"`
	CreatedAt     time.Time `xorm:"created comment('作成時間')" json:"created_at"`
	CreatedBy     string    `xorm:"comment('作成者')" json:"created_by"`
	ModifiedAt    time.Time `xorm:"updated comment('更新時間')" json:"modified_at"`
	ModifiedBy    string    `xorm:"comment('更新者')" json:"modified_by"`
	DeletedAt     time.Time `xorm:"deleted comment('削除時間')" json:"deleted_at"`
	DeletedBy     string    `xorm:"comment('削除者')" json:"deleted_by"`
	IsDelete      int64     `json:"is_delete"`
}

type InvoiceDetail struct {
	Id                  int64     `xorm:"pk unique autoincr" json:"id"`
	Index               *string   `json:"index"`
	InvoiceCode         *string   `xorm:"comment('請求ード')" json:"invoice_code"`
	InvoiceDetailsCode  *string   `xorm:"comment('請求詳細ード')" json:"invoice_details_code"`
	EstimateDetailsCode *string   `xorm:"comment('見積詳細ード')" json:"estimate_details_code"`
	EstimateCode        *string   `xorm:"comment('見積コード')" json:"estimate_code"`
	ProductCode         *string   `xorm:"comment('商品コード')" json:"product_code"`
	ProductName         *string   `xorm:"comment('商品名')" json:"product_name"`
	Quantity            *string   `xorm:"comment('数量')" json:"quantity"`
	Price               *string   `xorm:"comment('単価')" json:"price"`
	SubTotal            *string   `xorm:"comment('合計')" json:"sub_total"`
	Tax                 *string   `xorm:"comment('消費税')" json:"tax"`
	Total               *string   `xorm:"comment('税込合計')" json:"total"`
	Remarks             *string   `xorm:"comment('備考')" json:"remarks"`
	MainFlag            *bool     `xorm:"comment('保守フラグ')" json:"main_flag"`
	CreatedAt           time.Time `xorm:"created comment('作成時間')" json:"created_at"`
	CreatedBy           *string   `xorm:"comment('作成者')" json:"created_by"`
	ModifiedAt          time.Time `xorm:"updated comment('更新時間')" json:"modified_at"`
	ModifiedBy          *string   `xorm:"comment('更新者')" json:"modified_by"`
	DeletedAt           time.Time `xorm:"deleted comment('削除時間')" json:"deleted_at"`
	DeletedBy           *string   `xorm:"comment('削除者')" json:"deleted_by"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (invoice *Invoice) InvoiceToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":              invoice.Id,
		"invoice_code":    invoice.InvoiceCode,
		"delivery_code":   invoice.DeliveryCode,
		"estimate_code":   invoice.EstimateCode,
		"project_code":    invoice.ProjectCode,
		"project_name":    invoice.ProjectName,
		"customer_name":   invoice.CustomerName,
		"invoice_date":    invoice.InvoiceDate,
		"sub_total":       invoice.SubTotal,
		"tax":             invoice.Tax,
		"total":           invoice.Total,
		"bank_name":       invoice.BankName,
		"bank_number":     invoice.BankNumber,
		"bank_user":       invoice.BankUser,
		"remarks":         invoice.Remarks,
		"note":            invoice.Note,
		"invoice_pdf_num": invoice.InvoicePdfNum,
		"created_by":      invoice.CreatedBy,
		"modified_at":     invoice.ModifiedAt,
		"modified_by":     invoice.ModifiedBy,
		"deleted_at":      invoice.DeletedAt,
		"deleted_by":      invoice.DeletedBy,
		"is_delete":       invoice.IsDelete,
	}
	return
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (InvoiceDetail *InvoiceDetail) InvoiceDetailToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":                    InvoiceDetail.Id,
		"index":                 InvoiceDetail.Index,
		"invoice_code":          InvoiceDetail.InvoiceCode,
		"invoice_details_code":  InvoiceDetail.InvoiceDetailsCode,
		"estimate_details_code": InvoiceDetail.EstimateDetailsCode,
		"estimate_code":         InvoiceDetail.EstimateCode,
		"product_code":          InvoiceDetail.ProductCode,
		"product_name":          InvoiceDetail.ProductName,
		"quantity":              InvoiceDetail.Quantity,
		"price":                 InvoiceDetail.Price,
		"sub_total":             InvoiceDetail.SubTotal,
		"tax":                   InvoiceDetail.Tax,
		"total":                 InvoiceDetail.Total,
		"remarks":               InvoiceDetail.Remarks,
		"main_flag":             InvoiceDetail.MainFlag,
		"created_at":            InvoiceDetail.CreatedAt,
		"created_by":            InvoiceDetail.CreatedBy,
		"modified_at":           InvoiceDetail.ModifiedAt,
		"modified_by":           InvoiceDetail.ModifiedBy,
		"deleted_at":            InvoiceDetail.DeletedAt,
		"deleted_by":            InvoiceDetail.DeletedBy,
	}
	return
}
