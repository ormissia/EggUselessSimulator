// @File: model
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 装备实体

package model

type equipmentType uint8

const (
	attackEquipment equipmentType = iota
	defenseEquipment
	magicEquipment
)

//装备
//根据装备类型判断使用攻击力还是防御力
//当类型为魔法装备时，判断攻击力和防御力哪个是0值，如果都不是0值则随机造成伤害或者给自己加血
type Equipment struct {
	Id            string        //Id
	Name          string        //名字
	EquipmentType equipmentType //装备类型
	AttackValue   int           //攻击力
	DefenseValue  int           //防御力
}
