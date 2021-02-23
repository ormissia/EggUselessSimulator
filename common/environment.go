// @File: common
// @Date: 2021/2/22
// @Author: 安红豆
// @Description: 初始化战斗环境

package common

import (
	"EggUselessSimulator/model"
	"EggUselessSimulator/utils"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"math/rand"
	"time"
)

var (
	delimiter  string //配置文件中的分隔符
	attackTime []int  //最小攻击时间，最大攻击时间

	//名字数组
	roleNames = []string{"宋江", "卢俊义", "吴用", "公孙胜", "关胜", "林冲", "秦明", "呼延灼", "花荣", "柴进", "李应", "朱仝", "鲁智深", "武松", "董平", "张清", "杨志", "徐宁", "索超", "戴宗", "刘唐", "李逵", "史进", "穆弘", "雷横", "李俊", "阮小二", "张横", "阮小五", "张顺", "阮小七", "扬雄", "石秀", "解珍", "解宝", "燕青", "朱武", "黄信", "孙立", "宣赞", "郝思文", "韩滔", "彭玘", "单廷圭", "魏定国", "萧让", "裴宣", "欧鹏", "邓飞", "燕顺", "杨林", "凌振", "蒋敬", "吕方", "郭盛", "道全", "皇甫端", "王英", "扈三娘", "鲍旭", "樊瑞", "孔明", "孔亮", "项充", "李衮", "金大坚", "马麟", "童威", "童猛", "孟康", "侯健", "陈达", "杨春", "郑天寿", "陶宗旺", "宋清", "乐和", "龚旺", "丁得孙", "穆春", "曹正", "宋万", "杜迁", "薛永", "施恩", "李忠", "周通", "汤隆", "杜兴", "邹渊", "邹润", "朱贵", "朱富", "蔡福", "蔡庆", "李立", "李云", "焦挺", "石勇", "孙新", "顾大嫂", "张青", "孙二娘", "王定六", "郁保四", "白胜", "时迁", "段景住"}
	//最小攻击力，最大攻击力
	attackValue []int
	//最小血量，最大血量
	bloodValue []int

	attackEquipmentNames  []string //攻击装备名字数组
	defenseEquipmentNames []string //防御装备名字数组
	magicEquipmentNames   []string //魔法装备名字数组

	Roles            = make(map[string]*model.Role)      //要生成的角色map
	AttackEquipment  = make(map[string]*model.Equipment) //要生成的攻击装备map
	DefenseEquipment = make(map[string]*model.Equipment) //要生成的防御装备map
	MagicEquipment   = make(map[string]*model.Equipment) //要生成的魔法装备map
)

func init() {
	//初始化配置文件中的配置项
	initFileConfigs()
}

//初始化所有角色
func InitRoles() {
	//遍历名字数组，将名字作为key生成角色map
	for _, name := range roleNames {
		Roles[name] = &model.Role{
			Name:        name,
			AttackValue: utils.RandInt(10, 20),  //随机生成攻击力
			BloodValue:  utils.RandInt(50, 100), //随机生成血量
		}
	}

	//遍历生成的角色map，随机角色出场顺序
	for _, role := range Roles {
		//每个人登场之前随机睡眠100毫秒之内的时间
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		fmt.Printf("%s 大步来到殿堂之上\n", role.Name)
	}
}

//初始化配置文件中的配置项
func initFileConfigs() {
	//获取配置文件
	cfg, err := ini.Load("eggUseless.ini")
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
	}

	//初始化配置文件的分隔符
	delimiter = cfg.Section("programConfig").Key("delimiter").String()
	//初始化攻击时间
	attackTime = cfg.Section("programConfig").Key("attackTime").Ints(delimiter)
	//初始化角色名字数组
	roleNames = cfg.Section("role").Key("names").Strings(delimiter)
	//初始化攻击力
	attackValue = cfg.Section("role").Key("attackValue").Ints(delimiter)
	//初始化防御力
	bloodValue = cfg.Section("role").Key("bloodValue").Ints(delimiter)

	//初始化攻击装备名称数组
	attackEquipmentNames = cfg.Section("equipment").Key("attackNames").Strings(delimiter)
	//初始化防御装备名称数组
	defenseEquipmentNames = cfg.Section("equipment").Key("defenseNames").Strings(delimiter)
	//初始化魔法装备名称数组
	magicEquipmentNames = cfg.Section("equipment").Key("magicNames").Strings(delimiter)
}
