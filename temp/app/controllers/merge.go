package controllers


import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris"
	"golang-AppDistribution/lang"
)


/*
	app合并管理
*/
type Merge struct {
}

/**
* @api {Get} /merge  app列表获取
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
	* @api {Post} /merge 添加合并应用
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
	* @apiParam {String} ipa_id ios名称
	* @apiParam {String} ipa_name ios_id
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
func (Merge) AddMerge(w context.Context) {

}

/**
* @api {Post} /ploadIco  上传图片
* @apiDescription 上传图片
* @apiGroup merge
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiParam {String} file_name 文件名
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
func (Merge) UploadIco(w context.Context) {

}

/**
* @api {Get} /merge/{id} 获取app合并信息
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
* @api {Put} /merge/{id} 修改app合并信息
* @apiDescription 修改app合并信息
* @apiGroup merge
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
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
func (Merge) Update(ctx context.Context) {
	//操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}

/**
* @api {Delete} /merge/{id}  删除app合并信息
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
	//操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result":  "",
	})
}

