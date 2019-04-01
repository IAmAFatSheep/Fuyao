package main

import "time"

type Fuyao struct {
	ObjectType string    `json:"docType"`
	Name       string    `json:"Name"`     // 商品名称
	CodeNo     string    `json:"CodeNo"`   // 编码号
	BirthDay   time.Time `json:"BirthDay"` // 时间戳
	Deadline   int       `json:"SaveTime"` // 保存时间(天)
	Photo      string    `json:"Photo"`    // 照片
}
