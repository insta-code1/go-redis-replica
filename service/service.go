package service

import (
	"github.com/insta-code1/go-redis-replica/dao"
)

// Service service.
type Service struct {
	Dao *dao.Dao
}

// New new a service and return.
func New() (s *Service) {
	s = &Service{
		Dao: dao.New(),
	}
	return s
}
