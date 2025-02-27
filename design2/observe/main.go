package main

import (
	"fmt"
	"sync"
	"time"
)

type WeatherData struct {
	humidity    float64
	temperature float64
}

type Observer interface {
	Notify(data WeatherData)
}

type Display struct {
	Name string
}

func (d *Display) Notify(data WeatherData) {
	fmt.Printf("%s, %.2f, %.2f\n", d.Name, data.humidity, data.temperature)
}

type WeatherStation struct {
	observers map[Observer]chan WeatherData
	mu        sync.Mutex
}

func NewWeatherStation() *WeatherStation {
	return &WeatherStation{
		observers: make(map[Observer]chan WeatherData),
	}
}

func (w *WeatherStation) RegisterObserver(o Observer) {
	w.mu.Lock()
	defer w.mu.Unlock()
	ch := make(chan WeatherData, 10)
	w.observers[o] = ch

	go func() {
		for data := range ch {
			o.Notify(data)
		}
		fmt.Printf("Observer %s stopped\n", o.(*Display).Name)
	}()
}

func (w *WeatherStation) RemoveObserver(o Observer) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if ch, ok := w.observers[o]; ok {
		close(ch)
		delete(w.observers, o)
	}
}

func (w *WeatherStation) NotifyObservers(data WeatherData) {
	w.mu.Lock()
	defer w.mu.Unlock()

	for _, ch := range w.observers {
		ch <- data
	}
}

func (w *WeatherStation) Start() {
	for i := 0; i < 5; i++ {
		data := WeatherData{
			temperature: 20 + float64(i)*2,
			humidity:    50 + float64(i)*5,
		}
		w.NotifyObservers(data)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	weatherStation := NewWeatherStation()

	dis1 := &Display{Name: "dis 1"}
	dis2 := &Display{Name: "dis 2"}

	weatherStation.RegisterObserver(dis1)
	weatherStation.RegisterObserver(dis2)

	go weatherStation.Start()

	time.Sleep(6 * time.Second)

	weatherStation.RegisterObserver(dis1)

	go weatherStation.Start()

	time.Sleep(10 * time.Second)
}


// 观察者模式的核心思想是解耦，它将观察者和被观察者分离，使得它们可以独立变化

// 观察者模式的适用场景
//     当一个对象的状态变化需要通知其他对象，且这些对象的数量和类型可能随时变化时。
//     当一个对象需要将自己的变化通知给其他对象，但又希望与这些对象解耦时。
//     在事件驱动系统中，例如 GUI 框架、消息队列等

// 观察者模式的优点
//     解耦：
//         观察者和被观察者之间是松耦合的，它们可以独立变化。
//         主题不需要知道观察者的具体实现，只需要知道观察者接口。
//     动态关系：
//         可以在运行时动态添加或移除观察者。
//     支持广播通信：
//         主题可以一次性通知所有观察者。

// 观察者模式的缺点
//     性能问题：
//         如果观察者数量过多，通知所有观察者可能会导致性能问题。
//     循环依赖：
//         如果观察者和主题之间相互依赖，可能会导致循环调用。
//     更新顺序不确定：
//         观察者的更新顺序可能不确定，导致某些观察者获取到不一致的状态。
