package dao

import (
	"context"
	"fmt"
	"kit/log"
	"kit/xstr"
	"reply/model"
)

const (
	// add
	_addReplySQL = "INSERT INTO reply_%s (source_id,type_id,mid,comment,parent_id,path)VALUES(?,?,?,?,?,?)"
	// select
	_selReplysSQL     = "SELECT id,source_id,type_id,mid,comment,parent_id,path,created FROM reply_%s WHERE source_id=? AND type_id=? AND state=0 AND id IN(%s)"
	_selLimitReplySQL = "SELECT id,source_id,type_id,mid,comment,parent_id,path,created FROM reply_%s WHERE source_id=? AND type_id=? AND state=0 ORDER BY created DESC LIMIT ?,?"
	_selAllReplySQL   = "SELECT id,source_id,type_id,mid,comment,parent_id,path,created FROM reply_%s WHERE source_id=? AND type_id=? AND state=0"
	_countReplySQL    = "SELECT count(id) FROM reply_%s WHERE source_id=? AND type_id=? AND state=0 "
)

func (d *Dao) sharding(id int64) string {
	return fmt.Sprintf("%02d", id%10)
}

func (d *Dao) AddReply(c context.Context, reply *model.Reply) (affected int64, err error) {
	result, err := d.db.Exec(c, fmt.Sprintf(_addReplySQL, d.sharding(reply.SourceId)), reply.SourceId, reply.TypeId, reply.Mid, reply.Comment, reply.ParentId, reply.Path)
	if err != nil {
		log.Error("d.db.Exec(%+v) error(%v)", reply, err)
		return
	}
	return result.LastInsertId()
}

func (d *Dao) SelReplys(c context.Context, sourceId int64, typeId int8, ids []int64) (rs []*model.Reply, err error) {
	rows, err := d.db.Query(c, fmt.Sprintf(_selReplysSQL, d.sharding(sourceId), xstr.JoinInts(ids)), sourceId, typeId)
	if err != nil {
		log.Error("d.db.Query(%d,%d) error(%v)", sourceId, typeId, err)
		return
	}
	for rows.Next() {
		r := &model.Reply{}
		if err = rows.Scan(&r.Id, &r.SourceId, &r.TypeId, &r.Mid, &r.Comment, &r.ParentId, &r.Path, &r.Created); err != nil {
			log.Error("rows.Scan(%d,%d) rows.Scan error(%v)", sourceId, typeId, err)
			return
		}
		rs = append(rs, r)
	}
	return
}

func (d *Dao) SelLimitReply(c context.Context, sourceId int64, typeId int8, start, limit int) (rs []*model.Reply, err error) {
	rows, err := d.db.Query(c, fmt.Sprintf(_selLimitReplySQL, d.sharding(sourceId)), sourceId, typeId, start, limit)
	if err != nil {
		log.Error("d.db.Query(%d,%d) error(%v)", sourceId, typeId, err)
		return
	}
	for rows.Next() {
		r := &model.Reply{}
		if err = rows.Scan(&r.Id, &r.SourceId, &r.TypeId, &r.Mid, &r.Comment, &r.ParentId, &r.Path, &r.Created); err != nil {
			log.Error("rows.Scan(%d,%d) rows.Scan error(%v)", sourceId, typeId, err)
			return
		}
		rs = append(rs, r)
	}
	return
}

func (d *Dao) SelAllReply(c context.Context, sourceId int64, typeId int8) (rs []*model.Reply, err error) {
	rows, err := d.db.Query(c, fmt.Sprintf(_selAllReplySQL, d.sharding(sourceId)), sourceId, typeId)
	if err != nil {
		log.Error("d.db.Query(%d,%d) error(%v)", sourceId, typeId, err)
		return
	}
	for rows.Next() {
		r := &model.Reply{}
		if err = rows.Scan(&r.Id, &r.SourceId, &r.TypeId, &r.Mid, &r.Comment, &r.ParentId, &r.Path, &r.Created); err != nil {
			log.Error("rows.Scan(%d,%d) rows.Scan error(%v)", sourceId, typeId, err)
			return
		}
		rs = append(rs, r)
	}
	return
}

func (d *Dao) CountReply(c context.Context, sourceId int64, typeId int8) (count int, err error) {
	row := d.db.QueryRow(c, fmt.Sprintf(_selAllReplySQL, d.sharding(sourceId)), sourceId, typeId)
	if err = row.Scan(&count); err != nil {
		log.Error("row.Scan(%d,%d) error(%v)", sourceId, typeId, err)
	}
	return
}
