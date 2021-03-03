// @File: handler
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 开始战斗

package handler

import (
	"EggUselessSimulator/common"
	"EggUselessSimulator/model"
	"EggUselessSimulator/utils"
	"fmt"
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

		//获取攻击者能造成的伤害
		damage := role1.Attack()
		//计算被攻击者实际受到的伤害
		realDamage := role2.Defense(damage)

		//计算被攻击角色剩余血量
		role2.BloodValue = role2.BloodValue - realDamage

		//攻击招式
		a := "拳头"
		if role1.AttackEquipment.AttackValue != 0 {
			a = role1.AttackEquipment.Name
		}
		//受伤程度
		b := "打的摇摇晃晃"
		//当被攻击者血量小于等于0时，将其移出map
		if role2.BloodValue <= 0 {
			delete(aliveRoles, role2.Name)
			b = "打翻在地"
		}

		//根据攻击者输出和被攻击者受到的伤害分情况输出战斗信息
		fmt.Printf("%s 使出 %s ，造成了 %d 点伤害，把 %s %s\n", role1.Name, a, damage, role2.Name, b)
	}

	fmt.Printf("最终，%s 获胜", aliveRoles[utils.RandRoleName(aliveRoles)].Name)

}
