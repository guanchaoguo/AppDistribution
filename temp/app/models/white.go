package models

import (
	"github.com/go-xorm/xorm"
)

/*
	白名单操作model
*/
type White struct {
	Id             int32  `json:"id" xorm:"autoincr pk id"`
	Domain         string `json:"domain"`
	HallName       string `json:"hall_name"`
	Status         int32  `json:"status"`
	Channel        int32  `json:"channel"`
	Ips            string `json:"ips"`
	LockRemark     string `json:"lock_remark"`
	CreateDate     string `json:"create_date"`
	LastUpdateDate string `json:"last_update_date"`
}

func (White) TableName() string {
	return "whitelist"
}

func (White) GetObj() *xorm.Engine {
	return engine
}

func (White) GetEngine() *xorm.Engine {
	return engine
}

/*
	获取列表
*/
func (White) List(list []White, white *White, number int32, startPosition int32) ([]White, error) {
	err := engine.Limit(int(number), int(startPosition)).Find(&list, white)
	if err != nil {
		return nil, err
	}
	return list, nil
}

/*
	获取指定条件的总数
*/
func (White) GetCount(white *White) int32 {
	count, err := engine.Count(white)
	if err != nil {
		return 0
	}
	return int32(count)
}

/*
	添加一条数据
*/
func (White) Add(insertData *White) int64 {
	aff, err := engine.Insert(insertData)
	if aff == 0 || err != nil {
		return 0
	}
	return aff
}

/*
	批量添加数据（最大为150条数据）
*/
func (White) AddBatch(insertData *[]White) int64 {
	aff, err := engine.Insert(insertData)
	if aff == 0 || err != nil {
		return 0
	}
	return aff
}

/*
	获取单条信息
*/
func (White) GetOne(white *White, notInSlice []int) *White {
	var (
		has bool
		err error
	)
	if len(notInSlice) > 0 {
		has, err = engine.NotIn("id", notInSlice).Get(white)
	} else {
		has, err = engine.Get(white)
	}

	if !has || err != nil {
		return nil
	}
	return white
}

/*
	修改单条数据
*/
func (White) UpdateById(id int32, white *White) bool {
	aff, err := engine.Id(id).Update(white)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	删除单条操作
*/
func (White) Delete(white *White) bool {
	aff, err := engine.Delete(white)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	更新单条数据状态
*/
func (White) SaveStatusById(id int32, white *White) bool {
	_, err := engine.Id(id).Cols("status", "lock_remark").Update(white)
	if err != nil {
		return false
	}
	return true
}
