package models

import (
	"github.com/go-xorm/xorm"
	"fmt"
)

type Statistics struct {
	Id          int32  `json:"id" xorm:"autoincr pk id"`
	Upolader_id int32  `json:"upolader_id" `
	App_id      int32   `json:"app_id" `
	App_name    string `json:"app_name" `
	App_type    int8   `json:"app_type" `
	Down_count  int32   `json:"down_count" `//  下载次数
	Upload_count  int32   `json:"upload_count" `//  上传次数
	Scan_count  int32   `json:"scan_count" `// 浏览次数
	Created     string `json:"created" `
}

func (Statistics) TableName() string {
	return "app_statistics"
}

//app列表
func (Statistics) List() ([]Statistics, error) {
	list := make([]Statistics, 0)
	err := engine.Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

/*
	写入操作
*/
func (Statistics) AddOne(InsetData *Statistics) bool {
	aff, err := engine.Insert(InsetData)
	if aff <= 0 || err != nil {
		return false
	}
	return true
}

/*
	根据ID删除操作
*/
func (Statistics) DeleteById(id int32, data *Statistics) bool {
	aff, err := engine.Id(id).Delete(data)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	根据条件获取单条数据
*/
func (Statistics) GetOne(info *Statistics) (*Statistics, error) {
	has, err := engine.Get(info)
	if !has || err != nil {
		return nil, err
	}
	return info, nil
}

/*
	根据条件获取合并单条数据
*/
func (Statistics) GetMgergeOne(info *AppInfo) (*AppInfo, error) {
	has, err := engine.Get(info)
	if !has || err != nil {
		return nil, err
	}
	return info, nil
}

/*
	添加单条数据
*/
func (Statistics) Add(Data *Statistics) int32 {
	aff, err := engine.Insert(Data)
	if aff == 0 || err != nil {
		return 0
	}
	return int32(aff)
}

/*
	根据ID进行修改操作
*/
func (Statistics) UpdateById(id int32, Data *Statistics) bool {
	aff, err := engine.Id(id).Update(Data)
	fmt.Println(Data, id,err)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	根据ID进行状态操作
*/
func (Statistics) UpdateStatusById(id int32, Data *Statistics) bool {
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
func (Statistics) GetObj() *xorm.Engine {
	return engine
}
