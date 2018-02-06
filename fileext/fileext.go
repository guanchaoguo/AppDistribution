package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	fullFilename := "/Users/itfanr/Documents/test.txt"
	fmt.Println("fullFilename =", fullFilename)

	filename := path.Base(fullFilename) //获取文件名带后缀
	fmt.Println("filename =", filename)

	fileExt := path.Ext(filename) //获取文件后缀
	fmt.Println("fileExt =", fileExt)


	filenameOnly := strings.TrimSuffix(filename, fileExt)//获取文件名
	fmt.Println("filenameOnly =", filenameOnly)
}
