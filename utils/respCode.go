package utils

//请求状态码
const (
	RECODE_OK      = 1  //请求成功 正常
	RECODE_FAIL    = 0  //失败
	RECODE_UNLOGIN = -1 //未登录 没有权限
)

//业务逻辑状态码
const (
	RESPMSG_OK   = "1"
	RESPMSG_FAIL = "0"

	//管理员
	RESPMSG_SUCCESSLOGIN    = "SUCCESS_LOGIN"
	RESPMSG_FAILURELOGIN    = "FAILURE_LOGIN"
	RESPMSG_SUCCESSSESSION  = "SUCCESS_SESSION"
	RESPMSG_ERRORSESSION    = "ERROR_SESSION"
	RESPMSG_SIGNOUT         = "SUCCESS_SIGNOUT"
	RESPMSG_HASNOACCESS     = "HAS_NO_ACCESS"
	RESPMSG_ERRORADMINCOUNT = "ERROR_ADMINCOUNT"

	//用户
	RESPMSG_ERROR_USERLIST = "ERROR_USERS"
	RESPMSG_ERROR_USERINFO = "ERROR_USERINFO"

	//获取订单操作
	RESPMSG_ERROR_ORDERLIST  = "ERROR_ORDERS"
	RESPMSG_ERROR_ORDERCOUNT = "ERROR_ORDERCOUNT"
	RESPMSG_ERROR_ORDERINFO  = "ERROR_ORDERINFO"

	//商家
	RESPMSG_ERROR_RESTLIST       = "ERROR_RESTAURANTS"
	RESPMSG_SUCCESS_ADDREST      = "ADD_RESTUANT_SUCCESS"
	RESPMSG_FAIL_ADDREST         = "ADD_RESTUANT_FAIL"
	RESPMSG_ERROR_RESTAURANTINFO = "ERROR_RESTAURANTINFO"
	RESPMSG_SUCCESS_DELETESHOP   = "SUCCESS_DELETESHOP"
	RESPMSG_ERROR_SEARCHADDRESS  = "ERROR_SEARCHADDRESS"

	//食品
	RESPMSG_ERROR_FOODLIST   = "ERROR_FOODS"
	RESPMSG_ERROR_FOODADD    = "ERROR_ADDFOOD"
	RESPMSG_SUCCESS_FOODADD  = "SUCCESS_ADDFOOD"
	RESPMSG_ERROR_FOODDELE   = "ERROR_DELEFOOD"
	RESPMSG_SUCCESS_FOODDELE = "SUCCESS_DELEFOOD"

	//案件
	RESPMSG_SUCCESS_PROJECTADD    = "SUCCESS_PROJECTADD"
	RESPMSG_ERROR_PROJECTADD      = "ERROR_PROJECTADD"
	RESPMSG_SUCCESS_PROJECTUPDATE = "SUCCESS_PROJECTUPDATE"
	RESPMSG_ERROR_PROJECTUPDATE   = "ERROR_PROJECTUPDATE"
	RESPMSG_SUCCESS_PROJECTGET    = "SUCCESS_PROJECTGET"
	RESPMSG_ERROR_PROJECTGET      = "ERROR_PROJECTGET"

	//见积
	RESPMSG_SUCCESS_ESTIMATEADD    = "SUCCESS_ESTIMATEADD"
	RESPMSG_ERROR_ESTIMATEADD      = "ERROR_ESTIMATEADD"
	RESPMSG_SUCCESS_ESTIMATEUPDATE = "SUCCESS_ESTIMATEUPDATE"
	RESPMSG_ERROR_ESTIMATEUPDATE   = "ERROR_ESTIMATEUPDATE"
	RESPMSG_SUCCESS_ESTIMATEGET    = "SUCCESS_ESTIMATEGET"
	RESPMSG_ERROR_ESTIMATEGET      = "ERROR_ESTIMATEGET"
	RESPMSG_SUCCESS_ESTIMATEDELETE = "SUCCESS_ESTIMATEDELETE"
	RESPMSG_ERROR_ESTIMATEDELETE   = "ERROR_ESTIMATEDELETE"

	//见积详细
	RESPMSG_SUCCESS_ESTIMATEDETAILADD    = "SUCCESS_ESTIMATEDETAILADD"
	RESPMSG_ERROR_ESTIMATEDETAILADD      = "ERROR_ESTIMATEDETAILADD"
	RESPMSG_SUCCESS_ESTIMATEDETAILUPDATE = "SUCCESS_ESTIMATEDETAILUPDAT"
	RESPMSG_ERROR_ESTIMATEDETAILUPDATE   = "ERROR_ESTIMATEDETAILUPDATE"
	RESPMSG_SUCCESS_ESTIMATEDETAILGET    = "SUCCESS_ESTIMATEDETAILGET"
	RESPMSG_ERROR_ESTIMATEDETAILGET      = "ERROR_ESTIMATEDETAILGET"
	RESPMSG_SUCCESS_ESTIMATEDETAILDELETE = "SUCCESS_ESTIMATEDETAILDELETE"
	RESPMSG_ERROR_ESTIMATEDETAILDELETE   = "ERROR_ESTIMATEDETAILDELETE"

	//注文
	RESPMSG_SUCCESS_ORDERADD    = "SUCCESS_ORDERADD"
	RESPMSG_ERROR_ORDERADD      = "ERROR_ORDERADD"
	RESPMSG_SUCCESS_ORDERUPDATE = "SUCCESS_ORDERUPDATE"
	RESPMSG_ERROR_ORDERUPDATE   = "ERROR_ORDERUPDATE"
	RESPMSG_SUCCESS_ORDERGET    = "SUCCESS_ORDERGET"
	RESPMSG_ERROR_ORDERGET      = "ERROR_ORDERGET"
	RESPMSG_SUCCESS_ORDERDELETE = "SUCCESS_ORDERDELETE"
	RESPMSG_ERROR_ORDERDELETE   = "ERROR_ORDERDELETE"

	//客户
	RESPMSG_SUCCESS_CUSTOMERADD    = "SUCCESS_CUSTOMERADD"
	RESPMSG_ERROR_CUSTOMERADD      = "ERROR_CUSTOMERADD"
	RESPMSG_SUCCESS_CUSTOMERUPDATE = "SUCCESS_CUSTOMERUPDATE"
	RESPMSG_ERROR_CUSTOMERUPDATE   = "ERROR_CUSTOMERUPDATE"
	RESPMSG_SUCCESS_CUSTOMERGET    = "SUCCESS_CUSTOMERGET"
	RESPMSG_ERROR_CUSTOMERGET      = "ERROR_CUSTOMERGET"
	RESPMSG_SUCCESS_CUSTOMERDELETE = "SUCCESS_CUSTOMERDELETE"
	RESPMSG_ERROR_CUSTOMERDELETE   = "ERROR_CUSTOMERDELETE"

	//商品分类
	RESPMSG_SUCCESS_CATEGORYADD     = "SUCCESS_CATEGORYADD"
	RESPMSG_ERROR_CATEGORYADD       = "ERROR_CATEGORYADD"
	RESPMSG_SUCCESS_CATEGORYUPDATE  = "SUCCESS_CATEGORYUPDATE"
	RESPMSG_ERROR_CATEGORYUPDATE    = "ERROR_CATEGORYUPDATE"
	RESPMSG_SUCCESS_CATEGORIEGET    = "SUCCESS_CATEGORIEGET"
	RESPMSG_ERROR_CATEGORIEGET      = "ERROR_CATEGORIEGET"
	RESPMSG_SUCCESS_CATEGORIEDELETE = "SUCCESS_CATEGORIEDELETE"
	RESPMSG_ERROR_CATEGORIEDELETE   = "ERROR_CATEGORIEDELETE"

	//商品
	RESPMSG_SUCCESS_PRODUCTADD    = "SUCCESS_PRODUCTADD"
	RESPMSG_ERROR_PRODUCTADD      = "ERROR_PRODUCTADD"
	RESPMSG_SUCCESS_PRODUCTUPDATE = "SUCCESS_PRODUCTUPDATE"
	RESPMSG_ERROR_PRODUCTUPDATE   = "ERROR_PRODUCTUPDATE"
	RESPMSG_SUCCESS_PRODUCTGET    = "SUCCESS_PRODUCTGET"
	RESPMSG_ERROR_PRODUCTGET      = "ERROR_PRODUCTGET"
	RESPMSG_SUCCESS_PRODUCTDELETE = "SUCCESS_PRODUCTDELETE"
	RESPMSG_ERROR_PRODUCTDELETE   = "ERROR_PRODUCTDELETE"

	//用户
	RESPMSG_SUCCESS_USERADD    = "SUCCESS_USERADD"
	RESPMSG_ERROR_USERADD      = "ERROR_USERADD"
	RESPMSG_SUCCESS_USERUPDATE = "SUCCESS_USERUPDATE"
	RESPMSG_ERROR_USERUPDATE   = "ERROR_USERUPDATE"
	RESPMSG_SUCCESS_USERGET    = "SUCCESS_USERGET"
	RESPMSG_ERROR_USERGET      = "ERROR_USERGET"
	RESPMSG_SUCCESS_USERDELETE = "SUCCESS_USERDELETE"
	RESPMSG_ERROR_USERDELETE   = "ERROR_USERDELETE"
	RESPMSG_SUCCESS_USERLOGIN  = "SUCCESS_USERLOGIN"
	RESPMSG_ERROR_USERLOGIN    = "ERROR_USERLOGIN"
	RESPMSG_SUCCESS_USER       = "SUCCESS_USER"
	RESPMSG_ERROR_USER         = "ERROR_USER"
	RESPMSG_SUCCESS_PASSWORD   = "SUCCESS_PASSWORD"
	RESPMSG_ERROR_PASSWORD     = "ERROR_PASSWORD"

	//文件操作
	RESPMSG_ERROR_PICTUREADD  = "ERROR_PICTUREADD"
	RESPMSG_ERROR_PICTURETYPE = "ERROR_PICTURETYPE"
	RESPMSG_ERROR_PICTURESIZE = "ERROR_PICTURESIZE"

	//城市基础表
	RESPMSG_ERROR_CITYLIST = "ERRROR_CITYLIST"

	//未登陆
	EEROR_UNLOGIN = "ERROR_UNLOGIN"

	RECODE_UNKNOWERR = "8000"
)

