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
	//获取战斗开始时间
	startTime := time.Now()

	//初始化角色数据
	common.InitRoles()
	//开始战斗
	handler.StartFight()

	//计算战斗时间
	elapsed := time.Since(startTime)
	fmt.Printf("好家伙！经过了 %d 年，这一仗 打的那叫一个天昏地暗。", elapsed/1000000000)
}
