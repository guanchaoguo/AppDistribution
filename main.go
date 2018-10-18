package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"golang-AppDistribution/app/helper"
	"golang-AppDistribution/routes"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"fmt"
)

func newApp() *iris.Application {

	app := iris.New()

	//获取当前执行文件的路径
	file, _ := exec.LookPath(os.Args[0])
	AppPath, _ := filepath.Abs(file)
	losPath, _ := filepath.Split(AppPath)

	dirPath := losPath +"apidoc"

	app.Get("/apidoc", func(ctx iris.Context) {
		ctx.ServeFile(dirPath+"/index.html", false)
	})

	fileServer := app.StaticHandler(dirPath, false, false)

	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		path := r.URL.Path

		if !strings.Contains(path, ".") {
			router(w, r)
			return
		}
		ctx := app.ContextPool.Acquire(w, r)
		fileServer(ctx)
		app.ContextPool.Release(ctx)
	})

	return app
}

func main() {
	app := newApp()
	app.WrapRouter(cors.WrapNext(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
	}))

	// 访问日志处理
	r, close := helper.NewRequestLogger()
	defer close()
	app.Use(r)
	app.OnAnyErrorCode(r, func(ctx iris.Context) {
		ctx.Writef("(Unexpected) internal server error")
	})

	// 记录程序启动日志
	app.Use(func(this iris.Context) {
		this.Application().Logger().Infof("Begin request for path %s", this.Path())
		fmt.Println(this.Path())
		this.Next()
	})

	//开启路由监听
	routes.WebRoutes{}.StartRoute(app)

	//监听端口，并且输出启动日志，设置输出编码
	app.Run(iris.Addr(":8081"), iris.WithoutStartupLog, iris.WithCharset("UTF-8"), iris.WithoutVersionChecker)

}
