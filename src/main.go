package main

import (
	"gostudy/src/file"
	"os"
	"gostudy/src/common"
	"strings"
)

func main() {
	size := len(os.Args)
	if size <= 1 {
		common.Out("please input args")
		return
	} else if size <= 2 {
		common.Out("please input zip dest file")
		return
	}
	common.Out("main args : " + strings.Join(os.Args, ","))

	//build之后执行 src.exe G:\GoProject\gotest.go G:\GoProject\zipdest.zip
	file.CreateZip(os.Args[size - 1], os.Args[1:size - 1])

}
