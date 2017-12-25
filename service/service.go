package service

import (
	"kit/db/mysql"
	"reply/config"
	"reply/dao"
)

type service struct {
	dao *dao.Dao
}

func NewService(c *config.Config, db *mysql.DB) (s *service, err error) {
	s = &service{
		dao: dao.NewDao(c),
	}
	return
}

func Close() (err error) {
	return svr.dao.Close()
}
