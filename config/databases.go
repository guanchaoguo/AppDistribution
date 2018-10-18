package config

/*
	数据库配置文件
*/

const env  = "ONLINE" //LOCAL 开发环境 TEST 线上测试环境  ONLINE 正式环境

//获取当前执行文件的路径

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

//redis
type Redis struct {
	REDIS_HOST     string
	REDIS_PASSWORD string
	REDIS_PORT     string
}

//mysql配置
func GetMysqlConf() Mysql {
	switch env {
	case "LOCAL":
		return Mysql{
			DB_HOST:         "192.168.31.232",
			DB_DATABASE:     "AppDistribution",
			DB_USERNAME:     "root",
			DB_PASSWORD:     "123456",
			DB_PORT:         "3306",
			CHARSET:         "utf8",
			SetMaxIdleConns: 10,
			SetMaxOpenConns: 10,
		}
	case "TEST":
		return Mysql{
			DB_HOST:         "10.200.124.23",
			DB_DATABASE:     "AppDistribution",
			DB_USERNAME:     "Mysqladmin",
			DB_PASSWORD:     "Mysql1707!",
			DB_PORT:         "3306",
			CHARSET:         "utf8",
			SetMaxIdleConns: 10,
			SetMaxOpenConns: 10,
		}
	case "ONLINE":
		return Mysql{
			DB_HOST:         "10.200.66.96",
			DB_DATABASE:     "AppDistribution",
			DB_USERNAME:     "appfenfaadmin",
			DB_PASSWORD:     "AppFfa345^&*",
			DB_PORT:         "3306",
			CHARSET:         "utf8",
			SetMaxIdleConns: 10,
			SetMaxOpenConns: 10,
		}
	default:
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

}

//redis配置
func GetRedisConf() Redis {
	switch env {
	case "LOCAL":
		return Redis{
			REDIS_HOST:     "192.168.31.230",
			REDIS_PASSWORD: "bx123456",
			REDIS_PORT:     "6379",
		}
	case "TEST":
		return Redis{
			REDIS_HOST:     "10.200.124.21",
			REDIS_PASSWORD: "bx123456",
			REDIS_PORT:     "6379",
		}
	case "ONLINE":
		return Redis{
			REDIS_HOST:     "10.200.66.96",
			REDIS_PASSWORD: "AppLebo8102",
			REDIS_PORT:     "6379",
		}
	default:
		return Redis{
			REDIS_HOST:     "192.168.31.230",
			REDIS_PASSWORD: "bx123456",
			REDIS_PORT:     "6379",
		}


	}

}


// domain配置
func GetDomain() string {
	switch env {
	case "LOCAL":
		return "http://appdistribution-api.dev"
	case "TEST":
		return "https://test.appdapi.lggame.co"
	case "ONLINE":
		return "https://appdapi.lggame.co"
	default:
		return "http://appdistribution-api.dev"
	}
}


// 获取plist文件模板
func GetTempPlist() string {
	return`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>items</key>
		<array>
			<dict>
				<key>assets</key>
					<array>
						<dict>
							<key>kind</key>
							<string>software-package</string>
							<key>url</key>
							<string>%s</string>
						</dict>
					</array>
				<key>metadata</key>
				<dict>
					<key>bundle-identifier</key>
					<string>%s</string>
					<key>bundle-version</key>
					<string>%s</string>
					<key>kind</key>
					<string>software</string>
					<key>title</key>
					<string>%s</string>
				</dict>
			</dict>
		</array>
	</dict>
</plist>`
}


