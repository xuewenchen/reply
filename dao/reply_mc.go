package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"golang/gomemcache/memcache"
	"kit/log"
	"reply/model"
)

const (
	_replyKey = "r_%d_%d_%d" //sourceId_typeId_ReplyId
)

func (d *Dao) mcKey(sourceId, replyId int64, typeId int8) string {
	return fmt.Sprintf(_replyKey, sourceId, typeId, replyId)
}

func (d *Dao) AddReplysMc(c context.Context, sourceId int64, typeId int8, rs []*model.Reply) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	for _, r := range rs {
		var bs []byte
		if bs, err = json.Marshal(r); err != nil {
			log.Error("json.Marshal(%+v) error(%v)", r, err)
			return
		}
		key := d.mcKey(sourceId, r.Id, typeId)
		if err = conn.Store("set", key, bs, 0, int32(d.expireMc), 0); err != nil {
			log.Error("conn.Store(set,%s,%s) error(%v)", key, bs, err)
			return
		}
	}
	return
}

func (d *Dao) GetReplyMc(c context.Context, sourceId, ReplyId int64, typeId int8) (reply *model.Reply, err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := d.mcKey(sourceId, ReplyId, typeId)
	err = conn.Get("get", func(r *memcache.Reply) {
		reply = &model.Reply{}
		if err := json.Unmarshal(r.Value, r); err != nil {
			log.Error("json.Unmarshal(%v) error(%v)", r.Value, err)
			return
		}
	}, key)
	return
}

func (d *Dao) GetReplysMc(c context.Context, sourceId int64, typeId int8, ids []int64) (rs []*model.Reply, missed []int64, err error) {
	var (
		mm   = make(map[string]int64, len(ids))
		keys = make([]string, 0, len(ids))
	)
	for _, id := range ids {
		key := d.mcKey(sourceId, id, typeId)
		if _, ok := mm[key]; !ok {
			keys = append(keys, key)
			mm[key] = id
		}
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	err = conn.Get("get", func(r *memcache.Reply) {
		reply := &model.Reply{}
		if err := json.Unmarshal(r.Value, reply); err != nil {
			log.Error("json.Unmarshal(%s) error(%v)", r.Value, err)
			return
		}
		if reply != nil {
			if reply.Id != 0 {
				rs = append(rs, reply)
				delete(mm, r.Key)
			}
		}
	}, keys...)
	if err != nil {
		log.Error("conn.Get(%v) error(%v)", ids, err)
		return
	}
	for _, id := range mm {
		missed = append(missed, id)
	}
	return
}
