package models

import "fmt"

type App_info struct {
	Id          int32  `json:"id" xorm:"pk"`
	Logo        string `json:"logo"`
	Name        string `json:"name"`
	Type        int8   `json:"type"`
	Password    string `json:"-"`
	App_id      string `json:"app_id"`
	Status      int8   `json:"status"`
	Is_password int8   `json:"is_password"`
	Versions    string `json:"versions"`
	Updated     string `json:"-"`
	Created     string `json:"-"`
	Last_ip     string `json:"-"`
}

func (App_info) TableName() string {
	return "app_info"
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
func (App_info) DeleteById(id int32, menus *App_info) bool {
	aff, err := engine.Id(id).Delete(menus)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	根据条件获取单条数据
*/
func (App_info) GetOne(menus *App_info) *App_info {
	has, err := engine.Get(menus)
	if !has || err != nil {
		return nil
	}
	return menus
}

/*
	添加单条数据
*/
func (App_info) Add(menus *App_info) int32 {
	aff, err := engine.Insert(menus)
	if aff == 0 || err != nil {
		return 0
	}
	return int32(aff)
}

/*
	根据ID进行修改操作
*/
func (App_info) UpdateById(id int32, menus *App_info) bool {
	aff, err := engine.Id(id).Update(menus)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	获取列表
*/
func (App_info) List(list []App_info, condition *App_info) []App_info {
	err := engine.OrderBy("sort_id").Find(&list, condition)
	if err != nil {
		return nil
	}
	return list
}
