package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"golang-AppDistribution/lang"
	"golang-AppDistribution/app/models"
	"io"
	"os"
	"golang-AppDistribution/app/helper"
	"path"
	"strings"
	"github.com/phinexdaz/ipapk"
	"image"
	"image/png"
	"time"
)

/*
	app 发布
*/
type Release struct {
}


/**
* @api {Post} /release  app发布
* @apiDescription app发布
* @apiGroup release
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiHeader {String} Authorization token.
* @apiParam {String} locale 语言
* @apiParam {String} uploadfile 上传名称
* @apiSuccessExample {json} Success-Response:
	{
	  "code": 0,
	  "message": "操作成功",
	  "result":""
	}
*
*/
func (Release) Release(ctx context.Context) {
	file, info, err := ctx.FormFile("uploadfile")
	if err != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	defer file.Close()

	// 校验app 后缀是否正确
	fileName := info.Filename
	fileExt:=  path.Ext(fileName)
	filenameOnly := strings.TrimSuffix(fileName, fileExt)
	if(!strings.Contains(".ipa ,.apk",  fileExt)){
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "UploadFileTypeError_ipak"),
			"result":  "",
		})
		return
	}

	// 检查文件目录
	pathApp:= helper.NewUploadFile("uploads/app/")
	appPath:= pathApp+fileName

	pathIcon:= helper.NewUploadFile("uploads/iocn/")
	iconPath:= pathIcon+filenameOnly+".png"

	out, err := os.OpenFile(appPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "NetworkError"),
			"result":  "",
		})
		return
	}
	defer out.Close()

	// 复制文件到指定目录
	io.Copy(out, file)

	// 提取app基本信息
	timeDate:= time.Now().Format("2006-01-02 15:04:05")
	infoApp := GetAppInfo(appPath,iconPath)
	infoApp.Last_ip = ctx.RemoteAddr()
	infoApp.Logo = fileName
	infoApp.Name = filenameOnly+".png"
	infoApp.Created = timeDate
	infoApp.Updated = timeDate


	// 写入数据库
	insert := models.App_info{}.AddOne(infoApp)
	if !insert {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	// 操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}

func  GetAppInfo(appPath string,iconPath string) *models.App_info {
	apk, _ := ipapk.NewAppParser(appPath)

	appInfo:= new(models.App_info)
	appInfo.Versions = apk.Version
	appInfo.App_id = apk.BundleId

	SaveIcon(apk.Icon,iconPath)

	return  appInfo
}

func SaveIcon(icon image.Image ,iconPath string) error  {
	file, err := os.Create(iconPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, icon); err != nil {
		return err
	}
	return nil
}

func (Release) Upload(w context.Context) {
	str := `<html>
	<head>
		<title>Upload file</title>
	</head>
	<body>
		<form enctype="multipart/form-data"action="http://127.0.0.1:8080/release" method="POST">
			<input type="file" name="uploadfile" />
          <input type="submit" value="upload" />
		</form>
	</body>
	</html>`
	w.Header("Content-Type", "text/html")
	w.Write([]byte(str))
}
