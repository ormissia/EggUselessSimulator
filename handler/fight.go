// @File: handler
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 

package handler

import (
	"EggUselessSimulator/common"
	"EggUselessSimulator/utils"
)

var (
	roles = common.Roles
)

func StartFight() {
	for len(roles) > 0 {
		role1 := roles[utils.RandRoleName(roles)]
		role2 := roles[utils.RandRoleName(roles)]
		role1.Attack(role2)
		if role2.BloodValue <= 0 {
			delete(roles, role2.Name)
		}
	}
}
