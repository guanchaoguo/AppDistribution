package routes

import (
	"golang-AppDistribution/app/controllers"
	"golang-AppDistribution/app/middleware"
	"github.com/kataras/iris"
)

type WebRoutes struct {
}

func (WebRoutes) StartRoute(app *iris.Application) {
	app.Get("/", controllers.Test{}.Hello)


	//测试获取数据
	app.Get("/user", controllers.Test{}.GetUser)

	//测试redis
	app.Get("/redis", controllers.Test{}.GetRedis)

	//测试验证码
	app.Get("/test", controllers.Test{}.DemoCode)

	//新token
	app.Get("/token", controllers.Test{}.MyToken)

	//测试jwt
	app.Get("/ping", middleware.CheckJwt, controllers.Test{}.MyToken)

	// 测试上传页面
	app.Get("/upload", controllers.Release{}.Upload)
	app.Get("/index", controllers.Captcaha{}.Index)

	/*
		项目正式路由
	*/

	//获取验证码操作
	app.Get("/getCaptcha", controllers.Captcaha{}.GetCaptcha)

	//用户登录操作
	app.Post("/login", controllers.UserLogin{}.Login)

	//用户登出操作
	app.Post("/loginOut", middleware.CheckJwt, controllers.UserLogin{}.LoginOut)

	//app 上传
	app.Post("/release", controllers.Release{}.Release)

	/****--------------------app管理 start---------------****/
	//app 列表获取
	app.Get("/userApp", controllers.UserApp{}.Index)

	//app 信息获取
	app.Get("/userApp/{id}", controllers.UserApp{}.Show)

	//app 信息修改
	app.Put("/userApp/{id}", controllers.UserApp{}.Update)

	//app 信息下载密码设置
	app.Put("/userApp/downStats/{id}", controllers.UserApp{}.UpdateDownStats)

	//app 信息下载密码取消
	app.Get("/appDown/{code}", controllers.UserApp{}.AppDown)

	//app 信息删除
	app.Delete("/userApp", controllers.UserApp{}.Delete)
	/****--------------------app管理 end---------------****/

	/****--------------------app合并 start---------------****/
	//app合并列表获取
	app.Get("/indexMerge", controllers.Merge{}.Index)

	//app合并列表获取
	app.Post("/uploadIco", controllers.Merge{}.UploadIco)

	//app合并列表获取
	app.Post("/indexMerge", controllers.Merge{}.AddMerge)

	//app 合并信息获取
	app.Get("/indexMerge/{id}", controllers.Merge{}.Show)

	//app合并信息修改
	app.Put("/indexMerge/{id}", controllers.Merge{}.Update)

	//app合并信息删除
	app.Delete("/indexMerge", controllers.Merge{}.Delete)
	/****--------------------app合并 end---------------****/

	/****--------------------app下载管理 start---------------****/
	//app 统计数据图表
	app.Get("/statistics", controllers.Statistics{}.Index)

	//app 下载统计
	app.Get("/statistics/scan", controllers.Statistics{}.Scan)

	//app 下载统计
	app.Get("/statistics/down", controllers.Statistics{}.Down)

	/****--------------------app下载管理 end---------------****/

}
