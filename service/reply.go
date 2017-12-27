package service

import (
	"context"
	xtime "kit/time"
	"kit/xstr"
	"reply/model"
	"time"
)

func (s *service) Add(c context.Context, reply *model.Reply) (err error) {
	var affected int64
	if affected, err = s.dao.AddReply(c, reply); err != nil {
		return
	}
	reply.Id = affected
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
		start = pn * ps
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
	// 没有数据塞一个空的
	if len(rs) == 0 {
		if err = s.dao.AddReplyRedis(ctx, &model.Reply{Id: 0}); err != nil {
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
		tempIds, _ := xstr.SplitInts(r.Path)
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
