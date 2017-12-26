package dao

import (
	"kit/time"
	"reply/model"
	"testing"
)

func Test_ExpireReplyRedis(t *testing.T) {
	if ok, err = d.ExpireReplyRedis(ctx, SOURCEID, TYPEID); err != nil {
		t.Errorf("d.ExpireReplyRedis error(%v)", err)
	} else {
		t.Logf("ok(%v)", ok)
	}
}

func Test_AddReplyRedis(t *testing.T) {
	rs := []*model.Reply{
		&model.Reply{Id: 1, SourceId: SOURCEID, TypeId: TYPEID, Comment: "chen", ParentId: 0, Path: "/", Created: time.Time(1514254618)},
		&model.Reply{Id: 2, SourceId: SOURCEID, TypeId: TYPEID, Comment: "xue", ParentId: 0, Path: "/", Created: time.Time(1514254617)},
		&model.Reply{Id: 3, SourceId: SOURCEID, TypeId: TYPEID, Comment: "wen", ParentId: 0, Path: "/", Created: time.Time(1514254616)},
	}
	if err = d.AddReplyRedis(ctx, SOURCEID, TYPEID, rs); err != nil {
		t.Errorf("d.AddReplyRedis error(%v)", err)
	}
}

func Test_ListReplyRedis(t *testing.T) {
	var (
		start = 0
		end   = 20
		ids   []int64
		count int
	)
	if ids, count, err = d.ListReplyRedis(ctx, SOURCEID, TYPEID, start, end); err != nil {
		t.Errorf("d.ListReplyRedis error(%v)", err)
	} else {
		t.Log(ids, count)
	}
}
