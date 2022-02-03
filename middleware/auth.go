package middleware

import (
	"main/utils"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Author(context iris.Context) {
	requestPath := context.Path()
	println("Before the mainHandler: " + requestPath)
	token(context)
	// token := context.GetHeader("Authorization")

	// claim, err := utils.ParseToken(token)
	// if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
	// 	println("Before the mainHandler: " + "token error")
	// 	context.NextOrNotFound() //
	// } else {
	// 	context.Next() // 执行下一个处理器。
	// }
}

func token(context iris.Context) mvc.Result {
	token := context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		println("Before the mainHandler: " + "token error")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	} else {
		context.Next()
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_OK,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}
}
