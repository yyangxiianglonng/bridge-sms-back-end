package main

import (
	"main/config"
	"main/controller"
	"main/datasource"
	"main/service"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	app := newApp()

	//应用App设置
	configation(app)

	//路由设置
	mvcHandle(app)

	addr := config.InitConfig().Host + ":" + config.InitConfig().Port
	app.Run(iris.Addr(addr))
}

/**
*	构建App
 */
func newApp() (app *iris.Application) {

	app = iris.New()

	//设置日志级别 开发阶段为debug
	app.Logger().SetLevel("debug")

	//注册静态资源
	app.HandleDir("/static", "/Users/yangxianglong/Projects/WebstormProjects/bridgef/dist/static")
	//app.HandleDir("/static", "C:/inetpub/bridgesys/dist/static")

	//注册视图文件
	app.RegisterView(iris.HTML("/Users/yangxianglong/Projects/WebstormProjects/bridgef/dist", ".html"))
	//app.RegisterView(iris.HTML("C:/inetpub/bridgesys/dist", ".html"))
	app.Get("/", func(context iris.Context) {
		context.View("index.html")
	})

	return
}

/**
 * MVC 架构模式处理
 */
func mvcHandle(app *iris.Application) {
	sessManager := sessions.New(sessions.Config{
		Cookie:  "secessions",
		Expires: 1 * time.Minute,
	})

	//实例化mysql数据库引擎
	engine := datasource.NewMysqlEngine()

	//用户功能模块
	userService := service.NewUserService(engine)
	user := mvc.New(app.Party("/v1/login"))
	user.Register(
		userService,
		sessManager.Start,
	)
	user.Handle(new(controller.UserController))

	//案件功能模块
	projectService := service.NewProjectService(engine)
	project := mvc.New(app.Party("/v1/project"))
	project.Register(
		projectService,
		sessManager.Start,
	)
	project.Handle(new(controller.ProjectController))

	//见积功能模块
	estimateService := service.NewEstimateService(engine)
	estimate := mvc.New(app.Party("/v1/estimate"))
	estimate.Register(
		estimateService,
		sessManager.Start,
	)
	estimate.Handle(new(controller.EstimateController))

	//注文功能模块
	orderService := service.NewOrderService(engine)
	order := mvc.New(app.Party("/v1/order"))
	order.Register(
		orderService,
		sessManager.Start,
	)
	order.Handle(new(controller.OrderController))

	//纳品功能模块
	deliveryService := service.NewDeliveryService(engine)
	delivery := mvc.New(app.Party("/v1/delivery"))
	delivery.Register(
		deliveryService,
		sessManager.Start,
	)
	delivery.Handle(new(controller.DeliveryController))

	//检收功能模块
	acceptanceService := service.NewAcceptanceService(engine)
	acceptance := mvc.New(app.Party("/v1/acceptance"))
	acceptance.Register(
		acceptanceService,
		sessManager.Start,
	)
	acceptance.Handle(new(controller.AcceptanceController))

	//请求功能模块
	invoiceService := service.NewInvoiceService(engine)
	invoice := mvc.New(app.Party("/v1/invoice"))
	invoice.Register(
		invoiceService,
		sessManager.Start,
	)
	invoice.Handle(new(controller.InvoiceController))

	//客户功能模块
	customerService := service.NewCustomerService(engine)
	customer := mvc.New(app.Party("/v1/customer"))
	customer.Register(
		customerService,
		sessManager.Start,
	)
	customer.Handle(new(controller.CustomerController))

	//商品分类模块
	categoryService := service.NewCategoryService(engine)
	category := mvc.New(app.Party("/v1/category"))
	category.Register(
		categoryService,
		sessManager.Start,
	)
	category.Handle(new(controller.CategoryController))

	//商品模块
	productService := service.NewProductService(engine)
	product := mvc.New(app.Party("/v1/product"))
	product.Register(
		productService,
		sessManager.Start,
	)
	product.Handle(new(controller.ProductController))

}

/**
 * 项目设置
 */
func configation(app *iris.Application) {

	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context iris.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    " not found ",
			"data":   iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context iris.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    " interal error ",
			"data":   iris.Map{},
		})
	})
}
