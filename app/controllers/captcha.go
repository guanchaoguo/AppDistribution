package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/mojocn/base64Captcha"
	"golang-AppDistribution/app/helper"
)

type Captcaha struct{}

/**
* @api {Get} /getCaptcha 获取验证码
* @apiDescription 获取验证码
* @apiGroup captcha
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiHeader {String} Authorization token.
* @apiSuccessExample {json} Success-Response:
*{
*	code: 0,
*	message: "操作成功",
*	result: - {
*		data: - {
*			captcha: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAAAoCAYAAAAIeF9DAAADXklEQVR4nOxbvapUMRDOvcjCdltb2CkoKIi3WBRRtBCsfADBzgcQLYULlooPYCf4AL6AriLIFlcEBQXtLKy3W9hmJYebJY6TZCaZyclZ94PD5vxk5ku+zMw54d59s0NTGKwgd+7eW/fNQQN7fRPIgS/G61cvVcbw8PDJxsezw8fV5mmQgphjUWqI4UAR5eSZs5t+v398z+I22JSlJUYIl67diKZIXwzsnIrBClITs9ms+02JIoEmBTl98WDtjj78w/T06f3bPf9XE83VEEyEn5+PmuOJQaKGNDXQWEQMRZRSkFPWhctX16l2a9DgrD1e0qqjkPjy8UPRCrY+lstl1x6Pxwa2uRGiwdm3WTLe+w8ebey8eP70LzukCHHO7W+onUvOeAO1k28PrM2FBmesr+XuDkNYCL4Y2HmQlDXMIc19PtYfRktNHileUIxUf8gFCmBAlKDEQ6HJvZ6CxuTl8KDY82Ftc2qJzwUTZD6fb54hRwhGAN7nioHZKYWmyBiwOUjNlRXFijCdTjsxfPT2YShVf4ywuJx6AH3BWoUtaPvrRIBiGE7KiqnOmZBaKxj6wFYvtR5gqz43G6Tss4p6bOIpE62ZpkL3uPke6+/7yeWNCYrZI6UsZwwLRQ64aYqaOig2qJB8nYd2Kfb/uZBabdj13JQVa1P950x4Lm/qsyUIpqQQYqJwBpWL1IQ6HqfOnV9PJpPNPYnaJV3/MOz7zqikclNJ6us59jz0H4siK4ZtLxaL7rDn3A+6vtAJAlcapU74opQKRLlH5ePEgHDXc7lK1DMKgq+BkEwqv8f6+3awPqFXytTrN1Z/QoJY/Pr2Fd0CoaJGyjrhGtTVGhPGfxujko8tCMyO9oT0jS5l5e58hlIbFt7cr14qoC8XBRB+dFD4pPxoodqeD0RoBzVkMyYYZsNPXVCkllOWCKAYsOCXFH7K8xo2JftxULy5iE009rqqtbK4QreeskTA/YaRfn6bIiTpmNLOta39fC5f7vglhRLdG6Li4PrNzsdqtUKLe0nRLeUrtTMsCs1VYcWwByz42KHFXcJGtQjRhosOcxwho9HIHL17sycRIduAXv+214rhsO1f41du3SZFUXVBbDTEzrcRTgyqKDtUwE6MgaLJ/w/5n7ETpDH8CQAA//+M1lV7zXAz+QAAAABJRU5ErkJggg==",
*			captchaKey: "wkj4aVGWJKbmXafRQbraWcH0L3da5DZq"
*		}
* 	}
* }
*
 */
func (Captcaha) GetCaptcha(ctx context.Context) {
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeArithmetic,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}

	//创建字符公式验证码.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

	ctx.JSON(iris.Map{
		"code":    0,
		"message": "操作成功",
		"result": iris.Map{
			"data":iris.Map{
				"captchaKey": idKeyC,
				"captcha":    base64stringC,
			},
		},
	})
}

//检测验证码是否正确
func (Captcaha) CheckCaptcha(idkey string, verifyValue string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	if verifyResult {
		return true
	} else {
		return false
	}
}

func NewLen(length int) string {
	return helper.GetRandomString(length)
}

