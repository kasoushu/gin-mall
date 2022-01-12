package model

type Category struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	ParentId uint64 `json:"parent_id"`
	Level    uint   `json:"level"`
	Sort     uint   `json:"sort"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
