package service

import (
	"context"
	"kit/log"
	xtime "kit/time"
	"kit/xstr"
	"reply/model"
	"time"
)

func (s *service) Get(c context.Context, id int64) (r *model.Reply, err error) {

	return
}

func (s *service) Add(c context.Context, reply *model.Reply) (err error) {
	// db
	var affected int64
	if affected, err = s.dao.AddReply(c, reply); err != nil {
		return
	}
	reply.Id = affected
	reply.Created = xtime.Time(time.Now().Unix())

	// cache
	s.ch.Save(func() {
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
	})
	return
}

func (s *service) List(c context.Context, sourceId int64, typeId int8, pn, ps int) (rs []*model.Reply, count int, err error) {
	// cache
	var (
		ok    bool
		ids   []int64
		start = ps * ps
		end   = start + ps
	)
	if ok, err = s.dao.ExpireReplyRedis(c, sourceId, typeId); err != nil {
		return
	}
	if ok {
		if ids, count, err = s.dao.ListReplyRedis(c, sourceId, typeId, start, end-1); err != nil {
			return
		}
		if len(ids) == 0 {
			log.Error("List(%d,%d) len==0", sourceId, typeId)
			return
		}
		if rs, err = s.getReplys(c, sourceId, typeId, ids); err != nil {
			return
		}
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
	} else {
		// todo tomorrow
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

	}
	return
}
