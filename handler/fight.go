// @File: handler
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 开始战斗

package handler

import (
	"EggUselessSimulator/common"
	"EggUselessSimulator/model"
	"EggUselessSimulator/utils"
)

var (
	aliveRoles = make(map[string]*model.Role) //活着的角色map
	deadRoles  = make(map[string]*model.Role) //死亡的角色map
)

//将初始化得到的初始角色map赋值到活着的角色map中
func init() {
	aliveRoles = common.Roles
}

//开始大乱斗
func StartFight() {
	//当剩余角色大于1人时持续战斗
	for len(aliveRoles) > 1 {
		//TODO 判断剩余人数，当剩余人数10人、5人、2人时分别输出提示
		//随机获取两个角色
		role1 := aliveRoles[utils.RandRoleName(aliveRoles)]
		role2 := aliveRoles[utils.RandRoleName(aliveRoles)]
		role1.Attack(role2)

		//当被攻击者血量小于等于0时，将其移出map
		if role2.BloodValue <= 0 {
			delete(aliveRoles, role2.Name)
		}
	}

}
