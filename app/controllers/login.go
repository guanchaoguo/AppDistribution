package controllers

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"golang-AppDistribution/app/helper"
	"golang-AppDistribution/app/middleware"
	"golang-AppDistribution/app/models"
	"golang-AppDistribution/lang"
	"time"
)

/*
	用户登录操作
*/
type UserLogin struct {
}

/**
* @api {Post} /login 后台登录
* @apiDescription 后台登录
* @apiGroup login
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiHeader {String} Authorization token.
* @apiParam {String} locale 语言
* @apiParam {String} user_name 用户名
* @apiParam {String} password 密码
* @apiParam {String} code 验证码
* @apiParam {String} captchaKey 验证码key
* @apiSuccessExample {json} Success-Response:
	{
	  "code": 400,
	  "message": "登录成功",
	  "result": {
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50Ijoi5byA5Y-R55So5oi3IiwiZGF0ZSI6MTUxMDY0NjYyMSwiaWQiOjEsInVzZXJfbmFtZSI6ImFkbWluIn0.Nd1AuCJyD0CgDPkjp3lljxhyDCBatBMrcCO-lCnE6GE"
	  }
	}
*
*/
func (UserLogin) Login(w context.Context) {
	user_name := w.FormValue("user_name")   //用户名
	password := w.FormValue("password")     //密码
	captcha := w.FormValue("code")          //验证码
	captchaKey := w.FormValue("captchaKey") //验证码key

	//用户名不能为空
	if user_name == "" {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "login", "UsernameRequired"),
			"result":  "",
		})
		return
	}
	//密码不能为空
	if password == "" {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "login", "PasswordRequired"),
			"result":  "",
		})
		return
	}
	//验证码不能为空
	if captcha == "" {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "login", "CaptchaRequired"),
			"result":  "",
		})
		return
	}

	//验证码key不能为空
	if captchaKey == "" {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}
	//验证码错误
	cRes := Captcaha{}.CheckCaptcha(captchaKey, captcha)
	if !cRes {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "login", "CaptchaError"),
			"result":  "",
		})
		return
	}

	//进行用户数据校验
	userInfo, _ := models.User{}.GetUserByUserName(user_name)

	//用户不存在
	if userInfo == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "login", "UserNotExist"),
			"result":  "",
		})
		return
	}

	//密码比对，判断密码是否正确(先sha1加密，后MD5再加密)
	newPassword := password + userInfo.Salt
	h := sha1.New()
	h.Write([]byte(newPassword))
	enPassword := hex.EncodeToString(h.Sum(nil))
	m := md5.New()
	m.Write([]byte(enPassword))
	md5Password := hex.EncodeToString(m.Sum(nil))

	if md5Password != userInfo.Password {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "login", "PasswordError"),
			"result":  "",
		})
		return
	}

	//密码验证通过,生成token操作
	userData := map[string]interface{}{
		"id":        userInfo.Id,
		"account":   userInfo.Account,
		"user_name": userInfo.UserName,
	}
	token, err := middleware.NewToken(userData)
	if err != nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "login", "LoginFailure"),
			"result":  "",
		})
		return
	}

	//密码验证通过进行修改数据库最后登录时间
	updateUser := new(models.User)
	updateUser.LastLoginDate = time.Now().Format("2006-01-02 15:04:05")
	updateUser.LastLoginIp = w.RemoteAddr()
	res := models.User{}.UpdateById(userInfo.Id, updateUser)
	if !res {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "login", "LoginFailure"),
			"result":  "",
		})
		return
	}

	//登录成功获取菜单信息

	//返回登录成功提示
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "login", "LoginSuccess"),
		"result": iris.Map{
			"token":     token,
			"user_name": userInfo.Account,
		},
	})
	return

}

/**
	* @api {Post} /loginOut 退出登录
	* @apiDescription 退出登录
	* @apiGroup login
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (UserLogin) LoginOut(w context.Context) {
	tokenString := w.FormValue("token")
	h := md5.New()
	h.Write([]byte(tokenString))
	tokenKey := hex.EncodeToString(h.Sum(nil)) //MD5
	//删除redis数据
	r := helper.GetRedis()
	_, err := r.Do("DEL", tokenKey)
	if err != nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	//操作成功
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})

}
