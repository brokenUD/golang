// 建造者模式 是一种创建型设计模式，
// 其核心思想是 将一个复杂对象的构建与其表示分离，
// 使得同样的构建过程可以创建不同的表示。
// 通过逐步构造对象，建造者模式避免了冗长的构造函数参数列表，并支持灵活的对象配置。

// 建造者模式的使用场景
//     构造对象需要多个步骤（如配置复杂的对象）。
//     对象有多种表示形式（如不同品牌的电脑配置）。
//     避免构造函数参数过多（提高代码可读性）。

package main

import "fmt"

// --------------------- 产品：电脑 ---------------------
type Computer struct {
	CPU     string
	Memory  string
	Storage string
}

func (c *Computer) ShowSpecs() {
	fmt.Printf("CPU: %s, 内存: %s, 硬盘: %s\n", c.CPU, c.Memory, c.Storage)
}

// --------------------- 建造者接口 ---------------------
type ComputerBuilder interface {
	SetCPU() ComputerBuilder
	SetMemory() ComputerBuilder
	SetStorage() ComputerBuilder
	Build() *Computer
}

// --------------------- 具体建造者：游戏电脑 ---------------------
type GamingComputerBuilder struct {
	computer *Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{computer: &Computer{}}
}

func (b *GamingComputerBuilder) SetCPU() ComputerBuilder {
	b.computer.CPU = "Intel i9-12900K"
	return b
}

func (b *GamingComputerBuilder) SetMemory() ComputerBuilder {
	b.computer.Memory = "32GB DDR5"
	return b
}

func (b *GamingComputerBuilder) SetStorage() ComputerBuilder {
	b.computer.Storage = "1TB NVMe SSD + 2TB HDD"
	return b
}

func (b *GamingComputerBuilder) Build() *Computer {
	return b.computer
}

// --------------------- 具体建造者：办公电脑 ---------------------
type OfficeComputerBuilder struct {
	computer *Computer
}

func NewOfficeComputerBuilder() *OfficeComputerBuilder {
	return &OfficeComputerBuilder{computer: &Computer{}}
}

func (b *OfficeComputerBuilder) SetCPU() ComputerBuilder {
	b.computer.CPU = "Intel i5-12400"
	return b
}

func (b *OfficeComputerBuilder) SetMemory() ComputerBuilder {
	b.computer.Memory = "16GB DDR4"
	return b
}

func (b *OfficeComputerBuilder) SetStorage() ComputerBuilder {
	b.computer.Storage = "512GB SSD"
	return b
}

func (b *OfficeComputerBuilder) Build() *Computer {
	return b.computer
}

// --------------------- 指挥者（可选） ---------------------
type Director struct {
	builder ComputerBuilder
}

func (d *Director) Construct() *Computer {
	return d.builder.
		SetCPU().
		SetMemory().
		SetStorage().
		Build()
}

// --------------------- 客户端调用 ---------------------
func main() {
	// 方案1：直接使用建造者（无指挥者）
	gamingPC := NewGamingComputerBuilder().
		SetCPU().
		SetMemory().
		SetStorage().
		Build()
	fmt.Print("游戏电脑配置: ")
	gamingPC.ShowSpecs()

	// 方案2：通过指挥者构建
	officeBuilder := NewOfficeComputerBuilder()
	director := &Director{builder: officeBuilder}
	officePC := director.Construct()
	fmt.Print("办公电脑配置: ")
	officePC.ShowSpecs()
}
