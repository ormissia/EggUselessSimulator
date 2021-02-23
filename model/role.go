// @File: model
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 角色实体

package model

import (
	"fmt"
	"math/rand"
	"time"
)

//角色基础属性
type Role struct {
	Id          string      //Id
	Name        string      //名字
	Sex         bool        //性别
	AttackValue int         //攻击力
	BloodValue  int         //血量
	Equipments  []Equipment //装备
}

//角色能力接口定义
type RoleAbilityInterface interface {
	//TODO 修改攻击接口形参不传被攻击角色
	//角色攻击能力，计算并返回角色能造成的伤害
	Attack(role *Role) (damage int)
	//形参传入角色收到的伤害，计算并返回防御后受到的实际伤害
	Defense(damage int) (realDamage int)
}

//角色攻击其他角色
func (r *Role) Attack(role *Role) (damage int) {
	time.Sleep(time.Millisecond * 10 * time.Duration(rand.Intn(100)))
	role.BloodValue = role.BloodValue - r.AttackValue
	if r == role {
		fmt.Printf("%s 朝自己狠狠地来了一刀，对 自己 造成了 %d 点伤害，", r.Name, r.AttackValue)
		if role.BloodValue > 0 {
			fmt.Printf("把 自己 打的晃了几下，还剩 %d 点血量\n", role.BloodValue)
		} else {
			fmt.Printf("把 自己 打翻在地\n")
		}
	} else {
		fmt.Printf("%s 使出一股蛮力，对 %s 造成了 %d 点伤害，", r.Name, role.Name, r.AttackValue)
		if role.BloodValue > 0 {
			fmt.Printf("把 %s 打的晃了几下，还剩 %d 点血量\n", role.Name, role.BloodValue)
		} else {
			fmt.Printf("把 %s 打翻在地\n", role.Name)
		}
	}
	return damage
}

func (r Role) Defense(damage int) (realDamage int) {
	return realDamage
}
