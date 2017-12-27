package service

import (
	"context"
	"fmt"
	"kit/ecode"
	"kit/log"
	xtime "kit/time"
	"kit/xstr"
	"reply/model"
	"time"
)

var (
	// 防止缓存穿透的空列表
	_emptyRs = &model.Reply{Id: 0}
)

func (s *service) Add(c context.Context, reply *model.Reply) (err error) {
	if reply.ParentId != 0 {
		var parentReply *model.Reply
		if parentReply, err = s.getReply(c, reply.SourceId, reply.ParentId, reply.TypeId); err != nil {
			return
		}
		if parentReply == nil {
			err = ecode.ReplyNotExist
			return
		} else {
			reply.Path = parentReply.Path + fmt.Sprintf("%d", parentReply.Id) + ","
		}
	}

	var insertId int64
	if insertId, err = s.dao.AddReply(c, reply); err != nil {
		return
	}
	reply.Id = insertId
	reply.Created = xtime.Time(time.Now().Unix())
	// change cache
	s.changeCh.Save(func() {
		var (
			ok  bool
			ctx = context.Background()
		)
		if ok, err = s.dao.ExpireReplyRedis(ctx, reply.SourceId, reply.TypeId); err != nil {
			return
		}
		if ok {
			if err = s.dao.AddReplyRedis(ctx, reply); err != nil {
				return
			}
		}
		if err = s.dao.AddReplyMc(ctx, reply); err != nil {
			return
		}
	})
	return
}

func (s *service) List(c context.Context, sourceId int64, typeId int8, pn, ps int) (rs []*model.Reply, count int, err error) {
	// cache
	var (
		ok    bool
		ids   []int64
		start = (pn - 1) * ps
		end   = start + ps
	)
	if ok, err = s.dao.ExpireReplyRedis(c, sourceId, typeId); err != nil {
		return
	}
	// if redis has
	if ok {
		if ids, count, err = s.dao.ListReplyRedis(c, sourceId, typeId, start, end-1); err != nil {
			return
		}
		if len(ids) == 0 || (len(ids) == 1 && ids[0] == 0) {
			return
		}
		if rs, err = s.getReplys(c, sourceId, typeId, ids); err != nil {
			return
		}
		if err = s.getReplyMap(c, rs, sourceId, typeId); err != nil {
			return
		}
		return
	}
	if count, err = s.dao.CountReply(c, sourceId, typeId); err != nil {
		return
	}
	// start越界
	if count < start {
		return
	}
	if rs, err = s.dao.SelLimitReply(c, sourceId, typeId, start, end); err != nil {
		return
	}
	if len(rs) == 0 {
		return
	}
	if err = s.getReplyMap(c, rs, sourceId, typeId); err != nil {
		return
	}
	// cache回源
	s.loadCh.Save(func() {
		s.loadReply(sourceId, typeId)
	})
	return
}

func (s *service) loadReply(sourceId int64, typeId int8) (err error) {
	var (
		ctx = context.Background()
		rs  []*model.Reply
	)
	if rs, err = s.dao.SelAllReply(ctx, sourceId, typeId); err != nil {
		return
	}
	if len(rs) == 0 {
		if err = s.dao.AddReplyRedis(ctx, _emptyRs); err != nil {
			return
		}
		return
	}
	if err = s.dao.AddReplysRedis(ctx, sourceId, typeId, rs); err != nil {
		return
	}
	if err = s.dao.AddReplysMc(ctx, sourceId, typeId, rs); err != nil {
		return
	}
	return
}

func (s *service) getReplyMap(c context.Context, rs []*model.Reply, sourceId int64, typeId int8) (err error) {
	var (
		allIds   []int64
		allRs    []*model.Reply
		allIdMap = make(map[int64]struct{}, len(rs))
		allRsMap = make(map[int64]*model.Reply, len(rs))
	)
	for _, r := range rs {
		var (
			tempIds []int64
			path    string
		)
		if r.Path != "" {
			path = r.Path
			path = path[:len(path)-1]
			if tempIds, err = xstr.SplitInts(path); err != nil {
				log.Error("xstr.SplitInts(%s) error(%v)", path, err)
				return
			}
		}
		r.Rids = tempIds
		for _, id := range tempIds {
			if _, ok := allIdMap[id]; !ok {
				allIds = append(allIds, id)
				allIdMap[id] = struct{}{}
			}
		}
	}
	if len(allIds) == 0 {
		return
	}
	if allRs, err = s.getReplys(c, sourceId, typeId, allIds); err != nil {
		return
	}
	for _, r := range allRs {
		allRsMap[r.Id] = r
	}
	for _, r := range rs {
		for _, id := range r.Rids {
			if _, ok := allRsMap[id]; ok {
				r.Rs = append(r.Rs, allRsMap[id])
			}
		}
	}
	return
}

func (s *service) getReplys(c context.Context, sourceId int64, typeId int8, ids []int64) (rs []*model.Reply, err error) {
	var (
		missed []int64
		missRs []*model.Reply
	)
	if rs, missed, err = s.dao.GetReplysMc(c, sourceId, typeId, ids); err != nil {
		return
	}
	if len(missed) > 0 {
		if missRs, err = s.dao.SelReplys(c, sourceId, typeId, missed); err != nil {
			return
		}
		rs = append(rs, missRs...)
		if len(missRs) > 0 {
			s.changeCh.Save(func() {
				if err = s.dao.AddReplysMc(context.Background(), sourceId, typeId, missRs); err != nil {
					return
				}
			})
		}
	}
	return
}

func (s *service) getReply(c context.Context, sourceId, id int64, typeId int8) (r *model.Reply, err error) {
	if r, err = s.dao.GetReplyMc(c, sourceId, id, typeId); err != nil {
		return
	}
	if r != nil && r.Id != 0 {
		return
	}
	if r, err = s.dao.SelReply(c, sourceId, id); err != nil {
		return
	}
	if r != nil {
		s.changeCh.Save(func() {
			if err = s.dao.AddReplyMc(c, r); err != nil {
				return
			}
		})
	}
	return
}
