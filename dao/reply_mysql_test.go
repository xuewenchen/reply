package dao

import (
	"context"
	"reply/model"
	"testing"
)

func Test_AddReply(t *testing.T) {
	var err error
	reply := &model.Reply{
		SourceId: SOURCEID,
		TypeId:   TYPEID,
		Mid:      int64(123),
		Comment:  "这是一段文字",
		ParentId: int64(0),
		Path:     "",
	}
	if _, err = d.AddReply(context.Background(), reply); err != nil {
		t.Errorf("Test_AddReply fail error(%v)", err)
	}
	return
}

func Test_SelLimitReply(t *testing.T) {
	var (
		rs    []*model.Reply
		start = 0
		limit = 20
		err   error
	)
	if rs, err = d.SelLimitReply(context.Background(), SOURCEID, TYPEID, start, limit); err != nil {
		t.Errorf("d.SelLimitReply error(%v)", err)
	}
	t.Log(rs)
	return
}

func Test_SelAllReply(t *testing.T) {
	var (
		rs  []*model.Reply
		err error
	)
	if rs, err = d.SelAllReply(context.Background(), SOURCEID, TYPEID); err != nil {
		t.Errorf("d.SelAllReply error(%v)", err)
	}
	t.Log(rs)
	return
}

func Test_SelReplys(t *testing.T) {
	var (
		rs  []*model.Reply
		ids = []int64{1, 2, 3}
		err error
	)
	if rs, err = d.SelReplys(context.Background(), SOURCEID, TYPEID, ids); err != nil {
		t.Errorf("d.SelReplys error(%v)", err)
	}
	t.Log(rs)
	return
}

func Test_CountReply(t *testing.T) {
	var (
		count int
		err   error
	)
	if count, err = d.CountReply(context.Background(), SOURCEID, TYPEID); err != nil {
		t.Errorf("d.CountReply error(%v)", err)
	}
	t.Log(count)
	return
}

func Test_SelReply(t *testing.T) {
	var (
		r   *model.Reply
		id  = int64(1)
		err error
	)
	if r, err = d.SelReply(context.Background(), SOURCEID, id); err != nil {
		t.Errorf("d.SelReplys error(%v)", err)
	}
	t.Log(r)
	return
}
