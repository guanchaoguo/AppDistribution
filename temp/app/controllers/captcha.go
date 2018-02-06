package controllers

import (
	"bytes"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"golang-AppDistribution/app/helper"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
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
func (Captcaha) GetCaptcha(w context.Context) {
	d := make([]byte, 4)
	s := NewLen(4)
	urlCode := ""
	d = []byte(s)
	for v := range d {
		d[v] %= 10
		urlCode += strconv.FormatInt(int64(d[v]), 32)
	}

	fmt.Println(urlCode)

	//生成随机数 存放到redis中
	captchaKey := helper.GetRandomString(32)
	r := helper.GetRedis()
	_, err := r.Do("SET", captchaKey, urlCode)
	defer r.Close()
	if err != nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": "操作失败",
			"result":  "",
		})
		w.Request().Body.Close()
		return
	}

	var buf bytes.Buffer
	codeImg := NewImage(d, 100, 40)
	if err := png.Encode(&buf, codeImg); err != nil {
		panic(err.Error())
	}

	base64Captcha := fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(buf.Bytes()))

	w.JSON(iris.Map{
		"code":    0,
		"message": "操作成功",
		"result": iris.Map{
			"data":iris.Map{
				"captchaKey": captchaKey,
				"captcha":    base64Captcha,
			},
		},
	})
}

//检测验证码是否正确
func (Captcaha) CheckCaptcha(code string, captcha string) bool {
	//获取redis中的code
	r := helper.GetRedis()
	redis_captcha, _ := redis.String(r.Do("GET", code))
	//删除本次key(一次性)
	r.Do("DEL", code)
	defer r.Close() //关闭redis

	if redis_captcha != strings.ToUpper(captcha) {
		return false
	}
	return true
}

const (
	stdWidth  = 100
	stdHeight = 40
	maxSkew   = 2
)

const (
	fontWidth  = 5
	fontHeight = 8
	blackChar  = 1
)

var font = [][]byte{
	{ // 0
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
	},
	{ // 1
		0, 0, 1, 0, 0,
		0, 1, 1, 0, 0,
		1, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		1, 1, 1, 1, 1,
	},
	{ // 2
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 1, 1,
		0, 1, 1, 0, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 1,
	},
	{ // 3
		1, 1, 1, 1, 0,
		0, 0, 0, 0, 1,
		0, 0, 0, 1, 0,
		0, 1, 1, 1, 0,
		0, 0, 0, 1, 0,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
	},
	{ // 4
		1, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 1, 1, 1, 1,
		0, 0, 0, 1, 0,
		0, 0, 0, 1, 0,
		0, 0, 0, 1, 0,
	},
	{ // 5
		1, 1, 1, 1, 1,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
	},
	{ // 6
		0, 0, 1, 1, 1,
		0, 1, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
	},
	{ // 7
		1, 1, 1, 1, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 1, 0,
		0, 0, 1, 0, 0,
		0, 1, 0, 0, 0,
		0, 1, 0, 0, 0,
		0, 1, 0, 0, 0,
	},
	{ // 8
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
	},
	{ // 9
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 1, 0, 0, 1,
		0, 1, 1, 1, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
	},
}

type Image struct {
	*image.NRGBA
	color   *color.NRGBA
	width   int //a digit width
	height  int //a digit height
	dotsize int
}

func init() {
	rand.Seed(int64(time.Second))
}

func NewImage(digits []byte, width, height int) *Image {
	img := new(Image)
	r := image.Rect(img.width, img.height, stdWidth, stdHeight)
	img.NRGBA = image.NewNRGBA(r)

	img.color = &color.NRGBA{
		uint8(rand.Intn(129)),
		uint8(rand.Intn(129)),
		uint8(rand.Intn(129)),
		0xFF,
	}
	// Draw background (10 random circles of random brightness)
	img.calculateSizes(width, height, len(digits))
	img.fillWithCircles(10, img.dotsize)

	maxx := width - (img.width+img.dotsize)*len(digits) - img.dotsize
	maxy := height - img.height - img.dotsize*2

	x := rnd(img.dotsize*2, maxx)
	y := rnd(img.dotsize*2, maxy)

	// Draw digits.
	for _, n := range digits {
		img.drawDigit(font[n], x, y)
		x += img.width + img.dotsize
	}

	// Draw strike-through line.
	img.strikeThrough()
	return img
}

func (img *Image) WriteTo(w io.Writer) (int64, error) {
	return 0, png.Encode(w, img)
}

func (img *Image) calculateSizes(width, height, ncount int) {

	// Goal: fit all digits inside the image.
	var border int
	if width > height {
		border = height / 5
	} else {
		border = width / 5
	}
	// Convert everything to floats for calculations.
	w := float64(width - border*2)  //268
	h := float64(height - border*2) //48
	// fw takes into account 1-dot spacing between digits.

	fw := float64(fontWidth) + 1 //6

	fh := float64(fontHeight) //8
	nc := float64(ncount)     //7

	// Calculate the width of a single digit taking into account only the
	// width of the image.
	nw := w / nc //38
	// Calculate the height of a digit from this width.
	nh := nw * fh / fw //51

	// Digit too high?

	if nh > h {
		// Fit digits based on height.
		nh = h //nh = 44
		nw = fw / fh * nh
	}
	// Calculate dot size.
	img.dotsize = int(nh / fh)
	// Save everything, making the actual width smaller by 1 dot to account
	// for spacing between digits.
	img.width = int(nw)
	img.height = int(nh) - img.dotsize
}

