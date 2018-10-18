package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

type App_info struct {
	Id          int32  `json:"id" xorm:"autoincr pk id"`
	User_id     int32  `json:"-"`
	Logo        string `json:"logo"`
	Name        string `json:"name"`
	Type        int8   `json:"type"`
	Password    string `json:"password"`
	App_id      string `json:"app_id"`
	Status      int8   `json:"status"`
	Is_password int8   `json:"is_password"`
	Is_merger   int8   `json:"-"`
	Apk_id      int32  `json:"-"`
	Ipa_id      int32  `json:"-"`
	Apk_name    string `json:"-"`
	Ipa_name    string `json:"-"`
	Versions    string `json:"versions"`
	Size        int64  `json:"size"`
	App_url     string `json:"app_url"`
	Shot_url    string `json:"shot_url"`
	Desc        string `json:"desc"`
	Allow_count int32  `json:"allowDown"`
	Updated     string `json:"updated"`
	Created     string `json:"-"`
	Last_ip     string `json:"-"`
	Plist     string `json:"plist"`
}

// 合并应用
type AppInfo struct {
	Id       int32  `json:"id" xorm:"autoincr pk id"`
	Logo     string `json:"logo"`
	Type     int8   `json:"-"`
	Name     string `json:"name"`
	Apk_id   int32  `json:"apk_id"`
	Ipa_id   int32  `json:"ipa_id"`
	Apk_name string `json:"apk_name"`
	Ipa_name string `json:"ipa_name"`
	Shot_url string `json:"shot_url"`
	Desc     string `json:"desc"`
	Created  string `json:"-"`
	Updated  string `json:"updated"`
}

func (App_info) TableName() string {
	return "app_info"
}

// 应用列表显示
type StaticAndApp struct {
	App_info     `xorm:"extends"`
	Down_count   int32 `json:"down" `   //  下载次数
	Upload_count int32 `json:"upload" ` //  上传次数
	Scan_count   int32 `json:"scan" `   // 浏览次数
}

// 获取单个的app 信息
type OneStaticAndApp struct {
	App_info     `xorm:"extends"`
	Down_count   int32 `json:"downCount" `  //  下载次数
}

//app列表
func (App_info) List() ([]App_info, error) {
	list := make([]App_info, 0)
	err := engine.Find(&list)
	fmt.Println(list, err, 121)
	if err != nil {
		return nil, err
	}
	return list, nil
}

/*
	写入操作
*/
func (App_info) AddOne(InsetData *App_info) bool {
	aff, err := engine.Insert(InsetData)
	if aff <= 0 || err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/*
	根据ID删除操作
*/
func (App_info) DeleteById(id int32, data *App_info) bool {
	aff, err := engine.Id(id).Delete(data)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	根据条件获取单条数据
*/
func (App_info) GetOne(info *App_info) (*App_info, error) {
	has, err := engine.Get(info)
	if !has || err != nil {
		return nil, err
	}
	return info, nil
}

/*
	根据条件获取合并单条数据
*/
func (App_info) GetMgergeOne(info *AppInfo) (*AppInfo, error) {
	has, err := engine.Get(info)
	if !has || err != nil {
		return nil, err
	}
	return info, nil
}

/*
	添加单条数据
*/
func (App_info) Add(Data *App_info) int32 {
	aff, err := engine.Insert(Data)
	if aff == 0 || err != nil {
		return 0
	}
	return int32(aff)
}

/*
	根据ID进行修改操作
*/
func (App_info) UpdateById(id int32, Data *App_info) bool {
	aff, err := engine.Id(id).Update(Data)
	fmt.Println(Data, id)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	根据ID进行状态操作
*/
func (App_info) UpdateStatusById(id int32, Data *App_info) bool {
	aff, err := engine.Id(id).Cols("updated", "password", "is_password").Update(Data)
	fmt.Println(Data, id)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	获取会话句柄
*/
func (App_info) GetObj() *xorm.Engine {
	return engine
}
