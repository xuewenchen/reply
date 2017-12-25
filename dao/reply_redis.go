package dao

import (
	"context"
	"fmt"
	"golang/redigo/redis"
	"kit/log"
	"reply/model"
)

const (
	_replyIndexKey = "rs_%d_%d" // sourceId_typeId
)

func (d *Dao) ExpireReplyRedis(c context.Context, sourceId int64, typeId int8) (err error) {
	conn := d.redis.Get(c)
	defer conn.Close()
	key := fmt.Sprintf(_replyIndexKey, sourceId, typeId)
	if _, err = conn.Do(key, d.expireRedis); err != nil {
		log.Error("ExpireReplyRedis(%s) error(%v)", key, err)
	}
	return
}

func (d *Dao) AddReplyRedis(c context.Context, sourceId int64, typeId int8, rs []*model.Reply) (err error) {
	conn := d.redis.Get(c)
	defer conn.Close()
	key := fmt.Sprintf(_replyIndexKey, sourceId, typeId)
	for _, r := range rs {
		if err = conn.Send("ZADD", key, r.Created.Time().Unix(), r.Id); err != nil {
			log.Error("conn.Send(ZADD %s, %d, %d) error(%v)", key, r.Created.Time().Unix(), r.Id, err)
			return
		}
	}
	if err = conn.Send("EXPIRE", key, d.expireRedis); err != nil {
		log.Error("conn.Send(EXPIRE,%s) error(%v)", key, err)
		return
	}
	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush(%s) error(%v)", key, err)
		return
	}
	for i := 0; i < len(rs)+1; i++ {
		if _, err = conn.Receive(); err != nil {
			log.Error("conn.Receive(%s) error(%v)", key, err)
			return
		}
	}
	return
}

func (d *Dao) ListReplyRedis(c context.Context, sourceId int64, typeId int8, start, end int) (ids []int64, count int, err error) {
	conn := d.redis.Get(c)
	defer conn.Close()
	key := fmt.Sprintf(_replyIndexKey, sourceId, typeId)
	if err = conn.Send("ZREVRANGE", key, start, end); err != nil {
		log.Error("conn.Send(ZREVRANGE %s, %d, %d) error(%v)", key, start, end, err)
		return
	}
	if err = conn.Send("ZCARD", key); err != nil {
		log.Error("conn.Send(ZCARD %s) error(%v)", key, err)
		return
	}
	if ids, err = redis.Int64s(conn.Receive()); err != nil {
		log.Error("redis.Int64s(%s) error(%v)", key, err)
		return
	}
	if count, err = redis.Int(conn.Receive()); err != nil {
		log.Error("redis.Int(%s) error(%v)", key, err)
	}
	return
}
