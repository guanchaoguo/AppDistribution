package config

/*
	数据库配置文件
*/

//mysql
type Mysql struct {
	DB_HOST         string //数据库服务器
	DB_DATABASE     string //数据库名称
	DB_USERNAME     string //数据库登录名
	DB_PASSWORD     string //数据库密码
	DB_PORT         string //数据库端口
	CHARSET         string //字符集
	SetMaxIdleConns int    //默认打开数据库的连接数
	SetMaxOpenConns int    //最大打开数据库的连接数

}

//mongodb
type Mongodb struct {
	URL string
}

//RabbitMQ
type RabbitMQ struct {
	URL string
} 

//redis
type Redis struct {
	REDIS_HOST     string
	REDIS_PASSWORD string
	REDIS_PORT     string
}

//mysql配置
func GetMysqlConf() Mysql {
	return Mysql{
		DB_HOST:         "192.168.31.231",
		DB_DATABASE:     "browser",
		DB_USERNAME:     "root",
		DB_PASSWORD:     "123456",
		DB_PORT:         "3306",
		CHARSET:         "utf8",
		SetMaxIdleConns: 10,
		SetMaxOpenConns: 10,
	}
}

//mongodb配置
func GetMongodbConf() Mongodb {
	return Mongodb{
		URL: "mongodb://hhq163:bx123456@192.168.31.231:27017/live_game?connect=direct&maxPoolSize=10",
	}
}

//redis配置
func GetRedisConf() Redis {
	return Redis{
		REDIS_HOST:     "192.168.31.230",
		REDIS_PASSWORD: "bx123456",
		REDIS_PORT:     "6379",
	}
}

//RabbitMQ配置信息
func GetRabbitMQ() RabbitMQ  {
	return RabbitMQ{
		URL: "amqp://lebo2017:ljf12345@192.168.31.230:5672/test",
	}
}
