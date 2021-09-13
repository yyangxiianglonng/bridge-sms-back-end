package model

import "time"

/**
 * 案件信息结构体,用于生成商品信息表.
 */
type Product struct {
	Id          int64     `xorm:"pk unique autoincr" json:"id"`
	ProductName string    `json:"product_name"`
	Category1   string    `json:"category1"`
	Category2   string    `json:"category2"`
	Category3   string    `json:"category3"`
	Price       string    `json:"price"`
	CreatedAt   time.Time `xorm:"created" json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `xorm:"updated" json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
	DeletedAt   time.Time `xorm:"deleted" json:"deleted_at"`
	DeletedBy   string    `json:"deleted_by"`
	IsDelete    int64     `json:"is_delete"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (product *Product) ProductToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":           product.Id,
		"product_name": product.ProductName,
		"category1":    product.Category1,
		"category2":    product.Category2,
		"category3":    product.Category3,
		"price":        product.Price,
		"created_at":   product.CreatedAt,
		"created_by":   product.CreatedBy,
		"modified_at":  product.ModifiedAt,
		"modified_by":  product.ModifiedBy,
		"deleted_at":   product.DeletedAt,
		"deleted_by":   product.DeletedBy,
		"is_delete":    product.IsDelete,
	}
	return
}
