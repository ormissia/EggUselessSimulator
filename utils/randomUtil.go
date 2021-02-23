// @File: utils
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 随机生成工具

package utils

import (
	"EggUselessSimulator/model"
	"math/rand"
)

//获取min到max范围的随机整数
func RandInt(min, max int) (randNum int) {
	if min >= max || min == 0 || max == 0 {
		randNum = max
		return
	}
	randNum = rand.Intn(max-min) + min
	return
}

//用range的方式从map中获取第一个角色名字（伪随机）
func RandRoleName(roles map[string]*model.Role) (roleName string) {
	for _, role := range roles {
		roleName = role.Name
		return
	}
	return
}
