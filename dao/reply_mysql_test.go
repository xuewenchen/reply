package dao

import (
	"context"
	"os"
	"reply/config"
	"reply/model"
	"testing"
)

var (
	dao      *Dao
	SourceId = int64(1)
	TypeId   = model.NOTE_TYPE
)

func TestMain(m *testing.M) {
	dao = NewDao(config.Conf)
	os.Exit(m.Run())
}

// func Test_AddReply(t *testing.T) {
// 	var err error
// 	reply := &model.Reply{
// 		SourceId: SourceId,
// 		TypeId:   TypeId,
// 		Comment:  "这是一段文字",
// 		ParentId: int64(0),
// 		Path:     "/",
// 	}
// 	if _, err = dao.AddReply(context.Background(), reply); err != nil {
// 		t.Errorf("Test_AddReply fail error(%v)", err)
// 	}
// 	return
// }

func Test_SelLimitReply(t *testing.T) {
	var (
		rs  []*model.Reply
		err error
	)
	if rs, err = dao.SelLimitReply(context.Background(), SourceId, TypeId); err != nil {
		t.Errorf("dao.SelLimitReply error(%v)", err)
	}
	t.Log(rs)
	return
}

func Test_SelAllReply(t *testing.T) {
	var (
		rs  []*model.Reply
		err error
	)
	if rs, err = dao.SelAllReply(context.Background(), SourceId, TypeId); err != nil {
		t.Errorf("dao.SelAllReply error(%v)", err)
	}
	t.Log(rs)
	return
}
