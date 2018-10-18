package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/pkg/errors"
	"golang-AppDistribution/app/helper"
	"golang-AppDistribution/app/models"
	"golang-AppDistribution/lang"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"math"
	"fmt"
)

/*
	app合并管理
*/
type Merge struct {
}

/**
* @api {Get} /indexMerge  app列表获取
* @apiDescription app列表管理
* @apiGroup merge
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiParam {String} app_name    app名称
* @apiParam {String} start_date  开始时间
* @apiParam {String} end_date    结束时间
* @apiParam {Number} per_page 每页显示数据条数，默认15条
* @apiParam {Number} page 当前的所在页码，默认第1页
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 0,
*	  "message": "操作成功",
*	  "result": {
*		"current_page": 1,
*		"data": [
*		  {
*			"id": 1,// 全局唯一标识Id
*			"title": "豌豆荚",//应用名称
*			"logo": "{host}/uploads/icon/edsfd.png",// 下载地址
*			"apk_name": "豌豆荚", // 安卓名称
*			"apk_id": "2", // 安卓appId
*			"ipa_name": "豌豆荚", // ios名称
*			"ipa_id": "3", // iosid
*			"updated": "2017-12-27 10:13"
*		  },
*		],
*		"last_page": 1,
*		"per_page": 10,
*		"total": 7
*	  }
*	}
*
 */
func (Merge) Index(ctx context.Context) {

	app_name := ctx.FormValue("app_name")
	per_page := ctx.URLParamDefault("per_page", "10")
	page := ctx.URLParamDefault("page", "1")
	nowPage, _ := strconv.Atoi(page)
	start_date := ctx.FormValue("start_date")
	end_date := ctx.FormValue("end_date")

	number, _ := strconv.Atoi(per_page)
	startPosition := (int32(nowPage) - 1) * int32(number)

	list:= make([]models.AppInfo,0)
	whereCount := new(models.App_info)
	whereCount.Type = 3

	listObj := models.App_info{}.GetObj().NewSession().Where("type = ?",3)
	if app_name != "" {
		listObj = listObj.Where("name = ?",app_name)
		whereCount.Name = app_name
	}

	if start_date != "" &&  end_date != "" {
		listObj = listObj.Where("updated >=?",start_date).Where("updated <=?",end_date)
	}

	listErr := listObj.Desc("updated").Limit(number, int(startPosition)).Find(&list)
	count,_ := listObj.Count(whereCount)

	compute := float64(count) / float64(number)
	countPage := math.Ceil(compute)
	fmt.Println(count)

	if listErr != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Failure"),
			"result": "",
		})
		return
	}

	// 操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data":         list,
			"current_page": nowPage,
			"last_page":    countPage,
			"per_page":     number,
			"total":        count,
		},
	})
}

/**
	* @api {Post} /indexMerge 添加合并应用
	* @apiDescription 添加合并应用
	* @apiGroup merge
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiHeader {String} Authorization token.
	* @apiParam {String} title 应用名称
	* @apiParam {String} apk_name  安卓名称
	* @apiParam {String} apk_id 安卓appId
	* @apiParam {String} ipa_name ios名称
	* @apiParam {String} ipa_id  苹果名称应用的Id
	* @apiParam {String} ipa_name 苹果名称
	* @apiParam {String} shotUrl 短网址
	* @apiParam {String} logo  图片地址 通过上传icon接口获取
	* @apiParam {String} desc  合并后的应用描述
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (this Merge) AddMerge(ctx context.Context) {
	infoApp, errMsg := this.CheckUserForm(ctx)
	if errMsg != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": errMsg.Error(),
			"result":  "",
		})
		return
	}

	timeDate := time.Now().Format("2006-01-02 15:04:05")
	infoApp.Type = 3
	infoApp.Created = timeDate
	infoApp.Updated = timeDate
	infoApp.Last_ip = ctx.RemoteAddr()

	insert := models.App_info{}.AddOne(&infoApp)
	if !insert {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Failure"),
			"result":  "",
		})
		return
	}
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}

/**
* @api {Post} /uploadIco  上传图片
* @apiDescription 上传图片
* @apiGroup merge
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiParam {String} uploadfile 文件名
* @apiSuccessExample {json} Success-Response:
*  {
*   "code": 0,
*   "text": "操作成功",
*	  "result": {
*		"data": {
*		  "img_url": "{host}/uopload/icon/edf.png",
*		}
*  }
*
 */
func (Merge) UploadIco(ctx context.Context) {
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

	fileName := info.Filename
	fileExt := path.Ext(fileName)
	now := strconv.FormatInt(time.Now().Unix(), 10)
	if !strings.Contains(".png", fileExt) {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "UploadFileTypeError_png"),
			"result":  "",
		})
		return
	}

	pathIcon := helper.NewUploadFile("apidoc/uploads/iocn/")
	iconPath := pathIcon + now + ".png"

	out, err := os.OpenFile(iconPath, os.O_WRONLY|os.O_CREATE, 0666)
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

	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data": iris.Map{
				"img_url": "/uploads/iocn/" + now + ".png",
			},
		},
	})

}

