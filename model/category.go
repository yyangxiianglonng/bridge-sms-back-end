package model

import "time"

/**
 * 商品分类信息结构体,用于生成客户信息表.
 */
type Category struct {
	Id           int64     `xorm:"pk unique autoincr" json:"id"`
	CategoryName string    `json:"category_name"`
	ParentId     string    `json:"parent_id"`
	ParentFlg    bool      `json:"parent_flg"`
	Sorting      string    `json:"sorting"`
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
func (category *Category) CategoryToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":            category.Id,
		"category_name": category.CategoryName,
		"parent_id":     category.ParentId,
		"parent_flg":    category.ParentFlg,
		"sorting":       category.Sorting,
		"created_at":    category.CreatedAt,
		"created_by":    category.CreatedBy,
		"modified_at":   category.ModifiedAt,
		"modified_by":   category.ModifiedBy,
		"deleted_at":    category.DeletedAt,
		"deleted_by":    category.DeletedBy,
		"is_delete":     category.IsDelete,
	}
	return
}
