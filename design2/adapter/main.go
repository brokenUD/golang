// 定义
// 适配器模式是一种结构型设计模式，它允许将一个类的接口转换成客户端期望的另一个接口。适配器模式通常用于解决接口不兼容的问题。

// 适用场景
//     当你希望使用一个已经存在的类，但其接口与你的需求不匹配时。
//     当你希望创建一个可以复用的类，该类可以与不相关的类或不可预见的类协同工作。
//     当你希望统一多个类的接口时

// 对比装饰器模式和适配器模式
// 特性			装饰器模式						  适配器模式
// 目的			动态扩展对象的功能				   解决接口不兼容的问题
// 核心思想		包装对象，添加新行为				转换接口，使其兼容
// 适用场景		需要动态添加或移除功能				需要复用现有类，但接口不匹配
// 是否修改原对象	 不修改原对象，通过包装扩展功能		   不修改原对象，通过适配器转换接口
// 示例			为咖啡添加配料				       将旧打印机接口适配到新打印机接口

package main

import (
	"fmt"
)

type NewLogger interface{
	Log(message string)
}

type OldLogger struct{}
func (l *OldLogger) LogMessage(msg string, level int){
	fmt.Printf("[Old] level %d: %s\n", level, msg)
}

type LoggerAdapter struct{
	OldLogger *OldLogger
}
func (a *LoggerAdapter)Log(message string){
	a.OldLogger.LogMessage(message, 1)
}

func main(){
	newLogger := &LoggerAdapter{OldLogger: &OldLogger{}}
	newLogger.Log("Adapted successfully!")
}
