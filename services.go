package main

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FirstUserInfo(p *FirstUserInfo_P, user *USER, c *context.Context) error {
	return NewDBClient(c).Where(&USER{UserName: p.UserName}).First(&user).Error
}

func Register(p *Register_P, c *context.Context) error {
	if err := NewDBClient(c).Where(&USER{UserName: p.UserName}).First(&USER{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NewDBClient(c).Create(&USER{
				NickName:   p.NickName,
				UserName:   p.UserName,
				Password:   p.Password,
				CreateTime: time.Now().In(G_CONFIG.TIMELOCALTION).Format("2006-01-02 15:04:05"),
				Version:    uuid.NewString(),
			}).Error
		}
		return err
	}
	return errors.New("the account has been registered")
}

func Login(p *Login_P, user *USER, c *context.Context) error {
	if err := NewDBClient(c).Where(&USER{UserName: p.UserName}).First(&user).Error; err != nil {
		return errors.New("login failed")
	}
	if user.Password != p.Password {
		return errors.New("password is error")
	}
	return nil
}
