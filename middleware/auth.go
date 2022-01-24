package middleware

import (
	"main/utils"
	"time"

	"github.com/kataras/iris/v12"
)

func Author(context iris.Context) {
	requestPath := context.Path()
	println("Before the mainHandler: " + requestPath)

	token := context.GetHeader("Authorization")

	claim, err := utils.ParseToken(token)
	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		context.EndRequest()
	} else {
		context.Next() // 执行下一个处理器。
	}

	context.Next()
}
