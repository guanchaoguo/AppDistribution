package lang

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/kataras/iris/context"
)

type Lang struct {
}

func (Lang) GetLang(w context.Context,module string,langCode string) string  {

	locale := w.FormValue("locale")

	langPak :=  map[string]interface{} {
		//中文语言包
		"zh-cn" : map[string]interface{} {
			//公共模块语言包
			"comm"	: map[string]string {
				"Invalid" : "非法请求参数",
				"DataIsEmpty" : "数据为空",
				"Success"	: "操作成功",
				"Failure"	: "操作失败",
				"NetworkError":"网络错误,请重试",
				"Unlink"	: "链接已断开，请重新登录",
				"UploadFileTypeError_ipak" : "文件类型错误，只能上传 .apk/ipa 后缀的文件",
				"UploadFileTypeError_png" : "文件类型错误，只能上传 png 后缀的文件",
			},

			//登录模块语言包
			"login"	: map[string]string {
				"ReLogin" : "链接已断开，请重新登录",
				"UsernameRequired" : "用户名不能为空",
				"CaptchaRequired" : "验证码不能为空",
				"PasswordRequired" : "密码不能为空",
				"CaptchaError"  : "验证码错误",
				"UserNotExist" : "用户不存在",
				"PasswordError" : "密码错误",
				"LoginFailure" : "登录失败",
				"LoginSuccess" : "登录成功",
			},

			//系统用户管理模块语言包
			"account" : map[string]string {
				"UsernameRequired" : "用户名不能为空",
				"CaptchaRequired" : "密码不能为空",
				"AccountRequired" : "真实姓名不能为空",
				"UserNameExist" : "用户名已存在",
				"UserNameLength" : "用户名为6-50个字符长度的英文字母+数字组合",
				"PasswordLength" : "密码长度最少为6位",
			},
			//app语言包
			"appInfo" : map[string]string {
				"name_required" : "应用名称不能为空",
				"type_invalid" : "app类型错误",
				"versions_required" : "版本号不能为空",
				"shot_url_invalid" : "短网址必须为有效字符串",
				"ipa_name_required" : "安卓应用名不能为空",
				"apk_name_required" : "苹果应用名不能为空",
				"logo_required" : "logo不能为空",
				"ipaId_required" : "苹果id不能为空",
				"apkId_required" : "安卓id名不能为空",
			},
		},

		//英文语言包
		"en"	: map[string]interface{} {
			//公共模块语言包
			"comm"	: map[string]string {
				"Invalid" : "Illegal request parameters",
				"DataIsEmpty" : "Data is empty",
				"NetworkError":"network error",
				"Success"	: "Successful ",
				"Failure"	: "Pperation Failed",
				"Unlink"	: "The link has been disconnected. Please login again",
				"UploadFileTypeError_ipak" : "File type error, you can only upload the.xlsx suffix file",
			},

			//登录模块语言包
			"login"	: map[string]string {
				"ReLogin" : "The link has been disconnected. Please login again",
				"UsernameRequired" : "The user name cannot be empty",
				"CaptchaRequired" : "Verify that the code cannot be empty",
				"PasswordRequired" : "Password cant be empty",
				"CaptchaError"  : "Verification code error",
				"UserNotExist" : "User does not exists",
				"PasswordError" : "Password Error",
				"LoginFailure" : "Login failure",
				"LoginSuccess" : "Login successfully",
			},
		},
	}

	//判断键值是否存在，不存在默认为中文
	if _,ok := langPak[locale]; !ok{
		locale = "zh-cn"
	}
	jsonLang,_ := json.Marshal(langPak)
	return gjson.Get(string(jsonLang),locale + "." + module + "." +langCode).String()
}
