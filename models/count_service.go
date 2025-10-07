package models

type CountService struct {
	count int
}

func NewCountService() *CountService {
	return &CountService{
		count: 0,
	}
}

func (s *CountService) IncrementAndGet() int {
	s.count++
	return s.count
}
