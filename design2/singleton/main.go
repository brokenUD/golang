package main

import "sync"

type singleton struct{}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	if s1 == s2 {
		println("s1 ans s2 are the same")
	}
}
