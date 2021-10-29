package service

import (
	"main/model"

	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

type UserService interface {
	//通过管理员用户名+密码 获取用户实体 如果查询到，返回用户实体，并返回true 否则 返回 nil ，false
	GetByUserNameAndPassword(username, password string) (model.User, bool)
	//通过管理员用户名获取用户实体 如果查询到，返回用户实体，并返回true 否则 返回 nil ，false
	GetByUserName(userName string) (model.User, bool)
	//通过管理员邮箱地址获取用户实体 如果查询到，返回用户实体，并返回true 否则 返回 nil ，false
	GetByEmail(emailAddress string) (model.User, bool)
	//判断用户是否被激活
	IsActiveByUserName(userName string) (user []*model.User)
	//用户注册服务接口
	SaveUser(user model.User) bool
	//更新用户服务接口
	UpdateUser(activeCode string, user model.User) bool
	//修改用户密码接口
	ResetPasswor(email string, user model.User) bool
	//更新用户接口用于更新密码
	UpdateUserRandNum(email string, user model.User) bool
	//获取用户日增长统计数据
	GetUserDailyStatisCount(datetime string) int64
	//获取用户总数
	GetUserTotalCount() (int64, error)
	//用户列表
	GetUserList(offset, limit int) []*model.User
}

/**
 * 实例化用户服务结构实体对象
 */
func NewUserService(engine *xorm.Engine) UserService {
	return &userService{
		Engine: engine,
	}
}

/**
 * 用户服务实现结构体
 */
type userService struct {
	Engine *xorm.Engine
}

/**
 * 通过用户名和密码查询用户
 */
func (us *userService) GetByUserNameAndPassword(userName, passWord string) (model.User, bool) {
	var user model.User
	us.Engine.Where(" user_name = ? and pass_word = ? ", userName, passWord).Get(&user)

	return user, user.Id != 0
}

/**
 * 查询用户是否存在
 */
func (us *userService) GetByUserName(userName string) (model.User, bool) {
	var user model.User
	us.Engine.Where(" user_name = ? ", userName).Get(&user)

	return user, user.Id != 0
}

/**
 * 判断用户是否被激活
 */
func (us *userService) IsActiveByUserName(userName string) (user []*model.User) {
	err := us.Engine.Where(" user_name = ? ", userName).Find(&user)
	if err != nil {
		panic(err.Error())
	}
	return
}

/**
 * 邮箱地址是否被使用
 */
func (us *userService) GetByEmail(emailAddress string) (model.User, bool) {
	var user model.User
	us.Engine.Where(" email = ? ", emailAddress).Get(&user)
	return user, user.Id != 0
}

/**
 * 用户注册服务
 */
func (us *userService) SaveUser(user model.User) bool {
	_, err := us.Engine.Insert(&user)
	return err == nil
}

/**
 *激活用户服务
 */
func (us *userService) UpdateUser(email string, user model.User) bool {
	_, err := us.Engine.Where("email = ?", email).Cols("is_active").Update(&user)
	return err == nil
}

/**
 *修改密码服务
 */
func (us *userService) ResetPasswor(email string, user model.User) bool {
	_, err := us.Engine.Where("email = ?", email).Update(&user)
	return err == nil
}

/**
 *更新用户随机数用于修改密码
 */
func (us *userService) UpdateUserRandNum(email string, user model.User) bool {
	_, err := us.Engine.Where("email = ?", email).Update(&user)
	return err == nil
}

/**
 * 请求总的用户数量
 * 返回值：总用户数量
 */
func (us *userService) GetUserTotalCount() (int64, error) {

	//查询del_flag 为0 的用户的总数量；del_flag:0 正常状态；del_flag:1 用户注销或者被删除
	count, err := us.Engine.Where(" del_flag = ? ", 0).Count(new(model.User))
	if err != nil {
		panic(err.Error())
	}
	//用户总数
	return count, nil
}

/**
* 请求用户列表数据
* offset：偏移数量
* limit：一次请求获取的数据条数
 */
func (us *userService) GetUserList(offset, limit int) (userList []*model.User) {

	err := us.Engine.Where("del_flag = ?", 0).Limit(limit, offset).Find(&userList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
	return
}

/**
 * 获取用户日增长统计结果
 * datetime：某一个特殊的日期
 */
func (us *userService) GetUserDailyStatisCount(datetime string) int64 {

	result, err := us.Engine.Count(new(model.User))
	if err != nil {
		panic(err.Error())
	}
	return result
}
