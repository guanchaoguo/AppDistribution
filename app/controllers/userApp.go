package controllers

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris"
	"golang-AppDistribution/lang"
	"golang-AppDistribution/app/helper"
	"golang-AppDistribution/app/models"
	"strconv"
	"math"
	"errors"
	"time"
	"fmt"
	"strings"
)


/*
	app管理
*/
type UserApp struct {
}

/**
* @api {Get} /userApp  app列表获取
* @apiDescription app列表管理
* @apiGroup Application management
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
*			"down": 1 // 下载
*			"scan": 1 // 浏览
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

	app_name := ctx.FormValue("app_name")
	per_page := ctx.URLParamDefault("per_page", "10")
	page := ctx.URLParamDefault("page", "1")
	str_type :=  ctx.URLParamDefault("type", "0")
	app_type ,_:= strconv.Atoi(str_type)
	nowPage, _ := strconv.Atoi(page)
	start_date := ctx.FormValue("start_date")
	end_date := ctx.FormValue("end_date")

	number, _ := strconv.Atoi(per_page)
	var startPosition  = (int32(nowPage) - 1) * int32(number)
	var  whereCount = new(models.App_info)

	var list = make([]models.App_info, 0)
	var list_obj = models.App_info{}.GetObj().NewSession()

	if app_type == 0 {
		list_obj = list_obj.Where("type != ? ",3)
	}

	if app_type == 1 {
		list_obj = list_obj.Where("type = ? ",0)
	}

	if app_type == 2 {
		list_obj = list_obj.Where("type = ? ",1)
	}

	if app_name != "" {
		list_obj = list_obj.Where("name = ? ",app_name)

	}

	if start_date != "" &&  end_date != "" {
		fmt.Println(start_date,end_date)
		list_obj = list_obj.Where(" updated >= ? and updated <=? ",start_date,end_date)
	}

	// 获取列表显示
	//models.App_info{}.GetObj().ShowSQL(true)
	coountObj := list_obj.Clone()
	count,_ := coountObj.Count(whereCount)
	compute := float64(count) / float64(number)
	countPage := math.Ceil(compute)

	listErr := list_obj.Desc("updated").Limit(number, int(startPosition)).Find(&list)
	if listErr != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Failure"),
			"result": "",
		})
		return
	}

	// 加入统计字段
	type list_App struct {
		models.App_info
		Down_count   int32 `json:"down" `   //  下载次数
		Upload_count int32 `json:"upload" ` //  上传次数
		Scan_count   int32 `json:"scan" `   // 浏览次数
	}

	//获取查询的app的Id
	var ids_str = "( "
	var ids []string = []string{}
	var list_app = make([]list_App,len(list))
	for k, value := range list{
		list_app[k] = list_App{value,0,0,0}
		var  appId  = strconv.Itoa(int(value.Id))
		ids = append(ids,appId)
	}

	ids_str += strings.Join(ids,",")
	ids_str += " )"

	// 查询当前app 统计浏览 下载
	var sumList = make([]struct{
		App_id   int32 `json:"app_id" `   //  app_id
		Down_count   int32 `json:"down" `   //  下载次数
		Upload_count int32 `json:"upload" ` //  上传次数
		Scan_count   int32 `json:"scan" `   // 浏览次数
	},0)
	var sumSql = "SELECT app_id, sum(scan_count) as scan_count ,sum(down_count) as down_count FROM app_statistics WHERE  app_id in"+ids_str+" GROUP BY app_id"

	models.Statistics{}.GetObj().Sql(sumSql).Find(&sumList)
	models.Statistics{}.GetObj().ShowSQL(true)
	for _, static_v := range sumList{
		for k, list_v := range list_app{
			if list_v.Id == static_v.App_id {
				list_app[k].Down_count = static_v.Down_count
				list_app[k].Scan_count = static_v.Scan_count
				list_app[k].Upload_count = static_v.Upload_count
			}
		}
	}

	// 操作成功
	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data":         list_app,
			"current_page": nowPage,
			"last_page":    countPage,
			"per_page":     number,
			"total":        count,
		},
	})
}

/**
* @api {Get} /userApp/{id} 获取账号app信息
* @apiDescription 获APP信息
* @apiGroup Application management
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
	id := ctx.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	appInfo := new(models.App_info)
	appInfo.Id = int32(newId)
	app_info, err := models.App_info{}.GetOne(appInfo)
	if err != nil ||app_info ==nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	// 查询出来app的统计信息
	total,err:= models.App_info{}.GetObj().Where("app_id =?",appInfo.Id).SumInt(new(models.Statistics), "down_count")

	// 加入统计字段
	type Info struct {
		models.App_info
		Down_count   int32 `json:"down" `   //  下载次数
	}

	info:= &Info {*app_info,int32(total)}

	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data": info,
		},
	})
}

/**
* @api {Put} /userApp/{id} 修改保存账号
* @apiDescription 修改保存账号
* @apiGroup Application management
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} name 应用名称
* @apiParam {String} type  0 安卓 1 ipa
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
func (this UserApp) Update(ctx context.Context) {
	id := ctx.Params().Get("id")
	newId, _ := strconv.Atoi(id)
	appInfo := new(models.App_info)
	appInfo.Id = int32(newId)
	oneAppInfo, err := models.App_info{}.GetOne(appInfo)
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
* @api {Put} /userApp/downStats/{id} app下载密码设置
* @apiDescription  app下载密码设置
* @apiGroup Application management
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
	id := ctx.Params().Get("id")
	newId, _ := strconv.Atoi(id)
	appInfo := new(models.App_info)
	appInfo.Id = int32(newId)
	oneAppInfo, err := models.App_info{}.GetOne(appInfo)
	if err != nil || oneAppInfo == nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	var infoApp models.App_info
	timeDate := time.Now().Format("2006-01-02 15:04:05")
	infoApp.Updated = timeDate
	infoApp.Password =  helper.RandInt(4)
	infoApp.Is_password = int8(1)

	update := models.App_info{}.UpdateStatusById(appInfo.Id,&infoApp)
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
		"result": iris.Map{
			"data":iris.Map{
				"password":infoApp.Password,
			},
		},
	})
}

/**
* @api {Delete} /userApp/downStats/{id} 取消下载密码
* @apiDescription 取消下载密码
* @apiGroup Application management
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
	id := ctx.Params().Get("id")
	newId, _ := strconv.Atoi(id)
	appInfo := new(models.App_info)
	appInfo.Id = int32(newId)
	oneAppInfo, err := models.App_info{}.GetOne(appInfo)
	fmt.Println(oneAppInfo)
	if err != nil || oneAppInfo.Is_password == 0 {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	var infoApp models.App_info
	infoApp.Updated = time.Now().Format("2006-01-02 15:04:05")
	infoApp.Is_password = int8(0)

	update := models.App_info{}.UpdateStatusById(appInfo.Id,&infoApp)
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
		"result": "",
	})
}

/**
* @api {Get} /appDown/{code} 通过短连接下载app
* @apiDescription   下载app
* @apiGroup Application management
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiParam {String} locale 语言
*  {
*   "code": 0,
*   "text": "操作成功",
*   "result": {
*	 data:{
*		"shot_url" :"werrr"
*      }
*    },
*  }
*
 */
