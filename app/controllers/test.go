package controllers

import (
	"golang-AppDistribution/app/helper"
	"golang-AppDistribution/app/middleware"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"os"
	"io"
)

type Test struct {
}

type ApiLog struct {
	UserName string `json:"user_name" bson:"user_name"`
	ApiName  string `json:"apiName" bson:"apiName"`
}

func (Test) Hello(ctx context.Context) {
	ctx.JSON(iris.Map{"message": "Hello World!"})
}

func (Test) GetUser(ctx context.Context) {
	//user := models.User{}.GetUser()
	//ctx.JSON(iris.Map{"message":user})
}

func (Test) GetRedis(ctx context.Context) {
	c := helper.GetRedis()
	data, err := redis.String(c.Do("GET", "yzq222222_get_order_list"))
	defer c.Close() //用完记得关闭链接
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		ctx.JSON(iris.Map{"message": data})
	}
}


func (Test) MyHandler(ctx iris.Context) {
	//middleware.CheckJwt()
	//ctx.JSON(iris.Map{"token":"aa"})
	//token,_ := middleware.NewToken()
	//ctx.JSON(iris.Map{"token":token})
}

func (Test) MyToken(ctx iris.Context) {
	userData := map[string]interface{}{"user_name": "liangxinzhou"}
	token, _ := middleware.NewToken(userData)
	ctx.JSON(iris.Map{"token": token})
}

func (Test) Upload(ctx iris.Context) {
	//file, info, _:= ctx.FormFile("uploadfile")
	//fmt.Println(file, info)

	//ctx.ParseMultipartForm(8 << 20)
	//title := ctx.ParseFormValue["title"]
	//fhs := ctx.MultipartForm.File["radio[]"]

	ctx.Request().ParseForm()
	ctx.Request().ParseMultipartForm(100 << 20) //最大内存为100M


	mp := ctx.Request().MultipartForm
	if mp == nil {
		fmt.Println("not MultipartForm.")
		ctx.Write(([]byte)("不是MultipartForm格式"))
		return
	}

	fileHeaders, findFile := mp.File["uploadfile"]
	if !findFile || len(fileHeaders) == 0 {
		fmt.Println("file count == 0.")
		ctx.Write(([]byte)("没有上传文件"))
		return
	}

	for _, v := range fileHeaders {
		fileName := v.Filename
		file, err := v.Open()
		fmt.Println(err, "Open file error."+fileName)
		defer file.Close()

		outputFilePath := helper.NewUploadFile("apidoc/uploads/app/") + fileName
		writer, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0666)
		fmt.Println(err, "Open local file error")
		io.Copy(writer, file)
	}

	msg := fmt.Sprintf("成功上传了%d个文件", len(fileHeaders))
	ctx.Write(([]byte)(msg))

	/*fileDir, _ := exec.LookPath(os.Args[0])
	AppPath, _ := filepath.Abs(fileDir)
	losPath, _ := filepath.Split(AppPath)
	path_d := losPath +  "apidoc/uploads/app/"+info.Filename
	//path_d = "apidoc/uploads/app/"+info.Filename
	out, err:= os.OpenFile(path_d, os.O_WRONLY|os.O_CREATE, 0666)
	fmt.Println(err)
	defer out.Close()
	io.Copy(out, file)*/
	fmt.Println("ok")
}

func (Test) Up(ctx iris.Context) {
	str := `<html>
	<head>
		<title>Upload file</title>
	</head>
	<body>
		<form enctype="multipart/form-data"action="/ups" method="POST">
			<input type="file" name="uploadfile" multiple="multiple" />
          <input type="submit" value="upload" data-name =1 />
		</form>
	</body>
	</html>`
	ctx.Header("Content-Type", "text/html")
	ctx.Write([]byte(str))

}

