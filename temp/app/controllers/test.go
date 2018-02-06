package controllers

import (
	"golang-AppDistribution/app/helper"
	"golang-AppDistribution/app/middleware"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"gopkg.in/mgo.v2/bson"
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

func (Test) GetMongodb(ctx context.Context) {
	mon := helper.GetMongodb()
	defer mon.Close() //用完记得关闭链接
	//err := mon.DB("live_game").C("api_log").Insert(&ApiLog{"UserName","测试写入数据"})
	log := ApiLog{}
	err1 := mon.DB("live_game").C("api_log").Find(bson.M{"user_name": "h88888"}).One(&log)

	if err1 != nil {
		fmt.Println("mongodb get failed:", err1)
	} else {
		ctx.JSON(iris.Map{"message": &log})
		return
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

func (Test) DemoCode(ctx iris.Context) {

}

