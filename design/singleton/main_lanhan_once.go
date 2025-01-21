package singleton

import "sync"

var (
	s2 *singleton2
	once sync.Once
)

type singleton2 struct{}

func newSingleton2() *singleton2{
	return &singleton2{}
}

func (s1 *singleton2) Work(){
}

func GetInstance2() *singleton2{
	once.Do(func(){
		s2 = GetInstance2()
	})
	return s2
}
