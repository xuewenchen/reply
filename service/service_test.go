package service

import (
	"reply/model"
	"testing"
	"time"
)

// func Test_Add(t *testing.T) {
// 	reply := &model.Reply{
// 		SourceId: SOURCEID,
// 		TypeId:   TYPEID,
// 		Mid:      int64(123),
// 		Comment:  "这是一段文字",
// 		ParentId: int64(0),
// 	}
// 	if err = testSvr.Add(ctx, reply); err != nil {
// 		t.Errorf("testSvr.Add error(%v)", err)
// 	}
// }

func Test_List(t *testing.T) {
	var (
		rs    []*model.Reply
		count int
		pn    = 1
		ps    = 20
	)
	if rs, count, err = testSvr.List(ctx, SOURCEID, TYPEID, pn, ps); err != nil {
		t.Errorf("testSvr.List error(%v)", err)
	} else {
		t.Log(rs)
		t.Log(count)
	}
	time.Sleep(time.Second * 2)
}
