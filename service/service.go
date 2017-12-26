package service

import (
	"kit/db/channel"
	"kit/db/mysql"
	"reply/config"
	"reply/dao"
)

const (
	CHSIZE = 1024
)

type service struct {
	dao *dao.Dao
	ch  *channel.Cache
}

func NewService(c *config.Config, db *mysql.DB) (s *service, err error) {
	s = &service{
		dao: dao.NewDao(c),
		ch:  channel.New(CHSIZE),
	}
	return
}
