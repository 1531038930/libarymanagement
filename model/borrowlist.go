package model

import (
	"database/sql"
	"time"
)

type B_list struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uint
	Bookid    uint
	Starttime time.Time
	Endtime   sql.NullTime
}

func (B_list) TableName() string {
	return "BorrowList"
}
func (one *B_list) Add() error {
	//user:=User{UserId:one.UserId}  判断level
	book := Book{Bookid: one.Bookid}
	DB.Debug().Select("available").Where("bookid = ?", book.Bookid).Take(&book)
	if book.Available {
		book.Available = false
		one.Starttime = time.Now()
		work := DB.Begin() //开始事务
		if err := work.Debug().Select("available").Updates(&book).Error; err != nil {
			work.Rollback()
			return err
		}
		if err := work.Debug().Create(&one).Error; err != nil {
			work.Rollback()
			return err
		}
		work.Commit()
		return nil
	} else {
		return AddError{Message: "该书不可借出！"}
	}
}
