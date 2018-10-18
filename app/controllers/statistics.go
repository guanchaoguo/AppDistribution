package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"golang-AppDistribution/app/models"
	"golang-AppDistribution/lang"
	"strconv"
	"time"
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
*    "code": 0,
*    "message": "操作成功",
*    "result": {
*        "data": {
*            "date": ["2018-01-09","2018-01-08","2018-01-07","2018-01-06","2018-01-05","2018-01-04", "2018-01-03","2018-01-02","2018-01-01","2017-12-31"],
*            "upload": [0,0,0,0,1, 0,0, 0, 0,0],
*            "down": [ 0,1, 0,0, 0, 0, 0,0, 0, 0],
*            "scan": [ 0, 2,0,0,1, 0,0,0,0,0 ]
*        }
*    }
*}
 */
func (Statistics) Index(ctx context.Context) {

	// 获取uid
	//uid:=  int32(middleware.BaseInfo["id"].(float64));

	// 查询统计近十天的数据
	list := make([]models.Statistics, 0)
	endDate := time.Now().AddDate(0, 0, -10).Format("2006-01-02")
	Sql := "select DATE_FORMAT(created, '%Y-%m-%d' ) as created , sum(upload_count) upload_count , sum(down_count) down_count  , sum(scan_count) scan_count from app_statistics  where created >= "+endDate+"  group by created  order by created desc"
	err := models.Statistics{}.GetObj().Sql(Sql).Find(&list)
	if err != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	staticData  := struct {
		Date []string `json:"date"`
		Upload  []int32  `json:"upload"`
		Down    []int32  `json:"down"`
		Scan    []int32  `json:"scan"`
	}{
		make([]string, 10),make([]int32, 10), make([]int32, 10),make([]int32, 10),
	}

	// 拼装统计图数据格式
	for j := 0; j <= 9; j++ {
		staticData.Date[j] = time.Now().AddDate(0, 0, -j).Format("2006-01-02")
		for _, v := range list {
			// 上传  count_type 1 浏览 2 上传 3 下载
			if v.Created == staticData.Date[j] {
				staticData.Upload[j] = v.Upload_count
				staticData.Scan[j] = v.Scan_count
				staticData.Down[j] = v.Down_count
			}

		}
	}

	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data": staticData,
		},
	})
}

/**
* @api {Get} /statistics/scan/{id}  app浏览统计
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
	id := ctx.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	// app 是否合法
	appInfo := new(models.App_info)
	appInfo.Id = int32(newId)
	oneAppInfo, err := models.App_info{}.GetOne(appInfo)
	if err != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	sanInfo := new(models.Statistics)
	sanInfo.Created = time.Now().Format("2006-01-02")
	sanInfo.App_type = oneAppInfo.Type
	sanInfo.App_name = oneAppInfo.Name
	sanInfo.App_id = oneAppInfo.Id
	sanInfo.Scan_count = 1
	sanInfo.Upolader_id = oneAppInfo.User_id

	// 当天第一条数据是否存在
	var res bool
	staticInfo, err := models.Statistics{}.GetOne(sanInfo)
	if staticInfo == nil {
		res = models.Statistics{}.AddOne(sanInfo)
	} else {
		updateData := new(models.Statistics)
		updateData.Scan_count = sanInfo.Scan_count + 1
		res = models.Statistics{}.UpdateById(staticInfo.Id, updateData)
	}

	if !res {
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
		"result": iris.Map{
			"data": "",
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
	id := ctx.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	// app 是否合法
	appInfo := new(models.App_info)
	appInfo.Id = int32(newId)
	oneAppInfo, err := models.App_info{}.GetOne(appInfo)
	if err != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	sanInfo := new(models.Statistics)
	sanInfo.Created = time.Now().Format("2006-01-02")
	sanInfo.App_type = oneAppInfo.Type
	sanInfo.App_name = oneAppInfo.Name
	sanInfo.App_id = oneAppInfo.Id
	sanInfo.Down_count = 1
	sanInfo.Upolader_id = oneAppInfo.User_id

	// 当天第一条数据是否存在
	var res bool
	staticInfo, err := models.Statistics{}.GetOne(sanInfo)
	if staticInfo == nil {
		res = models.Statistics{}.AddOne(sanInfo)
	} else {
		updateData := new(models.Statistics)
		updateData.Down_count = sanInfo.Down_count + 1
		res = models.Statistics{}.UpdateById(staticInfo.Id, updateData)
	}

	if !res {
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
		"result": iris.Map{
			"data": "",
		},
	})

}
