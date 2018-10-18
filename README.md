## 项目框架说明

    http框架使用 iris  框架地址：https://iris-go.com/
    Mysql ORM 使用XORM,文档地址：https://www.kancloud.cn/kancloud/xorm-manual-zh-cn

    接口文档使用 apidoc  地址：https://www.kancloud.cn/kancloud/xorm-manual-zh-cn
    执行更新接口文档命令： apidoc -i ./app/controllers -o ./apidoc
## 生成在线文档apidoc命令
    apidoc -i app/controllers -o apidoc

##跨平台编译命令
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build

PS：go语言string、int、int64互相转换

//string到int
int,err:=strconv.Atoi(string)
//string到int64
int64, err := strconv.ParseInt(string, 10, 64)
//int到string
string:=strconv.Itoa(int)
//int64到string
string:=strconv.FormatInt(int64,10)
//string到float32(float64)
float,err := strconv.ParseFloat(string,32/64)
//float到string
string := strconv.FormatFloat(float32, 'E', -1, 32)
string := strconv.FormatFloat(float64, 'E', -1, 64)
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)

## 守护进程命令
 supervisorctl stop appDistribution