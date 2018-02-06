package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris"
	"github.com/tidwall/gjson"
	"golang-AppDistribution/app/helper"
	"golang-AppDistribution/lang"
	"time"
)

const SecretKey = "lebogame"

//验证jwt有效性中间件
func CheckJwt(ctx iris.Context) {
	tokenString := ctx.GetHeader("Authorization")

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SecretKey), nil
	})

	if token != nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//解析token后和缓存中的数据进行对比，判断是否合法
			h := md5.New()
			h.Write([]byte(tokenString))
			tokenKey := hex.EncodeToString(h.Sum(nil)) //MD5
			r := helper.GetRedis()
			redisUserInfo, er := redis.String(r.Do("GET", tokenKey))
			if er != nil {
				ctx.JSON(iris.Map{
					"mesage": lang.Lang{}.GetLang(ctx, "comm", "Unlink"),
					"code":   400,
				})
				ctx.Request().Body.Close()
				return
			}

			//进行用户比对
			if claims["user_name"] != gjson.Get(redisUserInfo, "user_name").String() {
				ctx.JSON(iris.Map{"mesage": lang.Lang{}.GetLang(ctx, "comm", "Unlink"), "code": 400})
				//删除redis数据
				r.Do("DEL", tokenKey)
				ctx.Request().Body.Close() //终端本次请求
				return
			}
			ctx.Next() //验证成功执行下一步操作
		} else {
			fmt.Println("error")
		}
	} else {
		ctx.JSON(iris.Map{"mesage": lang.Lang{}.GetLang(ctx, "comm", "Unlink"), "code": 400})
	}

}

//生成一个新的token
func NewToken(userData map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        userData["id"],
		"account":   userData["account"],
		"user_name": userData["user_name"],
		"date":      time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SecretKey))

	//把生成的token  MD5后作为redis中的键值，values为用户的信息
	redis := helper.GetRedis()
	h := md5.New()
	h.Write([]byte(tokenString))
	tokenKey := hex.EncodeToString(h.Sum(nil)) //MD5
	data, _ := json.Marshal(userData)
	_, err = redis.Do("SET", tokenKey, data) //token为key,把用户信息存到redis中
	defer redis.Close()                      //关闭redis

	if err != nil {
		return "", err
	}
	return tokenString, err
}
