package models

import (
	"github.com/go-xorm/xorm"
	"time"
)

/*
	代理服务器模型
*/
type Proxy struct {
	Id             int32  `json:"id" xorm:"autoincr pk id"`
	Addr           string `json:"addr"`
	Port           int32  `json:"port"`
	UserName       string `json:"user_name"`
	Password       string `json:"password"`
	LastUpdateDate string `json:"last_update_date"`
}

func (Proxy) TableName() string {
	return "proxy_info"
}

func (Proxy) GetObj() *xorm.Engine {
	return engine
}

/*
	获取列表数据
*/
func (Proxy) List(list []Proxy, proxy *Proxy, number int, startPosition int) ([]Proxy, error) {
	err := engine.Limit(number, startPosition).Find(&list, proxy)
	if err != nil {
		return nil, err
	}
	return list, nil
}

/*
	获取条件总数
*/
func (Proxy) GetCount(proxy *Proxy) int32 {
	count, err := engine.Count(proxy)
	if count <= 0 || err != nil {
		return 0
	}
	return int32(count)
}

/*
	根据ID删除数据
*/
func (Proxy) Delete(proxy *Proxy) bool {
	aff, err := engine.Delete(proxy)
	if aff == 0 || err != nil {
		return false
	}
	return true
}

/*
	添加单条数据
*/
func (Proxy) Add(proxy *Proxy) int64 {
	id, err := engine.Insert(proxy)
	if id == 0 || err != nil {
		return id
	}
	return id
}

/*
	根据条件获取单条数据
*/
func (Proxy) GetOne(proxy *Proxy) *Proxy {
	has, err := engine.Get(proxy)
	if !has || err != nil {
		return nil
	}
	return proxy
}

/*
	根据ID修改保存数据
*/
func (Proxy) UpdateById(id int32, proxy *Proxy) bool {
	proxy.LastUpdateDate = time.Now().Format("2006-01-02 15:04:05")
	up, err := engine.Id(id).Update(proxy)
	if up == 0 || err != nil {
		return false
	}
	return true
}
