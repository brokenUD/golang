// 装饰器模式的核心思想是动态地扩展对象的功能，而不改变其原始结构

package main

import "fmt"

type DataProcessor interface{
	Process(data string)string
}

type BasicProcessor struct{}
func (p *BasicProcessor)Process(data string)string{
	return "Processed: " + data
}

type EncryptDecorator struct{
	Processor DataProcessor
}
func (d *EncryptDecorator)Process(data string)string{
	encrypted := "Encrypted[" + data + "]"
	return d.Processor.Process(encrypted)
}

type CompressDecorator struct{
	Processor DataProcessor
}
func (d *CompressDecorator)Process(data string)string{
	compressed := "Compressed(" + data + ")"
	return d.Processor.Process(compressed)
}

func main(){
	processor := &BasicProcessor{}

	encryptedProcessor := &EncryptDecorator{Processor: processor}
	compressDecorator := &CompressDecorator{Processor: encryptedProcessor}

	result := compressDecorator.Process("Hello")
	fmt.Println(result)
}




// 为什么嵌套顺序会影响结果？
//     装饰器的执行顺序：
//         装饰器模式的核心是将对象包装在装饰器中，每一层装饰器都会在调用原始对象之前或之后执行一些额外的逻辑。
//         嵌套顺序决定了装饰器的执行顺序。外层的装饰器会先执行，然后依次向内层装饰器传递调用，最后到达原始对象。

//     功能的叠加方式：
//         每一层装饰器可能会对输入参数、输出结果或对象的状态进行修改。
//         如果嵌套顺序不同，装饰器对输入参数或输出结果的修改顺序也会不同，从而导致最终结果不同。

//     行为的组合方式：
//         装饰器的功能是叠加的，每一层装饰器都会在前一层的基础上增加新的行为。
//         如果嵌套顺序不同，行为的组合方式也会不同，最终的表现也会不同。