func (UserApp) AppDown(ctx context.Context) {
	code := ctx.Params().Get("code")
	if len(code) != 5 {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}


	appInfo := new(models.App_info)
	appInfo.Shot_url = code
	oneAppInfo, err := models.App_info{}.GetOne(appInfo)
	if err != nil || oneAppInfo == nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Failure"),
			"result":  "",
		})
		return
	}


	// 查询出来app的统计信息
	total,err:= models.Statistics{}.GetObj().Where("app_id =?",oneAppInfo.Id).SumInt(new(models.Statistics), "down_count")
	fmt.Println(345)
	list := make([]models.App_info,0)
	// 加入统计字段
	type Info struct {
		models.App_info
		Down_count   int32 `json:"down" `   //  下载次数
	}

	// 如果是合并的应用则同时返回安卓和苹果
	if oneAppInfo.Type == 3{
        models.App_info{}.GetObj().In("id",oneAppInfo.Apk_id,oneAppInfo.Ipa_id).Find(&list)
	}

	lists:= make([]Info,0)

	//  两条ios android
	if len(list) >1 {

		total_apk,_:= models.Statistics{}.GetObj().Where("app_id =?",oneAppInfo.Apk_id).SumInt(new(models.Statistics), "down_count")
		total_ipa,_:= models.Statistics{}.GetObj().Where("app_id =?",oneAppInfo.Ipa_id).SumInt(new(models.Statistics), "down_count")

		lists = append(lists, Info{list[0],int32(total_apk)})
		lists = append(lists,Info{list[1],int32(total_ipa)})
	}

	lists = append(lists,Info {*oneAppInfo,int32(total)})

	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data":lists,
		},
	})


}

