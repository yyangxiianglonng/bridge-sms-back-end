package utils

import (
	"log"
	"math/rand"
	"time"
	"unicode/utf8"

	uuid "github.com/iris-contrib/go.uuid"
	"golang.org/x/crypto/bcrypt"
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

func RandSeq() string {
	//letters := []rune("abcdefghijklmnopqrstuvwxyz")
	letters := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Uuid() uuid.UUID {
	uuid, _ := uuid.NewV4()
	return uuid
}

// 自动选择合适的位置截取字符串
func CutStringAsSize(str string) (cutSize int) {
	cutSize = 0
	for len(str) > 0 {
		_, size := utf8.DecodeRuneInString(str)
		cutSize += size
		if cutSize >= 72 {
			break
		}
		str = str[size:]
	}
	return
}