//业务逻辑状态信息描述
var recodeText = map[string]string{
	RESPMSG_OK:    "成功",
	RESPMSG_FAIL:  "失败",
	EEROR_UNLOGIN: "未登陆无操作权限，请先登陆", //未登陆 没有权限

	//管理员
	RESPMSG_SUCCESSLOGIN:    "管理员登陆成功",
	RESPMSG_FAILURELOGIN:    "管理员账号或密码错误，登陆失败",
	RESPMSG_SUCCESSSESSION:  "获取管理员信息成功",
	RESPMSG_ERRORSESSION:    "获取管理员信息失败",
	RESPMSG_HASNOACCESS:     "亲，您的权限不足",
	RESPMSG_SIGNOUT:         "退出成功",
	RESPMSG_ERRORADMINCOUNT: "获取管理员总数失败",

	//用户
	RESPMSG_ERROR_USERLIST: "查询用户失败",
	RESPMSG_ERROR_USERINFO: "查询用户信息失败",

	//订单
	RESPMSG_ERROR_ORDERLIST:  "获取订单失败",
	RESPMSG_ERROR_ORDERCOUNT: "获取用户订单数量失败",
	RESPMSG_ERROR_ORDERINFO:  "获取订单信息失败",

	//商家
	RESPMSG_ERROR_RESTLIST:       "查询商家店铺失败",
	RESPMSG_SUCCESS_ADDREST:      "添加商家店铺成功",
	RESPMSG_FAIL_ADDREST:         "添加商家店铺失败",
	RESPMSG_ERROR_RESTAURANTINFO: "获取商家信息失败",
	RESPMSG_SUCCESS_DELETESHOP:   "删除商家成功",
	RESPMSG_ERROR_SEARCHADDRESS:  "搜索地址失败",

	//食品
	RESPMSG_ERROR_FOODLIST:   "查询食品列表失败",
	RESPMSG_ERROR_FOODADD:    "添加食品失败",
	RESPMSG_SUCCESS_FOODADD:  "添加食品成功",
	RESPMSG_ERROR_FOODDELE:   "删除食品记录失败",
	RESPMSG_SUCCESS_FOODDELE: "删除食品记录成功",

	//案件
	RESPMSG_SUCCESS_PROJECTADD:    "案件追加成功",
	RESPMSG_ERROR_PROJECTADD:      "案件追加失敗",
	RESPMSG_SUCCESS_PROJECTUPDATE: "案件更新成功",
	RESPMSG_ERROR_PROJECTUPDATE:   "案件更新失敗",
	RESPMSG_SUCCESS_PROJECTGET:    "案件取得成功",
	RESPMSG_ERROR_PROJECTGET:      "案件取得失敗",

	//见积
	RESPMSG_SUCCESS_ESTIMATEADD:    "見積追加成功",
	RESPMSG_ERROR_ESTIMATEADD:      "見積追加失敗",
	RESPMSG_SUCCESS_ESTIMATEUPDATE: "見積更新成功",
	RESPMSG_ERROR_ESTIMATEUPDATE:   "見積更新失敗",
	RESPMSG_SUCCESS_ESTIMATEGET:    "見積取得成功",
	RESPMSG_ERROR_ESTIMATEGET:      "見積取得失敗",
	RESPMSG_SUCCESS_ESTIMATEDELETE: "見積削除成功",
	RESPMSG_ERROR_ESTIMATEDELETE:   "見積削除失敗",

	//见积详细
	RESPMSG_SUCCESS_ESTIMATEDETAILADD:    "見積詳細追加成功",
	RESPMSG_ERROR_ESTIMATEDETAILADD:      "見積詳細追加失敗",
	RESPMSG_SUCCESS_ESTIMATEDETAILUPDATE: "見積詳細更新成功",
	RESPMSG_ERROR_ESTIMATEDETAILUPDATE:   "見積詳細更新失敗",
	RESPMSG_SUCCESS_ESTIMATEDETAILGET:    "見積詳細取得成功",
	RESPMSG_ERROR_ESTIMATEDETAILGET:      "見積詳細取得失敗",
	RESPMSG_SUCCESS_ESTIMATEDETAILDELETE: "見積詳細削除成功",
	RESPMSG_ERROR_ESTIMATEDETAILDELETE:   "見積詳細削除失敗",

	//注文
	RESPMSG_SUCCESS_ORDERADD:    "注文追加成功",
	RESPMSG_ERROR_ORDERADD:      "注文追加失敗",
	RESPMSG_SUCCESS_ORDERUPDATE: "注文更新成功",
	RESPMSG_ERROR_ORDERUPDATE:   "注文更新失敗",
	RESPMSG_SUCCESS_ORDERGET:    "注文取得成功",
	RESPMSG_ERROR_ORDERGET:      "注文取得失敗",
	RESPMSG_SUCCESS_ORDERDELETE: "注文削除成功",
	RESPMSG_ERROR_ORDERDELETE:   "注文削除失敗",

	//客户
	RESPMSG_SUCCESS_CUSTOMERADD:    "顧客追加成功",
	RESPMSG_ERROR_CUSTOMERADD:      "顧客追加失敗",
	RESPMSG_SUCCESS_CUSTOMERUPDATE: "顧客更新成功",
	RESPMSG_ERROR_CUSTOMERUPDATE:   "顧客更新失敗",
	RESPMSG_SUCCESS_CUSTOMERGET:    "顧客取得成功",
	RESPMSG_ERROR_CUSTOMERGET:      "顧客取得失敗",
	RESPMSG_SUCCESS_CUSTOMERDELETE: "顧客削除成功",
	RESPMSG_ERROR_CUSTOMERDELETE:   "顧客削除失敗",

	//商品分类
	RESPMSG_SUCCESS_CATEGORYADD:     "カテゴリ追加成功",
	RESPMSG_ERROR_CATEGORYADD:       "カテゴリ追加失敗",
	RESPMSG_SUCCESS_CATEGORYUPDATE:  "カテゴリ更新成功",
	RESPMSG_ERROR_CATEGORYUPDATE:    "カテゴリ更新失敗",
	RESPMSG_SUCCESS_CATEGORIEGET:    "カテゴリ取得成功",
	RESPMSG_ERROR_CATEGORIEGET:      "カテゴリ取得失敗",
	RESPMSG_SUCCESS_CATEGORIEDELETE: "カテゴリ削除成功",
	RESPMSG_ERROR_CATEGORIEDELETE:   "カテゴリ削除失敗",

	//商品
	RESPMSG_SUCCESS_PRODUCTADD:    "商品追加成功",
	RESPMSG_ERROR_PRODUCTADD:      "商品追加失敗",
	RESPMSG_SUCCESS_PRODUCTUPDATE: "商品更新成功",
	RESPMSG_ERROR_PRODUCTUPDATE:   "商品更新失敗",
	RESPMSG_SUCCESS_PRODUCTGET:    "商品取得成功",
	RESPMSG_ERROR_PRODUCTGET:      "商品取得失敗",
	RESPMSG_SUCCESS_PRODUCTDELETE: "商品削除成功",
	RESPMSG_ERROR_PRODUCTDELETE:   "商品削除失敗",

	//用户
	RESPMSG_SUCCESS_USERADD:    "User ID追加成功",
	RESPMSG_ERROR_USERADD:      "User ID追加失敗",
	RESPMSG_SUCCESS_USERUPDATE: "User情報更新成功",
	RESPMSG_ERROR_USERUPDATE:   "User情報更新失敗",
	RESPMSG_SUCCESS_USERGET:    "User情報取得成功",
	RESPMSG_ERROR_USERGET:      "User情報取得失敗",
	RESPMSG_SUCCESS_USERDELETE: "User情報削除成功",
	RESPMSG_ERROR_USERDELETE:   "User削除失敗",
	RESPMSG_SUCCESS_USERLOGIN:  "ログイン成功",
	RESPMSG_ERROR_USERLOGIN:    "ログイン失敗",
	RESPMSG_SUCCESS_USER:       "User ID存在します",
	RESPMSG_ERROR_USER:         "User ID存在しません",
	RESPMSG_SUCCESS_PASSWORD:   "Password認証できました",
	RESPMSG_ERROR_PASSWORD:     "Passwordが間違っています",

	//图片操作
	RESPMSG_ERROR_PICTUREADD:  "图片上传失败",
	RESPMSG_ERROR_PICTURETYPE: "只支持PNG,JPG,JPEG格式的图片",
	RESPMSG_ERROR_PICTURESIZE: "图片尺寸太大,请保证在2M一下",

	//城市
	RESPMSG_ERROR_CITYLIST: "获取城市信息失败",

	//其他错误
	RECODE_UNKNOWERR: "服务器未知错误",
}

func Recode2Text(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
