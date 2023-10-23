package model

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	UserId     uint   `gorm:"primaryKey;autoIncrement"`
	UserName   string `gorm:"not null;unique;size:100" form:"username"`
	Pwd        string `gorm:"not null;size:100" form:"pwd"`
	Name       string `gorm:"not null;size:20" form:"name"`
	CreateDate int    `gorm:"not null"`
	DelDate    int
	Level      uint `gorm:"not null" form:"level"`
}

func (User) TableName() string {
	return "Users"
}

type AddError struct {
	Message string
}

func (e AddError) Error() string {
	return e.Message
}

func (one *User) Add() error {
	var query User
	res := DB.Debug().Where("user_name = ? ", one.UserName).Take(&query) //是否有重复的user_name
	if res.RowsAffected != 0 {
		return AddError{"已存在该用户名!"}
	} else {
		if len(one.Name) == 0 || len(one.Pwd) == 0 || len(one.UserName) == 0 || len(one.Name) == 0 {
			return AddError{"信息缺失"}
		}
		bytestr, err := bcrypt.GenerateFromPassword([]byte(one.Pwd), bcrypt.DefaultCost) //生成哈希
		if err != nil {
			return err
		}
		one.Pwd = string(bytestr) //密码使用加密后的存储
		fmt.Println(time.Now().Unix(), one.Pwd)
		one.CreateDate = int(time.Now().Unix()) //添加时间戳
		res = DB.Debug().Create(one)
		return nil
	}
}

func (one *User) Check() error {
	var ckuser User //通过user_id去数据库查询
	res := DB.Debug().Select("user_name", "user_id", "pwd", "level").Where("user_name = ? and del_date= 0", one.UserName).Take(&ckuser)
	one.UserId = ckuser.UserId
	if res.Error != nil {
		return res.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(ckuser.Pwd), []byte(one.Pwd)) //将密码与数据库加密后的匹配
	return err
}
func (one *User) CheckLevel() error {
	var ckuser User //通过user_id去数据库查询
	res := DB.Debug().Select("level").Where("user_id = ? and del_date= 0", one.UserId).Take(&ckuser)
	if res.Error != nil {
		return res.Error
	}
	if ckuser.Level != 1 { //1权限比2高
		return errors.New("权限不足")
	}
	return nil
}
