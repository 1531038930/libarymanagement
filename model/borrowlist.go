package model

import (
	"database/sql"
	"fmt"
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
	if book.Available { //书籍可借
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
func (one *B_list) End() error {
	DB.Debug().Where("endtime is null and user_id= ? and bookid= ?", one.UserId, one.Bookid).Take(&one) //SELECT * FROM `BorrowList` WHERE endtime is null and user_id= 5 and bookid= 2 LIMIT 1
	fmt.Println(one)
	fmt.Println(one.Starttime)
	return nil
}