/**
* @api {Get} /indexMerge/{id} 获取app合并信息
* @apiDescription 获取app合并信息
* @apiGroup merge
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 0,
*	  "message": "操作成功",
*	  "result": {
*		"data": {
*		    "id": 1,// 全局唯一标识Id
*           "type": "0", // 0 安卓 1 ipa
*			"versions": "6.0.0",// 版本号
*			"shotUrl": "biet",// 四位字符串短连接 下载地址 {host}/appDown/biet
*			"downCount": "1",// 已下载次数
*			"allowDown": "2",//允许下载次数
*			"logo": "{host}/uploads/icon/edsfd.png",// logo地址
*			"app_desc": "",// 应用基本介绍
*		}
*	  }
*	}
*
 */
func (Merge) Show(ctx context.Context) {
	id := ctx.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	appInfo := new(models.AppInfo)
	appInfo.Id = int32(newId)
	appInfo.Type = 3
	oneAppInfo, err := models.App_info{}.GetMgergeOne(appInfo)
	if err != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data":         oneAppInfo,
		},
	})
}

/**
* @api {Put} /indexMerge/{id} 修改app合并信息
* @apiDescription 修改app合并信息
* @apiGroup merge
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiParam {String} type  0 安卓 1 ipa
* @apiParam {String} name     标题
* @apiParam {String} versions  版本号
* @apiParam {String} shot_url  短网址
* @apiParam {String} allowDown 允许下载次数
* @apiParam {String} logo app 上传文件字段
* @apiParam {String} app_desc 应用基本介绍
* @apiSuccessExample {json} Success-Response:
*  {
*   "code": 0,
*   "text": "操作成功",
*   "result": "",
*  }
*
 */
func (this Merge) Update(ctx context.Context) {
	id := ctx.Params().Get("id")
	newId, _ := strconv.Atoi(id)
	appInfo := new(models.AppInfo)
	appInfo.Id = int32(newId)
	appInfo.Type = 3
	oneAppInfo, err := models.App_info{}.GetMgergeOne(appInfo)
	if err != nil || oneAppInfo == nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	infoApp, errMsg := this.CheckUserForm(ctx)
	if errMsg != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": errMsg.Error(),
			"result":  "",
		})
		return
	}

	timeDate := time.Now().Format("2006-01-02 15:04:05")
	infoApp.Updated = timeDate

	update := models.App_info{}.UpdateById(appInfo.Id,&infoApp)
	if !update {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}

/**
* @api {Delete} /indexMerge/{id}  删除app合并信息
* @apiDescription 删除app合并信息
* @apiGroup merge
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiSuccessExample {json} Success-Response:
*  {
*   "code": 0,
*   "text": "操作成功",
*   "result": "",
*  }
*
 */
func (Merge) Delete(ctx context.Context) {
	id := ctx.Params().Get("id")
	newId, _ := strconv.Atoi(id)
	appInfo := new(models.App_info)
	appInfo.Id = int32(newId)
	appInfo.Type = 3
	oneAppInfo, err := models.App_info{}.GetOne(appInfo)
	if err != nil  || oneAppInfo == nil{
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	app_info := new(models.App_info)
	delted := models.App_info{}.DeleteById(appInfo.Id,app_info)
	fmt.Println(app_info)
	if !delted {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}

/*
	表单验证
*/
func (Merge) CheckUserForm(ctx context.Context) (models.App_info, error) {
	title := ctx.FormValue("name")
	logo := ctx.FormValue("logo")
	apk_name := ctx.FormValue("apk_name")
	apk_id := ctx.FormValue("apk_id")
	ipa_name := ctx.FormValue("ipa_name")
	ipa_id := ctx.FormValue("ipa_id")
	desc := ctx.FormValue("desc")
	shotUrl := ctx.FormValue("shot_url")

	var appInfo models.App_info
	appInfo.Desc = desc
	if desc == "" {
		appInfo.Desc = ""
	}

	appInfo.Name = title
	if title == "" {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "name_required")
		return appInfo, errors.New(ErrorMsg)
	}

	appInfo.Logo = logo
	if logo == "" {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "logo_required")
		return appInfo, errors.New(ErrorMsg)
	}
	appInfo.Ipa_name = ipa_name
	if ipa_name == "" {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "ipa_name_required")
		return appInfo, errors.New(ErrorMsg)
	}

	appInfo.Apk_name = apk_name
	if apk_name == "" {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "apk_name_required")
		return appInfo, errors.New(ErrorMsg)
	}

	ipaId, _ := strconv.Atoi(ipa_id)
	appInfo.Ipa_id = int32(ipaId)
	if ipaId <=0  {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "ipaId_required")
		return appInfo, errors.New(ErrorMsg)
	}

	apkId, _ := strconv.Atoi(apk_id)
	appInfo.Apk_id = int32(apkId)
	if apkId <=0 {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "apkId_required")
		return appInfo, errors.New(ErrorMsg)
	}

	appInfo.Shot_url = shotUrl
	if len(shotUrl) != 5 {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "shot_url_invalid")
		return  appInfo,errors.New(ErrorMsg)
	}

	return appInfo, nil
}

func (Merge) upLoadCount(infoApp *models.App_info)  bool {
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