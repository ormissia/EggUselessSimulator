// @File: main
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 程序入口

package main

import (
	"EggUselessSimulator/common"
	"EggUselessSimulator/handler"
	"fmt"
	"time"
)

func main() {
	fmt.Println("这真的没什么卵用")
	startTime := time.Now()

	common.InitRoles()
	handler.StartFight()

	elapsed := time.Since(startTime)
	fmt.Printf("好家伙！经过了 %d 年，这一仗 打的那叫一个天昏地暗。", elapsed/1000000000)
}
