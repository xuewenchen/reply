package model

import (
	"kit/time"
)

const (
	NOTE_TYPE int8 = 0
)

type Reply struct {
	Id       int64     `json:"id"`
	SourceId int64     `json:"source_id"`
	TypeId   int8      `json:"type_id"`
	Mid      int64     `json:"mid"`
	Comment  string    `json:"comment"`
	ParentId int64     `json:"parent_id"`
	Path     string    `json:"path"`
	State    int8      `json:"state"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Rs       []*Reply  `json:"rs"`
	Rids     []int64   `json:"-"`
}
