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

	app.Get("/{code}", func(ctx iris.Context) {
		ctx.WriteString(ctx.FormValue("code"))
	})

	// 跨域测试
	app.Post("/send", func(ctx iris.Context) {
		ctx.WriteString("sent")
	})
	// 跨域测试
	app.Put("/send", func(ctx iris.Context) {
		ctx.WriteString("sent")
	})


	//测试获取数据
	app.Get("/user", controllers.Test{}.GetUser)

	//测试redis
	app.Get("/redis", controllers.Test{}.GetRedis)

	//新token
	app.Get("/token", controllers.Test{}.MyToken)

	//测试jwt
	app.Get("/ping", middleware.CheckJwt, controllers.Test{}.MyToken)

	// 测试上传页面
	app.Get("/upload", controllers.Release{}.Upload)
	app.Post("/ups", controllers.Test{}.Upload)
	app.Get("/up", controllers.Test{}.Up)

	app.Get("/pic", controllers.Release{}.UploadIoc)


	/*
		项目正式路由
	*/

	//获取验证码操作
	app.Get("/captcha", controllers.Captcaha{}.GetCaptcha)

	//用户登录操作
	app.Post("/login", controllers.UserLogin{}.Login)

	//用户登出操作
	app.Post("/loginOut", middleware.CheckJwt, controllers.UserLogin{}.LoginOut)

	//app 上传
	//app.Post("/release", middleware.CheckJwt,controllers.Release{}.Release)
	app.Post("/release", controllers.Release{}.Release)

	/****--------------------app管理 start---------------****/
	//app 列表获取
	app.Get("/userApp",  controllers.UserApp{}.Index)

	//app 信息获取
	app.Get("/userApp/{id}",controllers.UserApp{}.Show)

	//app 信息修改
	app.Put("/userApp/{id}", middleware.CheckJwt,controllers.UserApp{}.Update)

	//app 信息下载密码设置
	app.Put("/userApp/downStats/{id}",middleware.CheckJwt, controllers.UserApp{}.UpdateDownStats)

	//app 取消下载密码设置
	app.Delete("/userApp/downStats/{id}",middleware.CheckJwt, controllers.UserApp{}.DelteDownStats)

	//app 短网址下载
	app.Get("/appDown/{code}",controllers.UserApp{}.AppDown)

	//app 信息删除
	app.Delete("/userApp/{id}", middleware.CheckJwt,controllers.UserApp{}.Delete)

	//app 获取短链接
	app.Get("/shotUrl",middleware.CheckJwt, controllers.UserApp{}.GetShotUrl)
	/****--------------------app管理 end---------------****/

	/****--------------------app合并 start---------------****/
	//app合并列表获取
	app.Get("/indexMerge",middleware.CheckJwt, controllers.Merge{}.Index)

	//上传图片
	app.Post("/uploadIco", controllers.Merge{}.UploadIco)

	//app合并列表获取
	app.Post("/indexMerge",middleware.CheckJwt, controllers.Merge{}.AddMerge)

	//app 合并信息获取
	app.Get("/indexMerge/{id}",middleware.CheckJwt, controllers.Merge{}.Show)

	//app合并信息修改
	app.Put("/indexMerge/{id}",middleware.CheckJwt, controllers.Merge{}.Update)

	//app合并信息删除
	app.Delete("/indexMerge/{id}",middleware.CheckJwt, controllers.Merge{}.Delete)
	/****--------------------app合并 end---------------****/

	/****--------------------app下载管理 start---------------****/
	//app 统计数据图表
	//app.Get("/statistics",middleware.CheckJwt, controllers.Statistics{}.Index)
	app.Get("/statistics", middleware.CheckJwt,controllers.Statistics{}.Index)

	//app 浏览统计
	app.Get("/statistics/scan/{id}", controllers.Statistics{}.Scan)

	//app 下载统计
	app.Get("/statistics/down/{id}", controllers.Statistics{}.Down)

	/****--------------------app下载管理 end---------------****/

}
