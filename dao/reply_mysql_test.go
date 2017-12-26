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
		Comment:  "这是一段文字",
		ParentId: int64(0),
		Path:     "/",
	}
	if _, err = d.AddReply(context.Background(), reply); err != nil {
		t.Errorf("Test_AddReply fail error(%v)", err)
	}
	return
}

func Test_SelLimitReply(t *testing.T) {
	var (
		rs    []*model.Reply
		start = int64(1)
		limit = int64(20)
		err   error
	)
	if rs, err = d.SelLimitReply(context.Background(), SOURCEID, start, limit, TYPEID); err != nil {
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
