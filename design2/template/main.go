// 模板方法模式 是一种行为型设计模式，其核心思想是 定义一个算法的框架，
// 将某些步骤的具体实现延迟到子类中。
// 它通过父类（或接口）定义算法的骨架，子类在不改变算法结构的前提下重写特定步骤，实现代码复用和扩展。

// 模板方法模式的使用场景
//     多个类有相同流程，但某些步骤实现不同（如数据处理、游戏流程）。
//     需要固定算法流程，但允许部分步骤自定义（如文件解析、编译流程）。
//     避免代码冗余，将公共逻辑提取到父类。

// 适用场景总结
//     标准化流程：如数据导出（步骤：验证数据 → 生成文件 → 上传到服务器）。
//     框架设计：如 Web 请求处理（步骤：鉴权 → 参数解析 → 执行业务 → 返回结果）。
//     复杂算法分解：如编译器的编译流程（词法分析 → 语法分析 → 生成中间代码 → 优化）

package main

import "fmt"

// --------------------- 抽象类 ---------------------
type BeverageTemplate interface {
	BoilWater()     // 烧水（公共步骤）
	Brew()          // 冲泡（抽象步骤）
	PourInCup()     // 倒入杯子（公共步骤）
	AddCondiments() // 添加调料（抽象步骤）
}

// 模板方法（定义算法骨架）
func MakeBeverage(b BeverageTemplate) {
	b.BoilWater()
	b.Brew()
	b.PourInCup()
	b.AddCondiments()
}

// --------------------- 具体子类 ---------------------
// 制作咖啡
type Coffee struct{}

func (c *Coffee) BoilWater() {
	fmt.Println("烧水（100℃）")
}

func (c *Coffee) Brew() {
	fmt.Println("冲泡咖啡粉")
}

func (c *Coffee) PourInCup() {
	fmt.Println("倒入咖啡杯")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("加糖和牛奶")
}

// 制作茶
type Tea struct{}

func (t *Tea) BoilWater() {
	fmt.Println("烧水（80℃）")
}

func (t *Tea) Brew() {
	fmt.Println("浸泡茶叶")
}

func (t *Tea) PourInCup() {
	fmt.Println("倒入茶杯")
}

func (t *Tea) AddCondiments() {
	fmt.Println("加柠檬")
}

// --------------------- 客户端调用 ---------------------
func main() {
	fmt.Println("制作咖啡:")
	MakeBeverage(&Coffee{})

	fmt.Println("\n制作茶:")
	MakeBeverage(&Tea{})
}
