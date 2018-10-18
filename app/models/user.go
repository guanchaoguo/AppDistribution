package models

import "github.com/go-xorm/xorm"

type User struct {
	Id            int32  `json:"id" xorm:"autoincr pk id"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	Salt          string `json:"salt"`
	Account       string `json:"account"`
	LastDate      string `json:"last_date"`
	LastLoginDate string `json:"last_login_date"`
	LastLoginIp   string `json:"last_login_ip"`
}

func (*User) TableName() string {
	return "admin_user"
}

func (User) GetObj() *xorm.Engine {
	return engine
}

//根据用户登录名获取用户信息
func (User) GetUserByUserName(user_name string) (*User, error) {
	var user User
	has, err := engine.Where("user_name=?", user_name).Get(&user)

	if err != nil || !has {
		return nil, err
	}
	return &user, nil
}

//根据传入条件获取单条数据
func (User) GetAccountByOne(user *User, notInSlice []int) (*User, error) {

	var (
		has bool
		err error
	)
	if len(notInSlice) > 0 {
		has, err = engine.NotIn("id", notInSlice).Get(user)
	} else {
		has, err = engine.Get(user)
	}

	if err != nil || !has {
		return nil, err
	}
	return user, nil
}

//通过ID更新数据
func (User) UpdateById(id int32, user *User) bool {
	affected, err := engine.ID(id).Update(user)

	if affected == 0 || err != nil {
		return false
	}
	return true
}

//获取用户列表
func (User) List(condition []User, user *User, number int32, startPosition int32) ([]User, error) {
	err := engine.Limit(int(number), int(startPosition)).Find(&condition, user)
	if err != nil {
		return nil, err
	}
	return condition, nil
}

//添加新用户
func (User) Add(insertData *User) int32 {
	id, err := engine.Insert(insertData)

	if err != nil {
		return 0
	}
	return int32(id)
}

//根据ID删除用户
func (User) DeleteById(id int32) int32 {
	user := new(User)
	aff, err := engine.Id(id).Delete(user)
	if err != nil {
		return 0
	}
	return int32(aff)
}

//获取指定条件的总数据条数
func (User) GetCount(user *User) int64 {
	count, err := engine.Count(user)
	if err != nil {
		return 0
	}
	return count
}
