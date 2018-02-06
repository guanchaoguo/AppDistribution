package controllers

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris"
	"golang-AppDistribution/lang"
	"os/exec"
	"os"
	"path/filepath"
	"golang-AppDistribution/app/helper"
)


/*
	app管理
*/
type UserApp struct {
}


/**
* @api {Get} /userApp  app列表获取
* @apiDescription app列表管理
* @apiGroup userApp
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiParam {String}  type      0 表示 获取全部app类型  1 ：只获取安卓列表 2：只获取ios列表
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
*			"name": "豌豆荚",//应用名称
*			"logo": "{host}/uploads/icon/edsfd.png",// 下载地址
*			"type": "0", // 0 安卓 1 ipa
*			"is_password": "是否启动六位下载密码",
*			"password": "",// 下载密码
*			"app_id": "NetDragon.Mobile.iPhone.91Space	",// 应用ID号码
*			"versions": "6.0.0",// 版本号
*			"shotUrl": "kdaz", //短连接
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
func (UserApp) Index(ctx context.Context) {

	// 操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data":         "",
			"current_page": "",   //当前页码
			"last_page":    "",   //总共多少页
			"per_page":     "",    //每页多少条数据
			"total":        "",     //总的记录条数
		},
	})
}

/**
* @api {Get} /userApp/{id} 获取账号app信息
* @apiDescription 获APP信息
* @apiGroup userApp
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
func (UserApp) Show(ctx context.Context) {
	// 操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data":         "",
		},
	})
}

/**
* @api {Put} /userApp/{id} 修改保存账号
* @apiDescription 修改保存账号
* @apiGroup userApp
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} type  0 安卓 1 ipa
* @apiParam {String} versions  版本号
* @apiParam {String} shotUrl  短网址
* @apiParam {String} downCount 已下载次数
* @apiParam {String} allowDown 允许下载次数
* @apiParam {String} fileUpload app 上传文件字段
* @apiParam {String} app_desc 应用基本介绍
* @apiSuccessExample {json} Success-Response:
*  {
*   "code": 0,
*   "text": "操作成功",
*   "result": "",
*  }
*
*/
func (UserApp) Update(ctx context.Context) {
	//操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}

/**
* @api {Put} /userApp/downStats/{id} app下载密码设置
* @apiDescription  app下载密码设置
* @apiGroup userApp
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiSuccessExample {json} Success-Response:
*  {
*   "code": 0,
*   "text": "操作成功",
*   "result": {
*		"data": {
*			"password": "1234",// 四位下载密码
*		}
*  }
*
*/
func (UserApp) UpdateDownStats(ctx context.Context) {
	//操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}

/**
* @api {Delete} /userApp/downStats/{id} 取消下载密码
* @apiDescription 取消下载密码
* @apiGroup userApp
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
func (UserApp) DelteDownStats(ctx context.Context) {
	//操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}

/**
* @api {Get} /appDown/{code} 通过短连接下载app
* @apiDescription   下载app
* @apiGroup userApp
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiParam {String} locale 语言
*
 */
func (UserApp) AppDown(ctx context.Context) {
	file_name := ctx.FormValue("code")
	newFileName := helper.Substr(file_name, 20, 0)
	//获取当前执行文件的路径
	file, _ := exec.LookPath(os.Args[0])
	AppPath, _ := filepath.Abs(file)
	losPath, _ := filepath.Split(AppPath)
	filePath := losPath + "/uploads/" + file_name
	ctx.SendFile(filePath, newFileName)
}

/**
* @api {Delete} /userApp/{id}  删除app
* @apiDescription 删除app
* @apiGroup userApp
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
func (UserApp) Delete(ctx context.Context) {
	//操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}









