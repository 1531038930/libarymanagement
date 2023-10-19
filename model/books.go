package model

import "time"

type Book struct {
	Bookid     uint   `gorm:"primaryKey;autoIncrement"`
	Bookname   string `gorm:"not null;size:200" form:"bookname"`
	Author     string `gorm:"not null;size:100" form:"author"`
	Available  bool   `gorm:"not null;"`
	CreateDate int    `gorm:"not null"`
	Delflag    bool
}

func (Book) TableName() string {
	return "Books"
}
func (one *Book) Add() error { //增加book
	one.CreateDate = int(time.Now().Unix())
	one.Available = true
	if len(one.Bookname) == 0 || len(one.Author) == 0 { //若有必填项为空白
		return AddError{"信息缺失"}
	}
	res := DB.Debug().Create(one)
	return res.Error
}
