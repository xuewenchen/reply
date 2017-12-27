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
	dao      *dao.Dao
	changeCh *channel.Cache
	loadCh   *channel.Cache
}

func NewService(c *config.Config, db *mysql.DB) (s *service, err error) {
	s = &service{
		dao:      dao.NewDao(c),
		changeCh: channel.New(CHSIZE),
		loadCh:   channel.New(CHSIZE),
	}
	return
}
