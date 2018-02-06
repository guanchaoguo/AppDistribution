package models

/*
	文件导入模型管理
*/
type Upload struct {
	Id            int32  `json:"id" xorm:"pk"`
	FileName      string `json:"file_name"`
	SucceedNumber int64  `json:"succeed_number"`
	FailureNumber int64  `json:"failure_number"`
	CreateDate    string `json:"create_date"`
}

func (Upload) TableName() string {
	return "file_upload"
}

/*
	写入操作
*/
func (Upload) AddOne(upload *Upload) bool {
	aff, err := engine.Insert(upload)
	if aff <= 0 || err != nil {
		return false
	}
	return true
}

/*
	获取列表
*/
func (Upload) List(upload []Upload, condition *Upload, number int32, startPosition int32) ([]Upload, error) {
	err := engine.Limit(int(number), int(startPosition)).Desc("id").Find(&upload, condition)
	if err != nil {
		return nil, err
	}
	return upload, nil
}

/*
	获取指定条件的总条数
*/
func (Upload) GetCount(upload *Upload) int64 {
	count, err := engine.Count(upload)
	if err != nil || count <= 0 {
		return 0
	}
	return count
}
