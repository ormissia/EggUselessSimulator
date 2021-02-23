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
	Id          string //Id
	Name        string //名字
	Sex         bool   //性别
	AttackValue int    //攻击力
	BloodValue  int    //血量
}

//角色攻击其他角色
func (r1 *Role) Attack(r2 *Role) {
	time.Sleep(time.Millisecond * 10 * time.Duration(rand.Intn(100)))
	r2.BloodValue = r2.BloodValue - r1.AttackValue
	if r1 == r2 {
		fmt.Printf("%s 朝自己狠狠地来了一刀，对 自己 造成了 %d 点伤害，", r1.Name, r1.AttackValue)
		if r2.BloodValue > 0 {
			fmt.Printf("把 自己 打的晃了几下，还剩 %d 点血量\n", r2.BloodValue)
		} else {
			fmt.Printf("把 自己 打翻在地\n")
		}
	} else {
		fmt.Printf("%s 使出一股蛮力，对 %s 造成了 %d 点伤害，", r1.Name, r2.Name, r1.AttackValue)
		if r2.BloodValue > 0 {
			fmt.Printf("把 %s 打的晃了几下，还剩 %d 点血量\n", r2.Name, r2.BloodValue)
		} else {
			fmt.Printf("把 %s 打翻在地\n", r2.Name)
		}
	}

}