/**
* @api {Delete} /userApp/{id}  删除app
* @apiDescription 删除app
* @apiGroup Application management
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
	id := ctx.Params().Get("id")
	newId, _ := strconv.Atoi(id)
	appInfo := new(models.App_info)
	appInfo.Id = int32(newId)
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

/**
* @api {Get} /shotUrl  短网址校验唯一性
* @apiDescription 删除app
* @apiGroup Application management
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiHeader {String} Authorization token.
* @apiSuccessExample {json} Success-Response:
*  {
*   "code": 0,
*   "text": "操作成功",
*   "result": {
*	 data:{
*		"shot_url" :"werrr"
*      }
*    },
*  }
*
*/
func (UserApp) GetShotUrl(ctx context.Context) {
	Top:
    appInfo := make([]*models.App_info,0)
	Shot_url := NewLen(5)
	err := models.App_info{}.GetObj().Where("shot_url = ?",Shot_url).Find(&appInfo)
	if err != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	fmt.Println(appInfo)
	if len(appInfo) >0 {
		time.Sleep(300 * time.Millisecond)
		goto Top
	}

	ctx.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(ctx, "comm", "Success"),
		"result": iris.Map{
			"data":iris.Map{
				"shot_url":Shot_url,
				},
		},
	})

}

/*
	表单验证
*/
func (UserApp) CheckUserForm(ctx context.Context) (models.App_info,error) {
	name := ctx.FormValue("name")
	type_str := ctx.FormValue("type")
	versions := ctx.FormValue("versions")
	shotUrl := ctx.FormValue("shot_url")
	logo := ctx.FormValue("logo")
	app_desc := ctx.FormValue("desc")
	allow_count := ctx.FormValue("allowCount")

	var app_info models.App_info
	app_info.Desc = app_desc
	if app_desc == "" {
		app_info.Desc = ""
	}

	app_info.Name = name
	if name == ""{
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "name_required")
		return  app_info,errors.New(ErrorMsg)
	}

	type_int, _ := strconv.Atoi(type_str)
	app_info.Type = int8(type_int)
	if type_int < 0 ||  type_int >=2 {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "type_invalid")
		return  app_info,errors.New(ErrorMsg)
	}

	allow_count_int, _ := strconv.Atoi(allow_count)
	app_info.Allow_count = int32(allow_count_int)
	if type_int < 0 ||  type_int >=2 {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "type_invalid")
		return  app_info,errors.New(ErrorMsg)
	}

	app_info.Versions = versions
	if versions == ""{
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "versions_required")
		return  app_info,errors.New(ErrorMsg)
	}

	app_info.Shot_url = shotUrl
	if len(shotUrl) != 5  {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "shot_url_invalid")
		return  app_info,errors.New(ErrorMsg)
	}

	app_info.Logo = logo
	if logo == "" {
		ErrorMsg := lang.Lang{}.GetLang(ctx, "appInfo", "logo_required")
		return  app_info,errors.New(ErrorMsg)
	}
	return app_info, nil
}







