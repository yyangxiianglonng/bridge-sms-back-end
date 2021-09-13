package model

import (
	"time"
)

/**
 * 用户信息结构体,用于生成用户信息表
 */
type User struct {
	Id         int64     `xorm:"pk autoincr" json:"id"` //主键 用户ID
	UserName   string    `json:"user_name"`             //用户名称
	FullName   string    `json:"full_name"`             //全名
	Email      string    `json:"email"`                 //用户的邮箱
	PassWord   string    `json:"pass_word"`             //用户的账户密码
	IsActive   bool      `json:"is_active"`             //用户是否激活
	CreatedAt  time.Time `xorm:"created" json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `xorm:"updated" json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
	DeletedAt  time.Time `xorm:"deleted" json:"deleted_at"`
	DeletedBy  string    `json:"deleted_by"`
	IsDelete   int64     `json:"is_delete"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (user *User) UserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":          user.Id,
		"user_name":   user.UserName,
		"full_name":   user.FullName,
		"email":       user.Email,
		"pass_word":   user.PassWord,
		"is_active":   user.IsActive,
		"created_at":  user.CreatedAt,
		"created_by":  user.CreatedBy,
		"modified_at": user.ModifiedAt,
		"modified_by": user.ModifiedBy,
		"deleted_at":  user.DeletedAt,
		"deleted_by":  user.DeletedBy,
		"is_delete":   user.IsDelete,
	}
	return respInfo
}
