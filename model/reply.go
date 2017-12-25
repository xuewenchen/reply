package model

import (
	"kit/time"
)

type Reply struct {
	Id       int64     `json:"id"`
	TypeId   int8      `json:"type_id"`
	SourceId int64     `json:"source_id"`
	Comment  string    `json:"comment"`
	ParentId int64     `json:"parent_id"`
	Path     string    `json:"path"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
