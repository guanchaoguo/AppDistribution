package helper

/*
	项目公共函数包
*/

import (
	"golang-AppDistribution/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"log"
	"math/rand"
	"regexp"
	"time"
	"os/exec"
	"os"
	"path/filepath"
	"strconv"
)

var GlobalMongdbSession *mgo.Session

func Self_logger(myerr interface{}) {
	logfile := newLogFile()
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(myerr)

}

//获取redis操作对象
func GetRedis() redis.Conn {
	redisConf := config.GetRedisConf()
	c, err := redis.Dial("tcp", redisConf.REDIS_HOST+":"+redisConf.REDIS_PORT, redis.DialPassword(redisConf.REDIS_PASSWORD))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil
	}
	return c
}


//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXZY"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func InitMongodb(session *mgo.Session) {
	GlobalMongdbSession = session
	GlobalMongdbSession.SetPoolLimit(10)

}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}
	if end <= 0 {
		return string(rs[start:])
	}
	return string(rs[start:end])
}

//验证是否为IP格式
func CheckIp(ip string) bool {
	layout, err := regexp.MatchString("(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)", ip)
	if !layout || err != nil {
		return false
	}
	return true
}

func NewUploadFile(fpath string) string  {
	//获取当前执行文件的路径
	file, _ := exec.LookPath(os.Args[0])
	AppPath, _ := filepath.Abs(file)
	losPath, _ := filepath.Split(AppPath)
	path_d := losPath + fpath
	if !isDirExists(path_d) {
		fmt.Println("目录不存在")
		if err := os.MkdirAll(path_d, 0777); err != nil {
			fmt.Printf("%s", err)
		} else {
			fmt.Print("Create Directory OK!")
		}
	}
	return path_d

}

/**
 * 判断目录是否存在
 */
func isDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

	panic("not reached")
}

/**
 * 生成随机字符串
 */
func  RandomString(l int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


func RandInt(l int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := ""
	for i:=0; i<l; i++ {
		str += strconv.Itoa(r.Intn(10))
	}

	return str
}