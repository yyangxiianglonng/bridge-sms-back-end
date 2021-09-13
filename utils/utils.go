package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

/**
 * 格式化数据
 */
func FormatDatetime(time time.Time) string {
	return time.Format("2021-08-25 13:20:00")
}

// 加密密码
func HashAndSalt(password []byte) string {
	hashPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hashPassword)
}

// 验证密码
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
