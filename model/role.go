// @File: model
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 角色实体

package model

import (
	"math/rand"
	"time"
)

//角色能力接口定义
type RoleAbilityInterface interface {
	//角色攻击能力，计算并返回角色能造成的伤害
	Attack() (damage int)
	//形参传入角色收到的伤害，计算并返回防御后受到的实际伤害
	Defense(damage int) (realDamage int)
}

//角色基础属性
type Role struct {
	Id               string      //Id
	Name             string      //名字
	Sex              bool        //性别
	AttackValue      int         //攻击力
	BloodValue       int         //血量
	AttackEquipment  Equipment   //攻击装备
	DefenseEquipment Equipment   //防御装备
	MagicEquipment   []Equipment //魔法装备
}

//角色攻击其他角色
func (r *Role) Attack() (damage int) {
	time.Sleep(time.Millisecond * 10 * time.Duration(rand.Intn(100)))
	//如果没有攻击装备，值即为 +0 所以不需要对角色是否有攻击装备进行判断
	damage = r.AttackValue + r.AttackEquipment.AttackValue
	return
}

//角色防御技能，根据角色受到的伤害计算角色受到的真实伤害
func (r Role) Defense(damage int) (realDamage int) {
	//角色实际受到的伤害 = 攻击者造成的伤害 - 被攻击角色护甲值
	realDamage = damage - r.DefenseEquipment.DefenseValue

	//当被攻击角色护甲值大于等于攻击者造成的伤害时，判定造成最低伤害-1点伤害值
	if realDamage < 1 {
		realDamage = 1
	}

	return
}