func (img *Image) fillWithCircles(n, maxradius int) {
	color := img.color
	maxx := img.Bounds().Max.X
	maxy := img.Bounds().Max.Y
	for i := 0; i < n; i++ {
		setRandomBrightness(color, 255)
		r := rnd(1, maxradius)
		img.drawCircle(color, rnd(r, maxx-r), rnd(r, maxy-r), r)
	}
}

func (img *Image) drawHorizLine(color color.Color, fromX, toX, y int) {
	for x := fromX; x <= toX; x++ {
		img.Set(x, y, color)
	}
}

func (img *Image) drawCircle(color color.Color, x, y, radius int) {
	f := 1 - radius
	dfx := 1
	dfy := -2 * radius
	xx := 0
	yy := radius

	img.Set(x, y+radius, color)
	img.Set(x, y-radius, color)
	img.drawHorizLine(color, x-radius, x+radius, y)

	for xx < yy {
		if f >= 0 {
			yy--
			dfy += 2
			f += dfy
		}
		xx++
		dfx += 2
		f += dfx
		img.drawHorizLine(color, x-xx, x+xx, y+yy)
		img.drawHorizLine(color, x-xx, x+xx, y-yy)
		img.drawHorizLine(color, x-yy, x+yy, y+xx)
		img.drawHorizLine(color, x-yy, x+yy, y-xx)
	}
}

func (img *Image) strikeThrough() {
	r := 0
	maxx := img.Bounds().Max.X
	maxy := img.Bounds().Max.Y
	y := rnd(maxy/3, maxy-maxy/3)
	for x := 0; x < maxx; x += r {
		r = rnd(1, img.dotsize/3)
		y += rnd(-img.dotsize/2, img.dotsize/2)
		if y <= 0 || y >= maxy {
			y = rnd(maxy/3, maxy-maxy/3)
		}
		img.drawCircle(img.color, x, y, r)
	}
}

func (img *Image) drawDigit(digit []byte, x, y int) {
	skf := rand.Float64() * float64(rnd(-maxSkew, maxSkew))
	xs := float64(x)
	minr := img.dotsize / 2               // minumum radius
	maxr := img.dotsize/2 + img.dotsize/4 // maximum radius
	y += rnd(-minr, minr)
	for yy := 0; yy < fontHeight; yy++ {
		for xx := 0; xx < fontWidth; xx++ {
			if digit[yy*fontWidth+xx] != blackChar {
				continue
			}
			// Introduce random variations.
			or := rnd(minr, maxr)
			ox := x + (xx * img.dotsize) + rnd(0, or/2)
			oy := y + (yy * img.dotsize) + rnd(0, or/2)

			img.drawCircle(img.color, ox, oy, or)
		}
		xs += skf
		x = int(xs)
	}
}

func setRandomBrightness(c *color.NRGBA, max uint8) {
	minc := min3(c.R, c.G, c.B)
	maxc := max3(c.R, c.G, c.B)
	if maxc > max {
		return
	}
	n := rand.Intn(int(max-maxc)) - int(minc)
	c.R = uint8(int(c.R) + n)
	c.G = uint8(int(c.G) + n)
	c.B = uint8(int(c.B) + n)
}

func min3(x, y, z uint8) (o uint8) {
	o = x
	if y < o {
		o = y
	}
	if z < o {
		o = z
	}
	return
}

func max3(x, y, z uint8) (o uint8) {
	o = x
	if y > o {
		o = y
	}
	if z > o {
		o = z
	}
	return
}

func rnd(from, to int) int {
	return rand.Intn(to+1-from) + from
}

const (
	StdLen  = 16
	UUIDLen = 20
)

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func New() string {
	return NewLenChars(StdLen, StdChars)
}

func NewLen(length int) string {
	return NewLenChars(length, StdChars)
}

func NewLenChars(length int, chars []byte) string {
	b := make([]byte, length)
	r := make([]byte, length+(length/4))
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(crand.Reader, r); err != nil {
			panic("error reading from random source: " + err.Error())
		}
		for _, c := range r {
			if c >= maxrb {

				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
	panic("unreachable")
}

func (Captcaha) Pic(w context.Context) {
	d := make([]byte, 4)
	s := NewLen(4)
	urlCode := ""
	d = []byte(s)
	for v := range d {
		d[v] %= 10
		urlCode += strconv.FormatInt(int64(d[v]), 32)
	}

	//存放到redis中
	r := helper.GetRedis()
	_, err := r.Do("SET", urlCode, "")
	defer r.Close()
	if err != nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": "操作失败",
			"result":  "",
		})
		w.Request().Body.Close()
		return
	}

	var buf bytes.Buffer
	codeImg := NewImage(d, 100, 40)
	if err := png.Encode(&buf, codeImg); err != nil {
		panic(err.Error())
	}

	base64Captcha := fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(buf.Bytes()))
	w.JSON(iris.Map{
		"code":    0,
		"message": "操作成功",
		"result": iris.Map{
			"idkey":   urlCode,
			"captcha": base64Captcha,
		},
	})
}

func (Captcaha) Index(w context.Context) {
	str := "<meta charset=\"utf-8\"><h3>golang 图片验证码例子</h3><img border=\"1\" src=\"/pic\" alt=\"图片验证码\" onclick=\"this.src='/pic'\" />"
	w.Header("Content-Type", "text/html")
	w.Write([]byte(str))
}
