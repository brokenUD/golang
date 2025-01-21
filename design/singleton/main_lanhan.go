package singleton

import "sync"

var (
	s1 *singleton1
	sm sync.Mutex
)

type singleton1 struct{}

func newSingleton1() *singleton1{
	return &singleton1{}
}

func (s1 *singleton1) Work(){
}

func GetInstance1() *singleton1{
	if s1 != nil {
		return s1
	}
	sm.Lock()
	defer sm.Unlock()
	if s1 != nil {
		return s1
	}
	s1 = newSingleton1()
	return s1
}
