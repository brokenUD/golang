package singleton

var s *singleton

func init(){
	s = newSingleton()
}

type singleton struct{}

func newSingleton() *singleton{
	return &singleton{}
}

func (s *singleton) Work(){
}

func GetInstance() *singleton{
	return s
}
