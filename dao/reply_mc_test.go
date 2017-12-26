package dao

import (
	"kit/time"
	"reply/model"
	"testing"
)

func Test_AddReplysMc(t *testing.T) {
	rs := []*model.Reply{
		&model.Reply{Id: 1, SourceId: SOURCEID, TypeId: TYPEID, Comment: "chen", ParentId: 0, Path: "/", Created: time.Time(1514254618)},
		&model.Reply{Id: 2, SourceId: SOURCEID, TypeId: TYPEID, Comment: "xue", ParentId: 0, Path: "/", Created: time.Time(1514254617)},
		&model.Reply{Id: 3, SourceId: SOURCEID, TypeId: TYPEID, Comment: "wen", ParentId: 0, Path: "/", Created: time.Time(1514254616)},
	}
	if err = d.AddReplysMc(ctx, SOURCEID, TYPEID, rs); err != nil {
		t.Errorf("d.AddReplysMc error(%v)", err)
	}
}

func Test_GetReplyMc(t *testing.T) {
	var r *model.Reply
	if r, err = d.GetReplyMc(ctx, SOURCEID, int64(1), TYPEID); err != nil {
		t.Errorf("d.GetReplyMc error(%v)", err)
	} else {
		t.Log(r)
	}
}

func Test_GetReplysMc(t *testing.T) {
	var (
		rs     []*model.Reply
		ids    = []int64{1, 2, 3, 4}
		missed []int64
	)
	if rs, missed, err = d.GetReplysMc(ctx, SOURCEID, TYPEID, ids); err != nil {
		t.Errorf("d.GetReplysMc error(%v)", err)
	} else {
		for _, r := range rs {
			t.Log(r)
		}
		t.Log(missed)
	}
}
