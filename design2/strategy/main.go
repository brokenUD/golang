// 策略模式 是一种行为型设计模式，
// 其核心思想是 定义一系列算法，将每个算法封装起来，并使它们可以互相替换

// 策略模式的设计原则
//     开闭原则（OCP）：
//         新增算法时无需修改现有代码，只需扩展新的策略类。
//     单一职责原则（SRP）：
//         每个策略类只负责一个具体的算法实现。
//     组合优于继承：
//         通过组合策略对象实现行为变化，而不是通过继承。


package main

import "fmt"

// --------------------- 策略接口 ---------------------
type DiscountStrategy interface {
	CalculateDiscount(amount float64) float64
}

// --------------------- 具体策略类 ---------------------
// 满减策略：满 200 减 50
type FullReductionStrategy struct{}

func (s *FullReductionStrategy) CalculateDiscount(amount float64) float64 {
	if amount >= 200 {
		return amount - 50
	}
	return amount
}

// 折扣策略：9 折优惠
type PercentageDiscountStrategy struct{}

func (s *PercentageDiscountStrategy) CalculateDiscount(amount float64) float64 {
	return amount * 0.9
}

// 返现策略：返现 10%
type CashBackStrategy struct{}

func (s *CashBackStrategy) CalculateDiscount(amount float64) float64 {
	return amount - (amount * 0.1)
}

// --------------------- 上下文类 ---------------------
type Order struct {
	strategy DiscountStrategy
	amount   float64
}

func NewOrder(amount float64, strategy DiscountStrategy) *Order {
	return &Order{
		amount:   amount,
		strategy: strategy,
	}
}

// 动态切换策略
func (o *Order) SetStrategy(strategy DiscountStrategy) {
	o.strategy = strategy
}

// 计算最终金额
func (o *Order) Checkout() float64 {
	return o.strategy.CalculateDiscount(o.amount)
}

// --------------------- 客户端调用 ---------------------
func main() {
	order := NewOrder(300, &FullReductionStrategy{})
	fmt.Println("满减策略价格:", order.Checkout()) // 300 - 50 = 250

	order.SetStrategy(&PercentageDiscountStrategy{})
	fmt.Println("折扣策略价格:", order.Checkout()) // 300 * 0.9 = 270

	order.SetStrategy(&CashBackStrategy{})
	fmt.Println("返现策略价格:", order.Checkout()) // 300 - 30 = 270
}
