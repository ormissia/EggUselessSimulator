// @File: utils
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 

package utils

import (
	"EggUselessSimulator/model"
	"math/rand"
)

func RandInt(min, max int) (randNum int) {
	if min >= max || min == 0 || max == 0 {
		randNum = max
		return
	}
	randNum = rand.Intn(max-min) + min
	return
}

func RandRoleName(roles map[string]*model.Role) (roleName string) {
	for _, role := range roles {
		roleName = role.Name
		return
	}
	return
}
