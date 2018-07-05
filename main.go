package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

/**
 * description:   go调用phantomjs打开网页截屏保存
 * author:        lmw
 * date:          2018/07/05
 */
func main() {
	//保存文件名
	var picPath = "pic/" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
	//北京时间
	var url = "https://www.baidu.com/s?tn=ubuntuu_dg&ie=UTF-8&wd=%E5%8C%97%E4%BA%AC%E6%97%B6%E9%97%B4"

	//脚本位置
	var script = "script/ss.js"

	//phantomjs可执行文件位置
	var phantomjs = "./phantomjs"

	cmd := exec.Command(phantomjs, script, url, picPath)

	fmt.Println(strings.Join(cmd.Args, ","))

	_, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("--截图完成--")
		fmt.Println(err)
		return
	}

	defer cmd.Process.Kill()
}
