package dao

import (
	"context"
	"fmt"
	"kit/log"
	"reply/model"
)

const (
	_addReplySQL      = "INSERT INTO reply_%s (type_id,source_id,comment,parent_id,path)VALUES(?,?,?,?,?)"
	_selLimitReplySQL = "SELECT id,type_id,source_id,comment,parent_id,path,created FROM reply_%s WHERE source_id=? AND type_id=? ORDER BY created DESCl LIMIT 0,100"
	_selAllReplySQL   = "SELECT id,type_id,source_id,comment,parent_id,path,created FROM reply_%s WHERE source_id=? AND type_id=?"
)

func (d *Dao) sharding(id int64) string {
	return fmt.Sprintf("%02d", id%10)
}

func (d *Dao) AddReply(c context.Context, reply *model.Reply) (affected int64, err error) {
	result, err := d.db.Exec(c, fmt.Sprint(_addReplySQL, d.sharding(reply.SourceId)), reply.TypeId, reply.SourceId, reply.Comment, reply.ParentId, reply.Path)
	if err != nil {
		log.Error("d.db.Exec(%+v) error(%v)", reply, err)
		return
	}
	return result.LastInsertId()
}

func (d *Dao) SelLimitReply(c context.Context, sourceId int64, typeId int8) (rs []*model.Reply, err error) {
	rows, err := d.db.Query(c, fmt.Sprint(_selLimitReplySQL, d.sharding(sourceId)), sourceId, typeId)
	if err != nil {
		log.Error("d.db.Query(%d,%d) error(%v)", sourceId, typeId, err)
		return
	}
	for rows.Next() {
		r := &model.Reply{}
		if err = rows.Scan(); err != nil {
			log.Error("SelLimitReply(%d,%d) rows.Scan error(%v)", sourceId, typeId, err)
			return
		}
		rs = append(rs, r)
	}
	return
}

func (d *Dao) SelAllReply(c context.Context, sourceId int64, typeId int8) (rs []*model.Reply, err error) {
	rows, err := d.db.Query(c, fmt.Sprint(_selAllReplySQL, d.sharding(sourceId)), sourceId, typeId)
	if err != nil {
		log.Error("d.db.Query(%d,%d) error(%v)", sourceId, typeId, err)
		return
	}
	for rows.Next() {
		r := &model.Reply{}
		if err = rows.Scan(); err != nil {
			log.Error("rows.Scan(%d,%d) rows.Scan error(%v)", sourceId, typeId, err)
			return
		}
		rs = append(rs, r)
	}
	return
}
