// @File: common
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 

package common

import (
	"EggUselessSimulator/model"
	"EggUselessSimulator/utils"
	"fmt"
	"math/rand"
	"time"
)

var (
	Names      [108]string
	equipments [18]string
)

var (
	Roles = make(map[string]*model.Role)
)

func init() {
	//初始化角色名字数组
	Names = [108]string{"宋江", "卢俊义", "吴用", "公孙胜", "关胜", "林冲", "秦明", "呼延灼", "花荣", "柴进", "李应", "朱仝", "鲁智深", "武松", "董平", "张清", "杨志", "徐宁", "索超", "戴宗", "刘唐", "李逵", "史进", "穆弘", "雷横", "李俊", "阮小二", "张横", "阮小五", "张顺", "阮小七", "扬雄", "石秀", "解珍", "解宝", "燕青", "朱武", "黄信", "孙立", "宣赞", "郝思文", "韩滔", "彭玘", "单廷圭", "魏定国", "萧让", "裴宣", "欧鹏", "邓飞", "燕顺", "杨林", "凌振", "蒋敬", "吕方", "郭盛", "道全", "皇甫端", "王英", "扈三娘", "鲍旭", "樊瑞", "孔明", "孔亮", "项充", "李衮", "金大坚", "马麟", "童威", "童猛", "孟康", "侯健", "陈达", "杨春", "郑天寿", "陶宗旺", "宋清", "乐和", "龚旺", "丁得孙", "穆春", "曹正", "宋万", "杜迁", "薛永", "施恩", "李忠", "周通", "汤隆", "杜兴", "邹渊", "邹润", "朱贵", "朱富", "蔡福", "蔡庆", "李立", "李云", "焦挺", "石勇", "孙新", "顾大嫂", "张青", "孙二娘", "王定六", "郁保四", "白胜", "时迁", "段景住"}
}

func InitRoles() {
	for _, name := range Names {
		Roles[name] = &model.Role{
			Name:        name,
			AttackValue: utils.RandInt(10, 20),
			BloodValue:  utils.RandInt(50, 100),
		}
	}
	for _, role := range Roles {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		utils.PrintToScreen(fmt.Sprintf("%s 大步来到殿堂之上", role.Name))
	}
}
