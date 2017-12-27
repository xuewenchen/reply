package service

import (
	"kit/ecode"
	"kit/log"
	"kit/net/context"
	"reply/model"
	"strconv"
)

func hello(c context.Context) {
	res := c.Result()
	res["data"] = "hello"
	return
}

func add(c context.Context) {
	var (
		err      error
		mid      int64
		sourceId int64
		typeId   int64
		comment  string
		parentId int64
	)
	res := c.Result()
	params := c.Request().Form

	midStr := params.Get("mid")
	sourceIdStr := params.Get("source_id")
	typeIdStr := params.Get("type_id")
	comment = params.Get("comment")
	parentIdStr := params.Get("parent_id")

	if mid, err = strconv.ParseInt(midStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", midStr, err)
		res["code"] = ecode.RequestErr
		return
	}
	if sourceId, err = strconv.ParseInt(sourceIdStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", sourceIdStr, err)
		res["code"] = ecode.RequestErr
		return
	}
	if typeId, err = strconv.ParseInt(typeIdStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", typeIdStr, err)
		res["code"] = ecode.RequestErr
		return
	}
	if parentId, err = strconv.ParseInt(parentIdStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", parentIdStr, err)
	} else {
		parentId = 0
	}
	reply := &model.Reply{
		SourceId: sourceId,
		TypeId:   int8(typeId),
		Mid:      mid,
		Comment:  comment,
		ParentId: parentId,
	}
	if err = svr.Add(c, reply); err != nil {
		log.Error("svr.Add(%v) error(%v)", reply, err)
		res["code"] = err
	}
	return
}

func list(c context.Context) {
	var (
		err      error
		sourceId int64
		typeId   int64
		pn       int64
		ps       int64
		count    int
		rs       []*model.Reply
	)
	res := c.Result()
	params := c.Request().Form

	sourceIdStr := params.Get("source_id")
	typeIdStr := params.Get("type_id")
	pnStr := params.Get("pn")
	psStr := params.Get("ps")

	if sourceId, err = strconv.ParseInt(sourceIdStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", sourceIdStr, err)
		res["code"] = ecode.RequestErr
		return
	}
	if typeId, err = strconv.ParseInt(typeIdStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", typeIdStr, err)
		res["code"] = ecode.RequestErr
		return
	}
	if pn, err = strconv.ParseInt(pnStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", pnStr, err)
		res["code"] = ecode.RequestErr
		return
	}
	if ps, err = strconv.ParseInt(psStr, 10, 64); err != nil {
		log.Error("strconv.ParseInt(%s) error(%v)", psStr, err)
		res["code"] = ecode.RequestErr
		return
	}
	if rs, count, err = svr.List(c, sourceId, int8(typeId), int(pn), int(ps)); err != nil {
		res["code"] = err
		return
	}
	res["code"] = err
	res["data"] = rs
	res["count"] = count
	return
}
