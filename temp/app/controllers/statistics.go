package controllers

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris"
	"golang-AppDistribution/lang"
)


/*
	app统计管理
*/
type Statistics struct {
}


/**
* @api {Get} /statistics  app统计数据获取
* @apiDescription app统计数据获取
* @apiGroup statistics
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 0,
*	  "message": "操作成功",
*	  "result": {
*		"data": [
*		  {
*			"scan_count": "1",// 浏览次数
*			"download_count": "	2",//下载次数
*			"upload_count": "6",// 上传次数
*			"date": "2017-12-27 10:13"// 时间 只获取最近十天统计
*		  },
*		],
*	  }
*	}
*
 */
func (Statistics) Index(ctx context.Context) {

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
* @api {Get} /statistics/scan  app浏览统计
* @apiDescription app浏览统计
* @apiGroup statistics
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 0,
*	  "message": "操作成功",
*	  "result": ""
*	}
*
 */
func (Statistics) Scan(ctx context.Context) {

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
* @api {Get} /statistics/down  app下载统计
* @apiDescription app下载统计
* @apiGroup statistics
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 0,
*	  "message": "操作成功",
*	  "result": ""
*	}
*
 */
func (Statistics) Down(ctx context.Context) {

	// 操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data":         "",
		},
	})
}