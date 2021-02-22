# 没什么卵用模拟器🤡🤡🤡

[![Project Label](https://badgen.net/badge/github/GoLang/blue?label=Egg%20Useless%20Simulator)](https://github.com/ormissia/EggUselessSimulator)
[![Go Report Card](https://goreportcard.com/badge/github.com/ormissia/EggUselessSimulator)](https://goreportcard.com/report/github.com/ormissia/EggUselessSimulator)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/ormissia/EggUselessSimulator)](https://www.tickgit.com/browse?repo=github.com/ormissia/EggUselessSimulator)
[![licenses MIT](https://badgen.net/badge/license/MIT/#3DA639)](https://opensource.org/licenses/MIT)

> 用代码模拟一下水浒梁山内讧

## 项目演示
![](https://ormissia-blog.oss-cn-qingdao.aliyuncs.com/image-hosting/egguselesssimulator-1.png)

## 项目路线图
- [ ] **程序**
    - [ ] 注释完善
    - [ ] 通过`ini`配置文件生成人物，装备等属性
    - [ ] 使用线程安全的map，将战斗过程改为并发进行
    - [ ] 程序结束时输出统计数据，如共造成多少伤害、共使用了多少装备、共有多少人死于随机事件、共有多少人被自己杀死等等......
- [x] **战斗**
    - [x] 随机大乱斗
    - [x] 战斗至最后一个人
    - [ ] 拾取装备
      - [ ] 攻击属性的装备
      - [ ] 防御属性的装备。加护甲后伤害计算，实际造成的伤害 = 攻击方总伤害 - 护甲值（最低伤害不得小于1）
    - [ ] 增加技能
    - [ ] 增加物品
    - [ ] 增加随机事件（如房屋倒塌等）
    - [ ] 击败三人获得一次复活机会
