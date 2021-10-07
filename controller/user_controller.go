package controller

import (
	"main/model"
	"main/service"
	"main/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	//上下文对象
	Context iris.Context
	//UserService
	UserService service.UserService
	//session对象
	Session *sessions.Session
}

func (us *UserController) BeforeActivation(ba mvc.BeforeActivation) {
	//通过project_code获取对应的案件
	ba.Handle("PUT", "/signup/{active_code}", "PutSignup")
}

type UserEntity struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

/**
 * 用户登录功能
 * 接口：/v1/login
 * type：Post
 */
func (uc *UserController) Post() mvc.Result {
	const COMMENT = "method:Post url:/v1/login Controller:UserController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	var userEntity UserEntity
	uc.Context.ReadJSON(&userEntity)

	//数据参数检验
	if userEntity.UserName == "" || userEntity.PassWord == "" {
		iris.New().Logger().Error(COMMENT + "用户名或密码为空")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERLOGIN,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERLOGIN),
			},
		}
	}

	//根据用户名到数据库中查询对应的管理信息
	user, exist := uc.UserService.GetByUserName(userEntity.UserName)
	//用户不存在
	if !exist {
		iris.New().Logger().Error(COMMENT + "用户名不存在")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USER,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USER),
			},
		}
	}

	passwordMatch := utils.ComparePasswords(user.PassWord, []byte(userEntity.PassWord))
	if !passwordMatch {
		iris.New().Logger().Error(COMMENT + "密码不正确")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PASSWORD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PASSWORD),
			},
		}
	}

	userName := uc.UserService.IsActiveByUserName(userEntity.UserName)

	if userName[0].IsActive != true {
		iris.New().Logger().Error(COMMENT + " ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ACTIVE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ACTIVE),
			},
		}
	}

	token, err := utils.GenerateToken(userEntity.UserName, utils.HashAndSalt([]byte(userEntity.PassWord)))
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERLOGIN,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERLOGIN),
			},
		}
	}

	//管理员存在 设置session
	//userByte := admin.Encoder()
	//uc.Session.Set(token, token)

	iris.New().Logger().Info(COMMENT + "end")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":    utils.RECODE_OK,
			"type":      utils.RESPMSG_SUCCESS_USERLOGIN,
			"message":   utils.Recode2Text(utils.RESPMSG_SUCCESS_USERLOGIN),
			"token":     token,
			"full_name": user.FullName,
		},
	}
}

/*
* 即将注册的用户实体
 */
type AddUserEntity struct {
	Id         int64  `json:"id"`
	UserName   string `json:"user_name"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	PassWord   string `json:"pass_word"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	ActiveCode string `json:"active_code"`
}

/**
 * 注册用户功能
 * 接口：/v1/login/signup
 * type：Post
 */
func (uc *UserController) PostSignup() mvc.Result {
	const COMMENT = "method:Post url:/v1/login/signup Controller:UserController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	var userEntity AddUserEntity
	err := uc.Context.ReadJSON(&userEntity)

	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RESPMSG_FAIL,
				"type":    utils.RESPMSG_ERROR_USERADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERADD),
			},
		}

	}

	var userInfo model.User
	userInfo.Id = userEntity.Id
	userInfo.UserName = userEntity.UserName
	userInfo.FullName = userEntity.FullName
	userInfo.Email = userEntity.Email
	userInfo.PassWord = utils.HashAndSalt([]byte(userEntity.PassWord))
	userInfo.CreatedBy = userEntity.CreatedBy
	userInfo.ActiveCode = uuid.NewV4().String()

	//根据用户名到数据库中查询对应的管理信息
	_, existUserName := uc.UserService.GetByUserName(userEntity.UserName)

	//用户名已经被使用
	if existUserName {
		iris.New().Logger().Error(COMMENT + "用户名已经存在")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_SUCCESS_USER,
				"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_USER),
			},
		}
	}

	//根据邮箱地址到数据库中查询对应的管理信息
	_, existEmail := uc.UserService.GetByEmail(userEntity.Email)

	//邮箱已经被使用
	if existEmail {
		iris.New().Logger().Error(COMMENT + "邮箱已经被使用")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_SUCCESS_EMAIL,
				"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_EMAIL),
			},
		}
	}

	isSuccess := uc.UserService.SaveUser(userInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERADD),
			},
		}
	}

	// emailPara := &utils.EmailParam{
	// 	ServerHost: config.InitConfig().Email.ServerHost,
	// 	ServerPort: config.InitConfig().Email.ServerPort,
	// 	FromEmail:  config.InitConfig().Email.FromEmail,
	// 	FromPasswd: config.InitConfig().Email.FromPasswd,
	// 	Toers:      userInfo.Email,
	// 	CCers:      "yangxianglong@bridge.vc",
	// 	ActiveCode: userInfo.ActiveCode,
	// }

	// utils.InitEmail(emailPara)

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_USERADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_USERADD),
		},
	}
}

/**
 * 激活用户功能
 * 接口：/v1/login/signup/{active_code}
 * type：Put
 */
func (uc *UserController) PutSignup() mvc.Result {
	const COMMENT = "method:Put url:/v1/login/signup/{active_code} Controller:UserController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	var userEntity AddUserEntity
	err := uc.Context.ReadJSON(&userEntity)

	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RESPMSG_FAIL,
				"type":    utils.RESPMSG_ERROR_USERADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERADD),
			},
		}
	}
	var userInfo model.User
	userInfo.IsActive = true
	userInfo.ModifiedBy = "root"

	isSuccess := uc.UserService.UpdateUser(userEntity.ActiveCode, userInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_USERADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_USERADD),
		},
	}
}

func (uc *UserController) GetAll() mvc.Result {
	iris.New().Logger().Info("Get 请求,请求路径为UserAll")
	userList := uc.UserService.GetUserList(0, 5)

	if len(userList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERLIST,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERLIST),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, user := range userList {
		respList = append(respList, user.UserToRespDesc())
	}

	//返回用户列表
	return mvc.Response{
		Object: &respList,
	}
}

func (uc *UserController) GetInfo() mvc.Result {
	iris.New().Logger().Info("Get 请求,请求路径为info")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  1,
			"message": "请求成功!",
			"data":    "1002",
		},
	}
}
