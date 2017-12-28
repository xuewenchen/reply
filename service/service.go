package service

import (
	opentracing "github.com/opentracing/opentracing-go"
	"kit/db/channel"
	"reply/config"
	"reply/dao"
	"sync"
)

const (
	CHSIZE = 1024
)

type service struct {
	dao      *dao.Dao
	changeCh *channel.Cache
	loadCh   *channel.Cache
	wait     *sync.WaitGroup
	Tracer   opentracing.Tracer
}

func NewService(c *config.Config, tracer opentracing.Tracer) (s *service, err error) {
	s = &service{
		dao:    dao.NewDao(c),
		Tracer: tracer,
	}
	s.wait = &sync.WaitGroup{}
	s.wait.Add(1)
	s.changeCh = channel.New(CHSIZE, s.wait)
	s.wait.Add(1)
	s.loadCh = channel.New(CHSIZE, s.wait)
	return
}

func (s *service) Close() {
	s.changeCh.Close()
	s.loadCh.Close()
	s.wait.Wait()
}
