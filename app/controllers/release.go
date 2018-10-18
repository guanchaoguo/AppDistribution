package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/phinexdaz/ipapk"
	"golang-AppDistribution/app/helper"
	"golang-AppDistribution/app/models"
	"golang-AppDistribution/config"
	"golang-AppDistribution/lang"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
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
func (this Release) Release(ctx context.Context) {
	// 获取上传文件
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

	// 获取文件后缀名
	fileName := info.Filename
	fileExt := path.Ext(fileName)
	now := strconv.FormatInt(time.Now().Unix(), 10)
	if !strings.Contains(".ipa ,.apk", fileExt) {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "UploadFileTypeError_ipak"),
			"result":  "",
		})
		return
	}

	//设置app 类型
	var appType int8 = 0
	if ".ipa" == fileExt {
		appType = 1
	}

	// 设置app 保存路劲
	pathApp := helper.NewUploadFile("apidoc/uploads/app/")
	appName := now + fileExt
	app_save_Path := pathApp + appName

	// 设置iocn 保存路劲
	pathIcon := helper.NewUploadFile("apidoc/uploads/iocn/")
	iocnName := now + ".png"
	icon_save_Path := pathIcon + iocnName

	// 设置ios plist 保存路劲
	pathplist := helper.NewUploadFile("apidoc/uploads/app/")
	plistName := now + ".plist"
	list_save_pathp := pathplist + plistName

	//保存app文件
	out, err := os.OpenFile(app_save_Path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "NetworkError"),
			"result":  "",
		})
		return
	}
	defer out.Close()
	io.Copy(out, file)

	// 设置ionc下载路劲
	timeDate := time.Now().Format("2006-01-02 15:04:05")
	iocUrl := "/uploads/iocn/" + iocnName

	// 解析app基本信息属性
	infoApp := GetAppInfo(app_save_Path, icon_save_Path, iocUrl)

	//填空数据库保存信息
	infoApp.Last_ip = ctx.RemoteAddr()
	infoApp.Plist = "/uploads/app/" + plistName
	infoApp.App_url = "/uploads/app/" + appName
	infoApp.Shot_url = NewLen(5)
	infoApp.Type = appType
	infoApp.Created = timeDate
	infoApp.Updated = timeDate

	// 写入ios下载需要的plist文件
	if this.SavePlist(infoApp, list_save_pathp) != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Failure"),
			"result":  "",
		})
		return
	}

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

	// 统计该APP上传次数
	this.upLoadCount(infoApp)

	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}

func GetAppInfo(app_save_Path string, iconPath string, iconUrl string) *models.App_info {
	apk, _ := ipapk.NewAppParser(app_save_Path)

	appInfo := new(models.App_info)
	appInfo.Versions = apk.Version
	appInfo.App_id = apk.BundleId
	appInfo.Size = apk.Size
	appInfo.Name = apk.Name
	if apk.Icon != nil {
		SaveIcon(apk.Icon, iconPath)
		appInfo.Logo = iconUrl
	}

	return appInfo
}

func SaveIcon(icon image.Image, iconPath string) error {
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

func (Release) SavePlist(appInfo *models.App_info, list_save_pathp string) error {
	plist_down_path := config.GetDomain() + appInfo.App_url
	var xml_str = fmt.Sprintf(config.GetTempPlist(), plist_down_path, appInfo.App_id, appInfo.Versions, appInfo.Name)

	buf, err := ioutil.ReadAll(strings.NewReader(xml_str))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(list_save_pathp, buf, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (Release) upLoadCount(infoApp *models.App_info) bool {
	sanInfo := new(models.Statistics)
	sanInfo.Created = time.Now().Format("2006-01-02")
	sanInfo.App_type = infoApp.Type
	sanInfo.App_name = infoApp.Name
	sanInfo.App_id = infoApp.Id
	sanInfo.Upload_count = 1
	sanInfo.Upolader_id = infoApp.User_id

	// 当天第一条数据是否存在
	var res bool
	staticInfo, _ := models.Statistics{}.GetOne(sanInfo)
	if staticInfo == nil {
		res = models.Statistics{}.AddOne(sanInfo)
	} else {
		updateData := new(models.Statistics)
		updateData.Upload_count = sanInfo.Upload_count + 1
		res = models.Statistics{}.UpdateById(staticInfo.Id, updateData)
	}

	return res
}

func (Release) Upload(w context.Context) {
	str := `<html>
	<head>
		<title>Upload file</title>
	</head>
	<body>
		<form enctype="multipart/form-data"action="/release" method="POST">
			<input type="file" name="uploadfile" />
          <input type="submit" value="upload" data-name =1 />
		</form>
	</body>
	</html>`
	w.Header("Content-Type", "text/html")
	w.Write([]byte(str))
}

func (Release) UploadIoc(w context.Context) {
	str := `<html>
	<head>
		<title>Upload file</title>
	</head>
	<body>
		<form enctype="multipart/form-data"action="http://127.0.0.1:8081/uploadIco" method="POST">
			<input type="file" name="uploadfile" />
          <input type="submit" value="upload" data-name =1 />
		</form>
	</body>
	</html>`
	w.Header("Content-Type", "text/html")
	w.Write([]byte(str))
}
