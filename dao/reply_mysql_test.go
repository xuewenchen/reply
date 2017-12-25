package dao

import (
	"context"
	"os"
	"reply/config"
	"reply/model"
	"testing"
)

var (
	dao *Dao
)

func TestMain(m *testing.M) {
	dao = NewDao(config.Conf)
	os.Exit(m.Run())
}

func Test_AddReply(t *testing.T) {
	var err error
	reply := &model.Reply{
		SourceId: int64(1),
		TypeId:   model.NOTE_TYPE,
		Comment:  "这是一段文字",
		ParentId: int64(0),
		Path:     "/",
	}
	if _, err = dao.AddReply(context.Background(), reply); err != nil {
		t.Errorf("Test_AddReply fail error(%v)", err)
	}
	return
}
