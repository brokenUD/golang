// 代理模式 是一种结构型设计模式，
// 其核心思想是 通过代理对象控制对原始对象的访问。
// 代理对象与原始对象实现相同的接口，客户端通过代理间接操作原始对象，
// 从而在不修改原始对象代码的前提下，增强其功能或限制其访问。

// 代理模式的使用场景
//     延迟加载（Lazy Loading）：对象创建成本高时，代理延迟实例化。
//     访问控制（Protection Proxy）：限制对敏感对象的访问权限。
//     日志记录/监控：代理拦截请求并记录操作日志。
//     远程代理（Remote Proxy）：代理隐藏远程服务调用的复杂性（如 RPC）

package main

import (
	"fmt"
	"time"
)

// --------------------- 抽象接口 ---------------------
type Image interface {
	Display()
}

// --------------------- 真实对象 ---------------------
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	fmt.Printf("加载图片: %s\n", filename)
	// 模拟耗时操作
	time.Sleep(2 * time.Second)
	return &RealImage{filename: filename}
}

func (i *RealImage) Display() {
	fmt.Printf("显示图片: %s\n", i.filename)
}

// --------------------- 代理对象 ---------------------
type ProxyImage struct {
	filename  string
	realImage *RealImage // 延迟加载真实对象
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename) // 首次调用时加载真实对象
	}
	p.realImage.Display()
}

// --------------------- 客户端调用 ---------------------
func main() {
	// 创建代理（此时不加载真实图片）
	image := NewProxyImage("wallpaper.jpg")

	// 第一次调用时加载并显示图片
	image.Display() // 输出：加载图片 → 显示图片

	// 第二次直接显示已加载的图片
	image.Display() // 输出：显示图片
}
